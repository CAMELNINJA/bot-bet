package postgresrepo

import "github.com/CAMELNINJA/bot-bet.git/internal/models"

type gameRepo interface {
	CreateGame(game *models.Game) (int32, error)
	CreateGameUser(gameUser *models.GameUser) error
	CreateGameUserBet(gameUserBet *models.GameUserBet) error
	GetLastGame() (*models.Game, error)
	GetGameUsers(gameID int) ([]*models.GameUserBetWithUser, error)
	SetBet(gameUserBet *models.GameUserBet) error
}

func (r *PostrgesRepo) CreateGame(game *models.Game) (int32, error) {
	query := `INSERT INTO games (name ) VALUES ($1 ) RETURNING id`
	var id int32
	err := r.db.Get(&id, query, game.Name, game.IsActive)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostrgesRepo) CreateGameUser(gameUser *models.GameUser) error {
	query := `INSERT INTO game_users (game_id , name  ) VALUES ($1 , $2 )`
	_, err := r.db.Exec(query, gameUser.GameID, gameUser.Name, gameUser.IsWin)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostrgesRepo) CreateGameUserBet(gameUserBet *models.GameUserBet) error {
	query := `INSERT INTO game_user_bets (game_user_id , bet , user_id ) VALUES ($1 , $2 , $3 )`
	_, err := r.db.Exec(query, gameUserBet.SessionID, gameUserBet.Bet, gameUserBet.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostrgesRepo) GetLastGame() (*models.Game, error) {
	query := `SELECT id , name , is_active FROM games WHERE is_active=true ORDER BY id DESC LIMIT 1`
	game := &models.Game{}
	err := r.db.Get(game, query)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *PostrgesRepo) GetGameUsers(gameID int) ([]*models.GameUserBetWithUser, error) {
	query := `SELECT gu.id , gu.game_id , gu.name , gu.is_win , SUM(gub.bet) as sum_bet FROM game_users gu INNER JOIN game_user_bets gub ON gu.id = gub.game_user_id WHERE gu.game_id = $1 GROUP BY gu.id`
	gameUsers := []*models.GameUserBetWithUser{}
	err := r.db.Select(&gameUsers, query, gameID)
	if err != nil {
		return nil, err
	}
	return gameUsers, nil
}

func (r *PostrgesRepo) SetBet(gameUserBet *models.GameUserBet) error {
	query := `UPDATE game_user_bets SET bet = $1 WHERE id = $2`
	_, err := r.db.Exec(query, gameUserBet.Bet, gameUserBet.ID)
	if err != nil {
		return err
	}
	return nil
}
