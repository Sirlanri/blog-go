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
