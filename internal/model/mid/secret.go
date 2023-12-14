package mid

import "github.com/gogf/gf/v2/encoding/gjson"

type Secret struct {
	Name    string      `v:"required|max-length:30" json:"name"`
	Type    int         `v:"required" json:"type"`
	Content *gjson.Json `v:"required" json:"content"`
}

type UsernamePasswordContent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TextContent struct {
	Text string `json:"text"`
}

type DockerRegistryAuthContent struct {
	RegistryUrl string `json:"registryUrl"`
	UsernamePasswordContent
}
