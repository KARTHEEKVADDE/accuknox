package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/accuknox/knox-service/handlers"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", handlers.HealthyHome)
	router.HandleFunc("/cluster", handlers.CreateCluster).Methods("POST")
	router.HandleFunc("/node", handlers.CreateNode).Methods("POST")
	router.HandleFunc("/clusters", handlers.GetAllClusters).Methods("GET")
	router.HandleFunc("/nodes", handlers.GetAllNodes).Methods("GET")
	router.HandleFunc("/cluster/{id}", handlers.GetOneCluster).Methods("GET")
	router.HandleFunc("/node/{id}", handlers.GetOneNode).Methods("GET")
	router.HandleFunc("/cluster/{id}", handlers.UpdateCluster).Methods("PUT")
	router.HandleFunc("/node/{id}", handlers.UpdateNode).Methods("PUT")
	router.HandleFunc("/cluster/{id}", handlers.DeleteCluster).Methods("DELETE")
	router.HandleFunc("/node/{id}", handlers.DeleteNode).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
