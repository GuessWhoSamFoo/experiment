// cartorun
package cartorun

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/GuessWhoSamFoo/experiment/manifests/imports/cartorun/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/cartorun/internal"
)

type ClusterConfigTemplate interface {
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

// The jsii proxy struct for ClusterConfigTemplate
type jsiiProxy_ClusterConfigTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterConfigTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterConfigTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterConfigTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterConfigTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterConfigTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterConfigTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterConfigTemplate" API object.
func NewClusterConfigTemplate(scope constructs.Construct, id *string, props *ClusterConfigTemplateProps) ClusterConfigTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterConfigTemplate{}

	_jsii_.Create(
		"cartorun.ClusterConfigTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterConfigTemplate" API object.
func NewClusterConfigTemplate_Override(c ClusterConfigTemplate, scope constructs.Construct, id *string, props *ClusterConfigTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterConfigTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterConfigTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterConfigTemplate_Manifest(props *ClusterConfigTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterConfigTemplate",
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
func ClusterConfigTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterConfigTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterConfigTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterConfigTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterConfigTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterConfigTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterConfigTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterConfigTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterConfigTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterConfigTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterConfigTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterConfigTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the config template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/template/#clusterconfigtemplate
	Spec *ClusterConfigTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the config template.
//
// More info: https://cartographer.sh/docs/latest/reference/template/#clusterconfigtemplate
type ClusterConfigTemplateSpec struct {
	// ConfigPath is a path into the templated object's data that contains valid yaml.
	//
	// This is typically the information that will configure the components of the deployable image. ConfigPath is specified in jsonpath format, eg: .data
	ConfigPath *string `json:"configPath" yaml:"configPath"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterConfigTemplateSpecParams `json:"params" yaml:"params"`
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Template interface{} `json:"template" yaml:"template"`
	// Ytt defines a resource template written in `ytt` for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Ytt *string `json:"ytt" yaml:"ytt"`
}

type ClusterConfigTemplateSpecParams struct {
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner or Template does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Name of a parameter the template accepts from the Blueprint or Owner.
	Name *string `json:"name" yaml:"name"`
}

type ClusterDelivery interface {
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

// The jsii proxy struct for ClusterDelivery
type jsiiProxy_ClusterDelivery struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterDelivery) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDelivery) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDelivery) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDelivery) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDelivery) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDelivery) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterDelivery" API object.
func NewClusterDelivery(scope constructs.Construct, id *string, props *ClusterDeliveryProps) ClusterDelivery {
	_init_.Initialize()

	j := jsiiProxy_ClusterDelivery{}

	_jsii_.Create(
		"cartorun.ClusterDelivery",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterDelivery" API object.
func NewClusterDelivery_Override(c ClusterDelivery, scope constructs.Construct, id *string, props *ClusterDeliveryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterDelivery",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterDelivery".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterDelivery_Manifest(props *ClusterDeliveryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterDelivery",
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
func ClusterDelivery_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterDelivery",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterDelivery_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterDelivery",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterDelivery) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterDelivery) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterDelivery) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterDelivery) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterDelivery) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterDelivery) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterDelivery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterDeliveryProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the delivery.
	//
	// More info: https://cartographer.sh/docs/latest/reference/deliverable/#clusterdelivery
	Spec *ClusterDeliverySpec `json:"spec" yaml:"spec"`
}

// Spec describes the delivery.
//
// More info: https://cartographer.sh/docs/latest/reference/deliverable/#clusterdelivery
type ClusterDeliverySpec struct {
	// Resources that are responsible for deploying and validating the deliverable.
	Resources *[]*ClusterDeliverySpecResources `json:"resources" yaml:"resources"`
	// Specifies the label key-value pairs used to select deliverables See: https://cartographer.sh/docs/v0.1.0/architecture/#selectors.
	Selector *map[string]*string `json:"selector" yaml:"selector"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterDeliverySpecParams `json:"params" yaml:"params"`
	// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
	//
	// If not set, Cartographer will use serviceAccountName from supply chain.
	// If that is also not set, Cartographer will use the default service account in the workload's namespace.
	ServiceAccountRef *ClusterDeliverySpecServiceAccountRef `json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

type ClusterDeliverySpecParams struct {
	// Name of the parameter.
	//
	// Should match a template parameter name.
	Name *string `json:"name" yaml:"name"`
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Value of the parameter.
	//
	// If specified, owner properties are ignored.
	Value interface{} `json:"value" yaml:"value"`
}

type ClusterDeliverySpecResources struct {
	// Name of the resource.
	//
	// Used as a reference for inputs, as well as being the name presented in deliverable statuses to identify this resource.
	Name *string `json:"name" yaml:"name"`
	// TemplateRef identifies the template used to produce this resource.
	TemplateRef *ClusterDeliverySpecResourcesTemplateRef `json:"templateRef" yaml:"templateRef"`
	// Configs is a list of references to other 'config' resources in this list.
	//
	// A config resource has the kind ClusterConfigTemplate
	// In a template, configs can be consumed as: $(configs.<name>.config)$
	// If there is only one image, it can be consumed as: $(config)$
	Configs *[]*ClusterDeliverySpecResourcesConfigs `json:"configs" yaml:"configs"`
	// Deployment is a reference to a 'deployment' resource.
	//
	// A deployment resource has the kind ClusterDeploymentTemplate
	// In a template, the deployment can be consumed as: $(deployment.url)$ and $(deployment.revision)$
	Deployment *ClusterDeliverySpecResourcesDeployment `json:"deployment" yaml:"deployment"`
	// Params are a list of parameters to provide to the template in TemplateRef Template params do not have to be specified here, unless you want to force a particular value, or add a default value.
	//
	// Parameters are consumed in a template with the syntax: $(params.<name>)$
	Params *[]*ClusterDeliverySpecResourcesParams `json:"params" yaml:"params"`
	// Sources is a list of references to other 'source' resources in this list.
	//
	// A source resource has the kind ClusterSourceTemplate or ClusterDeploymentTemplate
	// In a template, sources can be consumed as: $(sources.<name>.url)$ and $(sources.<name>.revision)$
	// If there is only one source, it can be consumed as: $(source.url)$ and $(source.revision)$
	Sources *[]*ClusterDeliverySpecResourcesSources `json:"sources" yaml:"sources"`
}

type ClusterDeliverySpecResourcesConfigs struct {
	Name *string `json:"name" yaml:"name"`
	Resource *string `json:"resource" yaml:"resource"`
}

// Deployment is a reference to a 'deployment' resource.
//
// A deployment resource has the kind ClusterDeploymentTemplate
// In a template, the deployment can be consumed as: $(deployment.url)$ and $(deployment.revision)$
type ClusterDeliverySpecResourcesDeployment struct {
	Resource *string `json:"resource" yaml:"resource"`
}

type ClusterDeliverySpecResourcesParams struct {
	// Name of the parameter.
	//
	// Should match a template parameter name.
	Name *string `json:"name" yaml:"name"`
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Value of the parameter.
	//
	// If specified, owner properties are ignored.
	Value interface{} `json:"value" yaml:"value"`
}

type ClusterDeliverySpecResourcesSources struct {
	Name *string `json:"name" yaml:"name"`
	Resource *string `json:"resource" yaml:"resource"`
}

// TemplateRef identifies the template used to produce this resource.
type ClusterDeliverySpecResourcesTemplateRef struct {
	// Kind of the template to apply.
	Kind ClusterDeliverySpecResourcesTemplateRefKind `json:"kind" yaml:"kind"`
	// Name of the template to apply.
	Name *string `json:"name" yaml:"name"`
}

// Kind of the template to apply.
type ClusterDeliverySpecResourcesTemplateRefKind string

const (
	ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_SOURCE_TEMPLATE ClusterDeliverySpecResourcesTemplateRefKind = "CLUSTER_SOURCE_TEMPLATE"
	ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_DEPLOYMENT_TEMPLATE ClusterDeliverySpecResourcesTemplateRefKind = "CLUSTER_DEPLOYMENT_TEMPLATE"
	ClusterDeliverySpecResourcesTemplateRefKind_CLUSTER_TEMPLATE ClusterDeliverySpecResourcesTemplateRefKind = "CLUSTER_TEMPLATE"
)

// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
//
// If not set, Cartographer will use serviceAccountName from supply chain.
// If that is also not set, Cartographer will use the default service account in the workload's namespace.
type ClusterDeliverySpecServiceAccountRef struct {
	// Name of the service account being referred to.
	Name *string `json:"name" yaml:"name"`
	// Namespace of the service account being referred to if omitted, the Owner's namespace is used.
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type ClusterDeploymentTemplate interface {
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

// The jsii proxy struct for ClusterDeploymentTemplate
type jsiiProxy_ClusterDeploymentTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterDeploymentTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDeploymentTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDeploymentTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDeploymentTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDeploymentTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterDeploymentTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterDeploymentTemplate" API object.
func NewClusterDeploymentTemplate(scope constructs.Construct, id *string, props *ClusterDeploymentTemplateProps) ClusterDeploymentTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterDeploymentTemplate{}

	_jsii_.Create(
		"cartorun.ClusterDeploymentTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterDeploymentTemplate" API object.
func NewClusterDeploymentTemplate_Override(c ClusterDeploymentTemplate, scope constructs.Construct, id *string, props *ClusterDeploymentTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterDeploymentTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterDeploymentTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterDeploymentTemplate_Manifest(props *ClusterDeploymentTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterDeploymentTemplate",
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
func ClusterDeploymentTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterDeploymentTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterDeploymentTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterDeploymentTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterDeploymentTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterDeploymentTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterDeploymentTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterDeploymentTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterDeploymentTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterDeploymentTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterDeploymentTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterDeploymentTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the deployment template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/template/#clusterdeploymenttemplate
	Spec *ClusterDeploymentTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the deployment template.
//
// More info: https://cartographer.sh/docs/latest/reference/template/#clusterdeploymenttemplate
type ClusterDeploymentTemplateSpec struct {
	// ObservedCompletion describe the criteria for determining that the templated object completed configuration of environment.
	//
	// These criteria assert completion when metadata.Generation and status.ObservedGeneration match, AND success or failure criteria match. Cannot specify both ObservedMatches and ObservedCompletion.
	ObservedCompletion *ClusterDeploymentTemplateSpecObservedCompletion `json:"observedCompletion" yaml:"observedCompletion"`
	// ObservedMatches describe the criteria for determining that the templated object completed configuration of environment.
	//
	// These criteria assert completion when an output (usually a field in .status) matches an input (usually a field in .spec) Cannot specify both ObservedMatches and ObservedCompletion.
	ObservedMatches *[]*ClusterDeploymentTemplateSpecObservedMatches `json:"observedMatches" yaml:"observedMatches"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterDeploymentTemplateSpecParams `json:"params" yaml:"params"`
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Template interface{} `json:"template" yaml:"template"`
	// Ytt defines a resource template written in `ytt` for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Ytt *string `json:"ytt" yaml:"ytt"`
}

// ObservedCompletion describe the criteria for determining that the templated object completed configuration of environment.
//
// These criteria assert completion when metadata.Generation and status.ObservedGeneration match, AND success or failure criteria match. Cannot specify both ObservedMatches and ObservedCompletion.
type ClusterDeploymentTemplateSpecObservedCompletion struct {
	// SucceededCondition, when matched, indicates that the input was successfully deployed.
	Succeeded *ClusterDeploymentTemplateSpecObservedCompletionSucceeded `json:"succeeded" yaml:"succeeded"`
	// FailedCondition, when matched, indicates that the input did not deploy successfully.
	Failed *ClusterDeploymentTemplateSpecObservedCompletionFailed `json:"failed" yaml:"failed"`
}

// FailedCondition, when matched, indicates that the input did not deploy successfully.
type ClusterDeploymentTemplateSpecObservedCompletionFailed struct {
	// Key is a jsonPath expression pointing to the field to inspect on the templated object, eg: 'status.conditions[?(@.type=="Succeeded")].status'.
	Key *string `json:"key" yaml:"key"`
	// Value is the expected value that, when matching the key's actual value, makes this condition true.
	Value *string `json:"value" yaml:"value"`
}

// SucceededCondition, when matched, indicates that the input was successfully deployed.
type ClusterDeploymentTemplateSpecObservedCompletionSucceeded struct {
	// Key is a jsonPath expression pointing to the field to inspect on the templated object, eg: 'status.conditions[?(@.type=="Succeeded")].status'.
	Key *string `json:"key" yaml:"key"`
	// Value is the expected value that, when matching the key's actual value, makes this condition true.
	Value *string `json:"value" yaml:"value"`
}

type ClusterDeploymentTemplateSpecObservedMatches struct {
	// Input is a jsonPath to a value that is fulfilled before the templated object is reconciled.
	//
	// Usually a value in the .spec of the object
	Input *string `json:"input" yaml:"input"`
	// Output is a jsonPath to a value that is fulfilled after the templated object is reconciled.
	//
	// Usually a value in the .status of the object
	Output *string `json:"output" yaml:"output"`
}

type ClusterDeploymentTemplateSpecParams struct {
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner or Template does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Name of a parameter the template accepts from the Blueprint or Owner.
	Name *string `json:"name" yaml:"name"`
}

type ClusterImageTemplate interface {
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

// The jsii proxy struct for ClusterImageTemplate
type jsiiProxy_ClusterImageTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterImageTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterImageTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterImageTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterImageTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterImageTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterImageTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterImageTemplate" API object.
func NewClusterImageTemplate(scope constructs.Construct, id *string, props *ClusterImageTemplateProps) ClusterImageTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterImageTemplate{}

	_jsii_.Create(
		"cartorun.ClusterImageTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterImageTemplate" API object.
func NewClusterImageTemplate_Override(c ClusterImageTemplate, scope constructs.Construct, id *string, props *ClusterImageTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterImageTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterImageTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterImageTemplate_Manifest(props *ClusterImageTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterImageTemplate",
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
func ClusterImageTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterImageTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterImageTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterImageTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterImageTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterImageTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterImageTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterImageTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterImageTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterImageTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterImageTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterImageTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the image template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/template/#clusterimagetemplate
	Spec *ClusterImageTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the image template.
//
// More info: https://cartographer.sh/docs/latest/reference/template/#clusterimagetemplate
type ClusterImageTemplateSpec struct {
	// ImagePath is a path into the templated object's data that contains a valid image digest.
	//
	// This might be a URL or in some cases just a repository path and digest. The final spec for this field may change as we implement RFC-0016 https://github.com/vmware-tanzu/cartographer/blob/main/rfc/rfc-0016-validate-template-outputs.md ImagePath is specified in jsonpath format, eg: .status.artifact.image_digest
	ImagePath *string `json:"imagePath" yaml:"imagePath"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterImageTemplateSpecParams `json:"params" yaml:"params"`
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Template interface{} `json:"template" yaml:"template"`
	// Ytt defines a resource template written in `ytt` for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Ytt *string `json:"ytt" yaml:"ytt"`
}

type ClusterImageTemplateSpecParams struct {
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner or Template does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Name of a parameter the template accepts from the Blueprint or Owner.
	Name *string `json:"name" yaml:"name"`
}

type ClusterRunTemplate interface {
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

// The jsii proxy struct for ClusterRunTemplate
type jsiiProxy_ClusterRunTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterRunTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRunTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRunTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRunTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRunTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterRunTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterRunTemplate" API object.
func NewClusterRunTemplate(scope constructs.Construct, id *string, props *ClusterRunTemplateProps) ClusterRunTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterRunTemplate{}

	_jsii_.Create(
		"cartorun.ClusterRunTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterRunTemplate" API object.
func NewClusterRunTemplate_Override(c ClusterRunTemplate, scope constructs.Construct, id *string, props *ClusterRunTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterRunTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterRunTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterRunTemplate_Manifest(props *ClusterRunTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterRunTemplate",
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
func ClusterRunTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterRunTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterRunTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterRunTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterRunTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterRunTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterRunTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterRunTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterRunTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterRunTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterRunTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterRunTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the run template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/runnable/#clusterruntemplate
	Spec *ClusterRunTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the run template.
//
// More info: https://cartographer.sh/docs/latest/reference/runnable/#clusterruntemplate
type ClusterRunTemplateSpec struct {
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/
	Template interface{} `json:"template" yaml:"template"`
	// Outputs are a named list of jsonPaths that are used to gather results from the last successful object stamped by the template.
	//
	// E.g: 	my-output: .status.results[?(@.name=="IMAGE-DIGEST")].value
	Outputs *map[string]*string `json:"outputs" yaml:"outputs"`
}

type ClusterSourceTemplate interface {
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

// The jsii proxy struct for ClusterSourceTemplate
type jsiiProxy_ClusterSourceTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterSourceTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSourceTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSourceTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSourceTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSourceTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSourceTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterSourceTemplate" API object.
func NewClusterSourceTemplate(scope constructs.Construct, id *string, props *ClusterSourceTemplateProps) ClusterSourceTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterSourceTemplate{}

	_jsii_.Create(
		"cartorun.ClusterSourceTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterSourceTemplate" API object.
func NewClusterSourceTemplate_Override(c ClusterSourceTemplate, scope constructs.Construct, id *string, props *ClusterSourceTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterSourceTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterSourceTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterSourceTemplate_Manifest(props *ClusterSourceTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterSourceTemplate",
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
func ClusterSourceTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterSourceTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterSourceTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterSourceTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterSourceTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterSourceTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterSourceTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterSourceTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterSourceTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterSourceTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterSourceTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterSourceTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the source template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/template/#clustersourcetemplate
	Spec *ClusterSourceTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the source template.
//
// More info: https://cartographer.sh/docs/latest/reference/template/#clustersourcetemplate
type ClusterSourceTemplateSpec struct {
	// RevisionPath is a path into the templated object's data that contains a revision.
	//
	// The revision, along with the URL, represents the output of the Template. RevisionPath is specified in jsonpath format, eg: .status.artifact.revision
	RevisionPath *string `json:"revisionPath" yaml:"revisionPath"`
	// URLPath is a path into the templated object's data that contains a URL.
	//
	// The URL, along with the revision, represents the output of the Template. URLPath is specified in jsonpath format, eg: .status.artifact.url
	UrlPath *string `json:"urlPath" yaml:"urlPath"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterSourceTemplateSpecParams `json:"params" yaml:"params"`
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Template interface{} `json:"template" yaml:"template"`
	// Ytt defines a resource template written in `ytt` for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Ytt *string `json:"ytt" yaml:"ytt"`
}

type ClusterSourceTemplateSpecParams struct {
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner or Template does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Name of a parameter the template accepts from the Blueprint or Owner.
	Name *string `json:"name" yaml:"name"`
}

type ClusterSupplyChain interface {
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

// The jsii proxy struct for ClusterSupplyChain
type jsiiProxy_ClusterSupplyChain struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterSupplyChain) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSupplyChain) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSupplyChain) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSupplyChain) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSupplyChain) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterSupplyChain) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterSupplyChain" API object.
func NewClusterSupplyChain(scope constructs.Construct, id *string, props *ClusterSupplyChainProps) ClusterSupplyChain {
	_init_.Initialize()

	j := jsiiProxy_ClusterSupplyChain{}

	_jsii_.Create(
		"cartorun.ClusterSupplyChain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterSupplyChain" API object.
func NewClusterSupplyChain_Override(c ClusterSupplyChain, scope constructs.Construct, id *string, props *ClusterSupplyChainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterSupplyChain",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterSupplyChain".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterSupplyChain_Manifest(props *ClusterSupplyChainProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterSupplyChain",
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
func ClusterSupplyChain_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterSupplyChain",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterSupplyChain_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterSupplyChain",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterSupplyChain) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterSupplyChain) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterSupplyChain) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterSupplyChain) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterSupplyChain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterSupplyChain) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterSupplyChain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterSupplyChainProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the suppply chain.
	//
	// More info: https://cartographer.sh/docs/latest/reference/workload/#clustersupplychain
	Spec *ClusterSupplyChainSpec `json:"spec" yaml:"spec"`
}

// Spec describes the suppply chain.
//
// More info: https://cartographer.sh/docs/latest/reference/workload/#clustersupplychain
type ClusterSupplyChainSpec struct {
	// Resources that are responsible for bringing the application to a deliverable state.
	Resources *[]*ClusterSupplyChainSpecResources `json:"resources" yaml:"resources"`
	// Specifies the label key-value pairs used to select workloads See: https://cartographer.sh/docs/v0.1.0/architecture/#selectors.
	Selector *map[string]*string `json:"selector" yaml:"selector"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterSupplyChainSpecParams `json:"params" yaml:"params"`
	// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
	//
	// If not set, Cartographer will use serviceAccountName from supply chain.
	// If that is also not set, Cartographer will use the default service account in the workload's namespace.
	ServiceAccountRef *ClusterSupplyChainSpecServiceAccountRef `json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

type ClusterSupplyChainSpecParams struct {
	// Name of the parameter.
	//
	// Should match a template parameter name.
	Name *string `json:"name" yaml:"name"`
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Value of the parameter.
	//
	// If specified, owner properties are ignored.
	Value interface{} `json:"value" yaml:"value"`
}

type ClusterSupplyChainSpecResources struct {
	// Name of the resource.
	//
	// Used as a reference for inputs, as well as being the name presented in workload statuses to identify this resource.
	Name *string `json:"name" yaml:"name"`
	// TemplateRef identifies the template used to produce this resource.
	TemplateRef *ClusterSupplyChainSpecResourcesTemplateRef `json:"templateRef" yaml:"templateRef"`
	// Configs is a list of references to other 'config' resources in this list.
	//
	// A config resource has the kind ClusterConfigTemplate
	// In a template, configs can be consumed as: $(configs.<name>.config)$
	// If there is only one image, it can be consumed as: $(config)$
	Configs *[]*ClusterSupplyChainSpecResourcesConfigs `json:"configs" yaml:"configs"`
	// Images is a list of references to other 'image' resources in this list.
	//
	// An image resource has the kind ClusterImageTemplate
	// In a template, images can be consumed as: $(images.<name>.image)$
	// If there is only one image, it can be consumed as: $(image)$
	Images *[]*ClusterSupplyChainSpecResourcesImages `json:"images" yaml:"images"`
	// Params are a list of parameters to provide to the template in TemplateRef Template params do not have to be specified here, unless you want to force a particular value, or add a default value.
	//
	// Parameters are consumed in a template with the syntax: $(params.<name>)$
	Params *[]*ClusterSupplyChainSpecResourcesParams `json:"params" yaml:"params"`
	// Sources is a list of references to other 'source' resources in this list.
	//
	// A source resource has the kind ClusterSourceTemplate
	// In a template, sources can be consumed as: $(sources.<name>.url)$ and $(sources.<name>.revision)$
	// If there is only one source, it can be consumed as: $(source.url)$ and $(source.revision)$
	Sources *[]*ClusterSupplyChainSpecResourcesSources `json:"sources" yaml:"sources"`
}

type ClusterSupplyChainSpecResourcesConfigs struct {
	Name *string `json:"name" yaml:"name"`
	Resource *string `json:"resource" yaml:"resource"`
}

type ClusterSupplyChainSpecResourcesImages struct {
	Name *string `json:"name" yaml:"name"`
	Resource *string `json:"resource" yaml:"resource"`
}

type ClusterSupplyChainSpecResourcesParams struct {
	// Name of the parameter.
	//
	// Should match a template parameter name.
	Name *string `json:"name" yaml:"name"`
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Value of the parameter.
	//
	// If specified, owner properties are ignored.
	Value interface{} `json:"value" yaml:"value"`
}

type ClusterSupplyChainSpecResourcesSources struct {
	Name *string `json:"name" yaml:"name"`
	Resource *string `json:"resource" yaml:"resource"`
}

// TemplateRef identifies the template used to produce this resource.
type ClusterSupplyChainSpecResourcesTemplateRef struct {
	// Kind of the template to apply.
	Kind ClusterSupplyChainSpecResourcesTemplateRefKind `json:"kind" yaml:"kind"`
	// Name of the template to apply.
	Name *string `json:"name" yaml:"name"`
}

// Kind of the template to apply.
type ClusterSupplyChainSpecResourcesTemplateRefKind string

const (
	ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_SOURCE_TEMPLATE ClusterSupplyChainSpecResourcesTemplateRefKind = "CLUSTER_SOURCE_TEMPLATE"
	ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_IMAGE_TEMPLATE ClusterSupplyChainSpecResourcesTemplateRefKind = "CLUSTER_IMAGE_TEMPLATE"
	ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_TEMPLATE ClusterSupplyChainSpecResourcesTemplateRefKind = "CLUSTER_TEMPLATE"
	ClusterSupplyChainSpecResourcesTemplateRefKind_CLUSTER_CONFIG_TEMPLATE ClusterSupplyChainSpecResourcesTemplateRefKind = "CLUSTER_CONFIG_TEMPLATE"
)

// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
//
// If not set, Cartographer will use serviceAccountName from supply chain.
// If that is also not set, Cartographer will use the default service account in the workload's namespace.
type ClusterSupplyChainSpecServiceAccountRef struct {
	// Name of the service account being referred to.
	Name *string `json:"name" yaml:"name"`
	// Namespace of the service account being referred to if omitted, the Owner's namespace is used.
	Namespace *string `json:"namespace" yaml:"namespace"`
}

type ClusterTemplate interface {
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

// The jsii proxy struct for ClusterTemplate
type jsiiProxy_ClusterTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ClusterTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "ClusterTemplate" API object.
func NewClusterTemplate(scope constructs.Construct, id *string, props *ClusterTemplateProps) ClusterTemplate {
	_init_.Initialize()

	j := jsiiProxy_ClusterTemplate{}

	_jsii_.Create(
		"cartorun.ClusterTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ClusterTemplate" API object.
func NewClusterTemplate_Override(c ClusterTemplate, scope constructs.Construct, id *string, props *ClusterTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.ClusterTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

// Renders a Kubernetes manifest for "ClusterTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ClusterTemplate_Manifest(props *ClusterTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.ClusterTemplate",
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
func ClusterTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.ClusterTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ClusterTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.ClusterTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_ClusterTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_ClusterTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterTemplate) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_ClusterTemplate) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
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
func (c *jsiiProxy_ClusterTemplate) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (c *jsiiProxy_ClusterTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_ClusterTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ClusterTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the template.
	//
	// More info: https://cartographer.sh/docs/latest/reference/template/#clustertemplate
	Spec *ClusterTemplateSpec `json:"spec" yaml:"spec"`
}

// Spec describes the template.
//
// More info: https://cartographer.sh/docs/latest/reference/template/#clustertemplate
type ClusterTemplateSpec struct {
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*ClusterTemplateSpecParams `json:"params" yaml:"params"`
	// Template defines a resource template for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Template interface{} `json:"template" yaml:"template"`
	// Ytt defines a resource template written in `ytt` for a Kubernetes Resource or Custom Resource which is applied to the server each time the blueprint is applied.
	//
	// Templates support simple value interpolation using the $()$ marker format. For more information, see: https://cartographer.sh/docs/latest/templating/ You cannot define both Template and Ytt at the same time.
	Ytt *string `json:"ytt" yaml:"ytt"`
}

type ClusterTemplateSpecParams struct {
	// DefaultValue of the parameter.
	//
	// Causes the parameter to be optional; If the Owner or Template does not specify this parameter, this value is used.
	Default interface{} `json:"default" yaml:"default"`
	// Name of a parameter the template accepts from the Blueprint or Owner.
	Name *string `json:"name" yaml:"name"`
}

type Deliverable interface {
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

// The jsii proxy struct for Deliverable
type jsiiProxy_Deliverable struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Deliverable) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deliverable) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deliverable) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deliverable) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deliverable) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Deliverable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "Deliverable" API object.
func NewDeliverable(scope constructs.Construct, id *string, props *DeliverableProps) Deliverable {
	_init_.Initialize()

	j := jsiiProxy_Deliverable{}

	_jsii_.Create(
		"cartorun.Deliverable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Deliverable" API object.
func NewDeliverable_Override(d Deliverable, scope constructs.Construct, id *string, props *DeliverableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.Deliverable",
		[]interface{}{scope, id, props},
		d,
	)
}

// Renders a Kubernetes manifest for "Deliverable".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Deliverable_Manifest(props *DeliverableProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.Deliverable",
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
func Deliverable_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.Deliverable",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Deliverable_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.Deliverable",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (d *jsiiProxy_Deliverable) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (d *jsiiProxy_Deliverable) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
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
func (d *jsiiProxy_Deliverable) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (d *jsiiProxy_Deliverable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
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
func (d *jsiiProxy_Deliverable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (d *jsiiProxy_Deliverable) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_Deliverable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DeliverableProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the deliverable.
	//
	// More info: https://cartographer.sh/docs/latest/reference/workload/#deliverable
	Spec *DeliverableSpec `json:"spec" yaml:"spec"`
}

// Spec describes the deliverable.
//
// More info: https://cartographer.sh/docs/latest/reference/workload/#deliverable
type DeliverableSpec struct {
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*DeliverableSpecParams `json:"params" yaml:"params"`
	// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
	//
	// If not set, Cartographer will use serviceAccountName from supply chain.
	// If that is also not set, Cartographer will use the default service account in the workload's namespace.
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
	// The location of the source code for the workload.
	//
	// Specify one of `spec.source` or `spec.image`
	Source *DeliverableSpecSource `json:"source" yaml:"source"`
}

type DeliverableSpecParams struct {
	// Name of the parameter.
	//
	// Should match a blueprint or template parameter name.
	Name *string `json:"name" yaml:"name"`
	// Value of the parameter.
	Value interface{} `json:"value" yaml:"value"`
}

// The location of the source code for the workload.
//
// Specify one of `spec.source` or `spec.image`
type DeliverableSpecSource struct {
	// Source code location in a git repository.
	Git *DeliverableSpecSourceGit `json:"git" yaml:"git"`
	// OCI Image in a repository, containing the source code to be used throughout the supply chain.
	Image *string `json:"image" yaml:"image"`
	// Subpath inside the Git repository or Image to treat as the root of the application.
	//
	// Defaults to the root if left empty.
	SubPath *string `json:"subPath" yaml:"subPath"`
}

// Source code location in a git repository.
type DeliverableSpecSourceGit struct {
	Ref *DeliverableSpecSourceGitRef `json:"ref" yaml:"ref"`
	Url *string `json:"url" yaml:"url"`
}

type DeliverableSpecSourceGitRef struct {
	Branch *string `json:"branch" yaml:"branch"`
	Commit *string `json:"commit" yaml:"commit"`
	Tag *string `json:"tag" yaml:"tag"`
}

type Runnable interface {
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

// The jsii proxy struct for Runnable
type jsiiProxy_Runnable struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Runnable) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runnable) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runnable) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runnable) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runnable) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runnable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "Runnable" API object.
func NewRunnable(scope constructs.Construct, id *string, props *RunnableProps) Runnable {
	_init_.Initialize()

	j := jsiiProxy_Runnable{}

	_jsii_.Create(
		"cartorun.Runnable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Runnable" API object.
func NewRunnable_Override(r Runnable, scope constructs.Construct, id *string, props *RunnableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.Runnable",
		[]interface{}{scope, id, props},
		r,
	)
}

// Renders a Kubernetes manifest for "Runnable".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Runnable_Manifest(props *RunnableProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.Runnable",
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
func Runnable_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.Runnable",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Runnable_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.Runnable",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (r *jsiiProxy_Runnable) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (r *jsiiProxy_Runnable) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_Runnable) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (r *jsiiProxy_Runnable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_Runnable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (r *jsiiProxy_Runnable) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_Runnable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RunnableProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the runnable.
	//
	// More info: https://cartographer.sh/docs/latest/reference/runnable/#runnable
	Spec *RunnableSpec `json:"spec" yaml:"spec"`
}

// Spec describes the runnable.
//
// More info: https://cartographer.sh/docs/latest/reference/runnable/#runnable
type RunnableSpec struct {
	// RunTemplateRef identifies the run template used to produce resources for this runnable.
	RunTemplateRef *RunnableSpecRunTemplateRef `json:"runTemplateRef" yaml:"runTemplateRef"`
	// Inputs are key/values providing inputs to the templated object created for this runnable.
	//
	// Reference inputs in the template using the jsonPath: $(runnable.spec.inputs.<key>)$
	Inputs *map[string]interface{} `json:"inputs" yaml:"inputs"`
	// RetentionPolicy specifies how many successful and failed runs should be retained.
	//
	// Runs older than this (ordered by creation time) will be deleted.
	RetentionPolicy *RunnableSpecRetentionPolicy `json:"retentionPolicy" yaml:"retentionPolicy"`
	// Selector refers to an additional object that the template can refer to using: $(selected)$.
	Selector *RunnableSpecSelector `json:"selector" yaml:"selector"`
	// ServiceAccountName refers to the Service account with permissions to create resources submitted by the ClusterRunTemplate.
	//
	// If not set, Cartographer will use the default service account in the runnable's namespace.
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
}

// RetentionPolicy specifies how many successful and failed runs should be retained.
//
// Runs older than this (ordered by creation time) will be deleted.
type RunnableSpecRetentionPolicy struct {
	// MaxFailedRuns is the number of failed runs to retain.
	MaxFailedRuns *float64 `json:"maxFailedRuns" yaml:"maxFailedRuns"`
	// MaxSuccessfulRuns is the number of successful runs to retain.
	MaxSuccessfulRuns *float64 `json:"maxSuccessfulRuns" yaml:"maxSuccessfulRuns"`
}

// RunTemplateRef identifies the run template used to produce resources for this runnable.
type RunnableSpecRunTemplateRef struct {
	Name *string `json:"name" yaml:"name"`
	Kind *string `json:"kind" yaml:"kind"`
}

// Selector refers to an additional object that the template can refer to using: $(selected)$.
type RunnableSpecSelector struct {
	// MatchingLabels must match on a single target object, making the object available in the template as $(selected)$.
	MatchingLabels *map[string]*string `json:"matchingLabels" yaml:"matchingLabels"`
	// Resource is the GVK that must match the selected object.
	Resource *RunnableSpecSelectorResource `json:"resource" yaml:"resource"`
}

// Resource is the GVK that must match the selected object.
type RunnableSpecSelectorResource struct {
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	Kind *string `json:"kind" yaml:"kind"`
}

type Workload interface {
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

// The jsii proxy struct for Workload
type jsiiProxy_Workload struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Workload) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "Workload" API object.
func NewWorkload(scope constructs.Construct, id *string, props *WorkloadProps) Workload {
	_init_.Initialize()

	j := jsiiProxy_Workload{}

	_jsii_.Create(
		"cartorun.Workload",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Workload" API object.
func NewWorkload_Override(w Workload, scope constructs.Construct, id *string, props *WorkloadProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cartorun.Workload",
		[]interface{}{scope, id, props},
		w,
	)
}

// Renders a Kubernetes manifest for "Workload".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Workload_Manifest(props *WorkloadProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cartorun.Workload",
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
func Workload_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"cartorun.Workload",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Workload_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"cartorun.Workload",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (w *jsiiProxy_Workload) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (w *jsiiProxy_Workload) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
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
func (w *jsiiProxy_Workload) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (w *jsiiProxy_Workload) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		w,
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
func (w *jsiiProxy_Workload) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (w *jsiiProxy_Workload) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_Workload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WorkloadProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata" yaml:"metadata"`
	// Spec describes the workload.
	//
	// More info: https://cartographer.sh/docs/latest/reference/workload/#workload
	Spec *WorkloadSpec `json:"spec" yaml:"spec"`
}

// Spec describes the workload.
//
// More info: https://cartographer.sh/docs/latest/reference/workload/#workload
type WorkloadSpec struct {
	// Environment variables to be passed to the main container running the application.
	//
	// See https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
	Env *[]*WorkloadSpecEnv `json:"env" yaml:"env"`
	// Build configuration, for the build resources in the supply chain.
	Build *WorkloadSpecBuild `json:"build" yaml:"build"`
	// Image refers to a pre-built image in a registry.
	//
	// It is an alternative to specifying the location of source code for the workload. Specify one of `spec.source` or `spec.image`.
	Image *string `json:"image" yaml:"image"`
	// Additional parameters.
	//
	// See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy
	Params *[]*WorkloadSpecParams `json:"params" yaml:"params"`
	// Resource constraints for the application.
	//
	// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources *WorkloadSpecResources `json:"resources" yaml:"resources"`
	// ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain.
	//
	// If not set, Cartographer will use serviceAccountName from supply chain.
	// If that is also not set, Cartographer will use the default service account in the workload's namespace.
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
	// ServiceClaims to be bound through ServiceBindings.
	ServiceClaims *[]*WorkloadSpecServiceClaims `json:"serviceClaims" yaml:"serviceClaims"`
	// The location of the source code for the workload.
	//
	// Specify one of `spec.source` or `spec.image`
	Source *WorkloadSpecSource `json:"source" yaml:"source"`
}

// EnvVar represents an environment variable present in a Container.
type WorkloadSpecEnv struct {
	// Name of the environment variable.
	//
	// Must be a C_IDENTIFIER.
	Name *string `json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Value *string `json:"value" yaml:"value"`
	// Source for the environment variable's value.
	//
	// Cannot be used if value is not empty.
	ValueFrom *WorkloadSpecEnvValueFrom `json:"valueFrom" yaml:"valueFrom"`
}

// Source for the environment variable's value.
//
// Cannot be used if value is not empty.
type WorkloadSpecEnvValueFrom struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *WorkloadSpecEnvValueFromConfigMapKeyRef `json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	FieldRef *WorkloadSpecEnvValueFromFieldRef `json:"fieldRef" yaml:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *WorkloadSpecEnvValueFromResourceFieldRef `json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// Selects a key of a secret in the pod's namespace.
	SecretKeyRef *WorkloadSpecEnvValueFromSecretKeyRef `json:"secretKeyRef" yaml:"secretKeyRef"`
}

// Selects a key of a ConfigMap.
type WorkloadSpecEnvValueFromConfigMapKeyRef struct {
	// The key to select.
	Key *string `json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
type WorkloadSpecEnvValueFromFieldRef struct {
	// Path of the field to select in the specified API version.
	FieldPath *string `json:"fieldPath" yaml:"fieldPath"`
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
}

// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
type WorkloadSpecEnvValueFromResourceFieldRef struct {
	// Required: resource to select.
	Resource *string `json:"resource" yaml:"resource"`
	// Container name: required for volumes, optional for env vars.
	ContainerName *string `json:"containerName" yaml:"containerName"`
	// Specifies the output format of the exposed resources, defaults to "1".
	Divisor WorkloadSpecEnvValueFromResourceFieldRefDivisor `json:"divisor" yaml:"divisor"`
}

// Specifies the output format of the exposed resources, defaults to "1".
type WorkloadSpecEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkloadSpecEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkloadSpecEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkloadSpecEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkloadSpecEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkloadSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	var returns WorkloadSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkloadSpecEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkloadSpecEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	var returns WorkloadSpecEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Selects a key of a secret in the pod's namespace.
type WorkloadSpecEnvValueFromSecretKeyRef struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	Key *string `json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// Build configuration, for the build resources in the supply chain.
type WorkloadSpecBuild struct {
	// Env is an array of environment variables to propagate to build resources in the supply chain.
	//
	// See https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
	Env *[]*WorkloadSpecBuildEnv `json:"env" yaml:"env"`
}

// EnvVar represents an environment variable present in a Container.
type WorkloadSpecBuildEnv struct {
	// Name of the environment variable.
	//
	// Must be a C_IDENTIFIER.
	Name *string `json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Value *string `json:"value" yaml:"value"`
	// Source for the environment variable's value.
	//
	// Cannot be used if value is not empty.
	ValueFrom *WorkloadSpecBuildEnvValueFrom `json:"valueFrom" yaml:"valueFrom"`
}

// Source for the environment variable's value.
//
// Cannot be used if value is not empty.
type WorkloadSpecBuildEnvValueFrom struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *WorkloadSpecBuildEnvValueFromConfigMapKeyRef `json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	FieldRef *WorkloadSpecBuildEnvValueFromFieldRef `json:"fieldRef" yaml:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *WorkloadSpecBuildEnvValueFromResourceFieldRef `json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// Selects a key of a secret in the pod's namespace.
	SecretKeyRef *WorkloadSpecBuildEnvValueFromSecretKeyRef `json:"secretKeyRef" yaml:"secretKeyRef"`
}

// Selects a key of a ConfigMap.
type WorkloadSpecBuildEnvValueFromConfigMapKeyRef struct {
	// The key to select.
	Key *string `json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
type WorkloadSpecBuildEnvValueFromFieldRef struct {
	// Path of the field to select in the specified API version.
	FieldPath *string `json:"fieldPath" yaml:"fieldPath"`
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
}

// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
type WorkloadSpecBuildEnvValueFromResourceFieldRef struct {
	// Required: resource to select.
	Resource *string `json:"resource" yaml:"resource"`
	// Container name: required for volumes, optional for env vars.
	ContainerName *string `json:"containerName" yaml:"containerName"`
	// Specifies the output format of the exposed resources, defaults to "1".
	Divisor WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor `json:"divisor" yaml:"divisor"`
}

// Specifies the output format of the exposed resources, defaults to "1".
type WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor
type jsiiProxy_WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor_FromNumber(value *float64) WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	var returns WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor_FromString(value *string) WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor {
	_init_.Initialize()

	var returns WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecBuildEnvValueFromResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Selects a key of a secret in the pod's namespace.
type WorkloadSpecBuildEnvValueFromSecretKeyRef struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	Key *string `json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	Optional *bool `json:"optional" yaml:"optional"`
}

type WorkloadSpecParams struct {
	// Name of the parameter.
	//
	// Should match a blueprint or template parameter name.
	Name *string `json:"name" yaml:"name"`
	// Value of the parameter.
	Value interface{} `json:"value" yaml:"value"`
}

// Resource constraints for the application.
//
// See https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
type WorkloadSpecResources struct {
	// Limits describes the maximum amount of compute resources allowed.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]WorkloadSpecResourcesLimits `json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]WorkloadSpecResourcesRequests `json:"requests" yaml:"requests"`
}

type WorkloadSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for WorkloadSpecResourcesLimits
type jsiiProxy_WorkloadSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkloadSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkloadSpecResourcesLimits_FromNumber(value *float64) WorkloadSpecResourcesLimits {
	_init_.Initialize()

	var returns WorkloadSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkloadSpecResourcesLimits_FromString(value *string) WorkloadSpecResourcesLimits {
	_init_.Initialize()

	var returns WorkloadSpecResourcesLimits

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

type WorkloadSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for WorkloadSpecResourcesRequests
type jsiiProxy_WorkloadSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkloadSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func WorkloadSpecResourcesRequests_FromNumber(value *float64) WorkloadSpecResourcesRequests {
	_init_.Initialize()

	var returns WorkloadSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WorkloadSpecResourcesRequests_FromString(value *string) WorkloadSpecResourcesRequests {
	_init_.Initialize()

	var returns WorkloadSpecResourcesRequests

	_jsii_.StaticInvoke(
		"cartorun.WorkloadSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

type WorkloadSpecServiceClaims struct {
	Name *string `json:"name" yaml:"name"`
	Ref *WorkloadSpecServiceClaimsRef `json:"ref" yaml:"ref"`
}

type WorkloadSpecServiceClaimsRef struct {
	ApiVersion *string `json:"apiVersion" yaml:"apiVersion"`
	Kind *string `json:"kind" yaml:"kind"`
	Name *string `json:"name" yaml:"name"`
}

// The location of the source code for the workload.
//
// Specify one of `spec.source` or `spec.image`
type WorkloadSpecSource struct {
	// Source code location in a git repository.
	Git *WorkloadSpecSourceGit `json:"git" yaml:"git"`
	// OCI Image in a repository, containing the source code to be used throughout the supply chain.
	Image *string `json:"image" yaml:"image"`
	// Subpath inside the Git repository or Image to treat as the root of the application.
	//
	// Defaults to the root if left empty.
	SubPath *string `json:"subPath" yaml:"subPath"`
}

// Source code location in a git repository.
type WorkloadSpecSourceGit struct {
	Ref *WorkloadSpecSourceGitRef `json:"ref" yaml:"ref"`
	Url *string `json:"url" yaml:"url"`
}

type WorkloadSpecSourceGitRef struct {
	Branch *string `json:"branch" yaml:"branch"`
	Commit *string `json:"commit" yaml:"commit"`
	Tag *string `json:"tag" yaml:"tag"`
}

