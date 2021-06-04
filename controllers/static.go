package controllers

import (
	"admin/models"
	"admin/until"
	"github.com/astaxie/beego"
)

type StaticController struct {
	beego.Controller
}

// 上传图片
func (s *StaticController) UpdateImg() {
	f, h, err := s.GetFile("image")
	if err != nil {
		beego.Error("get image file err:", err)
		errRes := models.SystemError{
			Code:    models.GET_REGIIST_AVATARFILE_ERROR_CODE_,
			Message: models.SYSTEM_ERROR_MSG,
		}
		s.Data["json"] = errRes
		s.ServeJSON()
		return
	}
	defer f.Close()
	demain := models.ProCfg.String("demain")
	folderPath, err := until.MkUploadImgDir()
	if err != nil {
		beego.Error("create folder err:", err)
		errRes := models.SystemError{
			Code:    models.CREATE_FOLDER_ERROR_CODE,
			Message: models.SYSTEM_ERROR_MSG,
		}
		s.Data["json"] = errRes
		s.ServeJSON()
		return
	}

	path := folderPath + "/" + h.Filename

	err = s.SaveToFile("image", path)
	if err != nil {
		beego.Error("save img file err:", err)
		errRes := models.SystemError{
			Code:    models.UPLOAD_ERROR_CODE,
			Message: models.UPLOAD_FILE_ERROR_MSG,
		}
		s.Data["json"] = errRes
		s.ServeJSON()
		return
	}
	succRes := models.SuccessMsg{
		Code:    0,
		Data:    map[string]interface{}{"url": demain + path[1:]},
		Message: "上传成功",
	}
	s.Data["json"] = succRes
	s.ServeJSON()
	return
}
