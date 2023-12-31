package task

import (
	"devops-super/internal/model/mid"
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/logrusorgru/aurora"
	"os"
	"os/exec"
	"path/filepath"
)

func ShellExec(data *mid.ShellExecData) (err error) {
	fmt.Println(aurora.Blue(fmt.Sprintf("Shell 工作目录: %s", data.WorkDir)))
	shellFilePath := filepath.Join("/tmp", fmt.Sprintf("%s.sh", gfile.Temp()))
	if err = gfile.PutContents(shellFilePath, data.Content); err != nil {
		return
	}
	cmd := exec.Command("sh", "-x", shellFilePath)
	if !gutil.IsEmpty(data.WorkDir) {
		cmd.Dir = data.WorkDir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err = cmd.Run()
	return
}
