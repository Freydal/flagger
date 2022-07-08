/*
Copyright 2020, 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loadtester

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

const TaskTypeHelmv3 = "helmv3"
const TaskTypeHelm = "helm"

type HelmTaskv3 struct {
	TaskBase
	status       string
	command      string
	logCmdOutput bool
}

func (task *HelmTaskv3) Hash() string {
	return hash(task.canary + task.command)
}

func (task *HelmTaskv3) Run(ctx context.Context) (*TaskRunResult, error) {
	if task.status != "" {
		if out, err := task.checkStatus(ctx); err != nil {
			return out, fmt.Errorf("command %s failed: %s: %w", task.status, out, err)
		}
	}

	helmCmd := fmt.Sprintf("%s %s", TaskTypeHelmv3, task.command)
	task.logger.With("canary", task.canary).Infof("running command %v", helmCmd)

	cmd := exec.CommandContext(ctx, TaskTypeHelm, strings.Fields(task.command)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		task.logger.With("canary", task.canary).Errorf("command failed %s %v %s", task.command, err, out)
		return &TaskRunResult{false, out}, fmt.Errorf("command %s failed: %s: %w", task.command, out, err)
	} else {
		if task.logCmdOutput {
			fmt.Printf("%s\n", out)
		}
		task.logger.With("canary", task.canary).Infof("command finished %v", helmCmd)
	}
	return &TaskRunResult{true, out}, nil
}

func (task *HelmTaskv3) checkStatus(ctx context.Context) (*TaskRunResult, error) {
	helmCmd := fmt.Sprintf("%s %s", TaskTypeHelmv3, task.status)
	task.logger.With("canary", task.canary).Infof("running status command %v", helmCmd)

	cmd := exec.CommandContext(ctx, TaskTypeHelm, strings.Fields(task.status)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		task.logger.With("canary", task.canary).Errorf("status command failed %s %v %s", task.command, err, out)
		return &TaskRunResult{false, out}, fmt.Errorf("status command %s failed: %s: %w", task.command, out, err)
	} else {
		if task.logCmdOutput {
			fmt.Printf("%s\n", out)
		}
		task.logger.With("canary", task.canary).Infof("status command finished %v", helmCmd)
	}

	if !strings.Contains(string(out), "STATUS: deployed") {
		return &TaskRunResult{false, out}, fmt.Errorf("status is not deployed")
	}

	return &TaskRunResult{true, out}, nil
}

func (task *HelmTaskv3) String() string {
	return task.command
}
