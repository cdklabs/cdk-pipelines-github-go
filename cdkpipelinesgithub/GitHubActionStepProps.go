package cdkpipelinesgithub


// Experimental.
type GitHubActionStepProps struct {
	// The Job steps.
	// Experimental.
	JobSteps *[]*JobStep `field:"required" json:"jobSteps" yaml:"jobSteps"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Permissions for the GitHub Action step.
	// Default: The job receives 'contents: write' permissions. If you set additional permissions and require 'contents: write', it must be provided in your configuration.
	//
	// Experimental.
	Permissions *JobPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

