# \TenantApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantApi.md#CreateTenant) | **Post** /api/v1/tenant/{tenantName} | 创建租户
[**DeleteTenant**](TenantApi.md#DeleteTenant) | **Delete** /api/v1/tenant/{tenantName} | 删除租户
[**GetTenant**](TenantApi.md#GetTenant) | **Get** /api/v1/tenant/{tenantName} | 获取租户状态
[**ListTenants**](TenantApi.md#ListTenants) | **Get** /api/v1/tenant | 获取租户列表
[**UpdateTenant**](TenantApi.md#UpdateTenant) | **Put** /api/v1/tenant/{tenantName} | 更新租户信息


# **CreateTenant**
> CreateTenant(ctx, tenantName, body)
创建租户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantName** | **string**| 租户名字 | 
  **body** | [**TenantTenantParams**](TenantTenantParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenant**
> DeleteTenant(ctx, tenantName)
删除租户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantName** | **string**| 租户名字 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenant**
> TenantTenantInfo GetTenant(ctx, tenantName)
获取租户状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantName** | **string**| 租户名字 | 

### Return type

[**TenantTenantInfo**](tenant.TenantInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTenants**
> TenantTenantInfoList ListTenants(ctx, )
获取租户列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TenantTenantInfoList**](tenant.TenantInfoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTenant**
> UpdateTenant(ctx, tenantName, body)
更新租户信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantName** | **string**| 租户名字 | 
  **body** | [**TenantTenantParams**](TenantTenantParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

