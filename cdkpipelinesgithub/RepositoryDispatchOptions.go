// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub


// Repository dispatch options.
// Experimental.
type RepositoryDispatchOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}
