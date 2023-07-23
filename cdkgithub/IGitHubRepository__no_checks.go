//go:build no_runtime_type_checking

package cdkgithub

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IGitHubRepository) validateSetNameParameters(val *string) error {
	return nil
}

