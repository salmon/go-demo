package router

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	apitypes "transwarp/tos-app-market/router/api"
	"transwarp/tos-app-market/router/api/v1"
	"transwarp/tos-app-market/router/middleware"
)

var APIPATH = "/market/api/v1"

func InitRootRouter() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"root"}

	ws.Route(ws.GET("/readiness").To(readinessProbe).
		Doc("服务Ready状态检查").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}))

	ws.Route(ws.GET("/liveniess").To(livenessProbe).
		Doc("服务Live状态检查").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}))

	ws.Route(ws.GET("/stats").To(middleware.ServerStatsData).
		Doc("获取服务Stats").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}))

	return ws
}

func DockerImageRouter() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path(APIPATH + "/image").
		Doc("镜像相关操作").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"image"}

	ws.Route(ws.GET("/").
		To(v1.GetAllImages).
		Doc("获取所有镜像信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("name", "名字过滤").DataType("string").Required(false)).
		Writes(apitypes.ImageInfoListResponse{}).
		Returns(200, "OK", apitypes.ImageInfoListResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{image}/import").
		To(v1.ImportImage).
		Doc("导入镜像信息").
		Param(ws.PathParameter("image", "导入镜像名字").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(apitypes.ImportImageInfoRequest{}).
		Writes(nil).
		Returns(200, "OK", nil).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	return ws
}


func WapperWalmChartRouter() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path(APIPATH + "/walm-service").
		Doc("部署应用相关操作").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"walm-service"}

	ws.Route(ws.GET("/name/{walm-name}/release").
		To(v1.WalmListReleaseByNamespace).
		Doc("获取Namepaces下的所有Release列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Param(ws.PathParameter("namespace", "租户名字").DataType("string")).
		Writes(apitypes.WALMReleaseInfoResponse{}).
		Returns(200, "OK", apitypes.WALMReleaseInfoResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{walm-name}").
		To(v1.WalmInstallRelease).
		Doc("安装一个Release").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Reads(apitypes.WALMReleaseRequest{}).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.PUT("/name/{walm-name}/name/{release}").
		To(v1.WalmUpgradeRelease).
		Doc("升级&修改Release").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Param(ws.PathParameter("release", "Release名字").DataType("string")).
		Reads(apitypes.WALMReleaseRequest{}).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.DELETE("/name/{walm-name}/name/{release}").To(v1.WalmDeleteRelease).
		Doc("删除一个Release").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("namespace", "租户名字").DataType("string")).
		Param(ws.PathParameter("release", "Release名字").DataType("string")).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{walm-name}/name/{release}/restart").
		To(v1.WalmRestartRelease).
		Doc("Restart　Release关联的所有pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Param(ws.PathParameter("release", "Release名字").DataType("string")).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/name/{walm-name}/list").
		To(v1.WalmGetChartList).
		Doc("获取chart列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Writes(apitypes.WALMChartInfoListResponse{}).
		Returns(200, "OK", apitypes.WALMChartInfoListResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/name/{walm-name}/chart/{chart-name}").
		To(v1.WalmGetChartInfo).
		Doc("获取chart详细信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("walm-name", "WALM实例名字").DataType("string")).
		Param(ws.PathParameter("chart-name", "Chart名字").DataType("string")).
		Param(ws.QueryParameter("chart-version", "chart版本").DataType("string").DefaultValue("")).
		Writes(apitypes.WALMChartInfoResponse{}).
		Returns(200, "OK", apitypes.WALMChartInfoResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	return ws
}


func PackageChartRouter() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path(APIPATH + "/package-chart").
		Doc("打包应用相关操作").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"package-chart"}

	ws.Route(ws.POST("/name/{app-template}/version/{version}/online").
		To(v1.AddPackageChartToHarbor).
		Doc("上架应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Writes(nil).
		Returns(200, "OK", nil).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{app-template}/version/{version}/offline").
		To(v1.DeletePackageChartFromHarbor).
		Doc("下架应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Writes(nil).
		Returns(200, "OK", nil).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	return ws
}

func ApplicationDevTemplateRouter() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path(APIPATH + "/application-dev-template").
		Doc("应用上架相关操作").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"application-dev-template"}

	ws.Route(ws.GET("/").
		To(v1.GetAllApplicationDevTemplate).
		Doc("获取所有应用上架模板信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(apitypes.AppDevListResponse{}).
		Returns(200, "OK", apitypes.AppDevListResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/allapptypes").
		To(v1.GetAllApplicationTypes).
		Doc("获取所有应用的状态与描述").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]apitypes.ApplicationTypeResponse{}).
		Returns(200, "OK", []apitypes.ApplicationTypeResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/newcategory").
		To(v1.NewCategory).
		Doc("创建应用类别").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads("").
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/allcategories").
		To(v1.GetAllCategories).
		Doc("获取所有应用类别").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]apitypes.ApplicationTypeResponse{}).
		Returns(200, "OK", []apitypes.ApplicationTypeResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/name/{app-template}/version/{version}").
		To(v1.GetApplicationDevTemplate).
		Doc("获取保存的应用模板信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Writes(apitypes.AppDevResponse{}).
		Returns(200, "OK", apitypes.AppDevResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/new").
		To(v1.NewApplicationTemplate).
		Doc("创建应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(apitypes.AppInfoRequest{}).
		Writes(nil).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{app-template}/version/{version}").
		To(v1.CreateApplicationDevTemplate).
		Doc("保存应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Reads(apitypes.AppDevResponse{}).
		Writes(apitypes.AppDevRequest{}).
		Returns(200, "OK", apitypes.AppDevResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.PUT("/name/{app-template}/version/{version}").
		To(v1.PatchApplicationDevTemplate).
		Doc("修改保存的应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Writes(apitypes.AppDevResponse{}).
		Returns(200, "OK", apitypes.AppDevResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.DELETE("/name/{app-template}/version/{version}").
		To(v1.DeleteApplicationDevTemplateVersion).
		Doc("删除保存的应用模板版本").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Writes("").
		Returns(200, "OK", apitypes.AppDevResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.DELETE("/name/{app-template}").
		To(v1.DeleteApplicationDevTemplate).
		Doc("删除保存的应用模板").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Writes(nil).
		Returns(200, "OK", nil).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.POST("/name/{app-template}/version/{version}/validate").
		To(v1.ValidateApplicationDevTemplate).
		Doc("验证保存上架的应用").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Reads(apitypes.AppDevResponse{}).
		Writes(apitypes.AppDevRequest{}).
		Returns(200, "OK", apitypes.AppDevResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	ws.Route(ws.GET("/name/{app-template}/version/{version}/log").
		To(v1.GetApplicationDevTemplateLog).
		Doc("获取待上架应用的日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("app-template", "模板名字").DataType("string")).
		Param(ws.PathParameter("version", "模板版本").DataType("string")).
		Reads(apitypes.AppDevProcessLogResponse{}).
		Writes(apitypes.AppDevProcessLogResponse{}).
		Returns(200, "OK", apitypes.AppDevProcessLogResponse{}).
		Returns(404, "Not Found", apitypes.ErrorMessageResponse{}).
		Returns(500, "Internal Error", apitypes.ErrorMessageResponse{}),
	)

	return ws
}

func readinessProbe(request *restful.Request, response *restful.Response) {
	response.WriteEntity("OK")
}

func livenessProbe(request *restful.Request, response *restful.Response) {
	response.WriteEntity("OK")
}
