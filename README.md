# API Server for Playlist

## Endpoints
| Method    | URL Pattern               | Handler                   | Action                            |
|-----------|---------------------------|---------------------------|-----------------------------------|
| GET       | /v1/healthcheck           | healthcheckHandler        | Show application info             |
| POST      | /v1/games                 | createGameHandler         | create a new game                 |
| GET       | /v1/games/:id             | showGameHandler           | show details of a specific game   |
| PUT       | /v1/games/:id             | updateGameHandler         | update a specific game            |
| POST      | /v1/playlist              | createPlaylistHandler     | create a playlist                 |
| GET       | /v1/playlist/:id          | showPlaylistHandler       | show a specific playlist          |
| PUT       | /v1/playlist/:id          | updatePlaylistHandler     | update a specific playlist        |
| GET       | /v1/search                | searchHandler             | search for items                  |
