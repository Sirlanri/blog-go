package handler

import (
	"blog-go/serves"
	"blog-go/structs"
	"io"
	"io/ioutil"
	"os"

	"github.com/kataras/iris/v12"
)

//AddFriend -handler 添加一个友链
func AddFriend(ctx iris.Context) {
	var newFriend structs.ResFriend
	err := ctx.ReadJSON(&newFriend)
	if err != nil {
		println("错误-添加友链-前端传入数据错误", err.Error())
		return
	}
	println("AddFriend Handler处理完毕，转交Serve")
	serves.AddFriendServe(newFriend)
}

//UploadPic -handler 上传图片，并命名为UUID保存到uploadpics目录下，
//向前端返回图片的URL
func UploadPic(ctx iris.Context) {
	file, info, err := ctx.FormFile("pic")
	if err != nil {
		//status==500 上传失败
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("上传图片失败")
		println("上传图片失败", err.Error())
		return
	}
	defer file.Close()
	fname := serves.Createid() + info.Filename
	//图片保存目录 uploadpics
	out, err := os.OpenFile("./uploadpics/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("图片保存至服务器失败")
		println("图片保存至服务器失败", err.Error())
		return
	}
	defer out.Close()
	io.Copy(out, file)
	whole := "http://api.ri-co.cn/blog/getpics/" + fname
	ctx.WriteString(whole)
}

//GetPics 获取图片
func GetPics(ctx iris.Context) {
	name := ctx.Params().GetString("name")
	pic, err := os.Open("./uploadpics/" + name)
	if err != nil {
		println("查找图片出错 ", name, err.Error())
		return
	}
	defer pic.Close()
	buff, err := ioutil.ReadAll(pic)
	if err != nil {
		println("返回图片出错", err.Error())
	}
	ctx.Write(buff)
}

//Refreshms -handler 手动刷新友链服务器的延迟
func Refreshms(ctx iris.Context) {
	links := serves.GetAllAddress()
	newData := make(map[string]int)
	for _, link := range links {
		dropped := serves.DropHead(link)
		number := serves.Pingms(dropped)
		newData[link] = number
	}
	println("ping测量完毕，开始写入数据")
	serves.Updatems(newData)
	ctx.WriteString("ping完毕")
}
