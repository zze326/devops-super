package main

import (
	"devops-super/internal/consts"
	"devops-super/internal/logic/ci_pipeline/client/task"
	"devops-super/internal/model/mid"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
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
		fmt.Printf("【%s】开始\n", stage.Name)
		for _, taskRun := range stage.Tasks {
			switch taskRun.Type {
			case consts.PIPELINE_TASK_TYPE_GIT_PULL:
				if err := task.GitClone(taskRun.GitPullData); err != nil {
					log.Fatalf("git clone failed: err: %v", err)
				}
			case consts.PIPELINE_TASK_TYPE_SHELL_EXEC:
				if err := task.ShellExec(taskRun.ShellExecData); err != nil {
					log.Fatalf("exec shell failed, err: %v", err)
				}
			}
		}
		fmt.Printf("【%s】结束\n", stage.Name)
	}
}

func main() {
	initParams()
	runStages()
}
