//go:build no_runtime_type_checking

// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCredentialsProvider) validateCredentialStepsParameters(region *string) error {
	return nil
}

