package main

import (
	"flag"

	"github.com/dshmelev/kube-cfssl/pkg/kubecfssl"
)

var AppVersion = "unknown"
var AppGitCommit = ""
var AppGitState = ""

func Version() string {
	version := AppVersion
	if len(AppGitCommit) > 0 {
		version += "-"
		version += AppGitCommit[0:8]
	}
	if len(AppGitState) > 0 && AppGitState != "clean" {
		version += "-"
		version += AppGitState
	}
	return version
}

func main() {
	// parse standard command line arguments
	flag.Parse()

	kc := kubecfssl.New(Version())
	kc.Init()
}