package main

import (
	"github.com/shady831213/jarvism/core/loader"
	"github.com/shady831213/jarvism/core/errors"
	"regexp"
)

type compileChecker struct {
	loader.CheckerBase
}

func newCompileChecker() loader.Plugin {
	inst := new(compileChecker)
	inst.Init("compileChecker")
	//Errors
	inst.AddPats(errors.JVSRuntimeFail, false, regexp.MustCompile(`^Error((.+:)|(-\[.*\]))`))
	return inst
}

func init() {
	loader.RegisterChecker(newCompileChecker)
}
