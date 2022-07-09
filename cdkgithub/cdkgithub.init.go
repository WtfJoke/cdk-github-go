package cdkgithub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-github.ActionEnvironmentSecret",
		reflect.TypeOf((*ActionEnvironmentSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActionEnvironmentSecret{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-github.ActionEnvironmentSecretProps",
		reflect.TypeOf((*ActionEnvironmentSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-github.ActionSecret",
		reflect.TypeOf((*ActionSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ActionSecret{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-github.ActionSecretProps",
		reflect.TypeOf((*ActionSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-github.GitHubResource",
		reflect.TypeOf((*GitHubResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-github.GitHubResourceProps",
		reflect.TypeOf((*GitHubResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-github.IGitHubRepository",
		reflect.TypeOf((*IGitHubRepository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
		},
		func() interface{} {
			return &jsiiProxy_IGitHubRepository{}
		},
	)
}
