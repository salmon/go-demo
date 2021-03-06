/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmConfigMap struct {
	// config map data
	Data map[string]string `json:"data"`
	// resource kind
	Kind string `json:"kind"`
	// resource name
	Name string `json:"name"`
	// resource namespace
	Namespace string `json:"namespace"`
	// resource state
	State *AdaptorWalmState `json:"state"`
}
