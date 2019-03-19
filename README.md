# jarvism[![Build Status](https://travis-ci.org/shady831213/jarvism.svg?branch=master)](https://travis-ci.org/shady831213/jarvism)
Just A Really Very Impressive Simulation Manager

# install
instatll go

[go](https://github.com/golang/go)

install jarvism
```
go get -u github.com/shady831213/jarvism

jarvism help
```

# test
```
cd [JARVISM_DIR]
./script/clean.sh
go test -v ./...
./script/clean.sh
```

# Getting start
```
$ jarvism init -prj_dir jvs_prj
$ cd jvs_prj
$ ls 
algorithms  go	jarivsm  jvs_prj  src  testcases  work
$ jarvism help init
usage: jarvism init [-prj_dir DIR][-work_dir DIR]

. $prj_dir
|--- jarvism_cfg
|------ jarvism_cfg.yaml
|------ jarvism_setup.sh(export $JVS_PRJ_HOME;export $JVS_WORK_DIR)
|--- src
|--- testcases
. $work_dir
  -prj_dir string
    	assign prj dir, default is pwd
  -work_dir string
    	assign work dir, default is $prj_dir/work
```
Enjoy!

# Example

TBD

But if you have vcs, you can run tests in plugins/runner/host
```
cd $GOPATH/src/github.com/shady831213/jarvism/plugins/runner/host
go test -v
```

# Plugins

TBD

# godoc
[doc](https://godoc.org/github.com/shady831213/jarvism)

# usage
```
Usage:

	jarvism <command> [arguments]

The commands are:

	init        create a jarvism default project
	run_parse   only parse cfg(jarvism_cfg dir or jarvism_cfg.yaml file)
	run_test    run single test, build name must assigned
	run_group   run group
	run_build   run single build
	show_tests  list tests in corresponding build
	show_groups list all groups
	show_builds list all builds
	show_plugins list all plugins or reporter, simulator, runner, checker, testDiscoverer

Use "jarvsim help <command>" for more information about a command.

plugins:
all runners:
	 host
all testDiscoverers:
	 uvm_test
all simulators:
	 vcs
all checkers:
	 compileChecker
	 testChecker
all reporters:
	 junit

run options:

  -compile_args
    	compiling args pass to simulator (default false)
  -max_job int
    	limit of runtime coroutines, default is unlimited. (default -1)
  -repeat
    	run each testcase repeatly n times (default )
  -reporter
    	add reporter plugin, can apply multi times, default
  -seed
    	run testcase with specific seed
  -sim_args
    	simulation args pass to simulator (default false)
  -sim_only
    	bypass compile and only run simulation, default is false.
  -unique
    	if set jobId(timestamp) will be included in hash, then builds and testcases will have unique name and be in unique dir.default is false.
```
