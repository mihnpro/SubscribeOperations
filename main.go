package main

import (
	"log"
	"net/http"
	"test_task/db"
	_ "test_task/docs"
	"test_task/operations"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Subscription Service API
// @version 1.0
// @description API для управления подписками
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	db.Init()
	log.Println("Swagger docs available at: http://localhost:8080/swagger/index.html")

	operations.GetAllSubInfo(mux)
	operations.CreateSubInfo(mux)
	operations.UpdateSubInfo(mux)
	operations.DeleteSubInfo(mux)
	operations.GetSubInfo(mux)
	operations.GetFullSubPriceByPeriodInfo(mux)

	http.ListenAndServe(":8080", mux)
}
