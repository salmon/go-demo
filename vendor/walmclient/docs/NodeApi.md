# \NodeApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnotateNode**](NodeApi.md#AnnotateNode) | **Post** /api/v1/node/{nodename}/annotations | 修改节点Annotations
[**GetNode**](NodeApi.md#GetNode) | **Get** /api/v1/node/{nodename} | 获取节点详细信息
[**GetNodes**](NodeApi.md#GetNodes) | **Get** /api/v1/node | 获取节点列表
[**LabelNode**](NodeApi.md#LabelNode) | **Post** /api/v1/node/{nodename}/labels | 修改节点Labels


# **AnnotateNode**
> AnnotateNode(ctx, nodename, body)
修改节点Annotations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodename** | **string**| 节点名字 | 
  **body** | [**ApiAnnotateNodeRequestBody**](ApiAnnotateNodeRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNode**
> AdaptorWalmNode GetNode(ctx, nodename)
获取节点详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodename** | **string**| 节点名字 | 

### Return type

[**AdaptorWalmNode**](adaptor.WalmNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodes**
> AdaptorWalmNodeList GetNodes(ctx, optional)
获取节点列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetNodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelselector** | **optional.String**| 节点标签过滤 | 

### Return type

[**AdaptorWalmNodeList**](adaptor.WalmNodeList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelNode**
> LabelNode(ctx, nodename, body)
修改节点Labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodename** | **string**| 节点名字 | 
  **body** | [**ApiLabelNodeRequestBody**](ApiLabelNodeRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

