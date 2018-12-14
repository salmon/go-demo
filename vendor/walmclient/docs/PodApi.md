# \PodApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPodEvents**](PodApi.md#GetPodEvents) | **Get** /api/v1/pod/{namespace}/name/{pod}/events | 获取Pod对应的事件
[**GetPodLogs**](PodApi.md#GetPodLogs) | **Get** /api/v1/pod/{namespace}/name/{pod}/logs | 获取Pod对应的事件
[**RestartPod**](PodApi.md#RestartPod) | **Post** /api/v1/pod/{namespace}/name/{pod}/restart | 重启Pod


# **GetPodEvents**
> AdaptorWalmEventList GetPodEvents(ctx, namespace, pod)
获取Pod对应的事件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **pod** | **string**| pod名字 | 

### Return type

[**AdaptorWalmEventList**](adaptor.WalmEventList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPodLogs**
> string GetPodLogs(ctx, namespace, pod, optional)
获取Pod对应的事件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **pod** | **string**| pod名字 | 
 **optional** | ***GetPodLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPodLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **container** | **optional.String**| container名字 | 
 **tail** | **optional.Int32**| 最后几行 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartPod**
> RestartPod(ctx, namespace, pod)
重启Pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **pod** | **string**| pod名字 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

