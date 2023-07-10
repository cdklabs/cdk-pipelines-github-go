package cdkpipelinesgithub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the GitHubActionRole construct.
// Experimental.
type GitHubActionRoleProps struct {
	// A list of GitHub repositories you want to be able to access the IAM role.
	//
	// Each entry should be your GitHub username and repository passed in as a
	// single string.
	//
	// For example, `['owner/repo1', 'owner/repo2'].
	// Experimental.
	Repos *[]*string `field:"required" json:"repos" yaml:"repos"`
	// The GitHub OpenId Connect Provider. Must have provider url `https://token.actions.githubusercontent.com`. The audience must be `sts:amazonaws.com`.
	//
	// Only one such provider can be defined per account, so if you already
	// have a provider with the same url, a new provider cannot be created for you.
	// Experimental.
	Provider awsiam.IOpenIdConnectProvider `field:"optional" json:"provider" yaml:"provider"`
	// The name of the Oidc role.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Thumbprints of GitHub's certificates.
	//
	// Every time GitHub rotates their certificates, this value will need to be updated.
	//
	// Default value is up-to-date to June 27, 2023 as per
	// https://github.blog/changelog/2023-06-27-github-actions-update-on-oidc-integration-with-aws/
	// Experimental.
	Thumbprints *[]*string `field:"optional" json:"thumbprints" yaml:"thumbprints"`
}

