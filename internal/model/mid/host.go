package mid

type Host struct {
	Name        string `v:"required" json:"name"`
	Host        string `v:"required" json:"host"`
	Port        int64  `v:"required" json:"port"`
	Username    string `v:"required" json:"username"`
	Password    string `json:"password"`
	PrivateKey  string `json:"privateKey"`
	UseKey      bool   `v:"required" json:"useKey"`
	Desc        string `json:"desc"`
	SaveSession bool   `v:"required" json:"saveSession"`
}
