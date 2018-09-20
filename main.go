package main

import (
	_ "loveHome/routers"
	"github.com/astaxie/beego"
	_ "loveHome/models"
)

func main() {

	//groupname,fileid,err:=controllers.TestUploadByFilename("static/images/home01.jpg")
	//beego.Info("groupname = ",groupname,"fileid = ",fileid,"err = ",err)
	beego.Run()
	//beego.SetStaticPath("group1/M00/","/home/fastdfs/file/data/")

}



