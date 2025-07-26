package sloghelper

import (
	"context"
	"log/slog"

	"github.com/shibukawa/uuid62/v2"
)

// コンテキストにIDを設定
type _CtxKeyType struct{}

var ctxKey = _CtxKeyType{}

func WithTraceID(ctx context.Context) context.Context {
	traceID, _ := uuid62.V7()
	return context.WithValue(ctx, ctxKey, traceID)
}

// ハンドラーのラッパー
type WriteTraceIDHandler struct {
	parent slog.Handler
}

func WithWriteTraceIDHandler(parent slog.Handler) *WriteTraceIDHandler {
	return &WriteTraceIDHandler{
		parent: parent,
	}
}

// ログ出力に情報を付与するメソッド
func (h *WriteTraceIDHandler) Handle(ctx context.Context, record slog.Record) error {
	record.Add(slog.String("traceID", ctx.Value(ctxKey).(string)))
	return h.parent.Handle(ctx, record)
}

func (h *WriteTraceIDHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *WriteTraceIDHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &WriteTraceIDHandler{h.parent.WithAttrs(attrs)}
}

func (h *WriteTraceIDHandler) WithGroup(name string) slog.Handler {
	return &WriteTraceIDHandler{h.parent.WithGroup(name)}
}

var _ slog.Handler = (*WriteTraceIDHandler)(nil)
