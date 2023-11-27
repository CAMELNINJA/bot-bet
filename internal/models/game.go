package models

/*

CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_users (
    id SERIAL PRIMARY KEY,
    game_id INTEGER NOT NULL REFERENCES games(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    is_winner BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_user_bets (
    id SERIAL PRIMARY KEY,
    game_user_id INTEGER NOT NULL REFERENCES game_users(id),
    bet INTEGER NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
*/

type Game struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type GameUser struct {
	ID     int    `json:"id" db:"id"`
	GameID int    `json:"game_id" db:"game_id"`
	Name   string `json:"name" db:"name"`
	IsWin  bool   `json:"is_win" db:"is_win"`
}

type GameUserBet struct {
	ID        int `json:"-" db:"id"`
	SessionID int `json:"session_id" db:"game_user_id"`
	Bet       int `json:"sum_bet" db:"bet"`
	UserID    int `json:"user_id" db:"user_id"`
}

type GameWithUsers struct {
	Game
	Balance int                    `json:"balance" db:"balance"`
	Users   []*GameUserBetWithUser `json:"users"`
}

type GameUserBetWithUser struct {
	GameUser
	SumBet int `json:"sum_bet" db:"sum_bet"`
}
