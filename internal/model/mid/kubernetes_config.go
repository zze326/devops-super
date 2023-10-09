package mid

type KubernetesConfig struct {
	Name   string `v:"required|max-length:30" json:"name"`
	Config string `v:"required" json:"config"`
}
