package wallet

import (
	"course/internal/constants"
	"course/internal/entities"
	"course/internal/repositories/wallet"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	// Check basic auth exist
	err := basicAuth(r)
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	// Check query exist
	ids, ok := r.URL.Query()["id"]
	if !ok {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var wallets []*entities.Wallet

	for _, id := range ids {
		val, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if v := wallet.GetById(val); v != nil {
			wallets = append(wallets, v)
		}
	}

	jsonData, err := json.Marshal(wallets)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)

}

func basicAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		return errors.New("")
	}

	if username != constants.AdminName || password != constants.AdminPass {
		return errors.New("")
	}

	return nil
}
