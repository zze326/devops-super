package mid

type CiEnv struct {
	Name   string `v:"required|max-length:30" json:"name"`
	Image  string `v:"required" json:"image"`
	Secret string `json:"secret"`
}
