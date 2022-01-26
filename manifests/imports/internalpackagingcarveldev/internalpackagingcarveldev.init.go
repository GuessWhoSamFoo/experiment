package internalpackagingcarveldev

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"internalpackagingcarveldev.InternalPackage",
		reflect.TypeOf((*InternalPackage)(nil)).Elem(),
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
			j := jsiiProxy_InternalPackage{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"internalpackagingcarveldev.InternalPackageMetadata",
		reflect.TypeOf((*InternalPackageMetadata)(nil)).Elem(),
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
			j := jsiiProxy_InternalPackageMetadata{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageMetadataProps",
		reflect.TypeOf((*InternalPackageMetadataProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageMetadataSpec",
		reflect.TypeOf((*InternalPackageMetadataSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageMetadataSpecMaintainers",
		reflect.TypeOf((*InternalPackageMetadataSpecMaintainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageProps",
		reflect.TypeOf((*InternalPackageProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpec",
		reflect.TypeOf((*InternalPackageSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplate",
		reflect.TypeOf((*InternalPackageSpecTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpec",
		reflect.TypeOf((*InternalPackageSpecTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecCluster",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecClusterKubeconfigSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecClusterKubeconfigSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecDeploy",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecDeploy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecDeployKapp",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecDeployKapp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecDeployKappDelete",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecDeployKappDelete)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecDeployKappInspect",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecDeployKappInspect)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetch",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchGit",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchGitRefSelection",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchGitRefSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchGitRefSelectionSemver",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchGitRefSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchGitRefSelectionSemverPrereleases",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchGitRefSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchGitSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchGitSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchHelmChart",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchHelmChart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchHelmChartRepository",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchHelmChartRepository)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchHelmChartRepositorySecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchHelmChartRepositorySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchHttp",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchHttpSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchHttpSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImage",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImageSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImageSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImageTagSelection",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImageTagSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImageTagSelectionSemver",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImageTagSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImageTagSelectionSemverPrereleases",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImageTagSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImgpkgBundle",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImgpkgBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImgpkgBundleSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImgpkgBundleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelection",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemver",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemverPrereleases",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchInline",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchInline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchInlinePathsFrom",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchInlinePathsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchInlinePathsFromConfigMapRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchInlinePathsFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecFetchInlinePathsFromSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecFetchInlinePathsFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplate",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateHelmTemplate",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateHelmTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFrom",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromConfigMapRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateKbld",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateKbld)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateSops",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateSops)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateSopsAge",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateSopsAge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateSopsAgePrivateKeysSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateSopsAgePrivateKeysSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateSopsPgp",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateSopsPgp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateSopsPgpPrivateKeysSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateSopsPgpPrivateKeysSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYtt",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYtt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttInline",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttInline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttInlinePathsFrom",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttInlinePathsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromConfigMapRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttValuesFrom",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttValuesFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttValuesFromConfigMapRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttValuesFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecTemplateSpecTemplateYttValuesFromSecretRef",
		reflect.TypeOf((*InternalPackageSpecTemplateSpecTemplateYttValuesFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"internalpackagingcarveldev.InternalPackageSpecValuesSchema",
		reflect.TypeOf((*InternalPackageSpecValuesSchema)(nil)).Elem(),
	)
}
