package cdkgithub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

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
	// Default: - returns all the content stored in the Secrets Manager secret.
	//
	// Experimental.
	SourceSecretJsonField *string `field:"optional" json:"sourceSecretJsonField" yaml:"sourceSecretJsonField"`
}

