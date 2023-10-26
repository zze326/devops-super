package kubernetes

import (
	"context"
	"devops-super/utility/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cli *Client) GetPod(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	return cli.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (cli *Client) DeletePodForce(ctx context.Context, namespace, name string) error {
	return cli.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{
		TypeMeta:           metav1.TypeMeta{},
		GracePeriodSeconds: util.ToPointer(int64(1)),
		PropagationPolicy:  nil,
		DryRun:             nil,
	})
}
