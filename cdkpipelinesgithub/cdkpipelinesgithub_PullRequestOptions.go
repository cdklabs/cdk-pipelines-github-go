// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub


// Pull request options.
// Experimental.
type PullRequestOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

