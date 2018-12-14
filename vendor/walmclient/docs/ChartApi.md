# \ChartApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChartInfo**](ChartApi.md#GetChartInfo) | **Get** /api/v1/chart/{repo-name}/chart/{chart-name} | 获取chart详细信息
[**GetChartList**](ChartApi.md#GetChartList) | **Get** /api/v1/chart/{repo-name}/list | 获取chart列表
[**GetRepoList**](ChartApi.md#GetRepoList) | **Get** /api/v1/chart/repolist | 获取chart-repo列表


# **GetChartInfo**
> ReleaseChartInfo GetChartInfo(ctx, repoName, chartName, optional)
获取chart详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repo名字 | 
  **chartName** | **string**| Chart名字 | 
 **optional** | ***GetChartInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChartInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chartVersion** | **optional.String**| chart版本 | 

### Return type

[**ReleaseChartInfo**](release.ChartInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChartList**
> ReleaseChartInfoList GetChartList(ctx, repoName)
获取chart列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repo名字 | 

### Return type

[**ReleaseChartInfoList**](release.ChartInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoList**
> ReleaseRepoInfoList GetRepoList(ctx, )
获取chart-repo列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReleaseRepoInfoList**](release.RepoInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

