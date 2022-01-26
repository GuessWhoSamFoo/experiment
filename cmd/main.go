package main

import (
	"flag"
	"github.com/GuessWhoSamFoo/experiment/pkg"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
)

func init() {
	flag.StringVar(&pkg.Commit, "commit", "19769456b6b229b3e78f2b90eced15a353eb4e7c", "creates new TaskRuns via Runnable")
	flag.Parse()
}

func main() {
	app := cdk8s.NewApp(nil)
	pkg.NewExample(app, "test", nil)
	app.Synth()
}
