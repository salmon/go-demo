/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type AdaptorWalmResourceSet struct {
	// release configmaps
	Configmaps []AdaptorWalmConfigMap `json:"configmaps"`
	// release daemonsets
	Daemonsets []AdaptorWalmDaemonSet `json:"daemonsets"`
	// release deployments
	Deployments []AdaptorWalmDeployment `json:"deployments"`
	// release ingresses
	Ingresses []AdaptorWalmIngress `json:"ingresses"`
	// release instances
	Instances []AdaptorWalmApplicationInstance `json:"instances"`
	// release jobs
	Jobs []AdaptorWalmJob `json:"jobs"`
	// release secrets
	Secrets []AdaptorWalmSecret `json:"secrets"`
	// release services
	Services []AdaptorWalmService `json:"services"`
	// release statefulsets
	Statefulsets []AdaptorWalmStatefulSet `json:"statefulsets"`
}