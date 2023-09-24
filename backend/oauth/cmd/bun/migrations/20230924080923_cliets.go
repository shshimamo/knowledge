package migrations

import (
	"context"

	"github.com/shshimamo/knowledge/backend/oauth/internal/model/orm"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*orm.Client)(nil)).Exec(ctx)
		if err != nil {
			panic(err)
		}
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*orm.Client)(nil)).IfExists().Exec(ctx)
		if err != nil {
			panic(err)
		}
		return err
	})
}
