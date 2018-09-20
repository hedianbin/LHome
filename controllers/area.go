package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
	"encoding/json"
	"time"
	"loveHome/utils"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp:=make(map[string]interface{})

	//每次return都会执行一次这行返回json数据
	defer c.RetData(resp)
	//从redis拿数据
	redis_config_map:=map[string]string{
		"key":"lovehome",
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	redis_config,_:=json.Marshal(redis_config_map)
	cache_conn, err := cache.NewCache("redis", string(redis_config))

	if err!=nil{
		beego.Error("cache_conn err = ",err)
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		return
	}
	areaData:=cache_conn.Get("area")
	if areaData!=nil{
		beego.Info("get data from cache=============")
		//查询成功返回数据
		resp["errno"]=models.RECODE_OK
		resp["errmsg"]=models.RecodeText(models.RECODE_OK)
		//把从redis中取来的数据必须先解码才能在前台显示。
		var areas_info interface{}
		//解码数据并存到areas_info中
		json.Unmarshal(areaData.([]byte),&areas_info)
		resp["data"]=areas_info
		return
	}

	//第一步，从数据库中拿到数据。
	//声明一个数组用来存从数据库中查询到的所有城区数据。用数组
	o:=orm.NewOrm()
	var areas []models.Area

	//查询area表里的全部数据，存到areas缓存数组中，返回的是int64查询的条数，error错误信息
	num,err:=o.QueryTable("area").All(&areas)
	//如果没查询到返回的数据
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num==0{
		resp["errno"]=models.RECODE_NODATA
		resp["errmsg"]=models.RecodeText(models.RECODE_NODATA)
		return
	}

	resp["data"]=areas //把上面查询到的数据传给data

	//把取到的数据转换成json格式存入缓存
	json_str,err:=json.Marshal(areas)
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	json_err:=cache_conn.Put("area",json_str,time.Second*3600)
	if json_err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//查询成功返回数据
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	//第二步，把拿到的数据打包成json返回给前端。
	//上面已经defer c.RetData(resp)了，所以就不用写了
	beego.Info("query data succee ,resp = ",resp,"num = ",num)

}
