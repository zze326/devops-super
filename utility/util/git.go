package util

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

func GetRemoteBranchList(gitUrl string, auth *http.BasicAuth) ([]string, error) {
	remote := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		URLs:   []string{gitUrl},
		Mirror: false,
		Fetch:  nil,
	})

	refs, err := remote.List(&git.ListOptions{
		Auth:            auth,
		InsecureSkipTLS: true,
		Timeout:         0,
	})

	if err != nil {
		return nil, err
	}

	branches := make([]string, 0)

	// 遍历引用列表，过滤出分支引用并打印名称
	for _, ref := range refs {
		if name := ref.Name(); name.IsBranch() {
			branches = append(branches, name.Short())
		}
	}
	return branches, nil
}
