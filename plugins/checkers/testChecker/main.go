package main

import (
	"github.com/shady831213/jarvism/core/ast"
	"github.com/shady831213/jarvism/core/errors"
	"regexp"
)

type testChecker struct {
	ast.CheckerBase
}

func newTestChecker() ast.Checker {
	inst := new(testChecker)
	inst.Init("testChecker")

	//UVM ERROR and FATAL
	inst.AddPats(errors.JVSRuntimeFail, false, regexp.MustCompile(`^.*UVM_((ERROR)|(FATAL)) .*\@.*:`))
	//Errors
	inst.AddPats(errors.JVSRuntimeFail, false, regexp.MustCompile(`^Error((.+:)|(-\[.*\]))`))

	//UVM Warning
	inst.AddPats(errors.JVSRuntimeWarning, false, regexp.MustCompile(`^.*UVM_WARNING .*\@.*:`))
	//Timing violation
	inst.AddPats(errors.JVSRuntimeWarning, false, regexp.MustCompile(`.*Timing violation.*`))
	return inst
}

func init() {
	ast.RegisterChecker(newTestChecker)
}
