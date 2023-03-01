// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

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

