// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub


// Milestone options.
// Experimental.
type MilestoneOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

