package sqlite

import (
	"context"
	"errors"
	"fmt"
	"poseidon/internal/module/sqlite/model"
	"poseidon/internal/module/sqlite/repository"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func initSQLite(lc fx.Lifecycle, db *gorm.DB, settingRepo repository.SettingRepository) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var settingDst model.Setting
			if err := db.WithContext(ctx).AutoMigrate(&settingDst); err != nil {
				return fmt.Errorf("failed to migrate Setting model: %w", err)
			}

			_, err := settingRepo.FindOne(ctx)
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}

			if err := settingRepo.Upsert(ctx, &settingDst); err != nil {
				return fmt.Errorf("failed to initialize default setting: %w", err)
			}
			return nil
		},
	})
}

// Module wires the SQLite persistence layer into an Fx application.
var Module = fx.Module("sqlite",
	fx.Provide(
		fx.Private,
		NewSQLite,
	),
	fx.Provide(
		fx.Annotate(
			repository.NewSettingRepository,
			fx.As(new(repository.SettingRepository)),
		),
	),
	fx.Invoke(initSQLite),
)
