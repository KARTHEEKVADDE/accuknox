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
	"github.com/kartheekvadde/accuknox/models"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func createCluster(w http.ResponseWriter, r *http.Request) {
	var newCluster models.Cluster
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}

	res := json.Unmarshal(reqBody, &newCluster)
	fmt.Println(res, newCluster)
	// Call DB & Insert
	db := db.DbConn()
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO cluster(org_id,user_id,cluster_name,node_count,location,policy_id,status) VALUES( ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}

		res, err := insForm.Exec(newCluster.OrgID, newCluster.UserID, newCluster.ClusterName, newCluster.NodeCount, newCluster.Location, newCluster.PolicyID, newCluster.Status)
		if err != nil {
			log.Fatal(err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		log.Println("INSERT: Id: ", newCluster.ID, insForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCluster)
}
func createNode(w http.ResponseWriter, r *http.Request) {
	var newNode models.Node
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the node structure")
	}

	res := json.Unmarshal(reqBody, &newNode)
	fmt.Println(res, newNode)
	// Call DB & Insert
	db := db.DbConn()
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO node(org_id,user_id,node_name,cluster_name,node_count,location,policy_id,status) VALUES( ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}

		res, err := insForm.Exec(newNode.OrgID, newNode.UserID, newNode.NodeName, newNode.ClusterName, newNode.NodeCount, newNode.Location, newNode.PolicyID, newNode.Status)
		if err != nil {
			log.Fatal(err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		log.Println("INSERT: Id: ", newNode.ID, insForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNode)
}
func getOneCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)

	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM cluster WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var cluster models.Cluster
	for selDB.Next() {
		err = selDB.Scan(&cluster.ID, &cluster.OrgID, &cluster.UserID, &cluster.ClusterName, &cluster.NodeCount, &cluster.Location, &cluster.PolicyID, &cluster.Status)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(cluster)
}
func getOneNode(w http.ResponseWriter, r *http.Request) {
	nodeID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(nodeID)

	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM node WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var node models.Node
	for selDB.Next() {
		err = selDB.Scan(&node.ID, &node.OrgID, &node.UserID, &node.NodeName, &node.ClusterName, &node.NodeCount, &node.Location, &node.PolicyID, &node.Status)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(node)
}
func getAllClusters(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM cluster")
	if err != nil {
		panic(err.Error())
	}
	var cluster models.Cluster
	var result models.ResponseClusters
	for selDB.Next() {
		err = selDB.Scan(&cluster.ID, &cluster.OrgID, &cluster.UserID, &cluster.ClusterName, &cluster.NodeCount, &cluster.Location, &cluster.PolicyID, &cluster.Status)
		if err != nil {
			panic(err.Error())
		}
		result.Clusters = append(result.Clusters, cluster)
	}
	defer db.Close()
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}
func updateCluster(w http.ResponseWriter, r *http.Request) {
	var updatedCluster models.Cluster
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}
	json.Unmarshal(reqBody, &updatedCluster)
	fmt.Println(updatedCluster)
	// Call DB & Insert
	db := db.DbConn()
	if r.Method == "POST" {
		updForm, err := db.Prepare("UPDATE cluster SET org_id=?, user_id=?, cluster_name=?, node_count=?, location=?, policy_id=?, status=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		res, err := updForm.Exec(updatedCluster.OrgID, updatedCluster.UserID, updatedCluster.ClusterName, updatedCluster.NodeCount, updatedCluster.Location, updatedCluster.PolicyID, updatedCluster.Status, id)
		if err != nil {
			log.Fatal(err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		log.Println("UPDATE: Id: ", updatedCluster.ID, updForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedCluster)
}
func deleteCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)

	db := db.DbConn()
	delForm, err := db.Prepare("DELETE FROM cluster WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := delForm.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
	log.Println("UPDATE: Id: ", id)
}
func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/cluster", createCluster).Methods("POST")
	router.HandleFunc("/node", createNode).Methods("POST")
	router.HandleFunc("/clusters", getAllClusters).Methods("GET")
	router.HandleFunc("/cluster/{id}", getOneCluster).Methods("GET")
	router.HandleFunc("/node/{id}", getOneNode).Methods("GET")
	router.HandleFunc("/cluster/{id}", updateCluster).Methods("PUT")
	router.HandleFunc("/cluster/{id}", deleteCluster).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
