package init

import (
	"errors"
	"fmt"
	"github.com/shady831213/jarvism/cmd/base"
	"github.com/shady831213/jarvism/core/utils"
	"os"
	"path"
	"path/filepath"
)

var CmdInit = &base.Command{
	UsageLine: "jarvism init [-prj_dir DIR][-work_dir DIR]",
	Short:     "create a jarvism default project",
	Long: `
. $prj_dir
|--- jarvism_cfg
|------ jarvism_cfg.yaml
|------ jarvism_setup.sh(export $JVS_PRJ_HOME;export $JVS_WORK_DIR)
|--- src
|--- testcases
. $work_dir
`,
}

var (
	prjDir  string
	workDir string
)

func init() {
	CmdInit.Run = runInit
	CmdInit.Flag.StringVar(&prjDir, "prj_dir", "", "assign prj dir, default is pwd")
	CmdInit.Flag.StringVar(&workDir, "work_dir", "", "assign work dir, default is $prj_dir/work")
	base.Jarvism.AddCommand(CmdInit)
}

func runInit(cmd *base.Command, args []string) error {
	if prjDir == "" {
		prjDir = os.Getenv("PWD")
	}
	prjDir, err := filepath.Abs(os.ExpandEnv(prjDir))
	if err != nil {
		return errors.New(utils.Red(err.Error()))
	}

	if workDir == "" {
		workDir = path.Join(prjDir, "work")
	}
	workDir, err := filepath.Abs(os.ExpandEnv(workDir))
	if err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	//make dirs
	if err := os.MkdirAll(prjDir, os.ModePerm); err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	if err := os.MkdirAll(workDir, os.ModePerm); err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	if err := os.Mkdir(path.Join(prjDir, "jarvism_cfg"), os.ModePerm); err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	if err := os.Mkdir(path.Join(prjDir, "src"), os.ModePerm); err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	if err := os.Mkdir(path.Join(prjDir, "testcases"), os.ModePerm); err != nil {
		return errors.New(utils.Red(err.Error()))
	}

	//files
	setupContent := fmt.Sprintf("#!/bin/bash\nexport JVS_PRJ_HOME=%s\nexport JVS_WORK_DIR=%s\n", prjDir, workDir)
	if err := utils.WriteNewFile(path.Join(prjDir, "jarvism_cfg", "jarvism_setup.sh"), setupContent); err != nil {
		return errors.New(utils.Red(err.Error()))
	}
	yamlContent := "builds:\n\tbuild1:\n"
	if err := utils.WriteNewFile(path.Join(prjDir, "jarvism_cfg", "jarvism_cfg.yaml"), yamlContent); err != nil {
		return errors.New(utils.Red(err.Error()))
	}

	return nil
}
