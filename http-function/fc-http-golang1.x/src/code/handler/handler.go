package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type OrderMsg struct {
	Id      string `json:"id"`
	Address string `json:"address"`
	Package struct {
		PkgNo string `json:"pkgNo"`
	} `json:"package"`
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	w.Header().Add("Content-Type", "text/plain")
	var orderMsg OrderMsg
	err := json.NewDecoder(req.Body).Decode(&orderMsg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	log.Printf("Received JSON data: %+v\n", orderMsg)

	err = json.NewEncoder(w).Encode(orderMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
