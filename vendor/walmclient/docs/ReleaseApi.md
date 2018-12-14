# \ReleaseApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRelease**](ReleaseApi.md#DeleteRelease) | **Delete** /api/v1/release/{namespace}/name/{release} | 删除一个Release
[**GetRelease**](ReleaseApi.md#GetRelease) | **Get** /api/v1/release/{namespace}/name/{release} | 获取对应Release的详细信息
[**InstallRelease**](ReleaseApi.md#InstallRelease) | **Post** /api/v1/release/{namespace} | 安装一个Release
[**ListRelease**](ReleaseApi.md#ListRelease) | **Get** /api/v1/release | 获取所有Release列表
[**ListReleaseByNamespace**](ReleaseApi.md#ListReleaseByNamespace) | **Get** /api/v1/release/{namespace} | 获取Namepaces下的所有Release列表
[**RestartRelease**](ReleaseApi.md#RestartRelease) | **Post** /api/v1/release/{namespace}/name/{release}/restart | Restart　Release关联的所有pod
[**RollBackRelease**](ReleaseApi.md#RollBackRelease) | **Post** /api/v1/release/{namespace}/name/{release}/version/{version}/rollback | RollBack　Release版本
[**UpgradeRelease**](ReleaseApi.md#UpgradeRelease) | **Put** /api/v1/release/{namespace} | 升级一个Release


# **DeleteRelease**
> DeleteRelease(ctx, namespace, release)
删除一个Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **release** | **string**| Release名字 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelease**
> ReleaseReleaseInfo GetRelease(ctx, namespace, release)
获取对应Release的详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **release** | **string**| Release名字 | 

### Return type

[**ReleaseReleaseInfo**](release.ReleaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallRelease**
> InstallRelease(ctx, namespace, body)
安装一个Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **body** | [**ReleaseReleaseRequest**](ReleaseReleaseRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRelease**
> ReleaseReleaseInfoList ListRelease(ctx, )
获取所有Release列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReleaseReleaseInfoList**](release.ReleaseInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReleaseByNamespace**
> ReleaseReleaseInfoList ListReleaseByNamespace(ctx, namespace)
获取Namepaces下的所有Release列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 

### Return type

[**ReleaseReleaseInfoList**](release.ReleaseInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartRelease**
> RestartRelease(ctx, namespace, release)
Restart　Release关联的所有pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **release** | **string**| Release名字 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollBackRelease**
> RollBackRelease(ctx, namespace, release, version)
RollBack　Release版本

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **release** | **string**| Release名字 | 
  **version** | **string**| 版本号 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeRelease**
> UpgradeRelease(ctx, namespace, body)
升级一个Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **body** | [**ReleaseReleaseRequest**](ReleaseReleaseRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

