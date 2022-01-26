// internalpackagingcarveldev
package internalpackagingcarveldev

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/GuessWhoSamFoo/experiment/manifests/imports/internalpackagingcarveldev/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/internalpackagingcarveldev/internal"
)

type InternalPackage interface {
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

// The jsii proxy struct for InternalPackage
type jsiiProxy_InternalPackage struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_InternalPackage) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackage) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackage) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackage) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackage) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "InternalPackage" API object.
func NewInternalPackage(scope constructs.Construct, id *string, props *InternalPackageProps) InternalPackage {
	_init_.Initialize()

	j := jsiiProxy_InternalPackage{}

	_jsii_.Create(
		"internalpackagingcarveldev.InternalPackage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "InternalPackage" API object.
func NewInternalPackage_Override(i InternalPackage, scope constructs.Construct, id *string, props *InternalPackageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"internalpackagingcarveldev.InternalPackage",
		[]interface{}{scope, id, props},
		i,
	)
}

// Renders a Kubernetes manifest for "InternalPackage".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func InternalPackage_Manifest(props *InternalPackageProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"internalpackagingcarveldev.InternalPackage",
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
func InternalPackage_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"internalpackagingcarveldev.InternalPackage",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func InternalPackage_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"internalpackagingcarveldev.InternalPackage",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (i *jsiiProxy_InternalPackage) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (i *jsiiProxy_InternalPackage) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
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
func (i *jsiiProxy_InternalPackage) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (i *jsiiProxy_InternalPackage) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
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
func (i *jsiiProxy_InternalPackage) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (i *jsiiProxy_InternalPackage) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_InternalPackage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type InternalPackageMetadata interface {
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

// The jsii proxy struct for InternalPackageMetadata
type jsiiProxy_InternalPackageMetadata struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_InternalPackageMetadata) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackageMetadata) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackageMetadata) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackageMetadata) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackageMetadata) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InternalPackageMetadata) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "InternalPackageMetadata" API object.
func NewInternalPackageMetadata(scope constructs.Construct, id *string, props *InternalPackageMetadataProps) InternalPackageMetadata {
	_init_.Initialize()

	j := jsiiProxy_InternalPackageMetadata{}

	_jsii_.Create(
		"internalpackagingcarveldev.InternalPackageMetadata",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "InternalPackageMetadata" API object.
func NewInternalPackageMetadata_Override(i InternalPackageMetadata, scope constructs.Construct, id *string, props *InternalPackageMetadataProps) {
	_init_.Initialize()

	_jsii_.Create(
		"internalpackagingcarveldev.InternalPackageMetadata",
		[]interface{}{scope, id, props},
		i,
	)
}

// Renders a Kubernetes manifest for "InternalPackageMetadata".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func InternalPackageMetadata_Manifest(props *InternalPackageMetadataProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"internalpackagingcarveldev.InternalPackageMetadata",
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
func InternalPackageMetadata_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"internalpackagingcarveldev.InternalPackageMetadata",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func InternalPackageMetadata_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"internalpackagingcarveldev.InternalPackageMetadata",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (i *jsiiProxy_InternalPackageMetadata) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (i *jsiiProxy_InternalPackageMetadata) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
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
func (i *jsiiProxy_InternalPackageMetadata) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (i *jsiiProxy_InternalPackageMetadata) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
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
func (i *jsiiProxy_InternalPackageMetadata) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (i *jsiiProxy_InternalPackageMetadata) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_InternalPackageMetadata) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type InternalPackageMetadataProps struct {
	Spec *InternalPackageMetadataSpec `json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

type InternalPackageMetadataSpec struct {
	Categories *[]*string `json:"categories" yaml:"categories"`
	DisplayName *string `json:"displayName" yaml:"displayName"`
	IconSvgBase64 *string `json:"iconSvgBase64" yaml:"iconSvgBase64"`
	LongDescription *string `json:"longDescription" yaml:"longDescription"`
	Maintainers *[]*InternalPackageMetadataSpecMaintainers `json:"maintainers" yaml:"maintainers"`
	ProviderName *string `json:"providerName" yaml:"providerName"`
	ShortDescription *string `json:"shortDescription" yaml:"shortDescription"`
	SupportDescription *string `json:"supportDescription" yaml:"supportDescription"`
}

type InternalPackageMetadataSpecMaintainers struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageProps struct {
	Spec *InternalPackageSpec `json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

type InternalPackageSpec struct {
	CapacityRequirementsDescription *string `json:"capacityRequirementsDescription" yaml:"capacityRequirementsDescription"`
	Licenses *[]*string `json:"licenses" yaml:"licenses"`
	RefName *string `json:"refName" yaml:"refName"`
	ReleasedAt *time.Time `json:"releasedAt" yaml:"releasedAt"`
	ReleaseNotes *string `json:"releaseNotes" yaml:"releaseNotes"`
	Template *InternalPackageSpecTemplate `json:"template" yaml:"template"`
	// valuesSchema can be used to show template values that can be configured by users when a Package is installed in an OpenAPI schema format.
	ValuesSchema *InternalPackageSpecValuesSchema `json:"valuesSchema" yaml:"valuesSchema"`
	Version *string `json:"version" yaml:"version"`
}

type InternalPackageSpecTemplate struct {
	Spec *InternalPackageSpecTemplateSpec `json:"spec" yaml:"spec"`
}

type InternalPackageSpecTemplateSpec struct {
	// Canceled when set to true will stop all active changes.
	Canceled *bool `json:"canceled" yaml:"canceled"`
	Cluster *InternalPackageSpecTemplateSpecCluster `json:"cluster" yaml:"cluster"`
	Deploy *[]*InternalPackageSpecTemplateSpecDeploy `json:"deploy" yaml:"deploy"`
	Fetch *[]*InternalPackageSpecTemplateSpecFetch `json:"fetch" yaml:"fetch"`
	// When NoopDeletion set to true, App deletion should delete App CR but preserve App's associated resources.
	NoopDelete *bool `json:"noopDelete" yaml:"noopDelete"`
	// Paused when set to true will ignore all pending changes, once it set back to false, pending changes will be applied.
	Paused *bool `json:"paused" yaml:"paused"`
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
	// Controls frequency of app reconciliation.
	SyncPeriod *string `json:"syncPeriod" yaml:"syncPeriod"`
	Template *[]*InternalPackageSpecTemplateSpecTemplate `json:"template" yaml:"template"`
}

type InternalPackageSpecTemplateSpecCluster struct {
	KubeconfigSecretRef *InternalPackageSpecTemplateSpecClusterKubeconfigSecretRef `json:"kubeconfigSecretRef" yaml:"kubeconfigSecretRef"`
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type InternalPackageSpecTemplateSpecClusterKubeconfigSecretRef struct {
	Key *string `json:"key" yaml:"key"`
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecDeploy struct {
	Kapp *InternalPackageSpecTemplateSpecDeployKapp `json:"kapp" yaml:"kapp"`
}

type InternalPackageSpecTemplateSpecDeployKapp struct {
	Delete *InternalPackageSpecTemplateSpecDeployKappDelete `json:"delete" yaml:"delete"`
	Inspect *InternalPackageSpecTemplateSpecDeployKappInspect `json:"inspect" yaml:"inspect"`
	IntoNs *string `json:"intoNs" yaml:"intoNs"`
	MapNs *[]*string `json:"mapNs" yaml:"mapNs"`
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type InternalPackageSpecTemplateSpecDeployKappDelete struct {
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type InternalPackageSpecTemplateSpecDeployKappInspect struct {
	RawOptions *[]*string `json:"rawOptions" yaml:"rawOptions"`
}

type InternalPackageSpecTemplateSpecFetch struct {
	Git *InternalPackageSpecTemplateSpecFetchGit `json:"git" yaml:"git"`
	HelmChart *InternalPackageSpecTemplateSpecFetchHelmChart `json:"helmChart" yaml:"helmChart"`
	Http *InternalPackageSpecTemplateSpecFetchHttp `json:"http" yaml:"http"`
	Image *InternalPackageSpecTemplateSpecFetchImage `json:"image" yaml:"image"`
	ImgpkgBundle *InternalPackageSpecTemplateSpecFetchImgpkgBundle `json:"imgpkgBundle" yaml:"imgpkgBundle"`
	Inline *InternalPackageSpecTemplateSpecFetchInline `json:"inline" yaml:"inline"`
}

type InternalPackageSpecTemplateSpecFetchGit struct {
	LfsSkipSmudge *bool `json:"lfsSkipSmudge" yaml:"lfsSkipSmudge"`
	Ref *string `json:"ref" yaml:"ref"`
	RefSelection *InternalPackageSpecTemplateSpecFetchGitRefSelection `json:"refSelection" yaml:"refSelection"`
	// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
	SecretRef *InternalPackageSpecTemplateSpecFetchGitSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	Url *string `json:"url" yaml:"url"`
}

type InternalPackageSpecTemplateSpecFetchGitRefSelection struct {
	Semver *InternalPackageSpecTemplateSpecFetchGitRefSelectionSemver `json:"semver" yaml:"semver"`
}

type InternalPackageSpecTemplateSpecFetchGitRefSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *InternalPackageSpecTemplateSpecFetchGitRefSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type InternalPackageSpecTemplateSpecFetchGitRefSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
type InternalPackageSpecTemplateSpecFetchGitSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchHelmChart struct {
	// Example: stable/redis.
	Name *string `json:"name" yaml:"name"`
	Repository *InternalPackageSpecTemplateSpecFetchHelmChartRepository `json:"repository" yaml:"repository"`
	Version *string `json:"version" yaml:"version"`
}

type InternalPackageSpecTemplateSpecFetchHelmChartRepository struct {
	SecretRef *InternalPackageSpecTemplateSpecFetchHelmChartRepositorySecretRef `json:"secretRef" yaml:"secretRef"`
	Url *string `json:"url" yaml:"url"`
}

type InternalPackageSpecTemplateSpecFetchHelmChartRepositorySecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchHttp struct {
	// Secret may include one or more keys: username, password.
	SecretRef *InternalPackageSpecTemplateSpecFetchHttpSecretRef `json:"secretRef" yaml:"secretRef"`
	Sha256 *string `json:"sha256" yaml:"sha256"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	// URL can point to one of following formats: text, tgz, zip.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password.
type InternalPackageSpecTemplateSpecFetchHttpSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchImage struct {
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *InternalPackageSpecTemplateSpecFetchImageSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	TagSelection *InternalPackageSpecTemplateSpecFetchImageTagSelection `json:"tagSelection" yaml:"tagSelection"`
	// Example: username/app1-config:v0.1.0.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type InternalPackageSpecTemplateSpecFetchImageSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchImageTagSelection struct {
	Semver *InternalPackageSpecTemplateSpecFetchImageTagSelectionSemver `json:"semver" yaml:"semver"`
}

type InternalPackageSpecTemplateSpecFetchImageTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *InternalPackageSpecTemplateSpecFetchImageTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type InternalPackageSpecTemplateSpecFetchImageTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type InternalPackageSpecTemplateSpecFetchImgpkgBundle struct {
	Image *string `json:"image" yaml:"image"`
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *InternalPackageSpecTemplateSpecFetchImgpkgBundleSecretRef `json:"secretRef" yaml:"secretRef"`
	TagSelection *InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelection `json:"tagSelection" yaml:"tagSelection"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type InternalPackageSpecTemplateSpecFetchImgpkgBundleSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelection struct {
	Semver *InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemver `json:"semver" yaml:"semver"`
}

type InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type InternalPackageSpecTemplateSpecFetchImgpkgBundleTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type InternalPackageSpecTemplateSpecFetchInline struct {
	Paths *map[string]*string `json:"paths" yaml:"paths"`
	PathsFrom *[]*InternalPackageSpecTemplateSpecFetchInlinePathsFrom `json:"pathsFrom" yaml:"pathsFrom"`
}

type InternalPackageSpecTemplateSpecFetchInlinePathsFrom struct {
	ConfigMapRef *InternalPackageSpecTemplateSpecFetchInlinePathsFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	SecretRef *InternalPackageSpecTemplateSpecFetchInlinePathsFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type InternalPackageSpecTemplateSpecFetchInlinePathsFromConfigMapRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecFetchInlinePathsFromSecretRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplate struct {
	HelmTemplate *InternalPackageSpecTemplateSpecTemplateHelmTemplate `json:"helmTemplate" yaml:"helmTemplate"`
	// TODO implement jsonnet.
	Jsonnet interface{} `json:"jsonnet" yaml:"jsonnet"`
	Kbld *InternalPackageSpecTemplateSpecTemplateKbld `json:"kbld" yaml:"kbld"`
	// TODO implement kustomize.
	Kustomize interface{} `json:"kustomize" yaml:"kustomize"`
	Sops *InternalPackageSpecTemplateSpecTemplateSops `json:"sops" yaml:"sops"`
	Ytt *InternalPackageSpecTemplateSpecTemplateYtt `json:"ytt" yaml:"ytt"`
}

type InternalPackageSpecTemplateSpecTemplateHelmTemplate struct {
	Name *string `json:"name" yaml:"name"`
	Namespace *string `json:"namespace" yaml:"namespace"`
	Path *string `json:"path" yaml:"path"`
	ValuesFrom *[]*InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFrom `json:"valuesFrom" yaml:"valuesFrom"`
}

type InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFrom struct {
	ConfigMapRef *InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	Path *string `json:"path" yaml:"path"`
	SecretRef *InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromConfigMapRef struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateHelmTemplateValuesFromSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateKbld struct {
	Paths *[]*string `json:"paths" yaml:"paths"`
}

type InternalPackageSpecTemplateSpecTemplateSops struct {
	Age *InternalPackageSpecTemplateSpecTemplateSopsAge `json:"age" yaml:"age"`
	Paths *[]*string `json:"paths" yaml:"paths"`
	Pgp *InternalPackageSpecTemplateSpecTemplateSopsPgp `json:"pgp" yaml:"pgp"`
}

type InternalPackageSpecTemplateSpecTemplateSopsAge struct {
	PrivateKeysSecretRef *InternalPackageSpecTemplateSpecTemplateSopsAgePrivateKeysSecretRef `json:"privateKeysSecretRef" yaml:"privateKeysSecretRef"`
}

type InternalPackageSpecTemplateSpecTemplateSopsAgePrivateKeysSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateSopsPgp struct {
	PrivateKeysSecretRef *InternalPackageSpecTemplateSpecTemplateSopsPgpPrivateKeysSecretRef `json:"privateKeysSecretRef" yaml:"privateKeysSecretRef"`
}

type InternalPackageSpecTemplateSpecTemplateSopsPgpPrivateKeysSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateYtt struct {
	FileMarks *[]*string `json:"fileMarks" yaml:"fileMarks"`
	IgnoreUnknownComments *bool `json:"ignoreUnknownComments" yaml:"ignoreUnknownComments"`
	Inline *InternalPackageSpecTemplateSpecTemplateYttInline `json:"inline" yaml:"inline"`
	Paths *[]*string `json:"paths" yaml:"paths"`
	Strict *bool `json:"strict" yaml:"strict"`
	ValuesFrom *[]*InternalPackageSpecTemplateSpecTemplateYttValuesFrom `json:"valuesFrom" yaml:"valuesFrom"`
}

type InternalPackageSpecTemplateSpecTemplateYttInline struct {
	Paths *map[string]*string `json:"paths" yaml:"paths"`
	PathsFrom *[]*InternalPackageSpecTemplateSpecTemplateYttInlinePathsFrom `json:"pathsFrom" yaml:"pathsFrom"`
}

type InternalPackageSpecTemplateSpecTemplateYttInlinePathsFrom struct {
	ConfigMapRef *InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	SecretRef *InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromConfigMapRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateYttInlinePathsFromSecretRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateYttValuesFrom struct {
	ConfigMapRef *InternalPackageSpecTemplateSpecTemplateYttValuesFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	Path *string `json:"path" yaml:"path"`
	SecretRef *InternalPackageSpecTemplateSpecTemplateYttValuesFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type InternalPackageSpecTemplateSpecTemplateYttValuesFromConfigMapRef struct {
	Name *string `json:"name" yaml:"name"`
}

type InternalPackageSpecTemplateSpecTemplateYttValuesFromSecretRef struct {
	Name *string `json:"name" yaml:"name"`
}

// valuesSchema can be used to show template values that can be configured by users when a Package is installed in an OpenAPI schema format.
type InternalPackageSpecValuesSchema struct {
	OpenApIv3 interface{} `json:"openApIv3" yaml:"openApIv3"`
}

