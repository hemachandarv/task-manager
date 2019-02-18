package main

import (
	"github.com/hemv/task-manager/internal/app/task/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
