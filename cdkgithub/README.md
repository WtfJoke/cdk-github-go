[![npm version](https://badge.fury.io/js/cdk-github.svg)](https://badge.fury.io/js/cdk-github)
[![PyPI version](https://badge.fury.io/py/cdk-github.svg)](https://badge.fury.io/py/cdk-github)
[![NuGet version](https://badge.fury.io/nu/cdkgithub.svg)](https://badge.fury.io/nu/cdkgithub)
[![Maven Central](https://maven-badges.herokuapp.com/maven-central/io.github.wtfjoke/cdk-github/badge.svg)](https://maven-badges.herokuapp.com/maven-central/io.github.wtfjoke/cdk-github/)
[![release](https://github.com/wtfjoke/cdk-github/actions/workflows/release.yml/badge.svg)](https://github.com/wtfjoke/cdk-github/actions/workflows/release.yml)
![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge) [![View on Construct Hub](https://constructs.dev/badge?package=cdk-github)](https://constructs.dev/packages/cdk-github)

# CDK-GitHub

GitHub Constructs for use in [AWS CDK](https://aws.amazon.com/cdk/) .

This project aims to make GitHub's API accessible through CDK with various helper constructs to create resources in GitHub.
The target is to replicate most of the functionality of the official [Terraform GitHub Provider](https://registry.terraform.io/providers/integrations/github/latest/docs).

Internally [AWS CloudFormation custom resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) and [octokit](https://github.com/octokit/core.js) are used to manage GitHub resources (such as Secrets).

# 🔧 Installation

JavaScript/TypeScript:
`npm install cdk-github`

Python:
`pip install cdk-github`

Java

<details>
  <summary>Maven:</summary>

```xml
<dependency>
  <groupId>io.github.wtfjoke</groupId>
  <artifactId>cdk-github</artifactId>
  <version>VERSION</version>
</dependency>
```

</details>
<details>
  <summary>Gradle:</summary>

`implementation 'io.github.wtfjoke:cdk-github:VERSION'`

</details>
<details>
  <summary>Gradle (Kotlin):</summary>

`implementation("io.github.wtfjoke:cdk-github:VERSION")`

</details>

C#
See https://www.nuget.org/packages/CdkGithub

# 📚 Constructs

This library provides the following constructs:

* [ActionEnvironmentSecret](API.md#actionenvironmentsecret-) - Creates a [GitHub Action environment secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-an-environment) from a given AWS Secrets Manager secret.
* [ActionSecret](API.md#actionsecret-) - Creates a [GitHub Action (repository) secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-a-repository) from a given AWS Secrets Manager secret.
* [GitHubResource](API.md#githubresource-) - Creates an arbitrary GitHub resource. When no suitable construct fits your needs, this construct can be used to create most GitHub resources. It is an L1 construct.

# 🔓 Authentication

Currently the constructs only support authentication via a [GitHub Personal Access Token](https://github.com/settings/tokens/new). The token needs to be a stored in a AWS SecretsManager Secret and passed to the construct as parameter.

# 👩‍🏫 Examples

The API documentation and examples in different languages are available on [Construct Hub](https://constructs.dev/packages/cdk-github).
All (typescript) examples can be found in the folder [examples](src/examples/).

## ActionSecret

```go
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { ActionSecret } from 'cdk-github';

export class ActionSecretStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const githubTokenSecret = Secret.fromSecretNameV2(this, 'ghSecret', 'GITHUB_TOKEN');
    const sourceSecret = Secret.fromSecretNameV2(this, 'secretToStoreInGitHub', 'testcdkgithub');

    new ActionSecret(this, 'GitHubActionSecret', {
      githubTokenSecret,
      repository: { name: 'cdk-github', owner: 'wtfjoke' },
      repositorySecretName: 'A_RANDOM_GITHUB_SECRET',
      sourceSecret,
    });
  }
}
```

## ActionEnvironmentSecret

```go
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { ActionEnvironmentSecret } from 'cdk-github';

export class ActionEnvironmentSecretStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const githubTokenSecret = Secret.fromSecretNameV2(this, 'ghSecret', 'GITHUB_TOKEN');
    const sourceSecret = Secret.fromSecretNameV2(this, 'secretToStoreInGitHub', 'testcdkgithub');

    new ActionEnvironmentSecret(this, 'GitHubActionEnvironmentSecret', {
      githubTokenSecret,
      environment: 'dev',
      repository: { name: 'cdk-github', owner: 'wtfjoke' },
      repositorySecretName: 'A_RANDOM_GITHUB_SECRET',
      sourceSecret,
    });
  }
}
```

## GitHubResource

```go
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { StringParameter } from 'aws-cdk-lib/aws-ssm';
import { GitHubResource } from 'cdk-github';


export class GitHubResourceIssueStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const githubTokenSecret = Secret.fromSecretNameV2(this, 'ghSecret', 'GITHUB_TOKEN');
    // optional
    const writeResponseToSSMParameter = StringParameter.fromSecureStringParameterAttributes(this, 'responseBody', { parameterName: '/cdk-github/encrypted-response' });

    new GitHubResource(this, 'GitHubIssue', {
      githubTokenSecret,
      createRequestEndpoint: 'POST /repos/WtfJoke/dummytest/issues',
      createRequestPayload: JSON.stringify({ title: 'Testing cdk-github', body: "I'm opening an issue by using aws cdk 🎉", labels: ['bug'] }),
      createRequestResultParameter: 'number',
      deleteRequestEndpoint: 'PATCH /repos/WtfJoke/dummytest/issues/:number',
      deleteRequestPayload: JSON.stringify({ state: 'closed' }),
      writeResponseToSSMParameter,
    });
  }
}
```

# 💖 Contributing

Contributions of all kinds are welcome! Check out our [contributing guide](CONTRIBUTING.md).
