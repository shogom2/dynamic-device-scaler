package types

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ResourceClaimInfo struct {
	Name              string                `json:"name"`
	NodeName          string                `json:"node_name"`
	CreationTimestamp v1.Time               `json:"creation_timestamp"`
	Namespace         string                `json:"namespace"`
	Devices           []ResourceClaimDevice `json:"devices"`
}

type ResourceClaimDevice struct {
	Name      string `json:"name"`
	Model     string `json:"model"`
	State     string `json:"state"`
	UsedByPod bool   `json:"used_by_pod"`
}
