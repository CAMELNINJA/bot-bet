package usecase

import (
	"errors"
	"log/slog"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
)

type gameRepo interface {
	CreateGame(game *models.Game) (int32, error)
	CreateGameUser(gameUser *models.GameUser) error
	CreateGameUserBet(gameUserBet *models.GameUserBet) error
	GetLastGame() (*models.Game, error)
	GetGameUsers(gameID int) ([]*models.GameUserBetWithUser, error)
	SetBet(gameUserBet *models.GameUserBet) error

	// GetByTelegramID returns user by telegram id
	GetByTelegramID(id int) (*models.User, error)
}

type Game struct {
	repo gameRepo
	log  slog.Logger
}

func NewGame(repo gameRepo, log slog.Logger) *Game {
	return &Game{
		repo: repo,
		log:  log,
	}
}

func (g *Game) CreateGame(name string, users ...string) error {
	game := &models.Game{
		Name:     name,
		IsActive: true,
	}
	gameID, err := g.repo.CreateGame(game)
	if err != nil {
		return err
	}
	for _, user := range users {
		gameUser := &models.GameUser{
			GameID: int(gameID),
			Name:   user,
			IsWin:  false,
		}
		err := g.repo.CreateGameUser(gameUser)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) GetLastGame() (*models.GameWithUsers, error) {
	pgGame, err := g.repo.GetLastGame()
	if err != nil {
		return nil, err
	}

	pgGameUsers, err := g.repo.GetGameUsers(pgGame.ID)
	if err != nil {
		return nil, err
	}

	return &models.GameWithUsers{
		Game:  *pgGame,
		Users: pgGameUsers,
	}, nil
}

func (g *Game) SetBet(gameUserBet *models.GameUserBet) error {
	if gameUserBet.Bet < 0 {
		return nil
	}
	user, err := g.repo.GetByTelegramID(gameUserBet.SessionID)
	if err != nil {
		return err
	}

	if user.Balance < gameUserBet.Bet {
		return errors.New("not enough money")
	}
	return g.repo.SetBet(gameUserBet)
}
