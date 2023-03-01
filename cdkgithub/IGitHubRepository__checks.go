//go:build !no_runtime_type_checking

// AWS CDK Construct Library to interact with GitHub's API.
package cdkgithub

import (
	"fmt"
)

func (j *jsiiProxy_IGitHubRepository) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

