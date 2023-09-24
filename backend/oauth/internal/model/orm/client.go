package orm

import "time"

type Client struct {
	ID          int64 `bun:",autoincrement"`
	ClientId    int64
	Name        string
	RedirectUrl string
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
