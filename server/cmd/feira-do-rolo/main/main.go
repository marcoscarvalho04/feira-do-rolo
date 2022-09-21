package main

import (
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/config/db"
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/config/routes"
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/config/services"
)

func main() {
	db, err := db.ConfigureDatabase()
	if err != nil {
		panic(err)
	}
	services.ConfigureAllServices(db)
	routes.ConfigureServer()

}
