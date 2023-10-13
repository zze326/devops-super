package task

import (
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/util/gutil"
	"os"
	"os/exec"
)

func ShellExec(data *mid.ShellExecData) (err error) {
	cmd := exec.Command("sh", "-c", data.Content)
	if !gutil.IsEmpty(data.WorkDir) {
		cmd.Dir = data.WorkDir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err = cmd.Run()
	return
}
