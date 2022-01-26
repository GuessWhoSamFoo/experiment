// kappctrlk14sio
package kappctrlk14sio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/GuessWhoSamFoo/experiment/manifests/imports/kappctrlk14sio/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/kappctrlk14sio/internal"
)

type App interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_App) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "App" API object.
func NewApp(scope constructs.Construct, id *string, props *AppProps) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"kappctrlk14sio.App",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "App" API object.
func NewApp_Override(a App, scope constructs.Construct, id *string, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"kappctrlk14sio.App",
		[]interface{}{scope, id, props},
		a,
	)
}

// Renders a Kubernetes manifest for "App".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func App_Manifest(props *AppProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"kappctrlk14sio.App",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func App_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"kappctrlk14sio.App",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func App_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"kappctrlk14sio.App",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (a *jsiiProxy_App) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (a *jsiiProxy_App) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (a *jsiiProxy_App) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (a *jsiiProxy_App) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_App) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (a *jsiiProxy_App) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppProps struct {
	Spec *AppSpec `json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

type AppSpec struct {
	// Canceled when set to true will stop all active changes.
	Canceled *bool `json:"canceled" yaml:"canceled"`
	Cluster *AppSpecCluster `json:"cluster" yaml:"cluster"`
	Deploy *[]*AppSpecDeploy `json:"deploy" yaml:"deploy"`
	Fetch *[]*AppSpecFetch `json:"fetch" yaml:"fetch"`
	// When NoopDeletion set to true, App deletion should delete App CR but preserve App's associated resources.
	NoopDelete *bool `json:"noopDelete" yaml:"noopDelete"`
	// Paused when set to true will ignore all pending changes, once it set back to false, pending changes will be applied.
	Paused *bool `json:"paused" yaml:"paused"`
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
	// Controls frequency of app reconciliation.
	SyncPeriod *string `json:"syncPeriod" yaml:"syncPeriod"`
	Template *[]*AppSpecTemplate `json:"template" yaml:"template"`
}

type AppSpecCluster struct {
	KubeconfigSecretRef *AppSpecClusterKubeconfigSecretRef `json:"kubeconfigSecretRef" yaml:"kubeconfigSecretRef"`
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type AppSpecClusterKubeconfigSecretRef struct {
	Key *string `json:"key" yaml:"key"`
	Name *string `json:"name" yaml:"name"`
}

type AppSpecDeploy struct {
	Kapp *AppSpecDeployKapp `json:"kapp" yaml:"kapp"`
}

type AppSpecDeployKapp struct {
	Delete *AppSpecDeployKappDelete `json:"delete" yaml:"delete"`
	Inspect *AppSpecDeployKappInspect `json:"inspect" yaml:"inspect"`
	IntoNs *string `json:"intoNs" yaml:"intoNs"`
	MapNs *[]*string `json:"mapNs" yaml:"mapNs"`
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type AppSpecDeployKappDelete struct {
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type AppSpecDeployKappInspect struct {
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type AppSpecFetch struct {
	Git *AppSpecFetchGit `json:"git" yaml:"git"`
	HelmChart *AppSpecFetchHelmChart `json:"helmChart" yaml:"helmChart"`
	Http *AppSpecFetchHttp `json:"http" yaml:"http"`
	Image *AppSpecFetchImage `json:"image" yaml:"image"`
	ImgpkgBundle *AppSpecFetchImgpkgBundle `json:"imgpkgBundle" yaml:"imgpkgBundle"`
	Inline *AppSpecFetchInline `json:"inline" yaml:"inline"`
}

type AppSpecFetchGit struct {
	LfsSkipSmudge *bool `json:"lfsSkipSmudge" yaml:"lfsSkipSmudge"`
	Ref *string `json:"ref" yaml:"ref"`
	RefSelection *AppSpecFetchGitRefSelection `json:"refSelection" yaml:"refSelection"`
	// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
	SecretRef *AppSpecFetchGitSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	Url *string `json:"url" yaml:"url"`
}

type AppSpecFetchGitRefSelection struct {
	Semver *AppSpecFetchGitRefSelectionSemver `json:"semver" yaml:"semver"`
}

type AppSpecFetchGitRefSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *AppSpecFetchGitRefSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type AppSpecFetchGitRefSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
type AppSpecFetchGitSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchHelmChart struct {
	// Example: stable/redis.
	Name *string `json:"name" yaml:"name"`
	Repository *AppSpecFetchHelmChartRepository `json:"repository" yaml:"repository"`
	Version *string `json:"version" yaml:"version"`
}

type AppSpecFetchHelmChartRepository struct {
	SecretRef *AppSpecFetchHelmChartRepositorySecretRef `json:"secretRef" yaml:"secretRef"`
	Url *string `json:"url" yaml:"url"`
}

type AppSpecFetchHelmChartRepositorySecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchHttp struct {
	// Secret may include one or more keys: username, password.
	SecretRef *AppSpecFetchHttpSecretRef `json:"secretRef" yaml:"secretRef"`
	Sha256 *string `json:"sha256" yaml:"sha256"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	// URL can point to one of following formats: text, tgz, zip.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password.
type AppSpecFetchHttpSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchImage struct {
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *AppSpecFetchImageSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	TagSelection *AppSpecFetchImageTagSelection `json:"tagSelection" yaml:"tagSelection"`
	// Example: username/app1-config:v0.1.0.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type AppSpecFetchImageSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchImageTagSelection struct {
	Semver *AppSpecFetchImageTagSelectionSemver `json:"semver" yaml:"semver"`
}

type AppSpecFetchImageTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *AppSpecFetchImageTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type AppSpecFetchImageTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type AppSpecFetchImgpkgBundle struct {
	Image *string `json:"image" yaml:"image"`
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *AppSpecFetchImgpkgBundleSecretRef `json:"secretRef" yaml:"secretRef"`
	TagSelection *AppSpecFetchImgpkgBundleTagSelection `json:"tagSelection" yaml:"tagSelection"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type AppSpecFetchImgpkgBundleSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchImgpkgBundleTagSelection struct {
	Semver *AppSpecFetchImgpkgBundleTagSelectionSemver `json:"semver" yaml:"semver"`
}

type AppSpecFetchImgpkgBundleTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *AppSpecFetchImgpkgBundleTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type AppSpecFetchImgpkgBundleTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type AppSpecFetchInline struct {
	Paths *map[string]*string `json:"paths" yaml:"paths"`
	PathsFrom *[]*AppSpecFetchInlinePathsFrom `json:"pathsFrom" yaml:"pathsFrom"`
}

type AppSpecFetchInlinePathsFrom struct {
	ConfigMapRef *AppSpecFetchInlinePathsFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	SecretRef *AppSpecFetchInlinePathsFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type AppSpecFetchInlinePathsFromConfigMapRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type AppSpecFetchInlinePathsFromSecretRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplate struct {
	HelmTemplate *AppSpecTemplateHelmTemplate `json:"helmTemplate" yaml:"helmTemplate"`
	// TODO implement jsonnet.
	Jsonnet interface{} `json:"jsonnet" yaml:"jsonnet"`
	Kbld *AppSpecTemplateKbld `json:"kbld" yaml:"kbld"`
	// TODO implement kustomize.
	Kustomize interface{} `json:"kustomize" yaml:"kustomize"`
	Sops *AppSpecTemplateSops `json:"sops" yaml:"sops"`
	Ytt *AppSpecTemplateYtt `json:"ytt" yaml:"ytt"`
}

type AppSpecTemplateHelmTemplate struct {
	Name *string `json:"name" yaml:"name"`
	Namespace *string `json:"namespace" yaml:"namespace"`
	Path *string `json:"path" yaml:"path"`
	ValuesFrom *[]*AppSpecTemplateHelmTemplateValuesFrom `json:"valuesFrom" yaml:"valuesFrom"`
}

type AppSpecTemplateHelmTemplateValuesFrom struct {
	ConfigMapRef *AppSpecTemplateHelmTemplateValuesFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	Path *string `json:"path" yaml:"path"`
	SecretRef *AppSpecTemplateHelmTemplateValuesFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type AppSpecTemplateHelmTemplateValuesFromConfigMapRef struct {
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateHelmTemplateValuesFromSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateKbld struct {
	Paths *[]*string `json:"paths" yaml:"paths"`
}

type AppSpecTemplateSops struct {
	Age *AppSpecTemplateSopsAge `json:"age" yaml:"age"`
	Paths *[]*string `json:"paths" yaml:"paths"`
	Pgp *AppSpecTemplateSopsPgp `json:"pgp" yaml:"pgp"`
}

type AppSpecTemplateSopsAge struct {
	PrivateKeysSecretRef *AppSpecTemplateSopsAgePrivateKeysSecretRef `json:"privateKeysSecretRef" yaml:"privateKeysSecretRef"`
}

type AppSpecTemplateSopsAgePrivateKeysSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateSopsPgp struct {
	PrivateKeysSecretRef *AppSpecTemplateSopsPgpPrivateKeysSecretRef `json:"privateKeysSecretRef" yaml:"privateKeysSecretRef"`
}

type AppSpecTemplateSopsPgpPrivateKeysSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateYtt struct {
	FileMarks *[]*string `json:"fileMarks" yaml:"fileMarks"`
	IgnoreUnknownComments *bool `json:"ignoreUnknownComments" yaml:"ignoreUnknownComments"`
	Inline *AppSpecTemplateYttInline `json:"inline" yaml:"inline"`
	Paths *[]*string `json:"paths" yaml:"paths"`
	Strict *bool `json:"strict" yaml:"strict"`
	ValuesFrom *[]*AppSpecTemplateYttValuesFrom `json:"valuesFrom" yaml:"valuesFrom"`
}

type AppSpecTemplateYttInline struct {
	Paths *map[string]*string `json:"paths" yaml:"paths"`
	PathsFrom *[]*AppSpecTemplateYttInlinePathsFrom `json:"pathsFrom" yaml:"pathsFrom"`
}

type AppSpecTemplateYttInlinePathsFrom struct {
	ConfigMapRef *AppSpecTemplateYttInlinePathsFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	SecretRef *AppSpecTemplateYttInlinePathsFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type AppSpecTemplateYttInlinePathsFromConfigMapRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateYttInlinePathsFromSecretRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateYttValuesFrom struct {
	ConfigMapRef *AppSpecTemplateYttValuesFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	Path *string `json:"path" yaml:"path"`
	SecretRef *AppSpecTemplateYttValuesFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type AppSpecTemplateYttValuesFromConfigMapRef struct {
	Name *string `json:"name" yaml:"name"`
}

type AppSpecTemplateYttValuesFromSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

