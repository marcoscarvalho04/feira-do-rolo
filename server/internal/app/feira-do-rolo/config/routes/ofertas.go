package routes

import (
	"github.com/gorilla/mux"
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/ofertas"
)

func OfertasRouting(r *mux.Router) {
	ofertaController := ofertas.OfertaController{}
	r.HandleFunc("/offers", ofertaController.GetAll).Methods("GET")
	r.HandleFunc("/offers", ofertaController.Save).Methods("POST")

}
