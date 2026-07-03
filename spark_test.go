package computespark

import (
	"testing"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/togo"
)

func TestArgv(t *testing.T) {
	t.Setenv("SPARK_MASTER", "yarn")
	t.Setenv("SPARK_OPTS", "--deploy-mode cluster")
	s := newSpark(togo.New())
	got := s.argv(compute.Job{Cmd: []string{"app.py", "--date=today"}})
	want := []string{"--master", "yarn", "--deploy-mode", "cluster", "app.py", "--date=today"}
	if len(got) != len(want) {
		t.Fatalf("argv = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("argv[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
