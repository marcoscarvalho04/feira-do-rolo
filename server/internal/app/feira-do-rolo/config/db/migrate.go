package db

import (
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/ofertas"
	"gorm.io/gorm"
)

type AllMigrations interface {
	Migrate() error
}

func migrateAllModels(db *gorm.DB) error {
	allMigrationsToExecute := make([]AllMigrations, 0)
	allMigrationsToExecute = append(allMigrationsToExecute, ofertas.NewOfertaService(db))

	for _, value := range allMigrationsToExecute {
		err := value.Migrate()
		if err != nil {
			return err
		}
	}
	return nil

}
