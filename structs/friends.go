package structs

//Friend 数据库中的Friend类型
type Friend struct {
	ID           int    `json:"id"`
	PicAddress   string `json:"picaddress"`
	SiteName     string `json:"sitename"`
	Introduction string `json:"introduction"`
	URL          string `json:"url"`
	SSL          bool   `json:"ssl"`
	Ping         int    `json:"ping"`
}

//ResFriend 添加新友链，从前端接收的数据
type ResFriend struct {
	PicAddress   string
	Sitename     string
	Siteaddress  string
	Introduction string
	Ssl          bool
}

//ResUpdateFriend 修改友链信息的结构体
type ResUpdateFriend struct {
	ID           int
	PicAddress   string
	Sitename     string
	Siteaddress  string
	Introduction string
	Ssl          bool
	Ping         int
}
