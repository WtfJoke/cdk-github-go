// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/WtfJoke/cdk-github-go/cdkgithub/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/WtfJoke/cdk-github-go/cdkgithub/internal"
)

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

	if err := validateNewGitHubResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateGitHubResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

