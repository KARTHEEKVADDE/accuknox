package models

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
type ResponseClusters struct {
	Clusters []Cluster `json:"clusters"`
}
type Node struct {
	ID          int    `json:"id"`
	OrgID       int    `json:"org_id"`
	UserID      int    `json:"user_id"`
	NodeName    string `json:"node_name"`
	ClusterName string `json:"cluster_name"`
	NodeCount   int    `json:"node_count"`
	Location    string `json:"location"`
	PolicyID    int    `json:"policy_id"`
	Status      string `json:"status"`
}
type ResponseNodes struct {
	Nodes []Node `json:"nodes"`
}
type Pod struct {
	ID          int    `json:"id"`
	OrgID       int    `json:"org_id"`
	UserID      int    `json:"user_id"`
	PodName     string `json:"pod_name"`
	ClusterName string `json:"cluster_name"`
	NodeCount   int    `json:"node_count"`
	Location    string `json:"location"`
	PolicyID    int    `json:"policy_id"`
	Status      string `json:"status"`
}
type ResponsePods struct {
	Pods []Pod `json:"pods"`
}
