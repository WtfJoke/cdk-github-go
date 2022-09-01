// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

func (j *jsiiProxy_IGitHubRepository)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_IGitHubRepository)SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

