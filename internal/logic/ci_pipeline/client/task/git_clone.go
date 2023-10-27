package task

import (
	"devops-super/internal/model/mid"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
	"path"
	"path/filepath"
)

func GitClone(data *mid.GitPullData) error {
	fmt.Println(aurora.Blue(fmt.Sprintf("Git 地址: %s", data.GitUrl)))
	fmt.Println(aurora.Blue(fmt.Sprintf("分支: %s", data.Branch)))
	baseName := filepath.Base(data.GitUrl)
	extension := filepath.Ext(baseName)

	basicAuth := &http.BasicAuth{}
	if data.GitBasicAuth != nil {
		basicAuth.Username = data.GitBasicAuth.Username
		basicAuth.Password = data.GitBasicAuth.Password
	}
	fileNameWithoutExtension := baseName[0 : len(baseName)-len(extension)]
	r, err := git.PlainClone(path.Join(fileNameWithoutExtension), false, &git.CloneOptions{
		Auth:          basicAuth,
		URL:           data.GitUrl,
		Progress:      os.Stdout,
		SingleBranch:  true,
		ReferenceName: plumbing.ReferenceName(data.Branch),
	})
	if err != nil {
		return err
	}
	ref, err := r.Head()
	if err != nil {
		log.Printf("err: %v", err)
		return nil
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Printf("err: %v", err)
		return nil
	}
	fmt.Print(commit)
	return nil
}
