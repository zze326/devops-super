package mid

type Host struct {
	Name        string `v:"required" json:"name"`
	HostAddr    string `v:"required" json:"hostAddr"`
	Port        int    `v:"required" json:"port"`
	Username    string `v:"required" json:"username"`
	Password    string `json:"password"`
	PrivateKey  string `json:"privateKey"`
	UseKey      bool   `v:"required" json:"useKey"`
	Desc        string `json:"desc"`
	SaveSession bool   `v:"required" json:"saveSession"`
	HostGroupId int    `json:"hostGroupId"`
}

type DownloadFileIn struct {
	Id   int
	Path string
}
