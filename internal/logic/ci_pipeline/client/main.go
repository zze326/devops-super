package main

import (
	"devops-super/internal/consts"
	"devops-super/internal/logic/ci_pipeline/client/task"
	"devops-super/internal/model/mid"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
)

var (
	stages []*mid.CiPipelineConfigEnvStageItem
)

func initParams() {
	stagesJson := os.Getenv("STAGES")
	if err := gjson.DecodeTo([]byte(stagesJson), &stages); err != nil {
		log.Fatal(err)
	}
}

func runStages() {
	for _, stage := range stages {
		fmt.Println(aurora.Green(fmt.Sprintf("=====【%s】开始=====", stage.Name)))
		for _, taskRun := range stage.Tasks {
			switch taskRun.Type {
			case consts.PIPELINE_TASK_TYPE_GIT_PULL:
				if err := task.GitClone(taskRun.GitPullData); err != nil {
					fmt.Print(aurora.BgRed(fmt.Sprintf("Git 拉取失败，err: %v", err)))
					os.Exit(1)
				}
			case consts.PIPELINE_TASK_TYPE_SHELL_EXEC:
				if err := task.ShellExec(taskRun.ShellExecData); err != nil {
					fmt.Print(aurora.BgRed(fmt.Sprintf("Shell 执行失败，err: %v", err)))
					os.Exit(1)
				}
			}
		}
		fmt.Println(aurora.Green(fmt.Sprintf("=====【%s】结束=====", stage.Name)))
	}
}

func main() {
	initParams()
	runStages()
}
