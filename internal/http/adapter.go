package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(logger *slog.Logger, h *GameHandler) *chi.Mux {
	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(NewLoggerMidelware(logger))

	r.Get("/healthcheck", healthCheckHandler)
	r.Get("/api/v1/game", h.GetLastGame)
	r.Post("/api/v1/game/bet", h.SetBet)
	return r
}

// healthcheck
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func NewLoggerMidelware(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With(
			slog.String("component", "middleware/logger"),
		)
		log.Info("logger middleware enabled")
		// код самого обработчика
		fn := func(w http.ResponseWriter, r *http.Request) {
			// собираем исходную информацию о запросе
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			// создаем обертку вокруг `http.ResponseWriter`
			// для получения сведений об ответе
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			// Момент получения запроса, чтобы вычислить время обработки
			t1 := time.Now()

			// Запись отправится в лог в defer
			// в этот момент запрос уже будет обработан
			defer func() {
				entry.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("duration", time.Since(t1).String()),
				)
			}()
			// Передаем управление следующему обработчику в цепочке middleware
			next.ServeHTTP(ww, r)
		}
		// Возвращаем созданный выше обработчик, приведя его к типу http.HandlerFunc
		return http.HandlerFunc(fn)
	}
}
