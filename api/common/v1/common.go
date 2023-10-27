package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetGitBranchLstReq struct {
	g.Meta   `method:"get" path:"/common/git-branch-list" summary:"获取指定 Git 仓库的分支名称列表" tags:"通用"`
	GitUrl   string `v:"required" p:"gitUrl"`
	SecretId int    `v:"required" p:"secretId"`
}

type GetGitBranchLstRes struct {
	BranchLst []string `json:"branches"`
}
