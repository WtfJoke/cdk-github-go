package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/WtfJoke/cdk-github-go/cdkgithub/jsii"

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

	if err := validateNewActionEnvironmentSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateActionEnvironmentSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

