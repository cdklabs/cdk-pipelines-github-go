package cdkpipelinesgithub


// Job level settings applied to all docker asset publishing jobs in the workflow.
// Experimental.
type DockerAssetJobSettings struct {
	// Additional permissions to grant to the docker image publishing job.
	// Default: - no additional permissions.
	//
	// Experimental.
	Permissions *JobPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// GitHub workflow steps to execute before building and publishing the image.
	// Default: [].
	//
	// Experimental.
	SetupSteps *[]*JobStep `field:"optional" json:"setupSteps" yaml:"setupSteps"`
}

