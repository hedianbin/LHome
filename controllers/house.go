package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/astaxie/beego/cache"
	"loveHome/utils"
	"time"
	"path"
)
type HouseInfo struct {
	Area_id    string   `json:"area_id"`    //归属地的区域编号
	Title      string   `json:"title"`      //房屋标题
	Price      string   `json:"price"`      //单价,单位:分
	Address    string   `json:"address"`    //地址
	Room_count string   `json:"room_count"` //房间数目
	Acreage    string   `json:"acreage"`    //房屋总面积
	Unit       string   `json:"unit"`       //房屋单元,如 几室几厅
	Capacity   string   `json:"capacity"`   //房屋容纳的总人数
	Beds       string   `json:"beds"`       //房屋床铺的配置
	Deposit    string   `json:"deposit"`    //押金
	Min_days   string   `json:"min_days"`   //最好入住的天数
	Max_days   string   `json:"max_days"`   //最多入住的天数 0表示不限制
	Facilities []string `json:"facility"`   //房屋设施
}

type HouseController struct {
	beego.Controller
}

func (this *HouseController) RetData(resp map[string]interface{})  {
	this.Data["json"] = resp
	this.ServeJSON()
}
//请求当前用户已发布房源
func (this *HouseController) GetHouseData()  {
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)
	//1.从session获取用户的user_id
	user_id:=this.GetSession("user_id")

	if user_id==nil{
		resp["errno"]=models.RECODE_SESSIONERR
		resp["errmsg"]=models.RecodeText(models.RECODE_SESSIONERR)
		return
	}
	//2.从数据库中拿到user_id对应的house数据
	//这里先拿到house结构体对象
	/*
	这里需要注意，因为我们需要查询的是user_id所有的房屋信息，这个用户可能会有多套房，所以我们存房屋信息的结构体要用数组
	*/
	//select * from house where user.id=user_id
	//将house相关联的User和Area一并查询
	houses:=[]models.House{} //必须用数组
	o:=orm.NewOrm()
	//查询house表
	qs:=o.QueryTable("house")
	//查询user_id=user_id的人的all房子存在houses数组中,将house相关联的User和Area一并查询
	num,err:=qs.Filter("user__id",user_id.(int)).RelatedSel().All(&houses)
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
	//遍历所有房源。并添加到数组中
	var houses_rep []interface{}
	for _,houseinfo:=range houses{
		fmt.Printf("house.user = %+v\n", houseinfo.User)
		fmt.Printf("house.area = %+v\n", houseinfo.Area)
		houses_rep=append(houses_rep,houseinfo.To_house_info())
	}
	fmt.Printf("houses_rep = %+v\n", houses_rep)


	//3.返回打包好的json数据
	//创建一个map用来存房源数据
	respData:=make(map[string]interface{})
	//将数据库里查到的所有房子数组存到这个map中
	respData["houses"]=houses_rep

	//将这个map再传到data里，返回json
	resp["data"]=respData

}
//发布房源信息
func (this *HouseController) PostHouseData()  {
	//用来存json数据的
	resp:=make(map[string]interface{})
	defer this.RetData(resp)
	//1.解析用户发过来的房源数据，得到房源信息
	//先创建一个结构体用来放用户发过来的数据
	var reqData HouseInfo
	if err:=json.Unmarshal(this.Ctx.Input.RequestBody,&reqData);err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//2.判断数据的合法性
	fmt.Printf("%+v\n",reqData)

	if &reqData==nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3.把房源数据插入到house结构体中
	house:=models.House{}
	house.Title=reqData.Title
	house.Price,_=strconv.Atoi(reqData.Price)
	house.Price=house.Price*100
	house.Address=reqData.Address
	house.Room_count,_=strconv.Atoi(reqData.Room_count)
	house.Acreage,_=strconv.Atoi(reqData.Acreage)
	house.Unit=reqData.Unit
	house.Beds=reqData.Beds
	house.Capacity,_=strconv.Atoi(reqData.Capacity)
	house.Deposit,_=strconv.Atoi(reqData.Deposit)
	house.Deposit=house.Deposit*100
	house.Max_days,_=strconv.Atoi(reqData.Max_days)
	house.Min_days,_=strconv.Atoi(reqData.Min_days)
	//获取用户的id,通过GetSession方式
	user:=models.User{Id:this.GetSession("user_id").(int)}
	house.User=&user

	//4.处理Area城区
	//把取到的area_id转成int
	area_id,_:=strconv.Atoi(reqData.Area_id)
	//把area_id赋值到结构体Id字段中
	area:=models.Area{Id:area_id}
	//再把Area结构体数据赋值给house结构体中的Area
	//把结构体赋值必须用取地址符&
	//这是一对多操作，一个房子只能在一个地区，一个地区可以有多个房子
	house.Area=&area

	//5.获取到house_id
	//创建一个orm对象
	o:=orm.NewOrm()
	//把部分house数据插入到数据库中,得到house_id
	house_id,err:=o.Insert(&house)
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Debug("house insert id =", house_id, " succ!")

	//5.多对多m2m插入,将facilities 一起关联插入到表中
	//定义一个设施的结构体数组，先把用户选的多个设施获取到
	facilitles:=[]models.Facility{}

	//遍历用户发来的设施列表，取出fid.
	for _,fid:=range reqData.Facilities{
		f_id,_:=strconv.Atoi(fid) //把string转成int
		fac:=models.Facility{Id:f_id} //更新每个设备的id
		facilitles=append(facilitles,fac) //将每个设备id追加成设施数组
	}


	//注意，只要house里有house_id后才能用QueryM2M，第一个参数是需要修改的哪个表，我这次要改house表，首先house表里一定要有一个house.Id，然后house.Id没有关联的设施信息，第二个参数为要修改的数据。
	//这句的意思其实就是将房屋设施数据插放到house结构体中的Facilities字段所关联的表的字段中
	//看下面Facility关联着House，rel(m2m)多对多关系。自然而然的就会将数据插入到关联表中。而这个关联表就是facility_houses
	/*
	type Facility struct {
		Id 		int 		`json:"fid"`			//设施编号
		Name 	string		`orm:"size(32)"`		//设施名字
		Houses  [] *House	`orm:"rel(m2m)"`		//都有哪些房屋有此设施
	}
	*/
	m2m:=o.QueryM2M(&house,"Facilities")
	//得到m2m对象后，我们就可以把刚才获取到的用户设施数组facilitles加到facility_houses中了
	num,errM2M:=m2m.Add(facilitles)
	if errM2M!=nil||num==0{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Debug("house m2m facility insert num =", num, " succ!")

	//6.返回json和house_id,有id返回说明插入成功，0就说明没成功
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	//创建一个map用来存house_id
	respData:=make(map[string]interface{})
	respData["house_id"]=house_id
	//把house_id的map存到data中，再打包成json
	resp["data"]=respData

}
//请求房源详细信息
func (this *HouseController) GetDetailHouseData()  {

	//用来存json数据的
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	//1.从session获取user_id
	user_id:=this.GetSession("user_id")

	//2.从请求的url中得到房屋id
	//Param中的id值可以随便换，但要是router中的对应
	house_id:=this.Ctx.Input.Param(":id")
	//转换一下interface{}转成int
	h_id,err:=strconv.Atoi(house_id)
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//3.从redis缓存获取当前房屋的数据,如果有该房屋，则直接返回正确的json
	redis_config_map:=map[string]string{
		"key":"lovehome",
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	redis_config,_:=json.Marshal(redis_config_map)
	cache_conn, err := cache.NewCache("redis", string(redis_config))
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//先把house有的东西返回给house的json
	respData:=make(map[string]interface{})
	//设置一个变量，每个房子插入redis不能一样，容易覆盖，所以用house_id做为key，比如lovehome:1,lovehome:2
	house_page_key:=house_id
	house_info_value:=cache_conn.Get(house_page_key)
	if house_info_value!=nil{
		beego.Debug("======= get house info desc  from CACHE!!! ========")
		//返回json的user_id
		respData["user_id"]=user_id
		//返回json的house信息
		house_info:=make(map[string]interface{})
		//解码json并存到house_info里
		json.Unmarshal(house_info_value.([]byte),&house_info)
		//将house_info的map返回json的house给前端
		respData["house"]=house_info
		resp["data"]=respData
		return
	}
	//4.如果缓存没有房屋数据,那么从数据库中获取数据,再存入缓存中,然后返回给前端
	o:=orm.NewOrm()
	// --- 载入关系查询 -----
	house:=models.House{Id:h_id}
	//把房子信息读出来
	if err:= o.Read(&house);err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//5.关联查询area,user,images,facilities等表
	o.LoadRelated(&house,"Area")
	o.LoadRelated(&house,"User")
	o.LoadRelated(&house,"Images")
	o.LoadRelated(&house,"Facilities")


	//6.将房屋详细的json数据存放redis缓存中
	house_info_value,_=json.Marshal(house.To_one_house_desc())
	cache_conn.Put(house_page_key,house_info_value,3600*time.Second)

	//7.返回json数据给前端。
	respData["house"]=house.To_one_house_desc()
	respData["user_id"]=user_id
	resp["data"]=respData
}
//上传房源图片
func (this *HouseController) UploadHouseImage()  {
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	//1.从用户请求中获取到图片数据
	fileData,hd,err:=this.GetFile("house_image")
	defer fileData.Close() //获取完后等程序执行完后关掉连接
	//beego.Info("========",fileData,hd,err)
	//没拿到图片
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
	//2.将用户二进制数据存到fdfs中。得到fileid
	suffix:=path.Ext(hd.Filename)
	//判断上传文件的合法性
	if suffix!=".jpg"&&suffix!=".png"&&suffix!=".gif"&&suffix!=".jpeg"{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//去掉.
	suffixStr:=suffix[1:]
	//创建hd.Size大小的[]byte数组用来存放fileData.Read读出来的[]byte数据
	fileBuffer:=make([]byte,hd.Size)
	//读出的数据存到[]byte数组中
	_,err=fileData.Read(fileBuffer)
	if err!=nil{
		resp["errno"]=models.RECODE_IOERR
		resp["errmsg"]=models.RecodeText(models.RECODE_IOERR)
		return
	}
	//将图片上传到fdfs获取到fileid
	uploadResponse,err:=UploadByBuffer(fileBuffer,suffixStr)
	//3.从请求的url中获得house_id
	house_id:=this.Ctx.Input.Param(":id")
	//4.查看该房屋的index_image_url主图是否为空
	house:=models.House{} //打开house结构体
	//house结构体拿到houseid数据
	house.Id,_=strconv.Atoi(house_id)
	o:=orm.NewOrm() //创建orm
	errRead:=o.Read(&house) //读取house数据库where user.id
	if errRead!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//查询index_image_url是否为空
	//为空则更将fileid路径赋值给index_image_url
	if house.Index_image_url==""{
		house.Index_image_url=uploadResponse.RemoteFileId
	}
	//5.主图不为空，将该图片的fileid字段追加（关联查询）到houseimage字段中插入到house_image表中,并拿到了HouseImage，里面也有数据了
	//HouseImage功能就是如果主图有了，就追加其它图片的。
	house_image:=models.HouseImage{House:&house,Url:uploadResponse.RemoteFileId}
	//将house_image和house相关联,往house.Images里追加附加图片，可以追加多个
	house.Images=append(house.Images,&house_image)//向把HouseImage对象的数据添加到house.Images
	//将house_image入库，插入到house_image表中
	if _,err:=o.Insert(&house_image);err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//将house更新入库，插入到house中
	if _,err:=o.Update(&house);err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//6.拼接完整域名url_fileid
	respData:=make(map[string]string)
	respData["url"]=utils.AddDomain2Url(uploadResponse.RemoteFileId)

	//7.返回给前端json
	resp["data"]=respData
}
//请求首页房源
func (this *HouseController) GetHouseIndex() {
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)
	var respData []interface{}
	beego.Debug("Index Houses....")

	//1 从缓存服务器中请求 "home_page_data" 字段,如果有值就直接返回
	//先从缓存中获取房屋数据,将缓存数据返回前端即可
	//连接redis需要的参数信息
	redis_config_map:=map[string]string{
		"key":"lovehome",
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	//把参数信息转成json格式
	redis_config,_:=json.Marshal(redis_config_map)
	//连接redis
	cache_conn,err:=cache.NewCache("redis",string(redis_config))
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//设置key
	house_page_key:="home_page_data"
	//上传房源数据到指定key中
	house_page_value:=cache_conn.Get(house_page_key)
	//返回给前端json数据
	if house_page_value!=nil{
		beego.Debug("======= get house page info  from CACHE!!! ========")
		json.Unmarshal(house_page_value.([]byte),&respData)
		resp["data"]=respData
		return
	}

	//2 如果缓存没有,需要从数据库中查询到房屋列表
	//取出house对象
	houses:=[]models.House{}
	o:=orm.NewOrm()
	//查询数据库中所有房子信息
	if _,err:=o.QueryTable("house").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses);err==nil{
		//循环遍历这些房子及关联表查询
		for _,house:=range houses{
			//o.LoadRelated(&house, "Area")
			//o.LoadRelated(&house, "User")
			//o.LoadRelated(&house, "Images")
			//o.LoadRelated(&house, "Facilities")
			//用下面方法查到的部分房子信息追加到respData数组中
			respData=append(respData,house.To_house_info())
		}
	}
	//将data存入缓存中
	house_page_value,_=json.Marshal(respData)
	cache_conn.Put(house_page_key,house_page_value,3600*time.Second)

	//返回前端data
	resp["data"]=respData
	return
}

func (this *HouseController) GetHouseSearchData()  {
	resp:=make(map[string]interface{})
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)
	respData:=make(map[string]interface{})
	defer this.RetData(resp)
	//1.获取用户发来的参数，aid,sd,ed,sk,p
	var aid int
	this.Ctx.Input.Bind(&aid,"aid")
	var sd string
	this.Ctx.Input.Bind(&sd,"sd")
	var ed string
	this.Ctx.Input.Bind(&ed,"ed")
	var sk string
	this.Ctx.Input.Bind(&sk,"sk")
	var page int
	this.Ctx.Input.Bind(&page,"p")

	//fmt.Printf("aid = %d,sd = %s,ed =%s,sk =%s,p =%d,==============\n",aid,sd,ed,sk,page)
	//2.检验开始时间一定要早于结束时间
	//将日期转成指定格式
	start_time,_:=time.Parse("2006-01-02 15:04:05",sd+" 00:00:00")
	end_time,_:=time.Parse("2006-01-02 15:04:05",ed+" 00:00:00")
	if end_time.Before(start_time){ //如果end在start之前,返回错误信息
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="结束时间必须在开始时间之前"
		return
	}
	//fmt.Printf("##############start_date_time = %v,end_date_time = %v",start_time,end_time)
	//3.判断p的合法性，一定要大于0的整数
	if page<=0{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]="页数不能小于或等于0"
		return
	}
	//4.尝试从缓存中获取数据，返回查询结果json
	//定义一个key,注意这个存入redis中的key值拼接字符串，一定要用strconv.Itoa()转换，不要用string(),否则会出现\x01的效果,读取不了
	house_search_key:="house_search_"+strconv.Itoa(aid)
	//配置redis连接信息
	redis_config_map:=map[string]string{
		"key":"lovehome",
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	//转成json
	redis_config,_:=json.Marshal(redis_config_map)
	//连接redis
	cache_conn, err := cache.NewCache("redis", string(redis_config))
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return
	}
	//从redis中拿到数据
	house_search_info:=cache_conn.Get(house_search_key)

	if house_search_info!=nil{
		beego.Debug("======= get house_search_info  from CACHE!!! ========")
		//存解码后的数据
		house_info:=[]map[string]interface{}{}
		//解码json数据
		json.Unmarshal(house_search_info.([]byte),&house_info)
		//把解码后的数据打包成json传给前端
		respData["houses"]=house_info
		respData["total_page"]=10
		respData["current_page"]=1
		resp["data"]=respData
		return
	}
	//5.如果缓存中没有数据，从数据库中查询
	//（此处过于复杂，可以暂时以发布时间顺序查询）
	//指定查询的表
	houses:=[]models.House{}
	o:=orm.NewOrm()
	//查询house表
	qs:=o.QueryTable("house")
	//查询指定城区的所有房源，按发布时间降序排列
	num,err:=qs.Filter("area_id",aid).OrderBy("-ctime").All(&houses)
	if err!=nil{
		resp["errno"]=models.RECODE_DBERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DBERR)
		return
	}
	//求出所有分页
	total_page:=int(num)/models.HOUSE_LIST_PAGE_CAPACITY+1
	//起始页数
	house_page:=1
	//用来存遍历到的房屋数据
	var house_list []interface{}
	//遍历出上面查到的房屋数据，加到数组house_list中
	for _,house:=range houses  {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Images")
		o.LoadRelated(&house, "Facilities")
		house_list=append(house_list,house.To_house_info())
	}
	//拿到了houst_list数据
	fmt.Println("========house_list======",house_list)
	//6.将查询条件存储到缓存
	houst_search_list,_:=json.Marshal(house_list)
	cache_conn.Put(house_search_key,houst_search_list,3600*time.Second)

	//7.返回查询结果json数据给前端
	respData["houses"]=house_list
	respData["total_page"]=total_page
	respData["current_page"]=house_page
	resp["data"]=respData
	return
}