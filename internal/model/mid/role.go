package mid

type Role struct {
	Name string `v:"required|max-length:30" json:"name"`
	Code string `v:"required|max-length:30" json:"code"`
}
