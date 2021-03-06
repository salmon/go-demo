/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type ReleaseReleaseRequest struct {
	// chart name
	ChartName string `json:"chart_name"`
	// chart repo
	ChartVersion string `json:"chart_version"`
	// extra values added to the chart
	ConfigValues *interface{} `json:"config_values"`
	// map of dependency chart name and release
	Dependencies map[string]string `json:"dependencies"`
	// name of the release
	Name string `json:"name"`
	// pretty chart params for market
	ReleasePrettyParams *ReleasePrettyChartParams `json:"release_pretty_params"`
	// chart name
	RepoName string `json:"repo_name"`
}
