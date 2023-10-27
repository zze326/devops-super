// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"devops-super/api/common/v1"
)

type ICommonV1 interface {
	GetGitBranchLst(ctx context.Context, req *v1.GetGitBranchLstReq) (res *v1.GetGitBranchLstRes, err error)
}
