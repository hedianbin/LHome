package controllers

import (
	"github.com/weilaihui/fdfs_client"
	"github.com/astaxie/beego"
	"loveHome/models"
)

func TestUploadByFilename(fileName string) (groupName string,FileId string,err error) {
	fdfsClient,errClient:=fdfs_client.NewFdfsClient("conf/client.conf")
	if errClient!=nil{
		beego.Info("New FdfsClient error %s",errClient.Error())
		return "","",errClient
	}
	uploadResponse,errUpload:=fdfsClient.UploadByFilename(fileName)
	if errUpload!=nil{
		beego.Info("Upload FdfsClient error %s",errUpload.Error())
		return "","",errUpload
	}
	return uploadResponse.GroupName,uploadResponse.RemoteFileId,nil
}


func UploadByBuffer(fileBuffer []byte,suffixStr string) (uploadResp *fdfs_client.UploadFileResponse,err error)  {
	resp:=make(map[string]interface{})
	//连接到fdfs服务器
	fdfsClient,err:=fdfs_client.NewFdfsClient("conf/client.conf")
	if err!=nil{
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return nil,err
	}

	//把文件上传到fdfs上
	uploadResponse, err := fdfsClient.UploadByBuffer(fileBuffer, suffixStr)
	if err != nil {
		beego.Error("TestUploadByBuffer error %s", err.Error())
		resp["errno"]=models.RECODE_REQERR
		resp["errmsg"]=models.RecodeText(models.RECODE_REQERR)
		return nil,err
	}

	beego.Info(uploadResponse.GroupName)
	beego.Info(uploadResponse.RemoteFileId)
	return uploadResponse,nil
}