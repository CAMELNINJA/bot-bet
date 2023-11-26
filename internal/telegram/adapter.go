package telegram

import (
	"errors"
	"log/slog"

	"github.com/CAMELNINJA/bot-bet.git/internal/config"
	"github.com/CAMELNINJA/bot-bet.git/internal/models"
	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
)

type service interface {
	// GetByID returns user by id
	GetByID(id int) (*models.User, error)
	// GetByTelegramID returns user by telegram id
	GetByTelegramID(id int) (*models.User, error)
	// Create creates new user
	// Update updates user
	Update(user *models.User) error
	// Create creates new user
	Create(user *models.User) error
}

type adapter struct {
	log       slog.Logger
	service   service
	bot       *tgbotapi.BotAPI
	webAppUrl string
}

func NewAdapter(log slog.Logger, service service, webAppUrl string) *adapter {
	return &adapter{
		log:       log,
		service:   service,
		webAppUrl: webAppUrl,
	}
}

func StartNoProxy(c *config.Config) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(c.Telegram.BotToken)
}

const (
	ErrNoDataF          = "Произошла ошибка, попробуйте позже."
	ErrNoData           = "К сожалению, у нас пока нет данных. Информация может доставляться с задержкой"
	ErrHasFight         = "Ты уже участвуешь или запланировал другой бой. Прежде чем начать новый бой, заверши сначала предыдущий."
	ErrNoMonsters       = "Все монстры уже повержены."
	ErrNoHP             = "У тебя недостаточно жизней, чтобы начать бой."
	ErrNoAlivePlayers   = "Этот матч уже завершен."
	ErrNoButtonSelected = "Ошибка, для отправки файла необходимо сначала нажать кнопку"
)

var (
	ErrNoCommandSelected = errors.New("no command selected")
)

type Adapter interface {
	StartBot() error
	StopBot()
}
