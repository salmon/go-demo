# \SecretApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](SecretApi.md#CreateSecret) | **Post** /api/v1/secret/{namespace} | 创建一个Secret
[**DeleteSecret**](SecretApi.md#DeleteSecret) | **Delete** /api/v1/secret/{namespace}/name/{secretname} | 删除一个Secret
[**GetSecret**](SecretApi.md#GetSecret) | **Get** /api/v1/secret/{namespace}/name/{secretname} | 获取对应Secret的详细信息
[**GetSecrets**](SecretApi.md#GetSecrets) | **Get** /api/v1/secret/{namespace} | 获取Namepace下的所有Secret列表
[**UpdateSecret**](SecretApi.md#UpdateSecret) | **Put** /api/v1/secret/{namespace} | 更新一个Secret


# **CreateSecret**
> CreateSecret(ctx, namespace, body)
创建一个Secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **body** | [**ApiCreateSecretRequestBody**](ApiCreateSecretRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecret**
> DeleteSecret(ctx, namespace, secretname)
删除一个Secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **secretname** | **string**| Secret名字 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecret**
> AdaptorWalmSecret GetSecret(ctx, namespace, secretname)
获取对应Secret的详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **secretname** | **string**| secret名字 | 

### Return type

[**AdaptorWalmSecret**](adaptor.WalmSecret.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecrets**
> AdaptorWalmSecretList GetSecrets(ctx, namespace)
获取Namepace下的所有Secret列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 

### Return type

[**AdaptorWalmSecretList**](adaptor.WalmSecretList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecret**
> UpdateSecret(ctx, namespace, body)
更新一个Secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| 租户名字 | 
  **body** | [**ApiCreateSecretRequestBody**](ApiCreateSecretRequestBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

