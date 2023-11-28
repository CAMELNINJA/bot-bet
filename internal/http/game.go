package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/CAMELNINJA/bot-bet.git/internal/models"
	"github.com/CAMELNINJA/bot-bet.git/pkg/handlers"
)

type gameUsecase interface {
	GetLastGame(chatID int) (*models.GameWithUsers, error)
	SetBet(gameUserBet *models.GameUserBet) error
}

type GameHandler struct {
	gameUsecase gameUsecase

	log *slog.Logger
}

func NewGameHandler(gameUsecase gameUsecase, log *slog.Logger) *GameHandler {
	return &GameHandler{
		gameUsecase: gameUsecase,
		log:         log,
	}
}

func (h *GameHandler) GetLastGame(w http.ResponseWriter, r *http.Request) {
	chatIDSTR := r.URL.Query().Get("chat_id")
	if chatIDSTR == "" {
		h.log.Error("chat_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chatID, err := strconv.Atoi(chatIDSTR)
	if err != nil {
		h.log.Error("chat_id is not int", err)
		if err := handlers.AcseptError(w, http.StatusBadRequest, err); err != nil {
			h.log.Error("acsept error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	game, err := h.gameUsecase.GetLastGame(chatID)
	if err != nil {
		h.log.Error("get last game error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if game == nil {
		h.log.Error("game is nil")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = handlers.Acsept(w, http.StatusOK, game)
	if err != nil {
		h.log.Error("acsept error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *GameHandler) SetBet(w http.ResponseWriter, r *http.Request) {
	gameUserBet := &models.GameUserBet{}
	if err := json.NewDecoder(r.Body).Decode(gameUserBet); err != nil {
		h.log.Error("decode json body error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.gameUsecase.SetBet(gameUserBet); err != nil {
		h.log.Error("set bet error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := handlers.Acsept(w, http.StatusOK, nil); err != nil {
		h.log.Error("acsept error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
