//go:build no_runtime_type_checking

package cdkgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateActionEnvironmentSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewActionEnvironmentSecretParameters(scope constructs.Construct, id *string, props *ActionEnvironmentSecretProps) error {
	return nil
}

