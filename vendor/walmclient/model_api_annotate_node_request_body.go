/*
 * Walm
 *
 * Walm Web Server
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package walmclient

type ApiAnnotateNodeRequestBody struct {
	AddAnnotations map[string]string `json:"add_annotations"`
	RemoveAnnotations []string `json:"remove_annotations"`
}
