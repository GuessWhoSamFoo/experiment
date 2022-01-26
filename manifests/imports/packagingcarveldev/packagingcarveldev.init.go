package packagingcarveldev

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"packagingcarveldev.PackageInstall",
		reflect.TypeOf((*PackageInstall)(nil)).Elem(),
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
			j := jsiiProxy_PackageInstall{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallProps",
		reflect.TypeOf((*PackageInstallProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpec",
		reflect.TypeOf((*PackageInstallSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecCluster",
		reflect.TypeOf((*PackageInstallSpecCluster)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecClusterKubeconfigSecretRef",
		reflect.TypeOf((*PackageInstallSpecClusterKubeconfigSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecPackageRef",
		reflect.TypeOf((*PackageInstallSpecPackageRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecPackageRefVersionSelection",
		reflect.TypeOf((*PackageInstallSpecPackageRefVersionSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecPackageRefVersionSelectionPrereleases",
		reflect.TypeOf((*PackageInstallSpecPackageRefVersionSelectionPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecValues",
		reflect.TypeOf((*PackageInstallSpecValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageInstallSpecValuesSecretRef",
		reflect.TypeOf((*PackageInstallSpecValuesSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"packagingcarveldev.PackageRepository",
		reflect.TypeOf((*PackageRepository)(nil)).Elem(),
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
			j := jsiiProxy_PackageRepository{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositoryProps",
		reflect.TypeOf((*PackageRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpec",
		reflect.TypeOf((*PackageRepositorySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetch",
		reflect.TypeOf((*PackageRepositorySpecFetch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchGit",
		reflect.TypeOf((*PackageRepositorySpecFetchGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchGitRefSelection",
		reflect.TypeOf((*PackageRepositorySpecFetchGitRefSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchGitRefSelectionSemver",
		reflect.TypeOf((*PackageRepositorySpecFetchGitRefSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchGitRefSelectionSemverPrereleases",
		reflect.TypeOf((*PackageRepositorySpecFetchGitRefSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchGitSecretRef",
		reflect.TypeOf((*PackageRepositorySpecFetchGitSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchHttp",
		reflect.TypeOf((*PackageRepositorySpecFetchHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchHttpSecretRef",
		reflect.TypeOf((*PackageRepositorySpecFetchHttpSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImage",
		reflect.TypeOf((*PackageRepositorySpecFetchImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImageSecretRef",
		reflect.TypeOf((*PackageRepositorySpecFetchImageSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImageTagSelection",
		reflect.TypeOf((*PackageRepositorySpecFetchImageTagSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImageTagSelectionSemver",
		reflect.TypeOf((*PackageRepositorySpecFetchImageTagSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImageTagSelectionSemverPrereleases",
		reflect.TypeOf((*PackageRepositorySpecFetchImageTagSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImgpkgBundle",
		reflect.TypeOf((*PackageRepositorySpecFetchImgpkgBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImgpkgBundleSecretRef",
		reflect.TypeOf((*PackageRepositorySpecFetchImgpkgBundleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImgpkgBundleTagSelection",
		reflect.TypeOf((*PackageRepositorySpecFetchImgpkgBundleTagSelection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImgpkgBundleTagSelectionSemver",
		reflect.TypeOf((*PackageRepositorySpecFetchImgpkgBundleTagSelectionSemver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchImgpkgBundleTagSelectionSemverPrereleases",
		reflect.TypeOf((*PackageRepositorySpecFetchImgpkgBundleTagSelectionSemverPrereleases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchInline",
		reflect.TypeOf((*PackageRepositorySpecFetchInline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchInlinePathsFrom",
		reflect.TypeOf((*PackageRepositorySpecFetchInlinePathsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchInlinePathsFromConfigMapRef",
		reflect.TypeOf((*PackageRepositorySpecFetchInlinePathsFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"packagingcarveldev.PackageRepositorySpecFetchInlinePathsFromSecretRef",
		reflect.TypeOf((*PackageRepositorySpecFetchInlinePathsFromSecretRef)(nil)).Elem(),
	)
}
