# ReleaseReleaseInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmAdditionalValues** | [***TranswarpAppHelmValues**](transwarp.AppHelmValues.md) |  | [default to null]
**HelmExtraLabels** | [***ReleaseHelmExtraLabels**](release.HelmExtraLabels.md) |  | [default to null]
**ChartAppVersion** | **string** | jsonnet app version | [default to null]
**ChartName** | **string** | chart name | [default to null]
**ChartVersion** | **string** | chart version | [default to null]
**ConfigValues** | [***interface{}**](interface{}.md) | extra values added to the chart | [default to null]
**Dependencies** | **map[string]string** | map of dependency chart name and release | [default to null]
**Message** | **string** | why release is not ready | [default to null]
**Name** | **string** | name of the release | [default to null]
**Namespace** | **string** | namespace of release | [default to null]
**Ready** | **bool** | whether release is ready | [default to null]
**ReleasePrettyParams** | [***ReleasePrettyChartParams**](release.PrettyChartParams.md) | pretty chart params for market | [default to null]
**ReleaseStatus** | [***AdaptorWalmResourceSet**](adaptor.WalmResourceSet.md) | status of release | [default to null]
**RepoName** | **string** | chart name | [default to null]
**Version** | **int32** | version of the release | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


