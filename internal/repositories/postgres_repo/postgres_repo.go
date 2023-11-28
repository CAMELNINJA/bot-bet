package postgresrepo

import (
	"log/slog"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type PostrgesRepo struct {
	db      *sqlx.DB
	pgxConn *pgx.Conn
	logger  *slog.Logger
}

// NewPostrgesRepo  -.
func NewPostrgesRepo(db *sqlx.DB, pgxConn *pgx.Conn, logger *slog.Logger) *PostrgesRepo {

	return &PostrgesRepo{db: db, pgxConn: pgxConn, logger: logger}
}

func (r *PostrgesRepo) Ping() error {
	return r.db.Ping()
}

func (r *PostrgesRepo) GetByID(id int) (*models.User, error) {
	query := `SELECT id , login , chat_id , factbalance , balance FROM users WHERE id = $1`
	user := &models.User{}
	err := r.db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostrgesRepo) GetByTelegramID(id int) (*models.User, error) {
	query := `SELECT id , login , chat_id , factbalance , balance FROM users WHERE chat_id = $1`
	user := &models.User{}
	err := r.db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostrgesRepo) Create(user *models.User) error {
	query := `INSERT INTO users (login , chat_id , factbalance , balance) VALUES ($1 , $2 , $3 , $4)`
	_, err := r.db.Exec(query, user.UserName, user.ChatID, user.FactBalance, user.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostrgesRepo) Update(user *models.User) error {
	query := `UPDATE users SET login = $1 , chat_id = $2 , factbalance = $3 , balance = $4 WHERE id = $5`
	_, err := r.db.Exec(query, user.UserName, user.ChatID, user.FactBalance, user.Balance, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostrgesRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
