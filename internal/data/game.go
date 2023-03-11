package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/sebki/playlist/internal/validator"
)

type Game struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"-"`
	ModifiedAt    time.Time `json:"-"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	YearPublished int32     `json:"year_published"`
	Thumbnail     string    `json:"thumbnail,omitempty"`
	Image         string    `json:"image,omitempty"`
	MinPlayer     int32     `json:"min_player"`
	MaxPlayer     int32     `json:"max_plaxer"`
	MinPlaytime   int32     `json:"min_playtime"`
	MaxPlaytime   int32     `json:"max_playtime"`
	MinAge        int32     `json:"min_age"`
	MaxAge        int32     `json:"max_age"`
	// Tags          *[]Tag          `json:"tags,omitempty"`
	// Persons       *[]Person       `json:"persons"`
	// Orgs          *[]Organization `json:"organizations"`
	Version int32 `json:"version"`
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(game.Description == "", "description", "must be provided")
	v.Check(game.YearPublished >= 0, "year_published", "must be greater than 0")
	v.Check(game.YearPublished <= int32(time.Now().Year()+1), "year_published", "maximum one year ahead allowed")
	v.Check(game.MinPlayer >= 1, "min_player", "must not be less than 1")
	v.Check(game.MinPlayer <= 200, "min_player", "must be less than 200")
	v.Check(game.MaxPlayer >= 1, "max_player", "must not be less than 1")
	v.Check(game.MaxPlayer <= 200, "max_player", "must be less than 200")
	v.Check(game.MinPlaytime >= 0, "min_playtime", "must not be less than 0")
	v.Check(game.MinPlaytime <= 1000, "min_playtime", "must be less than 1000")
	v.Check(game.MaxPlaytime >= 0, "max_playtime", "must not be less than 0")
	v.Check(game.MaxPlaytime < 1000, "max_playtime", "must be less than 1000")
	v.Check(game.MinAge >= 0, "min_age", "must not be less than 0")
	v.Check(game.MinAge < 150, "min_age", "must be less than 150")
	v.Check(game.MaxAge >= 0, "max_age", "must not be less than 0")
	v.Check(game.MaxAge < 150, "max_age", "must be less than 150")
}

type GameModel struct {
	DB *sql.DB
}

func (m GameModel) Insert(game *Game) error {
	query := `
		INSERT INTO games (title, descr, year_published, thumbnail, image, min_player, max_player, min_playtime, max_playtime, min_age, max_age)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_at, modified_at, version
	`
	args := []any{
		game.Title,
		game.Description,
		game.YearPublished,
		game.Thumbnail,
		game.Image,
		game.MinPlayer,
		game.MaxPlayer,
		game.MinPlaytime,
		game.MaxPlaytime,
		game.MinAge,
		game.MaxAge,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&game.ID, &game.CreatedAt, &game.ModifiedAt, &game.Version)
}

func (m GameModel) Get(id int64) (*Game, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, modified_at, title, descr, year_published, thumbnail, image, min_player, max_player, min_playtime, max_playtime, min_age, max_age, version
		FROM games
		WHERE id = $1
	`

	var game Game

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		game.ID,
		game.CreatedAt,
		game.ModifiedAt,
		game.Title,
		game.Description,
		game.YearPublished,
		game.Thumbnail,
		game.Image,
		game.MinPlayer,
		game.MaxPlayer,
		game.MinPlaytime,
		game.MaxPlaytime,
		game.MinAge,
		game.MaxAge,
		game.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &game, nil
}

func (m GameModel) Update(game *Game) error {
	// check for version, update all editable fields, increment and return version
	query := `
		UPDATE game 
		SET modified_at = transaction_timestamp(), title = $1, desc = $2, year_published= $3, thumbnail = $4, image = $5, min_player = $6, max_player = $7, min_playtime = $8, max_playtime = $9, min_age = $10, max_age = $12, version = version + 1
		WHERE id = $12 AND version = $13
		RETURNING modified_at, version
	`
	args := []any{
		game.Title,
		game.Description,
		game.YearPublished,
		game.Thumbnail,
		game.Image,
		game.MinPlayer,
		game.MaxPlayer,
		game.MinPlaytime,
		game.MaxPlaytime,
		game.MinAge,
		game.MaxAge,
		game.ID,
		game.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&game.ModifiedAt, &game.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m GameModel) Delete(id int64) error {
	if id < 0 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM games
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m GameModel) GetAll(title string, bgType []string, filters Filters) ([]*Game, Metadata, error) {
	// TODO: filtering by tags
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, modified_at, title, descr, year_published, thumbnail, image, min_player, max_player, min_playtime, max_playtime, min_age, max_age, version
		FROM games
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{title, pq.Array(bgType), filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0

	games := []*Game{}

	for rows.Next() {
		var game Game
		err := rows.Scan(
			&game.ID,
			game.CreatedAt,
			game.ModifiedAt,
			game.Title,
			game.Description,
			game.YearPublished,
			game.Thumbnail,
			game.Image,
			game.MinPlayer,
			game.MaxPlayer,
			game.MinPlaytime,
			game.MaxPlaytime,
			game.MinAge,
			game.MaxAge,
			game.Version,
			totalRecords,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		games = append(games, &game)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return games, metadata, nil
}
