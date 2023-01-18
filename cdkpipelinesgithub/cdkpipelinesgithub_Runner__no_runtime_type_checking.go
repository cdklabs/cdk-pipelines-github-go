//go:build no_runtime_type_checking

// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateRunner_SelfHostedParameters(labels *[]*string) error {
	return nil
}

