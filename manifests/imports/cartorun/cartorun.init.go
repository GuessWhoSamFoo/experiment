package cartorun

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cartorun.ClusterConfigTemplate",
		reflect.TypeOf((*ClusterConfigTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterConfigTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterConfigTemplateProps",
		reflect.TypeOf((*ClusterConfigTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterConfigTemplateSpec",
		reflect.TypeOf((*ClusterConfigTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterConfigTemplateSpecParams",
		reflect.TypeOf((*ClusterConfigTemplateSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterDelivery",
		reflect.TypeOf((*ClusterDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterDelivery{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliveryProps",
		reflect.TypeOf((*ClusterDeliveryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpec",
		reflect.TypeOf((*ClusterDeliverySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecParams",
		reflect.TypeOf((*ClusterDeliverySpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResources",
		reflect.TypeOf((*ClusterDeliverySpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResourcesConfigs",
		reflect.TypeOf((*ClusterDeliverySpecResourcesConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResourcesDeployment",
		reflect.TypeOf((*ClusterDeliverySpecResourcesDeployment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResourcesParams",
		reflect.TypeOf((*ClusterDeliverySpecResourcesParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResourcesSources",
		reflect.TypeOf((*ClusterDeliverySpecResourcesSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecResourcesTemplateRef",
		reflect.TypeOf((*ClusterDeliverySpecResourcesTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cartorun.ClusterDeliverySpecResourcesTemplateRefKind",
		reflect.TypeOf((*ClusterDeliverySpecResourcesTemplateRefKind)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_SOURCE_TEMPLATE": ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_SOURCE_TEMPLATE,
			"CLUSTER_DEPLOYMENT_TEMPLATE": ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_DEPLOYMENT_TEMPLATE,
			"CLUSTER_TEMPLATE": ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_TEMPLATE,
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeliverySpecServiceAccountRef",
		reflect.TypeOf((*ClusterDeliverySpecServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterDeploymentTemplate",
		reflect.TypeOf((*ClusterDeploymentTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterDeploymentTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateProps",
		reflect.TypeOf((*ClusterDeploymentTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpec",
		reflect.TypeOf((*ClusterDeploymentTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpecObservedCompletion",
		reflect.TypeOf((*ClusterDeploymentTemplateSpecObservedCompletion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpecObservedCompletionFailed",
		reflect.TypeOf((*ClusterDeploymentTemplateSpecObservedCompletionFailed)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpecObservedCompletionSucceeded",
		reflect.TypeOf((*ClusterDeploymentTemplateSpecObservedCompletionSucceeded)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpecObservedMatches",
		reflect.TypeOf((*ClusterDeploymentTemplateSpecObservedMatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterDeploymentTemplateSpecParams",
		reflect.TypeOf((*ClusterDeploymentTemplateSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterImageTemplate",
		reflect.TypeOf((*ClusterImageTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterImageTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterImageTemplateProps",
		reflect.TypeOf((*ClusterImageTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterImageTemplateSpec",
		reflect.TypeOf((*ClusterImageTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterImageTemplateSpecParams",
		reflect.TypeOf((*ClusterImageTemplateSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterRunTemplate",
		reflect.TypeOf((*ClusterRunTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterRunTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterRunTemplateProps",
		reflect.TypeOf((*ClusterRunTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterRunTemplateSpec",
		reflect.TypeOf((*ClusterRunTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterSourceTemplate",
		reflect.TypeOf((*ClusterSourceTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterSourceTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSourceTemplateProps",
		reflect.TypeOf((*ClusterSourceTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSourceTemplateSpec",
		reflect.TypeOf((*ClusterSourceTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSourceTemplateSpecParams",
		reflect.TypeOf((*ClusterSourceTemplateSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterSupplyChain",
		reflect.TypeOf((*ClusterSupplyChain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterSupplyChain{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainProps",
		reflect.TypeOf((*ClusterSupplyChainProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpec",
		reflect.TypeOf((*ClusterSupplyChainSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecParams",
		reflect.TypeOf((*ClusterSupplyChainSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResources",
		reflect.TypeOf((*ClusterSupplyChainSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResourcesConfigs",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResourcesImages",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesImages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResourcesParams",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResourcesSources",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecResourcesTemplateRef",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cartorun.ClusterSupplyChainSpecResourcesTemplateRefKind",
		reflect.TypeOf((*ClusterSupplyChainSpecResourcesTemplateRefKind)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_SOURCE_TEMPLATE": ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_SOURCE_TEMPLATE,
			"CLUSTER_IMAGE_TEMPLATE": ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_IMAGE_TEMPLATE,
			"CLUSTER_TEMPLATE": ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_TEMPLATE,
			"CLUSTER_CONFIG_TEMPLATE": ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_CONFIG_TEMPLATE,
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterSupplyChainSpecServiceAccountRef",
		reflect.TypeOf((*ClusterSupplyChainSpecServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.ClusterTemplate",
		reflect.TypeOf((*ClusterTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterTemplateProps",
		reflect.TypeOf((*ClusterTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterTemplateSpec",
		reflect.TypeOf((*ClusterTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.ClusterTemplateSpecParams",
		reflect.TypeOf((*ClusterTemplateSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.Deliverable",
		reflect.TypeOf((*Deliverable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Deliverable{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableProps",
		reflect.TypeOf((*DeliverableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableSpec",
		reflect.TypeOf((*DeliverableSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableSpecParams",
		reflect.TypeOf((*DeliverableSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableSpecSource",
		reflect.TypeOf((*DeliverableSpecSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableSpecSourceGit",
		reflect.TypeOf((*DeliverableSpecSourceGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.DeliverableSpecSourceGitRef",
		reflect.TypeOf((*DeliverableSpecSourceGitRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.Runnable",
		reflect.TypeOf((*Runnable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Runnable{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableProps",
		reflect.TypeOf((*RunnableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableSpec",
		reflect.TypeOf((*RunnableSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableSpecRetentionPolicy",
		reflect.TypeOf((*RunnableSpecRetentionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableSpecRunTemplateRef",
		reflect.TypeOf((*RunnableSpecRunTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableSpecSelector",
		reflect.TypeOf((*RunnableSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.RunnableSpecSelectorResource",
		reflect.TypeOf((*RunnableSpecSelectorResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.Workload",
		reflect.TypeOf((*Workload)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Workload{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadProps",
		reflect.TypeOf((*WorkloadProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpec",
		reflect.TypeOf((*WorkloadSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnv",
		reflect.TypeOf((*WorkloadSpecEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnvValueFrom",
		reflect.TypeOf((*WorkloadSpecEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkloadSpecEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnvValueFromFieldRef",
		reflect.TypeOf((*WorkloadSpecEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkloadSpecEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.WorkloadSpecEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkloadSpecEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkloadSpecEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkloadSpecEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuild",
		reflect.TypeOf((*WorkloadSpecBuild)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnv",
		reflect.TypeOf((*WorkloadSpecBuildEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnvValueFrom",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnvValueFromFieldRef",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecBuildEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkloadSpecBuildEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecParams",
		reflect.TypeOf((*WorkloadSpecParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecResources",
		reflect.TypeOf((*WorkloadSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cartorun.WorkloadSpecResourcesLimits",
		reflect.TypeOf((*WorkloadSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkloadSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"cartorun.WorkloadSpecResourcesRequests",
		reflect.TypeOf((*WorkloadSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkloadSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecServiceClaims",
		reflect.TypeOf((*WorkloadSpecServiceClaims)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecServiceClaimsRef",
		reflect.TypeOf((*WorkloadSpecServiceClaimsRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecSource",
		reflect.TypeOf((*WorkloadSpecSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecSourceGit",
		reflect.TypeOf((*WorkloadSpecSourceGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cartorun.WorkloadSpecSourceGitRef",
		reflect.TypeOf((*WorkloadSpecSourceGitRef)(nil)).Elem(),
	)
}
