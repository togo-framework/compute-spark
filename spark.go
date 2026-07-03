// Package computespark is a togo compute backend that submits jobs to Apache
// Spark via spark-submit. Registers into the compute plugin's `compute` slot.
// Select with `togo provider:use compute spark`.
//
// Config (dynamic — CLI/.env/settings):
//   SPARK_SUBMIT_BIN  spark-submit binary   (default "spark-submit")
//   SPARK_MASTER      --master value        (e.g. yarn, k8s://…, local[*])
//   SPARK_OPTS        extra spark-submit opts (space-separated)
package computespark

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"os/exec"
	"strings"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("compute-spark", togo.PriorityService+1, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCompute, "spark", newSpark(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "compute-spark")
		}
		return nil
	})
}

type spark struct {
	bin    string
	master string
	opts   []string
}

func newSpark(k *togo.Kernel) *spark {
	return &spark{
		bin:    providers.Value(k, providers.CapCompute, "spark", "submit_bin", "spark-submit", false),
		master: providers.Value(k, providers.CapCompute, "spark", "master", "", false),
		opts:   strings.Fields(providers.Value(k, providers.CapCompute, "spark", "opts", "", false)),
	}
}

// Submit builds `spark-submit [--master X] [opts] <job.Cmd...>`.
func (s *spark) Submit(ctx context.Context, job compute.Job) (compute.Run, error) {
	var args []string
	if s.master != "" {
		args = append(args, "--master", s.master)
	}
	args = append(args, s.opts...)
	args = append(args, job.Cmd...)
	return runEngine(ctx, s.bin, args), nil
}

func (s *spark) argv(job compute.Job) []string { // exported-for-test seam
	var args []string
	if s.master != "" {
		args = append(args, "--master", s.master)
	}
	args = append(args, s.opts...)
	return append(args, job.Cmd...)
}

func runEngine(ctx context.Context, bin string, args []string) compute.Run {
	cmd := exec.CommandContext(ctx, bin, args...)
	var buf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &buf, &buf
	run := compute.Run{ID: id(), Status: "succeeded", Output: ""}
	if err := cmd.Run(); err != nil {
		run.Status = "failed"
	}
	run.Output = buf.String()
	return run
}

func id() string { b := make([]byte, 8); _, _ = rand.Read(b); return hex.EncodeToString(b) }
