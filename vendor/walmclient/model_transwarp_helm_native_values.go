/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type TranswarpHelmNativeValues struct {
	AppVersion string `json:"appVersion"`
	ChartName string `json:"chartName"`
	ChartVersion string `json:"chartVersion"`
	ReleaseName string `json:"releaseName"`
	ReleaseNamespace string `json:"releaseNamespace"`
}
