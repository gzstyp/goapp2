package controller

import (
	"com.fwtai/app2/common/toolClent"
	"github.com/gin-gonic/gin"
	"log"
	"path"
)

func UploadSingle(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		toolClent.Response199Msg(context, "读取文件失败")
		return
	}
	//dst := fmt.Sprintf("./%s", file.Filename)//也是可行的
	dst := path.Join("./", file.Filename)
	er := context.SaveUploadedFile(file, dst)
	if er != nil {
		toolClent.Response199Msg(context, "保存文件失败")
		return
	}
	toolClent.Response200Msg(context, "上传成功")
}

func UploadMulti(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		toolClent.Response199Msg(context, "读取文件失败")
		return
	}
	files := form.File //json对象的文件,{文件名作为key,二进制数据文件作为value}
	//bl := false
	for idx, f := range files {
		fs := form.File[idx] // 数组文件,下标0开始
		log.Println(f)       //文件的个数,它和上行是同样的效果*/
		log.Println(fs)

	}

	//files := form.File["file"]
	/*bl := false
	for index,file := range files{
		dst := fmt.Sprintf("C:/%s_%d",file.Filename,index)
		er := context.SaveUploadedFile(file, dst)
		if er != nil {
			toolClent.Response199Msg(context,"读取文件失败")
			bl = true
			break
		}
	}
	if bl {
		toolClent.Response199Msg(context,"上传文件失败")
		return
	}
	toolClent.Response200Msg(context,fmt.Sprintf("%d上传文件成功,共计:"+string(len(files))))*/
}
