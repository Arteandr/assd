package main

import (
	"course/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/", user.GetUserListHandler).Methods("GET")
	userRouter.HandleFunc("/", user.CreateUserHandler).Methods("POST")
	userRouter.HandleFunc("/{id:[0-9]+}", user.GetUserHandler).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", user.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":9001", r)
}
