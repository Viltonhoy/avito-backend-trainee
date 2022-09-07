package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	Store  Storager
}

type NewAdRequest struct {
	name        string
	description string
	links       []string
	price       decimal.Decimal
}

type NewAdResponse struct {
	id   int
	code int ``
}

type GetAdRequest struct {
	id     int
	fields *string
}

type GetAdResponse struct {
	name          string
	description   *string
	price         decimal.Decimal
	mainPhotoLink string
	links         *[]string
}

type AdListRequest struct {
	command string
}

type AdListResponse struct {
	result []GetAdResponse
}

func (h *Handler) GetAdList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAd(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateNewAd(w http.ResponseWriter, r *http.Request) {
	var hand *NewAdRequest

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {
		// http.Error(w, "", http.StatusBadRequest)
		return
	}

	switch {
	case len(hand.name) > 200:
		//
		return
	case len(hand.description) > 1000:
		//
		return
	case len(hand.links) > 3:
		//
		return
	}

}
