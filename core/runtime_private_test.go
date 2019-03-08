package core

import (
	"os"
	"testing"
)

func setUp(name string, cfg map[interface{}]interface{}) (*runTime, error) {
	group := newAstGroup("Jarvis")
	if err := group.Parse(cfg); err != nil {
		return nil, err
	}
	if err := group.Link(); err != nil {
		return nil, err
	}
	return newRunTime(name, group), nil
}

func setUpGroup(group *astGroup, args []string) (*runTime, error) {
	return setUp(group.Name, map[interface{}]interface{}{"args": convertArgs(args), "groups": []interface{}{group.Name}})
}

func setUpTest(testName, buildName string, args []string) (*runTime, error) {
	return setUp(testName, map[interface{}]interface{}{"build": buildName,
		"args":  convertArgs(args),
		"tests": map[interface{}]interface{}{testName: nil}})
}

func setUpOnlyBuild(buildName string, args []string) (*runTime, error) {
	return setUp(buildName, map[interface{}]interface{}{"build": buildName,
		"args": convertArgs(args)})
}

func TestGroupSetup(t *testing.T) {
	if r, err := setUpGroup(GetJvsAstRoot().GetGroup("group1"), nil); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if len(r.runFlow) != 1 {
			t.Error("expect 1 runFlow but get " + string(len(r.runFlow)))
			t.FailNow()
		}
	}

	if r, err := setUpGroup(GetJvsAstRoot().GetGroup("group2"), []string{}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if len(r.runFlow) != 2 {
			t.Error("expect 2 runFlow but get " + string(len(r.runFlow)))
			t.FailNow()
		}
	}
	if r, err := setUpGroup(GetJvsAstRoot().GetGroup("group3"), []string{}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if len(r.runFlow) != 2 {
			t.Error("expect 2 runFlow but get " + string(len(r.runFlow)))
			t.FailNow()
		}
	}
}

func TestSingleTestSetup(t *testing.T) {
	if r, err := setUpTest("test1", "build1", []string{"-seed 1"}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if r.cmdStdout != os.Stdout {
			t.Error("when running single test, expect stdout is open but closed")
			t.FailNow()
		}
	}
}

func TestRunOnlyBuildSetup(t *testing.T) {
	if r, err := setUpOnlyBuild("build1", []string{"-test_phase jarvis"}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if r.cmdStdout != os.Stdout {
			t.Error("when running only build, expect stdout is open but closed")
			t.FailNow()
		}
	}
}