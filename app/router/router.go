package router

import (
	"app/api"
	"app/middlerware"
	"app/utils"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	rr := gin.Default()
	r := rr.Group("app")

	{
		r.Use(middlerware.Cors())
		r.GET("login", api.Login)
		r.GET("login/getJWT", api.GetJwt)
	}

	us := r.Group("user")

	{
		us.GET("getUserInfoList", api.U.GetUserInfoList)
		us.GET("getUserInfoByid", api.U.FindUserInfo)
		us.Use(middlerware.JWTCheck())
		us.POST("updateUserInfo", api.U.UpdateUserInfo)
		us.DELETE("deleteUserInfo", api.U.DeleteUserInfo)
	}

	cs := r.Group("com")

	{

		cs.GET("getComInfoList", api.C.GetComInfoList)
		cs.GET("getComInfoByid", api.C.FindComInfo)
	}

	ds := r.Group("dis")

	{

		ds.GET("getDisInfoList", api.D.GetDisInfoList)
		ds.GET("getDisInfoByid", api.D.FindDisInfo)
		ds.Use(middlerware.JWTCheck())
		ds.POST("createDisInfo", api.D.CreateDisInfo)
		ds.DELETE("deleteDisInfo", api.D.DeleteDisInfo)
	}

	cms := r.Group("comment")

	{
		cms.GET("getCommentInfoList", api.CC.GetCommentInfoList)
		cms.GET("getCommentInfoByid", api.CC.FindCommentInfo)
		cms.Use(middlerware.JWTCheck())
		cms.POST("createCommentInfo", api.CC.CreateCommentInfo)
		cms.DELETE("deleteCommentInfo", api.CC.DeleteCommentInfo)
		cms.POST("createCommentSonInfo", api.CC.CreateCommentSonInfo)
		cms.DELETE("deleteCommentSonInfo", api.CC.DeleteCommentSonInfo)
	}

	upload := r.Group("upload")
	{
		upload.POST("file", utils.UploadFileUseAliyun)
	}

	return rr
}
