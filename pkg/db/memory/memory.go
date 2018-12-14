package memory

import "transwarp/tos-app-market/pkg"

var TOSApplications []pkg.AppAdditionalInfo
var TOSAppDevTemplates []pkg.AppDevTemplate
var TOSCategory []string

func init() {
	TOSApplications = make([]pkg.AppAdditionalInfo, 0)
	TOSAppDevTemplates = make([]pkg.AppDevTemplate, 0)

	appInfo := pkg.AppAdditionalInfo {
		AppName: "WORDPRESS",
		Description: "WordPress是一个以PHP和MySQL为平台的自由开源的博客软件和内容管理系统[5]。WordPress具有插件架构和模板系统。截至2018年4月，排名前1000万的网站中有超过30.6%的网站使用WordPress[6]。WordPress是最受欢迎的网站内容管理系统[7]。WordPress是目前因特网上最流行的博客系统。[8][9]WordPress在最着名的网络发布阶段中脱颖而出。如今，它控制着超过7000万个站点。",
		EncodedIcon: "",
		AppCategory: "网站建设",
	}

	appDevTemplate := pkg.AppDevTemplate {
		AppName: "WORDPRESS",
		AppId: "wordpress",
		Name: "wordpress",
		Version: "1.0",
		AppArchitectureDescription: "",
		Dependencies: make(map[string]string, 0),
		MicroServiceList: make([]pkg.AppMicroServiceTemplate, 0),
		Owner: "admin",
		Tenant: "admin",
		CreateTime: "2018-12-05 11:55",
		ModifyTime: "2018-12-05 11:55",
	}
	wordpressMicroServer := pkg.AppMicroServiceTemplate {
		MicroServiceName: "wordpress",
		Replica: 1,
		ContainerList: make([]pkg.AppContainerTemplate, 0),
	}
	container := pkg.AppContainerTemplate{
		ContainerName: "wordpress",
		ImageName: "172.16.1.41:5000/gold/wordpress:latest",
		AppPorts: make([]pkg.AppContainerPortTemplate, 0),
	}
	container.AppPorts = append(container.AppPorts, pkg.AppContainerPortTemplate{
		Port: 8080,
		WebIngressEnable: true,
		WebIngressName: "web",
		WebIngressPath: "/",
	})
	container.StartType = pkg.AppContainerStartTypeTemplate{
		StartType: "SHELL",
		Content: "tomcat start",
	}
	container.ResourceType = pkg.AppContainerResourceTypeTemplate{
		CpuLimit: 1,
		CpuRequest: 1,
		MemoryLimit: 2,
		MemoryRequest: 2,
	}
	wordpressMicroServer.ContainerList = append(wordpressMicroServer.ContainerList, container)
	appDevTemplate.MicroServiceList = append(appDevTemplate.MicroServiceList, wordpressMicroServer)

	TOSApplications = append(TOSApplications, appInfo)
	TOSAppDevTemplates = append(TOSAppDevTemplates, appDevTemplate)
	TOSCategory = make([]string, 0)
	TOSCategory = append(TOSCategory, "网站建设", "大数据", "人工智能")
}
