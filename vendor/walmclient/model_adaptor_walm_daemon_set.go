/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmDaemonSet struct {
	// desired number scheduled
	DesiredNumberScheduled int32 `json:"desired_number_scheduled"`
	// resource kind
	Kind string `json:"kind"`
	// resource name
	Name string `json:"name"`
	// resource namespace
	Namespace string `json:"namespace"`
	// number available
	NumberAvailable int32 `json:"number_available"`
	// daemon set pods
	Pods []AdaptorWalmPod `json:"pods"`
	// resource state
	State *AdaptorWalmState `json:"state"`
	// updated number scheduled
	UpdatedNumberScheduled int32 `json:"updated_number_scheduled"`
}
