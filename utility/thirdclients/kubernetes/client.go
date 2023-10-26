package kubernetes

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	*kubernetes.Clientset
	*rest.Config
	Ctx context.Context
}

func NewClient(ctx context.Context, config string) (*Client, error) {
	var (
		restConf  *rest.Config
		err       error
		clientSet *kubernetes.Clientset
	)
	if restConf, err = clientcmd.RESTConfigFromKubeConfig([]byte(config)); err != nil {
		return nil, err
	}

	if clientSet, err = kubernetes.NewForConfig(restConf); err != nil {
		return nil, err
	}
	return &Client{
		Clientset: clientSet,
		Config:    restConf,
		Ctx:       ctx,
	}, nil
}

func IsPodNotFoundError(err error) bool {
	if statusErr, ok := err.(*errors.StatusError); ok {
		if statusErr.ErrStatus.Reason == metav1.StatusReasonNotFound {
			return true
		}
	}
	return false
}
