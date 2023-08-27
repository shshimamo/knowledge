package middlewares

import (
	"github.com/shshimamo/knowledge-main/middlewares/sloghelper"
	"log/slog"
	"net/http"
	"os"
)

func NewSlogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// ハンドラとロガーの初期化
		handler := sloghelper.WithWriteTraceIDHandler(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		))
		slog.SetDefault(slog.New(handler))

		// コンテキストに情報を載せてslogを呼ぶ
		ctx = sloghelper.WithTraceID(ctx)
		slog.InfoContext(ctx, "Assign traceID")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
