package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/accuknox/knox-service/db"
	"github.com/accuknox/knox-service/models"
)

//HealthyHome checks Health
func HealthyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Kartheek!")
}

//CreateCluster creates a single cluster
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	var newCluster models.Cluster
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}

	res := json.Unmarshal(reqBody, &newCluster)
	fmt.Println(res, newCluster)
	// Call DB & Insert
	db := db.Conn()
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO clusters(org_id,user_id,cluster_name,node_count,location,policy_id,status) VALUES( ?, ?, ?, ?, ?, ?, ?)")
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
		newCluster.ID = int(lastID)
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		log.Println("INSERT: Id: ", newCluster.ID, insForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCluster)
}

//CreateNode creates a single node
func CreateNode(w http.ResponseWriter, r *http.Request) {
	newNode := models.Node{}
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the node structure")
	}

	res := json.Unmarshal(reqBody, &newNode)
	fmt.Println(res, newNode)
	// Call DB & Insert
	db := db.Conn()
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO nodes(org_id,user_id,node_name,cluster_name,node_count,location,policy_id,status) VALUES( ?, ?, ?, ?, ?, ?, ?, ?)")
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
		newNode.ID = int(lastID)
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		log.Println("INSERT: Id: ", lastID, insForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNode)
}

//GetOneCluster reads a single cluster
func GetOneCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)

	db := db.Conn()
	selDB, err := db.Query("SELECT * FROM clusters WHERE id=?", id)
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

//GetOneNode reads a single node
func GetOneNode(w http.ResponseWriter, r *http.Request) {
	nodeID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(nodeID)

	db := db.Conn()
	selDB, err := db.Query("SELECT * FROM nodes WHERE id=?", id)
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

//GetAllClusters reads all the clusters
func GetAllClusters(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	selDB, err := db.Query("SELECT * FROM clusters")
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

//GetAllNodes reads all the nodes
func GetAllNodes(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	selDB, err := db.Query("SELECT * FROM nodes")
	if err != nil {
		panic(err.Error())
	}
	var node models.Node
	var result models.ResponseNodes
	for selDB.Next() {
		err = selDB.Scan(&node.ID, &node.OrgID, &node.UserID, &node.NodeName, &node.ClusterName, &node.NodeCount, &node.Location, &node.PolicyID, &node.Status)
		if err != nil {
			panic(err.Error())
		}
		result.Nodes = append(result.Nodes, node)
	}
	defer db.Close()
	fmt.Println(result)
	json.NewEncoder(w).Encode(result)
}

//UpdateCluster updates a single cluster
func UpdateCluster(w http.ResponseWriter, r *http.Request) {
	var updatedCluster models.Cluster
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the cluster structure")
	}
	json.Unmarshal(reqBody, &updatedCluster)
	fmt.Println(updatedCluster, updatedCluster.OrgID)
	// Call DB & Update
	db := db.Conn()
	if r.Method == "PUT" {
		updForm, err := db.Prepare("UPDATE clusters SET org_id=?, user_id=?, cluster_name=?, node_count=?, location=?, policy_id=?, status=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		res, err := updForm.Exec(updatedCluster.OrgID, updatedCluster.UserID, updatedCluster.ClusterName, updatedCluster.NodeCount, updatedCluster.Location, updatedCluster.PolicyID, updatedCluster.Status, id)
		fmt.Println(res, err)
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

//UpdateNode updates a single node
func UpdateNode(w http.ResponseWriter, r *http.Request) {
	var updatedNode models.Node
	nodeID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(nodeID)
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly check the node structure")
	}
	json.Unmarshal(reqBody, &updatedNode)
	fmt.Println(updatedNode)
	// Call DB & Update
	db := db.Conn()
	if r.Method == "PUT" {
		updForm, err := db.Prepare("UPDATE nodes SET org_id=?, user_id=?, node_name=?, cluster_name=?, node_count=?, location=?, policy_id=?, status=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		res, err := updForm.Exec(updatedNode.OrgID, updatedNode.UserID, updatedNode.NodeName, updatedNode.ClusterName, updatedNode.NodeCount, updatedNode.Location, updatedNode.PolicyID, updatedNode.Status, id)
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
		log.Println("UPDATE: Id: ", updatedNode.ID, updForm)
	}
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedNode)
}

//DeleteCluster deletes a single cluster
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	clusterID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(clusterID)

	db := db.Conn()
	delForm, err := db.Prepare("DELETE FROM clusters WHERE id=?")
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
	log.Println("DELETE: Id: ", id)
}

//DeleteNode deletes a single node
func DeleteNode(w http.ResponseWriter, r *http.Request) {
	nodeID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(nodeID)

	db := db.Conn()
	delForm, err := db.Prepare("DELETE FROM nodes WHERE id=?")
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
