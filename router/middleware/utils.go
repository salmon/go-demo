package middleware

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
	"transwarp/tos-app-market/router/api"
)

const (
	TID_HEADER_NAME = "x-tenant-uid"
	UID_HEADER_NAME = "x-app-uid"
)

func GetUserInfo(request *restful.Request) (string, string, error) {
	tenantId := request.HeaderParameter(TID_HEADER_NAME)
	userId := request.HeaderParameter(UID_HEADER_NAME)

	if false && (len(tenantId) == 0 || len(userId) == 0) {
		return tenantId, userId, fmt.Errorf("can not get valid %s %s from header", TID_HEADER_NAME, UID_HEADER_NAME)
	}

	return tenantId, userId, nil
}

func WriteErrorResponse(response *restful.Response, code int, errMsg string) error {
	return response.WriteHeaderAndEntity(http.StatusInternalServerError, api.ErrorMessageResponse{code, errMsg})
}

func WriteNotFoundResponse(response *restful.Response, code int, errMsg string) error {
	return response.WriteHeaderAndEntity(http.StatusNotFound, api.ErrorMessageResponse{code, errMsg})
}
