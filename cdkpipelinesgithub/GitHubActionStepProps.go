package cdkpipelinesgithub


// Experimental.
type GitHubActionStepProps struct {
	// The Job steps.
	// Experimental.
	JobSteps *[]*JobStep `field:"required" json:"jobSteps" yaml:"jobSteps"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// The GitHub Environment for the GitHub Action step.
	//
	// To set shell-level environment variables, use `env`.
	// See: https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment
	//
	// Default: No GitHub Environment is selected.
	//
	// Experimental.
	GithubEnvironment *string `field:"optional" json:"githubEnvironment" yaml:"githubEnvironment"`
	// Permissions for the GitHub Action step.
	// Default: The job receives 'contents: write' permissions. If you set additional permissions and require 'contents: write', it must be provided in your configuration.
	//
	// Experimental.
	Permissions *JobPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

