package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
	"path"
	"loveHome/utils"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}
//用户注册
func (this *UserController) Reg() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	/*
	这是传过来的数据
	mobile: "111"
	password: "111"
	sms_code: "111"
	*/
	beego.Info(`resp["mobile"] = `,resp["mobile"])
	beego.Info(`resp["password"] = `,resp["password"])
	beego.Info(`resp["sms_code"] =`,resp["sms_code"])
	//插入数据库
	o:=orm.NewOrm()
	user:=models.User{} //啥数据不用传
	user.Password_hash=resp["password"].(string)
	user.Name=resp["mobile"].(string)
	user.Mobile=resp["mobile"].(string)
	//插入
	id,err:=o.Insert(&user)
	if err!=nil{
		resp["errno"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	beego.Info("reg succee ,id = ",id)
	//设置一个session,用来登录后显示用户名
	this.SetSession("name",user.Name)
	this.SetSession("user_id",user.Id)
	this.SetSession("mobile",user.Mobile)

}
//上传头像
func (this *UserController) PostAvatar()  {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)

	//1.用户上传头像，发起请求给后台，POST

	//2.我们得到数据，就是this.GetFile("avatar"),返回fileData是数据,hd为form操作句柄
	fileData,hd,err:=this.GetFile("avatar")
	defer fileData.Close()
	beego.Info("========",fileData,hd,err)
	if fileData==nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//3.通过某方法得到文件后缀。
	/*
	一般我们通过方法headerData.Filename，能得到我们上传的文件的名字比如2.jpg,然后我们需要去得到这个文件的后缀
	一般如果用户上传的文件是a.jpg.avi.mp3这样，我们用go原生的字符串截取，只能得到.jpg这样的用户名。但是beego提供
	了一个专门用来截取后缀的现成的函数path.Ext()
	*/

	suffix:=path.Ext(hd.Filename) //获取到的是.jpg，有点。下面把.去掉
	//判断是否为图片
	if suffix!=".jpg"&&suffix!=".png"&&suffix!=".gif"&&suffix!=".jpeg"{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//去掉.jpg前面的.，变成jpg
	suffixStr:=suffix[1:]
	//创建hd.Size大小的[]byte数组用来存放fileData.Read读出来的[]byte数据
	fileBuffer:=make([]byte,hd.Size)
	//读出的数据存到[]byte数组中
	_,err=fileData.Read(fileBuffer)
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}

	//4.得到的文件存储在fastDFS上，得到fileid-url路径，
	uploadResponse,err:=UploadByBuffer(fileBuffer,suffixStr)
	//RemoteFileId:=strings.Replace(uploadResponse.RemoteFileId,`\`,"/",-1)

	//5.通过session得到user_id,因为我们在登录和注册的时候设置了user_id。只需通过session拿到user_id
	user_id:=this.GetSession("user_id")
	//6.然后把fileid-url存到mysql数据库的user表对应的字段中。
	//创建一个user对象，用来往结构体中放数据
	var user models.User
	//获取数据库操作句柄
	o:=orm.NewOrm()
	//查询user表
	qs:=o.QueryTable("user")
	//查询id=user.ID的，放到&user结构体中
	qs.Filter("id",user_id).One(&user)
	//把图片的远程路径存到结构体中user.Avatar_url用户图片路径中
	user.Avatar_url =uploadResponse.RemoteFileId
	//将user结构体更新到数据库中
	_,errUpdate := o.Update(&user)
	if errUpdate!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}

	//7.把fileid-url和服务器域名拼接成完整的url路径
	avatar_url:=make(map[string]string)
	avatar_url["avatar_url"]=utils.AddDomain2Url(uploadResponse.RemoteFileId)
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	resp["data"]=avatar_url

	//8.打包成json返回给前端。

}
//获取用户数据
func (this UserController) GetUserData() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//1.从session获取用户的user_id
	user_id:=this.GetSession("user_id")
	//name:=this.GetSession("name")
	//mobile:=this.GetSession("mobile")
	if user_id==nil{
		resp["errno"]=models.RECODE_SESSIONERR
		resp["errmsg"]=models.RecodeText(models.RECODE_SESSIONERR)
		return
	}
	//2.从数据库中拿到user_id对应的user数据
	user:=models.User{Id:user_id.(int)}
	o:=orm.NewOrm()

	err:=o.Read(&user)
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}

	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	user.Avatar_url=utils.AddDomain2Url(user.Avatar_url)
	resp["data"]=&user

}

//更新用户名
func (this UserController) UpdateName() {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//1.从session中得到user_id
	user_id:=this.GetSession("user_id")
	//2.拿到用户请求修改的name值
	//创建一个map,用来存用户请求的name和name的值
	UserName:=make(map[string]string)
	//获取表单数据，存到UserName中
	json.Unmarshal(this.Ctx.Input.RequestBody,&UserName)

	//3.更新数据库对应user_id的name值
	o:=orm.NewOrm()
	//设置查询条件id=user_id
	user:=models.User{Id:user_id.(int)}
	//根据查询条件读取用户信息
	if err:=o.Read(&user);err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//更新用户名
	user.Name=UserName["name"]
	//更新数据库
	if _,err:=o.Update(&user);err!=nil {
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//4.更新session中name字段
	this.SetSession("name",UserName["name"])
	//5.返回成功json
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	resp["data"]=UserName

}

//实名认证
func (this *UserController)PostRealName()  {
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//1.先从session中获取用户user_id
	user_id:=this.GetSession("user_id")
	//2.然后获取到用户发过来的ResponseBody数据
	//创建一个map,用来存用户请求的name和name的值
	RealName:=make(map[string]string)
	//获取表单数据，存到UserName中
	json.Unmarshal(this.Ctx.Input.RequestBody,&RealName)
	//3.检验一下数据是否合法

	//4.把数据更新到user表对应字段中
	o:=orm.NewOrm()
	//设置查询条件id=user_id
	user:=models.User{Id:user_id.(int)}
	//根据查询条件读取用户信息
	if err:=o.Read(&user);err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//更新用户名
	//id_card:"222424112345675433"
	//real_name:"何殿斌"
	user.Id_card=RealName["id_card"]
	user.Real_name=RealName["real_name"]
	//更新数据库
	if _,err:=o.Update(&user);err!=nil {
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//5.更新session中的user_id字段确保过期时间刷新
	this.SetSession("user_id",user_id)
	//6.返回成功的json信息
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)

}