/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmApplicationInstance struct {
	// instance events
	Events []AdaptorWalmEvent `json:"events"`
	// instance id
	InstanceId string `json:"instance_id"`
	// resource kind
	Kind string `json:"kind"`
	// instance modules
	Modules *AdaptorWalmInstanceResourceSet `json:"modules"`
	// resource name
	Name string `json:"name"`
	// resource namespace
	Namespace string `json:"namespace"`
	// resource state
	State *AdaptorWalmState `json:"state"`
}
