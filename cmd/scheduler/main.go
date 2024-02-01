package main

import (
	"Scheduler/pkg/defalut"
	"Scheduler/pkg/filternode"
	"Scheduler/pkg/filterpod"
	"Scheduler/pkg/randscore"
	"fmt"
	"os"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(defalut.Name, defalut.New),
		app.WithPlugin(filterpod.Name, filterpod.New),
		app.WithPlugin(filternode.Name, filternode.New),
		app.WithPlugin(randscore.Name, randscore.New),
	)
	logs.InitLogs()
	defer logs.FlushLogs()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
