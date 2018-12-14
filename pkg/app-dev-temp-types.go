package pkg

type AppContainerPortTemplate struct {
	Port int `json:"port" description:"端口"`
	WebIngressEnable bool `json:"web_ingress_enable" description:"是否开启web入口"`
	WebIngressName string `json:"web_ingress_name" description:"Web Ingress名字"`
	WebIngressPath string `json:"web_ingress_path" description:"Web Ingress路径"`
}

type AppContainerStartTypeTemplate struct {
	StartType string `json:"start_type" description:"启动方式. COMMAND/SHELL"`
	Content string `json:"content" description:"启动内容"`
}

type AppContainerStorageTypeTemplate struct {
	Name string `json:"name" description:"磁盘名字"`
	SizeGB int `json:"size_gb" description:"磁盘容量"`
	MountPoint string `json:"mount_point" description:"挂载目录"`
	StorageClass string `json:"storage_class" description:"存储类型"`
}

type AppContainerResourceTypeTemplate struct {
	CpuRequest float64 `json:"cpu_request" description:"cpu request. 最小使用cpu数量"`
	CpuLimit float64 `json:"cpu_limit" description:"cpu limit. 最大使用cpu数量"`
	MemoryRequest int `json:"memory_request" description:"memory request. 最小memory使用. 单位MB"`
	MemoryLimit int `json:"memory_limit" description:"memory request. 最大memory使用 单位MB"`
	Volumes []AppContainerStorageTypeTemplate `json:"volumes" description:"volumes配置"`
}

type AppEnvParamterTemplate struct {
	EnvName string `json:"env_name" description:"环境变量名字"`
	ParamName string `json:"param_name" description:"参数名字"`
	Description string `json:"description" description:"参数描述"`
	DefaultValue string `json:"default_value" description:"默认值"`
}

type AppConfigMapParamterTemplate struct {
	MountPoint string `json:"mount_point" description:"挂载目录"`
	Content string `json:"content" description:"配置文件内容"`
}

type AppDependencyParamterTemplate struct {
	ContainerInnerName string `json:"inner_name" description:"容器内部名字"`
	DependencyType string `json:"dependency_type" description:"依赖类型. 参数/微服务地址/配置文件"`
	DependencyName string `json:"dependency_outer_name" description:"依赖外部对象名字"`
}

type AppParamtersTemplate struct {
	EnvParamterList []AppEnvParamterTemplate `json:"env_paramters" description:"参数设置"`
	ConfigMapParamterList []AppConfigMapParamterTemplate `json:"configmap_paramters" description:"配置文件"`
	DependencyParamterList []AppDependencyParamterTemplate `json:"dependency_paramters" description:"依赖关系"`
}

type AppContainerTemplate struct {
	ContainerName string `json:"container_name" description:"容器名称"`
	ImageName string `json:"image_name" description:"镜像地址"`
	AppPorts []AppContainerPortTemplate `json:"ports" description:"端口列表"`
	StartType AppContainerStartTypeTemplate `json:"start_type" description:"启动方式"`
	ResourceType AppContainerResourceTypeTemplate `json:"resource_type" description:"资源类型"`
}

type AppMicroServiceTemplate struct {
	MicroServiceName string `json:"micro_service_ame" description:"微服务名称"`
	ContainerList []AppContainerTemplate `json:"containers" description:"容器列表"`
	Replica int `json:"replica" description:"微服务副本数量"`
	HealthCheck string `json:"health_check" description:"健康检查"`
}

type AppDevTemplate struct {
	AppName string `json:"app_name" description:"应用名称"`
	AppId string `json:"app_id" description:"应用唯一ID"`
	Name string `json:"name" description:"chart名字.如果不配置同app_id"`
	Version string `json:"version" description:"chart版本"`
	AppArchitectureDescription string `json:"app_architecture_description" description:"应用版本描述/应用架构图"`
	Dependencies map[string]string `json:"dependencies" description:"定义应用间依赖"`
	MicroServiceList []AppMicroServiceTemplate `json:"micro_services" description:"微服务定义"`
	Owner string `json:"owner" description:"创建人"`
	Tenant string `json:"tenant" description:"所属租户"`
	CreateTime string `json:"create_time" description:"创建时间"`
	ModifyTime string `json:"modify_time" description:"修改时间"`
}

type AppAdditionalInfo struct {
	Description string `json:"description" description:"应用简介"`
	AppName string `json:"app_name" description:"应用名称"`
	EncodedIcon string `json:"encoded_icon" description:"应用图标"`
	AppCategory string `json:"app_category" description:"应用分类"`
}
