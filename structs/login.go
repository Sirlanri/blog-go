package structs

//NamePw 从前端post的用户名+密码结构体
type NamePw struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
