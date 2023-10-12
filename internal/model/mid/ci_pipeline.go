package mid

type CiPipeline struct {
	Name               string `v:"required|max-length:30" json:"name"`
	KubernetesConfigId int    `v:"required" json:"kubernetesConfigId"`
	Desc               string `json:"desc"`
}

type CiPipelineConfig []*CiPipelineConfigEnvItem

func (c CiPipelineConfig) GetEnvIds() []int {
	var envIds []int
	for _, envItem := range c {
		envIds = append(envIds, envItem.Id)
	}
	return envIds
}

type CiPipelineConfigEnvItem struct {
	Id     int                             `json:"id"`
	Stages []*CiPipelineConfigEnvStageItem `json:"stages"`
}

type CiPipelineConfigEnvStageItem struct {
	Name string `json:"name"`
}
