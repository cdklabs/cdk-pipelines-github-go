// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub


// Check suite options.
// Experimental.
type CheckSuiteOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

