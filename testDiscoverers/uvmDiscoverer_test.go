package testDiscoverers

import (
	"fmt"
	"github.com/shady831213/jarvisSim"
	"github.com/shady831213/jarvisSim/parser"
	_ "github.com/shady831213/jarvisSim/simulators"
	"math/rand"
	"os"
	"path"
	"testing"
)

func TestUvmDiscoverer(t *testing.T) {
	parser.SetRand(rand.New(rand.NewSource(1)))
	os.Setenv("JVS_PRJ_HOME", path.Join(jarivsSim.TestDiscoverersPath(), "testFiles"))
	cfg, err := parser.Lex("testFiles/test_discover.yaml")
	if err != nil {
		t.Error(err)
	}
	err = parser.Parse(cfg)
	if err != nil {
		t.Error(err)
	}
	build1 := parser.GetJvsAstRoot().GetBuild("build1")
	build2 := parser.GetJvsAstRoot().GetBuild("build2")
	compare(t, "discoverer of build1 name", "uvm_test", build1.GetTestDiscoverer().Name())
	compare(t, "testDir of build1 name", path.Join(jarivsSim.TestDiscoverersPath(), "testFiles", "build1_testcases"), build1.GetTestDiscoverer().TestDir())
	compare(t, "testList of build1 name", fmt.Sprint([]string{"test2"}), fmt.Sprint(build1.GetTestDiscoverer().TestList()))
	compare(t, "discoverer of build2 name", "uvm_test", build2.GetTestDiscoverer().Name())
	compare(t, "testDir of build2 name", path.Join(jarivsSim.TestDiscoverersPath(), "testFiles", "testcases"), build2.GetTestDiscoverer().TestDir())
	compare(t, "testList of build2 name", fmt.Sprint([]string{"test1"}), fmt.Sprint(build2.GetTestDiscoverer().TestList()))
}

func compare(t *testing.T, fields, exp, res string) {
	if exp != res {
		t.Error(fields + " expect " + exp + " but get " + res + "!")
	}
}