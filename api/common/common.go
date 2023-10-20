// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"devops-super/api/common/v1"
)

type ICommonV1 interface {
	TestConnectKubernetes(ctx context.Context, req *v1.TestConnectKubernetesReq) (res *v1.TestConnectKubernetesRes, err error)
}
