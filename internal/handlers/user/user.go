package user

import (
	"course/internal/entities"
	"course/internal/repositories/user"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u entities.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if err := u.Validate(); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if err := user.Create(u); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	u, err := user.GetUser(id)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	users := user.GetList()

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if err = user.Delete(id); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
