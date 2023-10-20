// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IKubernetes interface {
		TestConnect(ctx context.Context, config string) (err error)
		GetNamespaces(ctx context.Context, config string) ([]string, error)
	}
)

var (
	localKubernetes IKubernetes
)

func Kubernetes() IKubernetes {
	if localKubernetes == nil {
		panic("implement not found for interface IKubernetes, forgot register?")
	}
	return localKubernetes
}

func RegisterKubernetes(i IKubernetes) {
	localKubernetes = i
}
