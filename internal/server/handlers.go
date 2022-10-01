package server

import (
	"avito-backend-trainee/internal/generated"
	"avito-backend-trainee/internal/storage"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Logger *zap.Logger
	Store  Storager
}

func (h *Handler) GetAdList(w http.ResponseWriter, r *http.Request) {
	var hand *generated.AdListRequest

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {

		switch {
		case errors.Is(err, storage.ErrBadOrderType):
			http.Error(w, "wrong value of \"ordBy\" type", http.StatusBadRequest)
			return

		case errors.Is(err, storage.ErrBadPaginationFields):
			http.Error(w, "wrong value of \"PaginationFields\" type", http.StatusBadRequest)
			return

		default:
			http.Error(w, "malformed request body", http.StatusBadRequest)
			return
		}

	}

	result, err := h.Store.ReadAllAds(r.Context(), hand.Order, hand.Offset, storage.PaginationFields(*hand.Pagination))
	if err != nil {
		//
		return
	}

	marshalledRequest, err := json.Marshal(result)
	if err != nil {
		//
		return
	}
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(marshalledRequest)
	if err != nil {
		h.Logger.Error("", zap.Error(writeErr))
	}

}

func (h *Handler) GetAd(w http.ResponseWriter, r *http.Request) {
	var hand *generated.GetAdRequest

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {
		//
		return
	}

	result, err := h.Store.ReadAd(r.Context(), hand.Id)
	if err != nil {
		//
		return
	}

	marshalledRequest, err := json.Marshal(result)
	if err != nil {
		//
		return
	}
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(marshalledRequest)
	if err != nil {
		h.Logger.Error("", zap.Error(writeErr))
	}
}

func (h *Handler) CreateNewAd(w http.ResponseWriter, r *http.Request) {
	var hand *generated.NewAdRequest

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {
		// http.Error(w, "", http.StatusBadRequest)
		return
	}

	switch {
	case len(hand.Name) > 200:
		http.Error(w, "", http.StatusBadRequest)
		return
	case len(hand.Description) > 1000:
		http.Error(w, "", http.StatusBadRequest)
		return
	case len(hand.Links) > 3:
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	res, err := h.Store.AddNewAd(r.Context(), hand.Name, hand.Description, hand.Links, hand.Price)
	if err != nil {
		//
		return
	}

	var result = generated.NewAdResponse{
		Id: res,
		// Code:
	}

	marshalledRequest, err := json.Marshal(result)
	if err != nil {
		//
		return
	}
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(marshalledRequest)
	if err != nil {
		h.Logger.Error("", zap.Error(writeErr))
	}
}
