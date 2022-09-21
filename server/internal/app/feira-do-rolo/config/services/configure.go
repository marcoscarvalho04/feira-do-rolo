package services

import (
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/ofertas"
	"gorm.io/gorm"
)

func ConfigureAllServices(db *gorm.DB) {
	ofertas.Service = ofertas.NewOfertaService(db)

}
