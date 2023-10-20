package kubernetes

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func (cli *Client) GetNamespaces() ([]string, error) {
	list, err := cli.CoreV1().Namespaces().List(cli.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nsArr []string

	for _, nsItem := range list.Items {
		nsArr = append(nsArr, nsItem.Name)
	}

	return nsArr, nil
}
