package v1

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"time"
	"transwarp/tos-app-market/pkg"
	"transwarp/tos-app-market/pkg/db/memory"
	"transwarp/tos-app-market/router/middleware"
	apitypes "transwarp/tos-app-market/router/api"
)

func GetAllApplicationTypes(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Errorf("GetApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}
	applicationTypes := make([]apitypes.ApplicationTypeResponse, 0)

	PendingType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.PENDING,
		Description: "待验证",
	}
	ValidatingType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.VALIDATING,
		Description: "验证中",
	}
	ValidateFailedType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.VALIDATE_FAILED,
		Description: "验证失败",
	}
	OnlinePendingType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.ONLINE_PENDING,
		Description: "待验证",
	}
	OnlineType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.ONLINE,
		Description: "已上架",
	}
	OfflineType := apitypes.ApplicationTypeResponse{
		ApplicationType: apitypes.OFFLINE,
		Description: "已下架",
	}

	applicationTypes = append(applicationTypes, PendingType, ValidatingType, ValidateFailedType, OnlinePendingType, OnlineType, OfflineType)
	response.WriteEntity(applicationTypes)
}

func NewCategory(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Errorf("NewCategory %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	var appCategory string
	request.ReadEntity(&appCategory)
	if err != nil {
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("failed to read request body : %s", err.Error()))
		return
	}

	memory.TOSCategory = append(memory.TOSCategory, appCategory)
}

func GetAllCategories(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Errorf("GetAllCategories %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appCategoriesResponse := apitypes.AppCategoriesResponse{
		Categories: memory.TOSCategory,
	}

	response.WriteEntity(appCategoriesResponse)
}

func NewApplicationTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Errorf("GetApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appInfoRequest := new(apitypes.AppInfoRequest)
	err = request.ReadEntity(appInfoRequest)
	if err != nil {
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("failed to read request body : %s", err.Error()))
		return
	}

	t := time.Now()
	newAppAdditionalInfo := pkg.AppAdditionalInfo{
		AppName: appInfoRequest.AppName,
		Description: appInfoRequest.Description,
		EncodedIcon: appInfoRequest.EncodedIcon,
		AppCategory: appInfoRequest.AppCategory,
	}
	newAppDevTemplate := pkg.AppDevTemplate {
		AppName: appInfoRequest.AppName,
		AppId: appInfoRequest.AppName,
		Name: appInfoRequest.AppName,
		Version: "1.0",
		Dependencies: make(map[string]string, 0),
		MicroServiceList: make([]pkg.AppMicroServiceTemplate, 0),
		Owner: userId,
		Tenant: tenantId,
		CreateTime: t.Format("2006-01-02 15:04:05"),
		ModifyTime: t.Format("2006-01-02 15:04:05"),
	}

	memory.TOSApplications = append(memory.TOSApplications, newAppAdditionalInfo)
	memory.TOSAppDevTemplates = append(memory.TOSAppDevTemplates, newAppDevTemplate)
}

func GetAllApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("GetAllApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appDevList := apitypes.AppDevListResponse{}
	appDevList.AppDevList = make([]apitypes.AppDevResponse, 0)

	for _, appDevTemplate := range memory.TOSAppDevTemplates {
		newAppDevTemplate := apitypes.AppDevResponse{}
		newAppDevTemplate.AppDevTemplate = appDevTemplate

		for _, appInfo := range memory.TOSApplications {
			if appInfo.AppName == appDevTemplate.AppName {
				newAppDevTemplate.AppCategory = appInfo.AppCategory
				newAppDevTemplate.Description = appInfo.Description
				newAppDevTemplate.EncodedIcon = appInfo.EncodedIcon
				// ToDo
				newAppDevTemplate.AppStatus = apitypes.PENDING
				break
			}
		}

		appDevList.AppDevList = append(appDevList.AppDevList, newAppDevTemplate)
	}

	response.WriteEntity(appDevList)
}

func GetApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("GetApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appName := request.PathParameter("app-template")
	appVersion := request.PathParameter("version")
	newAppDevTemplate := apitypes.AppDevResponse{}
	for _, appDevTemplate := range memory.TOSAppDevTemplates {
		if appDevTemplate.AppName == appName && appDevTemplate.Version == appVersion {
			newAppDevTemplate.AppDevTemplate = appDevTemplate

			for _, appInfo := range memory.TOSApplications {
				if appInfo.AppName == appDevTemplate.AppName {
					newAppDevTemplate.AppCategory = appInfo.AppCategory
					newAppDevTemplate.Description = appInfo.Description
					newAppDevTemplate.EncodedIcon = appInfo.EncodedIcon
					// ToDo
					newAppDevTemplate.AppStatus = apitypes.PENDING
					break
				}
			}

			response.WriteEntity(newAppDevTemplate)
			return
		}
	}

	middleware.WriteNotFoundResponse(response, -1, fmt.Sprintf("can not found template %s %s", appName, appVersion))
}

func CreateApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("CreateApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appName := request.PathParameter("app-template")
	appVersion := request.PathParameter("version")
	logrus.Infof("CreateApplicationDevTemplate %s %s", appName, appVersion)
	appDevRequest := new(apitypes.AppDevRequest)
	err = request.ReadEntity(appDevRequest)
	if err != nil {
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("failed to read request body : %s", err.Error()))
		return
	}

	memory.TOSAppDevTemplates = append(memory.TOSAppDevTemplates, appDevRequest.AppDevTemplate)
}

func PatchApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("PatchApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}

	appName := request.PathParameter("app-template")
	appVersion := request.PathParameter("version")

	middleware.WriteNotFoundResponse(response, -1, fmt.Sprintf("can not found template %s %s", appName, appVersion))
}

func DeleteApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("DeleteApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}
}

func DeleteApplicationDevTemplateVersion(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("DeleteApplicationDevTemplateVersion %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}
}

func ValidateApplicationDevTemplate(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("ValidateApplicationDevTemplate %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}
}

func GetApplicationDevTemplateLog(request *restful.Request, response *restful.Response) {
	tenantId, userId, err := middleware.GetUserInfo(request)
	if err != nil {
		logrus.Infof("GetApplicationDevTemplateLog %s %s %+v", tenantId, userId, err)
		middleware.WriteErrorResponse(response, -1, fmt.Sprintf("GetUserInfo %s %s %+v", tenantId, userId, err))
		return
	}
}
