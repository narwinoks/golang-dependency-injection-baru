package main

import (
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {
	serve := InitializeServer()
	err := serve.ListenAndServe()
	helper.PanicIfError(err)
}
