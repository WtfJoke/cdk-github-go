//go:build no_runtime_type_checking

package cdkgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateActionSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewActionSecretParameters(scope constructs.Construct, id *string, props *ActionSecretProps) error {
	return nil
}

