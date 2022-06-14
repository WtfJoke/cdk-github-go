// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/WtfJoke/cdk-github-go/cdkgithub/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
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
	// The GitHub repository name.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The GitHub repository owner.
	// Experimental.
	RepositoryOwner *string `field:"required" json:"repositoryOwner" yaml:"repositoryOwner"`
	// The GitHub secret name to be stored.
	// Experimental.
	RepositorySecretName *string `field:"required" json:"repositorySecretName" yaml:"repositorySecretName"`
	// This AWS secret value will be stored in GitHub as a secret (under the name of repositorySecretName).
	// Experimental.
	SourceSecret awssecretsmanager.ISecret `field:"required" json:"sourceSecret" yaml:"sourceSecret"`
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
	// The GitHub repository name.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The GitHub repository owner.
	// Experimental.
	RepositoryOwner *string `field:"required" json:"repositoryOwner" yaml:"repositoryOwner"`
	// The GitHub secret name to be stored.
	// Experimental.
	RepositorySecretName *string `field:"required" json:"repositorySecretName" yaml:"repositorySecretName"`
	// This AWS secret value will be stored in GitHub as a secret (under the name of repositorySecretName).
	// Experimental.
	SourceSecret awssecretsmanager.ISecret `field:"required" json:"sourceSecret" yaml:"sourceSecret"`
}

