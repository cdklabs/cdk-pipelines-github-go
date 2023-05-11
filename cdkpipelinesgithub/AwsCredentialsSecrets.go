package cdkpipelinesgithub


// Names of secrets for AWS credentials.
// Experimental.
type AwsCredentialsSecrets struct {
	// Experimental.
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Experimental.
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Experimental.
	SessionToken *string `field:"optional" json:"sessionToken" yaml:"sessionToken"`
}

