package main

import (
	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	getBirthday()
}

func main() {
	wf.Run(run)
}
