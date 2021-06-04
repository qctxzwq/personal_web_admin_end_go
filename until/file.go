package until

import (
	"admin/models"
	"github.com/astaxie/beego"
	"os"
	"time"
)

// 循环创建文件夹
func MkUploadImgDir() (filePath string, err error) {
	uploadPath := models.ProCfg.String("uploadpath")
	folderName := time.Now().Format("2006-01-02")
	filePath = uploadPath + folderName
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, 777)
		if err != nil {
			beego.Error("create folder err:", err)
			return "",err
		}
		return "",err
	}
	return filePath,nil
}

// 判断文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
