package models

type User struct {
	ID          int    `json:"id"`
	UserName    string `json:"user_name" db:"login"`
	ChatID      int    `json:"chat_id" db:"chat_id"`
	FactBalance int    `json:"factbalance"`
	Balance     int    `json:"balance"`
}
