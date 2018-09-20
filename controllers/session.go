package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionData() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	name:=this.GetSession("name")
	user:=models.User{}
	if name!=nil{
		user.Name=name.(string)
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		resp["data"]=user
	}else{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
	}
}

func (this *SessionController) DeleteSessionData() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("user_id")
	this.DelSession("name")
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
}

func (this *SessionController) Login() {
	//1.得到用户信息
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	//2.判断是否合法
	if resp["mobile"]==nil || resp["password"]==nil{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		return
	}
	//判断手机号是否为11位
	//if len(resp["mobile"].(string))!=11{
	//	resp["errno"]=models.RECODE_DATAERR
	//	resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
	//	return
	//}

	//3.与数据库匹配，判断帐号密码是否正确
	o:=orm.NewOrm()
	user:=models.User{Mobile:resp["mobile"].(string)}
	//查询user表
	qs:=o.QueryTable("user")
	//过滤只查询mobile==user.Name的，One(&user)返回数据到user结构体中，记得用取地址
	err:=qs.Filter("mobile",user.Mobile).One(&user)
	if err!=nil{
		beego.Info("o.Read(&user) err = ",err)
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		return
	}
	if user.Password_hash!=resp["password"]{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		return
	}
	//4.添加session
	this.SetSession("name",user.Name)
	this.SetSession("mobile",resp["mobile"])
	this.SetSession("user_id",user.Id)
	//5.返回json数据给前端
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
}
