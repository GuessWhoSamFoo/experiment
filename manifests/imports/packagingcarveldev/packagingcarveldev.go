// packagingcarveldev
package packagingcarveldev

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/GuessWhoSamFoo/experiment/manifests/imports/packagingcarveldev/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/packagingcarveldev/internal"
)

type PackageInstall interface {
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

// The jsii proxy struct for PackageInstall
type jsiiProxy_PackageInstall struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_PackageInstall) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageInstall) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageInstall) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageInstall) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageInstall) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageInstall) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "PackageInstall" API object.
func NewPackageInstall(scope constructs.Construct, id *string, props *PackageInstallProps) PackageInstall {
	_init_.Initialize()

	j := jsiiProxy_PackageInstall{}

	_jsii_.Create(
		"packagingcarveldev.PackageInstall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "PackageInstall" API object.
func NewPackageInstall_Override(p PackageInstall, scope constructs.Construct, id *string, props *PackageInstallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"packagingcarveldev.PackageInstall",
		[]interface{}{scope, id, props},
		p,
	)
}

// Renders a Kubernetes manifest for "PackageInstall".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func PackageInstall_Manifest(props *PackageInstallProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"packagingcarveldev.PackageInstall",
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
func PackageInstall_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"packagingcarveldev.PackageInstall",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func PackageInstall_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"packagingcarveldev.PackageInstall",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (p *jsiiProxy_PackageInstall) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (p *jsiiProxy_PackageInstall) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PackageInstall) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (p *jsiiProxy_PackageInstall) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PackageInstall) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (p *jsiiProxy_PackageInstall) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PackageInstall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PackageInstallProps struct {
	Spec *PackageInstallSpec `json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

type PackageInstallSpec struct {
	// Canceled when set to true will stop all active changes.
	Canceled *bool `json:"canceled" yaml:"canceled"`
	Cluster *PackageInstallSpecCluster `json:"cluster" yaml:"cluster"`
	// When NoopDelete set to true, PackageInstall deletion should delete PackageInstall/App CR but preserve App's associated resources.
	NoopDelete *bool `json:"noopDelete" yaml:"noopDelete"`
	PackageRef *PackageInstallSpecPackageRef `json:"packageRef" yaml:"packageRef"`
	// Paused when set to true will ignore all pending changes, once it set back to false, pending changes will be applied.
	Paused *bool `json:"paused" yaml:"paused"`
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
	// Controls frequency of App reconciliation in time + unit format.
	//
	// Always >= 30s. If value below 30s is specified, 30s will be used.
	SyncPeriod *string `json:"syncPeriod" yaml:"syncPeriod"`
	Values *[]*PackageInstallSpecValues `json:"values" yaml:"values"`
}

type PackageInstallSpecCluster struct {
	KubeconfigSecretRef *PackageInstallSpecClusterKubeconfigSecretRef `json:"kubeconfigSecretRef" yaml:"kubeconfigSecretRef"`
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type PackageInstallSpecClusterKubeconfigSecretRef struct {
	Key *string `json:"key" yaml:"key"`
	Name *string `json:"name" yaml:"name"`
}

type PackageInstallSpecPackageRef struct {
	RefName *string `json:"refName" yaml:"refName"`
	VersionSelection *PackageInstallSpecPackageRefVersionSelection `json:"versionSelection" yaml:"versionSelection"`
}

type PackageInstallSpecPackageRefVersionSelection struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *PackageInstallSpecPackageRefVersionSelectionPrereleases `json:"prereleases" yaml:"prereleases"`
}

type PackageInstallSpecPackageRefVersionSelectionPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type PackageInstallSpecValues struct {
	SecretRef *PackageInstallSpecValuesSecretRef `json:"secretRef" yaml:"secretRef"`
}

type PackageInstallSpecValuesSecretRef struct {
	Key *string `json:"key" yaml:"key"`
	Name *string `json:"name" yaml:"name"`
}

type PackageRepository interface {
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

// The jsii proxy struct for PackageRepository
type jsiiProxy_PackageRepository struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_PackageRepository) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageRepository) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageRepository) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageRepository) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageRepository) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PackageRepository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "PackageRepository" API object.
func NewPackageRepository(scope constructs.Construct, id *string, props *PackageRepositoryProps) PackageRepository {
	_init_.Initialize()

	j := jsiiProxy_PackageRepository{}

	_jsii_.Create(
		"packagingcarveldev.PackageRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "PackageRepository" API object.
func NewPackageRepository_Override(p PackageRepository, scope constructs.Construct, id *string, props *PackageRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"packagingcarveldev.PackageRepository",
		[]interface{}{scope, id, props},
		p,
	)
}

// Renders a Kubernetes manifest for "PackageRepository".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func PackageRepository_Manifest(props *PackageRepositoryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"packagingcarveldev.PackageRepository",
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
func PackageRepository_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"packagingcarveldev.PackageRepository",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func PackageRepository_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"packagingcarveldev.PackageRepository",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (p *jsiiProxy_PackageRepository) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (p *jsiiProxy_PackageRepository) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PackageRepository) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (p *jsiiProxy_PackageRepository) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PackageRepository) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (p *jsiiProxy_PackageRepository) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PackageRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PackageRepositoryProps struct {
	Spec *PackageRepositorySpec `json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
}

type PackageRepositorySpec struct {
	Fetch *PackageRepositorySpecFetch `json:"fetch" yaml:"fetch"`
	// Paused when set to true will ignore all pending changes, once it set back to false, pending changes will be applied.
	Paused *bool `json:"paused" yaml:"paused"`
	// Controls frequency of PackageRepository reconciliation.
	SyncPeriod *string `json:"syncPeriod" yaml:"syncPeriod"`
}

type PackageRepositorySpecFetch struct {
	Git *PackageRepositorySpecFetchGit `json:"git" yaml:"git"`
	Http *PackageRepositorySpecFetchHttp `json:"http" yaml:"http"`
	Image *PackageRepositorySpecFetchImage `json:"image" yaml:"image"`
	ImgpkgBundle *PackageRepositorySpecFetchImgpkgBundle `json:"imgpkgBundle" yaml:"imgpkgBundle"`
	Inline *PackageRepositorySpecFetchInline `json:"inline" yaml:"inline"`
}

type PackageRepositorySpecFetchGit struct {
	LfsSkipSmudge *bool `json:"lfsSkipSmudge" yaml:"lfsSkipSmudge"`
	Ref *string `json:"ref" yaml:"ref"`
	RefSelection *PackageRepositorySpecFetchGitRefSelection `json:"refSelection" yaml:"refSelection"`
	// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
	SecretRef *PackageRepositorySpecFetchGitSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	Url *string `json:"url" yaml:"url"`
}

type PackageRepositorySpecFetchGitRefSelection struct {
	Semver *PackageRepositorySpecFetchGitRefSelectionSemver `json:"semver" yaml:"semver"`
}

type PackageRepositorySpecFetchGitRefSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *PackageRepositorySpecFetchGitRefSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type PackageRepositorySpecFetchGitRefSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

// Secret may include one or more keys: ssh-privatekey, ssh-knownhosts.
type PackageRepositorySpecFetchGitSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type PackageRepositorySpecFetchHttp struct {
	// Secret may include one or more keys: username, password.
	SecretRef *PackageRepositorySpecFetchHttpSecretRef `json:"secretRef" yaml:"secretRef"`
	Sha256 *string `json:"sha256" yaml:"sha256"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	// URL can point to one of following formats: text, tgz, zip.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password.
type PackageRepositorySpecFetchHttpSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type PackageRepositorySpecFetchImage struct {
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *PackageRepositorySpecFetchImageSecretRef `json:"secretRef" yaml:"secretRef"`
	SubPath *string `json:"subPath" yaml:"subPath"`
	TagSelection *PackageRepositorySpecFetchImageTagSelection `json:"tagSelection" yaml:"tagSelection"`
	// Example: username/app1-config:v0.1.0.
	Url *string `json:"url" yaml:"url"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type PackageRepositorySpecFetchImageSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type PackageRepositorySpecFetchImageTagSelection struct {
	Semver *PackageRepositorySpecFetchImageTagSelectionSemver `json:"semver" yaml:"semver"`
}

type PackageRepositorySpecFetchImageTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *PackageRepositorySpecFetchImageTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type PackageRepositorySpecFetchImageTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type PackageRepositorySpecFetchImgpkgBundle struct {
	Image *string `json:"image" yaml:"image"`
	// Secret may include one or more keys: username, password, token.
	//
	// By default anonymous access is used for authentication. TODO support docker config formated secret
	SecretRef *PackageRepositorySpecFetchImgpkgBundleSecretRef `json:"secretRef" yaml:"secretRef"`
	TagSelection *PackageRepositorySpecFetchImgpkgBundleTagSelection `json:"tagSelection" yaml:"tagSelection"`
}

// Secret may include one or more keys: username, password, token.
//
// By default anonymous access is used for authentication. TODO support docker config formated secret
type PackageRepositorySpecFetchImgpkgBundleSecretRef struct {
	// Object is expected to be within same namespace.
	Name *string `json:"name" yaml:"name"`
}

type PackageRepositorySpecFetchImgpkgBundleTagSelection struct {
	Semver *PackageRepositorySpecFetchImgpkgBundleTagSelectionSemver `json:"semver" yaml:"semver"`
}

type PackageRepositorySpecFetchImgpkgBundleTagSelectionSemver struct {
	Constraints *string `json:"constraints" yaml:"constraints"`
	Prereleases *PackageRepositorySpecFetchImgpkgBundleTagSelectionSemverPrereleases `json:"prereleases" yaml:"prereleases"`
}

type PackageRepositorySpecFetchImgpkgBundleTagSelectionSemverPrereleases struct {
	Identifiers *[]*string `json:"identifiers" yaml:"identifiers"`
}

type PackageRepositorySpecFetchInline struct {
	Paths *map[string]*string `json:"paths" yaml:"paths"`
	PathsFrom *[]*PackageRepositorySpecFetchInlinePathsFrom `json:"pathsFrom" yaml:"pathsFrom"`
}

type PackageRepositorySpecFetchInlinePathsFrom struct {
	ConfigMapRef *PackageRepositorySpecFetchInlinePathsFromConfigMapRef `json:"configMapRef" yaml:"configMapRef"`
	SecretRef *PackageRepositorySpecFetchInlinePathsFromSecretRef `json:"secretRef" yaml:"secretRef"`
}

type PackageRepositorySpecFetchInlinePathsFromConfigMapRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

type PackageRepositorySpecFetchInlinePathsFromSecretRef struct {
	DirectoryPath *string `json:"directoryPath" yaml:"directoryPath"`
	Name *string `json:"name" yaml:"name"`
}

