// A generated module for DaggerTest functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
)

type DaggerTest struct{}

func (m *DaggerTest) GenerateArgoSecrets(ctx context.Context) (string, error) {
	return dag.Container().
		From("amazonlinux:2").
		WithWorkdir("/exec").
		WithExec([]string{"yum", "install", "-y", "curl", "tar", "gzip", "unzip"}).
		WithExec([]string{"yum", "remove", "awscli"}).
		WithExec([]string{"curl", "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip", "-o", "awscliv2.zip"}).
		WithExec([]string{"unzip", "awscliv2.zip"}).
		WithExec([]string{"./aws/install", "--bin-dir", "/usr/local/bin", "--install-dir", "/usr/local/aws-cli", "--update"}).
		WithExec([]string{"curl", "-LO", "https://dl.k8s.io/release/v1.29.0/bin/linux/amd64/kubectl"}).
		WithExec([]string{"chmod", "+x", "kubectl"}).
		WithExec([]string{"mv", "kubectl", "/usr/local/bin/kubectl"}).
		WithExec([]string{"aws", "--version"}).Stdout(ctx)
}
