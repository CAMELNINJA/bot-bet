package usecase

import (
	"log/slog"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
)

type gameRepo interface {
	CreateGame(game *models.Game) (int32, error)
	CreateGameUser(gameUser *models.GameUser) error
	CreateGameUserBet(gameUserBet *models.GameUserBet) error
	GetLastGame() (*models.GameWithUsers, error)
	SetBet(gameUserBet *models.GameUserBet) error
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
	g.repo.GetLastGame()
}
