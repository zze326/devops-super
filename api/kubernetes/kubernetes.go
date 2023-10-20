// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package kubernetes

import (
	"context"

	"devops-super/api/kubernetes/v1"
)

type IKubernetesV1 interface {
	TestConnect(ctx context.Context, req *v1.TestConnectReq) (res *v1.TestConnectRes, err error)
}
