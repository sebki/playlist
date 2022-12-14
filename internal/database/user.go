package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/sebki/playlist/internal/errors"
	"github.com/sebki/playlist/internal/models"
)

func (db *db) CreateNewUser(username, email, password string) (models.User, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	user := models.User{}

	err := db.checkUserOrEmail(username, email)
	if err != nil {
		if errors.IsValidationError(err) {
			return user, err
		}
		log.Fatal(err)
	}

	u, err := json.Marshal(&user)
	if err != nil {
		return user, err
	}

	mu := &api.Mutation{
		SetJson:   u,
		CommitNow: true,
	}

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		return user, err
	}

	user.Username = username
	user.Email = email
	for _, v := range assigned.Uids {
		user.Uid = v
		break
	}
	return user, nil
}

func (db *db) checkUserOrEmail(username, email string) error {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		user(func: eq(username, %q)) {
			username
		}
		email(func: eq(email, %q)) {
			email
		}
	}`, username, email)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return err
	}

	var data struct {
		Username []struct {
			Name string `json:"username"`
		} `json:"user"`
		Email []struct {
			Email string `json:"email"`
		} `json:"email"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return err
	}

	if len(data.Username) == 0 && len(data.Email) == 0 {
		return nil
	} else {
		for _, v := range data.Username {
			if v.Name == username {
				return errors.ErrUsernameExists
			}
		}
		for _, v := range data.Email {
			if v.Email == email {
				return errors.ErrEmailExists
			}
		}
	}

	return nil
}

func (db *db) LoginByEmail(email, password string) (models.User, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	u := models.User{}

	q := fmt.Sprintf(`
	{
		email(func: eq(email, %q)) {
			uid
			username
			email
			checkpwd(password, %q)
		}
	}`, email, password)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return u, err
	}

	var data struct {
		Email []struct {
			Uid      string `json:"uid"`
			Username string `json:"username"`
			Email    string `json:"email"`
			CheckPwd bool   `json:"checkpwd(password)"`
		} `json:"email"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return u, err
	}

	if len(data.Email) == 0 {
		return u, errors.ErrCredentialsIncorrect
	} else {
		for _, v := range data.Email {
			if v.CheckPwd {
				u.Uid = v.Uid
				u.Username = v.Username
				u.Email = v.Email
				return u, nil
			}
		}
	}

	return u, errors.ErrCredentialsIncorrect
}

func (db *db) LoginByUsername(username, password string) (models.User, error) {
	ctx := context.Background()
	txn := db.Client.NewTxn()
	defer txn.Discard(ctx)

	u := models.User{}

	q := fmt.Sprintf(`
	{
		user(func: eq(username, %q)) {
			uid
			username
			email
			checkpwd(password, %q)
		}
	}`, username, password)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return u, err
	}

	var data struct {
		User []struct {
			Uid      string `json:"uid"`
			Username string `json:"username"`
			Email    string `json:"email"`
			CheckPwd bool   `json:"checkpwd(password)"`
		} `json:"user"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return u, err
	}

	if len(data.User) == 0 {
		return u, errors.ErrCredentialsIncorrect
	} else {
		for _, v := range data.User {
			if v.CheckPwd {
				u.Uid = v.Uid
				u.Username = v.Username
				u.Email = v.Email
				return u, nil
			}
		}
	}

	return u, errors.ErrCredentialsIncorrect
}
