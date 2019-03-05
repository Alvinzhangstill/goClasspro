package controllers

import (
	"class/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"path"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*  1.插入数据
		//1.有ORM对象
		o := orm.NewOrm()
		//2.有一个要插入数据的结构体对象
		user := models.User{}
		//3.对结构体对象赋值
		user.Name = "alvin"
		user.Pwd = "0313"
		//插入
		_,err := o.Insert(&user)
		if err != nil {
			beego.Info("插入失败")
			return
		}
	*/

	/*  2.查询相关
		//1.有ORM对象
		o := orm.NewOrm()

		//2.查询的对象
		user := models.User{}
		//3.指定查询对象字段

		////根据ID查询
		//user.Id = 2
		////4.查询
		//err := o.Read(&user)


		//根据字段来查询
		user.Name = "alvin"
		err := o.Read(&user,"Name")

		if err != nil{
			beego.Info("查询失败",err)
			return
		}

		beego.Info("查询成功",user)
	*/

	/*   3.更新相关
		//1.要有ORM对象
		o := orm.NewOrm()

		//2.需要更新的结构体对象
		user := models.User{}
		//3.查到需要更新的数据
		user.Id = 1
		err := o.Read(&user)

		//4.给数据重新赋值
		if err == nil {
			user.Name = "Alvin"
			user.Pwd = "0313zj"
			//5.更新
			_,err = o.Update(&user)
			if err != nil{
				beego.Info("更新失败",err)
				return
			}else{
				beego.Info("更新成功!")
			}
		}
	*/

	/*   4.删除相关
		//1.有ORM对象
		o := orm.NewOrm()
		//2.删除的对象
		user := models.User{}
		//3.指定删除哪一条数据
		user.Id = 1
		//4.删除
		_,err := o.Delete(&user)
		if err != nil{
			beego.Info("删除失败",err)
			return
		}
	*/

	c.TplName = "register.html"

}

func (c *MainController) Post() {
	//1.拿到数据
	username := c.GetString("userName")
	pwd := c.GetString("password")
	//2.对数据进行校验
	if username == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/reg", 302)
		return
	}
	//3.插入数据库
	o := orm.NewOrm()

	user := models.User{}

	user.Name = username
	user.Pwd = pwd

	_, err := o.Insert(&user)

	if err != nil {
		beego.Info("插入数据库出错", err)
		c.Redirect("/reg", 302)
		return
	}
	//4.注册成功 —> 返回登陆界面
	//c.TplName = "login.html"    //可以传递数据
	c.Redirect("/login", 302) //不可以传递数据，但是速度快
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

func (c *MainController) HandleLogin() {
	//1.拿到数据
	username := c.GetString("userName")
	pwd := c.GetString("password")
	//2.判断数据是否合法
	if username == "" || pwd == "" {
		beego.Info("用户名或密码不能为空! ")
		c.TplName = "login.html"
		return
	}

	//3.查询账号密码是否正确
	o := orm.NewOrm()

	user := models.User{}
	user.Name = username

	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户不存在", err)
		c.TplName = "login.html"
		return
	} else {
		fmt.Println("登陆成功")
	}

	//4.跳转
	//c.Ctx.WriteString("个人主页")
	//c.TplName = "index.html"
	c.Redirect("/index", 302)

}

//显示首页内容
func (c *MainController) ShowIndex() {

	o := orm.NewOrm()
	var articles []models.Article //结构体数组
	_, err := o.QueryTable("Article").All(&articles)
	if err != nil {
		beego.Info("查询所有文章信息出错")
	}
	beego.Info(articles)
	c.Data["articles"] = articles

	c.TplName = "index.html"
	//c.Data[""]

}

func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}

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
	beego.Info("fileext:", fileext)

	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("上传数据格式错误")
		return
	}
	//2.限制文件大小
	if h.Size > 3132313212333333345 { // 1B为单位
		beego.Info("上传文件过大")
		return
	} else {
		beego.Info("文件大小合适")
	}
	//3.防止文件名重复
	filename := GetRandomString(10) + fileext
	beego.Info("filename:", filename)

	if err != nil {
		beego.Info("文件上传失败")
		return
	} else {
		//这里不是很懂
		beego.Info("savetofile begin")
		c.SaveToFile("uploadname", "./static/img/"+filename)
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

func (c *MainController) ShowContent() {
	//1.获取id
	id, err := c.GetInt("id")
	beego.Info("id is :", id)
	if err != nil {
		beego.Info("获取文章ID错误", err)
		return
	}
	//2.查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询数据失败", err)
		return
	}
	beego.Info("arti:", arti)

	//3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "content.html"
}

func (c *MainController) ShowUpdate() {
	//1.拿到要更新的页面Id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取id失败")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	//3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "update.html"
}

func (c *MainController) HandleUpdate() {
	//1.拿到数据
	id, _ := c.GetInt("id")
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
		} else {
			beego.Info("文件大小合适")
		}
		//3.防止文件名重复
		filename = GetRandomString(10) + fileext
		beego.Info("filename:", filename)
		//这里不是很懂
		beego.Info("savetofile begin")
		c.SaveToFile("uploadname", "./static/img/"+filename)
	}
	if artiName == "" || artiContent == "" {
		beego.Info("更新数据获取失败")
		return
	}
	//3.更新操作
	o := orm.NewOrm()
	//需要更新的结构体对象
	arti := models.Article{Id: id}
	//查到需要更新的数据
	err = o.Read(&arti)
	if err == nil {
		arti.Aname = artiName
		arti.Acontent = artiContent
		beego.Info("filename 修改后:", filename)
		arti.Aimg = "/static/img/" + filename

		_, err = o.Update(&arti, "Aname", "Acontent", "Aimg")
		if err != nil {
			beego.Info("更新数据显示错误")
			return
		} else {
			beego.Info("更新成功!!!!")
		}
	}

	//4.返回列表页面
	c.Redirect("/index", 302)

}

func (c *MainController) HandleDelete(){
	id,err := c.GetInt("id")
	if err != nil{
		beego.Info("获取id失败")
		return
	}

	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询失败")
		return
	}
	_,err = o.Delete(&arti)
	if err != nil{
		beego.Info("删除成功!")
	}

	//返回文章列表页面
	c.Redirect("/index",302)



}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
