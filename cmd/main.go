package main

import (
	"course/internal/handlers/user"
	"course/internal/routers/gorilla"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/", user.GetUserListHandler).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", user.GetUserHandler).Methods("GET")
	userRouter.Handle("/", gorilla.AuthMiddleware(http.HandlerFunc(user.CreateUserHandler))).Methods("POST")
	userRouter.Handle("/{id:[0-9]+}", gorilla.AuthMiddleware(http.HandlerFunc(user.DeleteUserHandler))).Methods("DELETE")

	http.ListenAndServe(":9001", r)
}
