//go:build no_runtime_type_checking

package cdkgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateGitHubResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGitHubResourceParameters(scope constructs.Construct, id *string, props *GitHubResourceProps) error {
	return nil
}

