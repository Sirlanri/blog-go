package structs

//ResMP 从前端post的用户名+密码结构体
type ResMP struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

//Gotnp -sql.ConnectDB 从数据库中获取的 用户名、密码、权限组合
type Gotnp struct {
	Name     string `json:"name"`
	Password string `json:"paswword"`
	Power    int    `json:"power"`
}
