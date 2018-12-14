package api

import (
	"transwarp/tos-app-market/pkg"
	"walmclient"
)

type ErrorMessageResponse struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
}

type ApplicationStatusType int32

const (
	PENDING ApplicationStatusType = 0
	VALIDATING ApplicationStatusType = 1
	VALIDATE_FAILED ApplicationStatusType = 2
	ONLINE_PENDING ApplicationStatusType = 3
	ONLINE ApplicationStatusType = 4
	OFFLINE ApplicationStatusType = 5
)

type ApplicationTypeResponse struct {
	ApplicationType ApplicationStatusType `json:"application_type_value" description:"应用状态"`
	Description string `json:"description" description:"应用状态描述"`
}

type ApplicationCategoriesResponse struct {
	CategoryNameList []string `json:"categories" description:"所有应用分类"`
}

type ImageInfoListResponse struct {
	ImageInfo []string `json:"images" description:"镜像名字信息列表"`
}

type ImportImageInfoRequest struct {
	SrcImageFullName string `json:"src_full_image" description:"原镜像名字信息"`
	UserName string `json:"username" description:"拉取镜像用户名"`
	Password string `json:"password" description:"拉取镜像密码"`
	DestImageTagName string `json:"dest_image_tag" description:"目的镜像Tag名字"`
}

type AppDevListResponse struct {
	AppDevList []AppDevResponse `json:"app_dev_templates" description:"应用模板列表"`
}

type AppDevRequest struct {
	pkg.AppDevTemplate
}

type AppDevResponse struct {
	pkg.AppDevTemplate
	AppCategory string `json:"app_category" description:"应用分类"`
	EncodedIcon string `json:"encoded_icon" description:"应用图标"`
	Description string `json:"description" description:"应用简介"`
	AppStatus ApplicationStatusType `json:"app_status" description:"应用状态信息 已上架/已下架/验证中/验证失败/待上架"`
}

type AppInfoRequest struct {
	AppCategory string `json:"app_category" description:"应用分类"`
	AppName string `json:"app_name" description:"应用名称"`
	EncodedIcon string `json:"encoded_icon" description:"应用图标"`
	Description string `json:"description" description:"应用简介"`
}

type AppDevProcessLogResponse struct {
	AppName string `json:"app_name" description:"应用名称"`
	AppLog string `json:"app_log" description:"应用日志"`
	AppStatus ApplicationStatusType `json:"app_status" description:"应用状态信息 已上架/已下架/验证中/验证失败/待上架"`
}

type AppCategoriesResponse struct {
	Categories[]string `json:"categories" description:"应用类别信息"`
}

type WALMReleaseInfoListResponse struct {
	walmclient.ReleaseReleaseInfoList
}

type WALMReleaseInfoResponse struct {
	walmclient.ReleaseReleaseInfo
}

type WALMReleaseRequest struct {
	walmclient.ReleaseReleaseRequest
}

type WALMChartInfoListResponse struct {
	walmclient.ReleaseChartInfoList
}

type WALMChartInfoResponse struct {
	walmclient.ReleaseChartInfo
}
