/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmServicePort struct {
	// service port name
	Name string `json:"name"`
	// node port
	NodePort int32 `json:"node_port"`
	// service port
	Port int32 `json:"port"`
	// service port protocol
	Protocol string `json:"protocol"`
	// backend pod port
	TargetPort string `json:"target_port"`
}
