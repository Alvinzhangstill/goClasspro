func (c *MainController) HandleAdd() {
//1.获取前端来的数据
artiName := c.GetString("articleName")
artiContent := c.GetString("content")
//返回三个参数：文件的二进制流、文件头、错误信息
f, h, err := c.GetFile("uploadname")
defer f.Close()

//1.要限定格式
//获取文件后缀
fileext := path.Ext(h.Filename)
//fmt.Printf("%T , %v", fileext, fileext)
beego.Info("fileext:",fileext)

if fileext != ".jpg" && fileext != ".png" {
beego.Info("上传数据格式错误")
return
}
//2.限制文件大小
if h.Size > 3132313212333333345 { // 1B为单位
beego.Info("上传文件过大")
return
}else{
beego.Info("文件大小合适")
}
//3.防止文件名重复
filename := GetRandomString(10) + fileext
beego.Info("filename:",filename)

if err != nil {
beego.Info("文件上传失败")
return
} else {
//这里不是很懂
beego.Info("savetofile begin")
c.SaveToFile("uploadname", "./static/img/"+ filename)
}
beego.Info(artiName, artiContent)

//2.判断数据是否合法
if artiName == "" || artiContent == "" {
beego.Info("请输入文章详情或内容")
return
}

//3.插入数据
o := orm.NewOrm()
arti := models.Article{}
arti.Aname = artiName
arti.Acontent = artiContent
arti.Aimg = "/static/img/" + filename

_, err = o.Insert(&arti)
if err != nil {
beego.Info("插入数据库失败")
return
}
//4.返回文章界面
c.Redirect("/index", 302)

}