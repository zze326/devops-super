package mid

type CiPipeline struct {
	Name               string `v:"required|max-length:30" json:"name"`
	KubernetesConfigId int    `v:"required" json:"kubernetesConfigId"`
	Desc               string `json:"desc"`
}
