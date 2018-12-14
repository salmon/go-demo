/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmNode struct {
	// resource allocatable
	Allocatable map[string]string `json:"allocatable"`
	// node annotations
	Annotations map[string]string `json:"annotations"`
	// resource capacity
	Capacity map[string]string `json:"capacity"`
	// resource kind
	Kind string `json:"kind"`
	// node labels
	Labels map[string]string `json:"labels"`
	// limits resource allocated
	LimitsAllocated map[string]string `json:"limits_allocated"`
	// resource name
	Name string `json:"name"`
	// resource namespace
	Namespace string `json:"namespace"`
	// ip of node
	NodeIp string `json:"node_ip"`
	// requests resource allocated
	RequestsAllocated map[string]string `json:"requests_allocated"`
	// resource state
	State *AdaptorWalmState `json:"state"`
}
