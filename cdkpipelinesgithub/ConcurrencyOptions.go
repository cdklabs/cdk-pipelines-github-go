package cdkpipelinesgithub


// Concurrency options at workflow level.
// See: https://docs.github.com/en/actions/using-jobs/using-concurrency
//
// Experimental.
type ConcurrencyOptions struct {
	// The concurrency group to use for the job.
	// Experimental.
	Group *string `field:"required" json:"group" yaml:"group"`
	// Conditionally cancel currently running jobs or workflows in the same concurrency group.
	// Default: false.
	//
	// Experimental.
	CancelInProgress *bool `field:"optional" json:"cancelInProgress" yaml:"cancelInProgress"`
}

