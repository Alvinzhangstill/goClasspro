func (c *MainController) HandleUpdate(){
//1.拿到数据
id,_ := c.GetInt("id")
artiName := c.GetString("articleName")
artiContent := c.GetString("content")
f, h, err := c.GetFile("uploadname")

var filename string
if err != nil {
beego.Info("文件上传失败")
return
} else {
defer f.Close()
//1.要限定格式
fileext := path.Ext(h.Filename)
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
filename = GetRandomString(10) + fileext
beego.Info("filename:",filename)
//这里不是很懂
beego.Info("savetofile begin")
c.SaveToFile("uploadname", "./static/img/"+ filename)
}
if artiName == "" || artiContent == ""{
beego.Info("更新数据获取失败")
return
}
//3.更新操作
o := orm.NewOrm()
//需要更新的结构体对象
arti := models.Article{Id:id}
//查到需要更新的数据
err = o.Read(&arti)
if err == nil{
arti.Aname = artiName
arti.Acontent = artiContent
beego.Info("filename 修改后:",filename)
arti.Aimg = "/static/img" + filename

_,err = o.Update(&arti,"Aname","Acontent","Aimg")
if err != nil{
beego.Info("更新数据显示错误")
return
}else{
beego.Info("更新成功!!!!")
}
}


//4.返回列表页面
c.Redirect("/index",302)

}
