package ofertas

import (
	"encoding/json"
	"fmt"
	"github.com/marcoscarvalho04/feira-do-rolo/internal/app/feira-do-rolo/requests"
	"io/ioutil"
	"net/http"
)

//  Feira do Rolo Api:
//   version: 1.0.0
//   title: Feira do rolo API
//  Schemes: http, https
//  Host: localhost:5000
//  BasePath: /
//  Produces:
//    - application/json
//
// swagger:meta

// TODO: Implements all swagger tags

var Service *OfertaService

type OfertaController struct {
}

func (s OfertaController) GetAll(w http.ResponseWriter, r *http.Request) {
	allOffers, err := Service.GetPaginated(r)
	if len(allOffers) == 0 {
		w.WriteHeader(requests.ERROR_NOT_FOUND)
		return
	}
	if err != nil {
		w.WriteHeader(requests.BAD_REQUEST)
		return
	}
	parsedContent, err := json.Marshal(allOffers)
	if err != nil {
		w.WriteHeader(requests.INTERNAL_ERROR_SERVER)
	}

	w.WriteHeader(requests.OK)
	w.Write(parsedContent)

}

func (s OfertaController) Save(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	offerToSave := &OfertaDTO{}

	if err != nil {
		msg := fmt.Sprintf("error reading body request: %s", err.Error())
		w.WriteHeader(requests.BAD_REQUEST)
		w.Write([]byte(msg))
		return
	}
	err = json.Unmarshal(requestBody, &offerToSave)
	if err != nil {
		msg := fmt.Sprintf("error unmarshelling body request: %s", err.Error())
		w.WriteHeader(requests.BAD_REQUEST)
		w.Write([]byte(msg))
		return
	}

	offerResult, err := Service.Save(*offerToSave)
	if err != nil {
		msg := fmt.Sprintf("error saving new offer: %s", err.Error())
		w.WriteHeader(requests.INTERNAL_ERROR_SERVER)
		w.Write([]byte(msg))
		return
	}
	offerResultBytes, err := json.Marshal(&offerResult)
	if err != nil {
		msg := fmt.Sprintf("error marshelling response: %s", err.Error())
		w.WriteHeader(requests.INTERNAL_ERROR_SERVER)
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(requests.CREATED)
	w.Write(offerResultBytes)
}
