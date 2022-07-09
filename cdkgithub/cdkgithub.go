// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/WtfJoke/cdk-github-go/cdkgithub/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/WtfJoke/cdk-github-go/cdkgithub/internal"
)

// Experimental.
type ActionEnvironmentSecret interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActionEnvironmentSecret
type jsiiProxy_ActionEnvironmentSecret struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ActionEnvironmentSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewActionEnvironmentSecret(scope constructs.Construct, id *string, props *ActionEnvironmentSecretProps) ActionEnvironmentSecret {
	_init_.Initialize()

	j := jsiiProxy_ActionEnvironmentSecret{}

	_jsii_.Create(
		"cdk-github.ActionEnvironmentSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewActionEnvironmentSecret_Override(a ActionEnvironmentSecret, scope constructs.Construct, id *string, props *ActionEnvironmentSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-github.ActionEnvironmentSecret",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActionEnvironmentSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-github.ActionEnvironmentSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionEnvironmentSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ActionEnvironmentSecretProps struct {
	// The GithHub environment name which the secret should be stored in.
	// Experimental.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The AWS secret in which the OAuth GitHub (personal) access token is stored.
	// Experimental.
	GithubTokenSecret awssecretsmanager.ISecret `field:"required" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// The GitHub repository information (owner and name).
	// Experimental.
	Repository IGitHubRepository `field:"required" json:"repository" yaml:"repository"`
	// The GitHub secret name to be stored.
	// Experimental.
	RepositorySecretName *string `field:"required" json:"repositorySecretName" yaml:"repositorySecretName"`
	// This AWS secret value will be stored in GitHub as a secret (under the name of repositorySecretName).
	// Experimental.
	SourceSecret awssecretsmanager.ISecret `field:"required" json:"sourceSecret" yaml:"sourceSecret"`
	// The key of a JSON field to retrieve in sourceSecret.
	//
	// This can only be used if the secret stores a JSON object.
	// Experimental.
	SourceSecretJsonField *string `field:"optional" json:"sourceSecretJsonField" yaml:"sourceSecretJsonField"`
}

// Experimental.
type ActionSecret interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActionSecret
type jsiiProxy_ActionSecret struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ActionSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewActionSecret(scope constructs.Construct, id *string, props *ActionSecretProps) ActionSecret {
	_init_.Initialize()

	j := jsiiProxy_ActionSecret{}

	_jsii_.Create(
		"cdk-github.ActionSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewActionSecret_Override(a ActionSecret, scope constructs.Construct, id *string, props *ActionSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-github.ActionSecret",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActionSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-github.ActionSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ActionSecretProps struct {
	// The AWS secret in which the OAuth GitHub (personal) access token is stored.
	// Experimental.
	GithubTokenSecret awssecretsmanager.ISecret `field:"required" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// The GitHub repository information (owner and name).
	// Experimental.
	Repository IGitHubRepository `field:"required" json:"repository" yaml:"repository"`
	// The GitHub secret name to be stored.
	// Experimental.
	RepositorySecretName *string `field:"required" json:"repositorySecretName" yaml:"repositorySecretName"`
	// This AWS secret value will be stored in GitHub as a secret (under the name of repositorySecretName).
	// Experimental.
	SourceSecret awssecretsmanager.ISecret `field:"required" json:"sourceSecret" yaml:"sourceSecret"`
	// The key of a JSON field to retrieve in sourceSecret.
	//
	// This can only be used if the secret stores a JSON object.
	// Experimental.
	SourceSecretJsonField *string `field:"optional" json:"sourceSecretJsonField" yaml:"sourceSecretJsonField"`
}

// Experimental.
type GitHubResource interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GitHubResource
type jsiiProxy_GitHubResource struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_GitHubResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubResource(scope constructs.Construct, id *string, props *GitHubResourceProps) GitHubResource {
	_init_.Initialize()

	j := jsiiProxy_GitHubResource{}

	_jsii_.Create(
		"cdk-github.GitHubResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubResource_Override(g GitHubResource, scope constructs.Construct, id *string, props *GitHubResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-github.GitHubResource",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func GitHubResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-github.GitHubResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type GitHubResourceProps struct {
	// The GitHub api endpoint url for creating resources in format: `POST /repos/OWNER/REPO/issues`.
	//
	// This is called when the GitHubResource is created.
	//
	// Example:
	// ```
	// const createRequestEndpoint = 'POST /repos/octocat/Hello-World/issues'
	// ```.
	// Experimental.
	CreateRequestEndpoint *string `field:"required" json:"createRequestEndpoint" yaml:"createRequestEndpoint"`
	// The GitHub api endpoint url to delete this resource in format: `POST /repos/OWNER/REPO/issues`.
	//
	// This is called when the GitHubResource is deleted/destroyed.
	//
	// Example:
	// ```
	// const deleteRequestEndpoint = 'PATCH repos/octocat/Hello-World/issues/1'
	// ```
	// If you want to use the  @see {@link GitHubResourceProps#createRequestResultParameter}, you can use the following syntax (assuming you have set createRequestResultParameter to `"number"`):
	// ```
	// const deleteRequestEndpoint = 'PATCH repos/octocat/Hello-World/:number'
	// ```.
	// Experimental.
	DeleteRequestEndpoint *string `field:"required" json:"deleteRequestEndpoint" yaml:"deleteRequestEndpoint"`
	// The AWS secret in which the OAuth GitHub (personal) access token is stored.
	// Experimental.
	GithubTokenSecret awssecretsmanager.ISecret `field:"required" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// The GitHub api request payload for creating resources. This is a JSON parseable string.
	//
	// Used for  @see {@link GitHubResourceProps#createRequestEndpoint}.
	//
	// Example:
	// ```
	// const createRequestPayload = JSON.stringify({ title: 'Found a bug', body: "I'm having a problem with this.", assignees: ['octocat'], milestone: 1, labels: ['bug'] })
	// ```.
	// Experimental.
	CreateRequestPayload *string `field:"optional" json:"createRequestPayload" yaml:"createRequestPayload"`
	// Used to extract a value from the result of the createRequest(Endpoint) to be used in update/deleteRequests.
	//
	// Example: `"number"` (for the issue number)
	//
	// When this parameter is set and can be extracted from the result, the extracted value will be used for the PhyscialResourceId of the CustomResource.
	// Changing the parameter once the stack is deployed is not supported.
	// Experimental.
	CreateRequestResultParameter *string `field:"optional" json:"createRequestResultParameter" yaml:"createRequestResultParameter"`
	// The GitHub api request payload to delete this resource. This is a JSON parseable string.
	//
	// Used for  @see {@link GitHubResourceProps#deleteRequestEndpoint}.
	//
	// Example:
	// ```
	// const deleteRequestPayload = JSON.stringify({ state: 'closed' })
	// ```.
	// Experimental.
	DeleteRequestPayload *string `field:"optional" json:"deleteRequestPayload" yaml:"deleteRequestPayload"`
	// The GitHub api endpoint url to update this resource in format: `POST /repos/OWNER/REPO/issues`.
	//
	// This is called when the GitHubResource is updated.
	//
	// In most of the cases you want to either omit this or use the same value as createRequestEndpoint.
	//
	// Example:
	// ```
	// const updateRequestEndpoint = 'PATCH repos/octocat/Hello-World/issues/1'
	// ```
	// If you want to use the  @see {@link GitHubResourceProps#createRequestResultParameter}, you can use the following syntax (assuming you have set createRequestResultParameter to `"number"`):
	// ```
	// const updateRequestEndpoint = 'PATCH repos/octocat/Hello-World/:number'
	// ```.
	// Experimental.
	UpdateRequestEndpoint *string `field:"optional" json:"updateRequestEndpoint" yaml:"updateRequestEndpoint"`
	// The GitHub api request payload to update this resources. This is a JSON parseable string.
	//
	// Used for  @see {@link GitHubResourceProps#createRequestEndpoint}.
	//
	// Example:
	// ```
	// const updateRequestPayload = JSON.stringify({ title: 'Found a bug', body: "I'm having a problem with this.", assignees: ['octocat'], milestone: 1, state: 'open', labels: ['bug'] })
	// ```.
	// Experimental.
	UpdateRequestPayload *string `field:"optional" json:"updateRequestPayload" yaml:"updateRequestPayload"`
	// The response body of the last GitHub api request will be written to this ssm parameter.
	// Experimental.
	WriteResponseToSSMParameter awsssm.IParameter `field:"optional" json:"writeResponseToSSMParameter" yaml:"writeResponseToSSMParameter"`
}

// Experimental.
type IGitHubRepository interface {
	// The GitHub repository name.
	// Experimental.
	Name() *string
	// Experimental.
	SetName(n *string)
	// The GitHub repository owner.
	// Experimental.
	Owner() *string
	// Experimental.
	SetOwner(o *string)
}

// The jsii proxy for IGitHubRepository
type jsiiProxy_IGitHubRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_IGitHubRepository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepository) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IGitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepository) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

