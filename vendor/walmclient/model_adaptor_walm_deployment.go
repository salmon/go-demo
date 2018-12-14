/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmDeployment struct {
	// available replicas
	AvailableReplicas int32 `json:"available_replicas"`
	// current replicas
	CurrentReplicas int32 `json:"current_replicas"`
	// expected replicas
	ExpectedReplicas int32 `json:"expected_replicas"`
	// resource kind
	Kind string `json:"kind"`
	// resource name
	Name string `json:"name"`
	// resource namespace
	Namespace string `json:"namespace"`
	// deployment pods
	Pods []AdaptorWalmPod `json:"pods"`
	// resource state
	State *AdaptorWalmState `json:"state"`
	// updated replicas
	UpdatedReplicas int32 `json:"updated_replicas"`
}