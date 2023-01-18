// GitHub Workflows support for CDK Pipelines
package cdkpipelinesgithub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-pipelines-github-go/cdkpipelinesgithub/jsii"
)

// Represents a Yaml File.
// Experimental.
type YamlFile interface {
	// A comment to be added to the top of the YAML file.
	//
	// Can be multiline. All non-empty line are pefixed with '# '. Empty lines are kept, but not commented.
	//
	// For example:
	// ```ts
	// pipeline.workflowFile.commentAtTop =
	// `AUTOGENERATED FILE, DO NOT EDIT!
	// See ReadMe.md
	// `;
	// ```
	//
	// Results in YAML:
	// ```yaml
	// # AUTOGENERATED FILE, DO NOT EDIT!
	// # See ReadMe.md
	//
	// name: deploy
	// ...
	// ```.
	// Experimental.
	CommentAtTop() *string
	// Experimental.
	SetCommentAtTop(val *string)
	// Applies an RFC 6902 JSON-patch to the synthesized object file. See https://datatracker.ietf.org/doc/html/rfc6902 for more information.
	//
	// For example, with the following yaml file
	// ```yaml
	// name: deploy
	// on:
	//    push:
	//      branches:
	//        - main
	//    workflow_dispatch: {}
	// ...
	// ```
	//
	// modified in the following way:
	//
	// ```ts
	// pipeline.workflowFile.patch(JsonPatch.add("/on/workflow_call", "{}"));
	// pipeline.workflowFile.patch(JsonPatch.remove("/on/workflow_dispatch"));
	// ```
	//
	// would result in the following yaml file:
	//
	// ```yaml
	// name: deploy
	// on:
	//    push:
	//      branches:
	//        - main
	//    workflow_call: {}
	// ...
	// ```.
	// Experimental.
	Patch(patches ...JsonPatch)
	// Returns the patched yaml file.
	// Experimental.
	ToYaml() *string
	// Update the output object.
	// Experimental.
	Update(obj interface{})
	// Write the patched yaml file to the specified location.
	// Experimental.
	WriteFile()
}

// The jsii proxy struct for YamlFile
type jsiiProxy_YamlFile struct {
	_ byte // padding
}

func (j *jsiiProxy_YamlFile) CommentAtTop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentAtTop",
		&returns,
	)
	return returns
}


// Experimental.
func NewYamlFile(filePath *string, options *YamlFileOptions) YamlFile {
	_init_.Initialize()

	if err := validateNewYamlFileParameters(filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_YamlFile{}

	_jsii_.Create(
		"cdk-pipelines-github.YamlFile",
		[]interface{}{filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewYamlFile_Override(y YamlFile, filePath *string, options *YamlFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-pipelines-github.YamlFile",
		[]interface{}{filePath, options},
		y,
	)
}

func (j *jsiiProxy_YamlFile)SetCommentAtTop(val *string) {
	_jsii_.Set(
		j,
		"commentAtTop",
		val,
	)
}

func (y *jsiiProxy_YamlFile) Patch(patches ...JsonPatch) {
	args := []interface{}{}
	for _, a := range patches {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		y,
		"patch",
		args,
	)
}

func (y *jsiiProxy_YamlFile) ToYaml() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (y *jsiiProxy_YamlFile) Update(obj interface{}) {
	if err := y.validateUpdateParameters(obj); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		y,
		"update",
		[]interface{}{obj},
	)
}

func (y *jsiiProxy_YamlFile) WriteFile() {
	_jsii_.InvokeVoid(
		y,
		"writeFile",
		nil, // no parameters
	)
}

