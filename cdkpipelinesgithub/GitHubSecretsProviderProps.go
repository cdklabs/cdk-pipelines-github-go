package cdkpipelinesgithub


// Locations of GitHub Secrets used to authenticate to AWS.
// Experimental.
type GitHubSecretsProviderProps struct {
	// Experimental.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Experimental.
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Experimental.
	SessionToken *string `field:"optional" json:"sessionToken" yaml:"sessionToken"`
}

