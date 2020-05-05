package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kartheekvadde/accuknox/db"

	"github.com/gorilla/mux"
)

type Cluster struct {
	ID          int    `json:"id"`
	OrgID       int    `json:"org_id"`
	UserID      int    `json:"user_id"`
	ClusterName string `json:"cluster_name"`
	NodeCount   int    `json:"node_count"`
	Location    string `json:"location"`
	PolicyID    int    `json:"policy_id"`
	Status      string `json:"status"`
}
type ResponseBody struct {
	Clusters []Cluster `json:"clusters"`
}

var clusters = []Cluster{
	{
		ID: 1, OrgID: 121, UserID: 001, ClusterName: "Cluster-001", NodeCount: 001, Location: "Hyderabad", PolicyID: 001, Status: "Active",
	},
}
var result = ResponseBody{clusters}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createCluster(w http.ResponseWriter, r *http.Request) {
	var newCluster Cluster
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}

	res := json.Unmarshal(reqBody, &newCluster)
	fmt.Println(res, newCluster)
	result.Clusters = append(result.Clusters, newCluster)
	// Call DB & Insert
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCluster)
}

func getOneCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)
	for _, singleCluster := range result.Clusters {
		if singleCluster.ID == id {
			json.NewEncoder(w).Encode(singleCluster)
		}
	}
}

func getAllClusters(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM cluster")
	if err != nil {
		panic(err.Error())
	}
	cluster := Cluster{}
	for selDB.Next() {
		var id, org_id, user_id, node_count, policy_id int
		var cluster_name, location, status string
		err = selDB.Scan(&id, &org_id, &user_id, &cluster_name, &node_count, &location, &policy_id, &status)
		if err != nil {
			panic(err.Error())
		}
		cluster.ID = id
		cluster.OrgID = org_id
		cluster.UserID = user_id
		cluster.ClusterName = cluster_name
		cluster.NodeCount = node_count
		cluster.Location = location
		cluster.PolicyID = policy_id
		cluster.Status = status
		result.Clusters = append(result.Clusters, cluster)
	}
	defer db.Close()
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

func updateCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)
	var updatedCluster Cluster

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}
	json.Unmarshal(reqBody, &updatedCluster)
	fmt.Println(updatedCluster)
	for i, singleCluster := range result.Clusters {
		if singleCluster.ID == id {
			singleCluster.OrgID = updatedCluster.OrgID
			singleCluster.UserID = updatedCluster.UserID
			singleCluster.ClusterName = updatedCluster.ClusterName
			singleCluster.NodeCount = updatedCluster.NodeCount
			singleCluster.Location = updatedCluster.Location
			singleCluster.PolicyID = updatedCluster.PolicyID
			singleCluster.Status = updatedCluster.Status
			result.Clusters = append(result.Clusters[:i], singleCluster)
			json.NewEncoder(w).Encode(singleCluster)
		}
	}
}

func deleteCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)

	for i, singleCluster := range result.Clusters {
		if singleCluster.ID == id {
			result.Clusters = append(result.Clusters[:i], result.Clusters[i+1:]...)
			fmt.Fprintf(w, "The cluster with ID %v has been deleted successfully", clusterID)
		}
	}
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/cluster", createCluster).Methods("POST")
	router.HandleFunc("/clusters", getAllClusters).Methods("GET")
	router.HandleFunc("/cluster/{id}", getOneCluster).Methods("GET")
	router.HandleFunc("/clusters/{id}", updateCluster).Methods("PUT")
	router.HandleFunc("/clusters/{id}", deleteCluster).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
