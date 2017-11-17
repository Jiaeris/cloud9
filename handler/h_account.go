package handler

import (
	"cloud9/models"
	"cloud9/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
	"cloud9/utils"
	"strconv"
	"time"
)

//登录
func Login(c *gin.Context) {
	account := c.PostForm("account")
	md5password := c.PostForm("md5password")
	passKey := c.PostForm("salt")

	//验证
	va := &models.VerifyAccount{
		Account:     account,
		Md5Password: md5password,
		Salt:        passKey,
	}
	msgCode := va.Verify()
	if msgCode != msg.OK {
		c.JSON(http.StatusOK, msg.ResponseMessage{
			Code: msgCode,
			Data: "account validation failed .",
		})
		return
	}

	//验证成功 生成token
	tokenStr, err := utils.TokenCreate(account)
	if err != nil {
		panic(err)
	}

	//更新登录时间
	am := &models.Account{
		Name: account,
	}
	am.UpdateLastTime()

	//设置cookie
	c.Header("Access-Control-Allow-Origin", "*")

	cookie := &http.Cookie{
		Name:  "cloud",
		Value: tokenStr,

		Expires: time.Now().Add(6.5 * 3600 * time.Second),
	}

	http.SetCookie(c.Writer, cookie)

	//c.SetCookie("cloud", tokenStr, int(time.Now().Unix()+7*3600), "/", "127.0.0.1", false, false)

	//返回成功
	c.JSON(http.StatusOK, msg.ResponseMessage{
		Code: msgCode,
		Data: tokenStr,
	})
}

//登出
func Logout(c *gin.Context) {

}

//未登录提示
func UnLogin(c *gin.Context) {
	c.JSON(http.StatusOK, &msg.ResponseMessage{
		Code: msg.NotLoggedIn,
		Data: "user not login",
	})
}

//获取个人资料
func Detail(c *gin.Context) {
	value, isExist := c.Get("name")
	if isExist {
		account := &models.Account{
			Name: value.(string),
		}
		c.JSON(http.StatusOK, account.GetDetails())
		return
	}
	UnLogin(c)
}

//更新个人信息
func DetailUpdate(c *gin.Context) {
	value, isExist := c.Get("name")
	if !isExist {
		UnLogin(c)
		return
	}
	account := &models.Account{
		Name: value.(string),
	}
	ud := models.UserDetail{}

	nickname := c.PostForm("nickname")
	sex := c.PostForm("sex")
	birthday := c.PostForm("birthday")
	location := c.PostForm("location")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	about := c.PostForm("about")

	var fields []string
	if nickname != "" {
		ud.Nickname = nickname
		fields = append(fields, "nickname")
	}
	if sex != "" {
		sexInt, _ := strconv.Atoi(sex)
		ud.Sex = sexInt
		fields = append(fields, "sex")
	}
	if birthday != "" {
		birthdayObj := utils.String2Time(birthday)
		ud.Birthday = birthdayObj
		fields = append(fields, "birthday")
	}
	if location != "" {
		ud.Location = location
		fields = append(fields, "location")
	}
	if phone != "" {
		ud.Phone = phone
		fields = append(fields, "phone")
	}
	if email != "" {
		ud.Email = email
		fields = append(fields, "email")
	}
	if about != "" {
		ud.About = about
		fields = append(fields, "about")
	}

	account.UpdateDetails(ud, fields)
}

func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	account := models.Account{
		Name:     name,
		Password: password,
	}
	errorCode := account.CreateAccount()
	if errorCode != msg.OK {
		switch errorCode {
		case msg.AccountExist:
			c.JSON(http.StatusOK, msg.ResponseMessage{
				Code: errorCode,
				Data: "account already exist",
			})
			break
		}
		return
	}
	c.JSON(http.StatusOK, msg.ResponseMessage{
		Code: errorCode,
		Data: "create account successful",
	})
}
