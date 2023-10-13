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
	Id         int                             `json:"id" yaml:"id"`
	Image      string                          `json:"image" yaml:"image"`
	SecretName string                          `json:"secretName" yaml:"secretName"`
	Stages     []*CiPipelineConfigEnvStageItem `json:"stages" yaml:"stages"`
}

type CiPipelineConfigEnvStageItem struct {
	Name  string                              `json:"name" yaml:"name"`
	Tasks []*CiPipelineConfigEnvStageTaskItem `json:"tasks" yaml:"tasks"`
}
type CiPipelineConfigEnvStageTaskItem struct {
	Type          int            `json:"type" yaml:"type"`
	GitPullData   *GitPullData   `json:"gitPullData,omitempty"`
	ShellExecData *ShellExecData `json:"shellExecData,omitempty"`
}

type GitPullData struct {
	GitUrl       string                   `json:"gitUrl,omitempty"`
	Branch       string                   `json:"branch,omitempty"`
	SecretId     int                      `json:"secretId,omitempty"`
	GitBasicAuth *UsernamePasswordContent `json:"GitBasicAuth,omitempty"`
}

type ShellExecData struct {
	WorkDir string `json:"workDir,omitempty"`
	Content string `json:"content"`
}
