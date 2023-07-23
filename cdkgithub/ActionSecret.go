package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/WtfJoke/cdk-github-go/cdkgithub/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/WtfJoke/cdk-github-go/cdkgithub/internal"
)

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

	if err := validateNewActionSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateActionSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

