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
type ResponseBody struct {
	Clusters []Cluster `json:"clusters"`
}
