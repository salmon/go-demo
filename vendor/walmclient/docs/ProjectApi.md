# \ProjectApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInstanceInProject**](ProjectApi.md#DeleteInstanceInProject) | **Delete** /api/v1/project/{namespace}/name/{project}/instance/{release} | 删除一个Project组件
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /api/v1/project/{namespace}/name/{project} | 删除一个Project
[**DeployInstanceInProject**](ProjectApi.md#DeployInstanceInProject) | **Post** /api/v1/project/{namespace}/name/{project}/instance | 添加一个Project组件
[**DeployProject**](ProjectApi.md#DeployProject) | **Post** /api/v1/project/{namespace}/name/{project} | 创建一个Project
[**DeployProjectInProject**](ProjectApi.md#DeployProjectInProject) | **Post** /api/v1/project/{namespace}/name/{project}/project | 添加多个Project组件
[**GetProjectInfo**](ProjectApi.md#GetProjectInfo) | **Get** /api/v1/project/{namespace}/name/{project} | 获取Project的详细信息
[**ListProjectAllNamespaces**](ProjectApi.md#ListProjectAllNamespaces) | **Get** /api/v1/project | 获取所有Project列表
[**ListProjectByNamespace**](ProjectApi.md#ListProjectByNamespace) | **Get** /api/v1/project/{namespace} | 获取对应租户的Project列表


# **DeleteInstanceInProject**
> DeleteInstanceInProject(ctx, namespace, project, release, optional)
删除一个Project组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 
  **release** | **string**| Release名字 | 
 **optional** | ***DeleteInstanceInProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteInstanceInProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **optional.Bool**| 异步与否 | 
 **timeoutSec** | **optional.Int32**| 超时时间 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, namespace, project, optional)
删除一个Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 
 **optional** | ***DeleteProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **optional.Bool**| 异步与否 | 
 **timeoutSec** | **optional.Int32**| 超时时间 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployInstanceInProject**
> DeployInstanceInProject(ctx, namespace, project, body, optional)
添加一个Project组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 
  **body** | [**ReleaseReleaseRequest**](ReleaseReleaseRequest.md)|  | 
 **optional** | ***DeployInstanceInProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployInstanceInProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **optional.Bool**| 异步与否 | 
 **timeoutSec** | **optional.Int32**| 超时时间 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployProject**
> DeployProject(ctx, namespace, project, body, optional)
创建一个Project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 
  **body** | [**ReleaseProjectParams**](ReleaseProjectParams.md)|  | 
 **optional** | ***DeployProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **optional.Bool**| 异步与否 | 
 **timeoutSec** | **optional.Int32**| 超时时间 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployProjectInProject**
> DeployProjectInProject(ctx, namespace, project, body, optional)
添加多个Project组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 
  **body** | [**ReleaseProjectParams**](ReleaseProjectParams.md)|  | 
 **optional** | ***DeployProjectInProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeployProjectInProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **optional.Bool**| 异步与否 | 
 **timeoutSec** | **optional.Int32**| 超时时间 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectInfo**
> ReleaseProjectInfo GetProjectInfo(ctx, namespace, project)
获取Project的详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **project** | **string**| Project名字 | 

### Return type

[**ReleaseProjectInfo**](release.ProjectInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectAllNamespaces**
> ReleaseProjectInfoList ListProjectAllNamespaces(ctx, )
获取所有Project列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReleaseProjectInfoList**](release.ProjectInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectByNamespace**
> ReleaseProjectInfoList ListProjectByNamespace(ctx, namespace)
获取对应租户的Project列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 

### Return type

[**ReleaseProjectInfoList**](release.ProjectInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

