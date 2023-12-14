package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (cli *Client) GetConfigMap(namespace, name string) (*corev1.ConfigMap, error) {
	configMap, err := cli.CoreV1().ConfigMaps(namespace).Get(cli.Ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return configMap, nil
}

func (cli *Client) CreateConfigMap(namespace, name string, data map[string]string) error {
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
	_, err := cli.CoreV1().ConfigMaps(namespace).Create(cli.Ctx, configMap, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client) UpdateConfigMap(namespace string, configMap *corev1.ConfigMap) error {
	_, err := cli.CoreV1().ConfigMaps(namespace).Update(cli.Ctx, configMap, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}
