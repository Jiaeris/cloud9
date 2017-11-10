package models

import (
	"cloud9/utils"
	"time"
	"cloud9/msg"
)

var (
	PermissionAdmin       = 7
	PermissionVipUser     = 4
	PermissionGeneralUser = 2
	PermissionGuestUser   = 1
)

type AccountManager interface {
	UpdateLastTime()
	GetDetails() interface{}
	CreateAccount() (errorCode int)
}

type Account struct {
	id         int64
	Name       string
	Password   string
	CreateTime time.Time
	LastTime   time.Time
}

type VerifyAccount struct {
	Account     string
	Md5Password string
	Salt        string
}

type UserDetail struct {
	Name     string    `json:"name,omitempty"`
	Nickname string    `json:"nickname,omitempty""`
	Sex      int       `json:"sex,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
	Location string    `json:"location,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Email    string    `json:"email,omitempty"`
	About    string    `json:"about,omitempty"`
}

func (a *Account) UpdateLastTime() {
	Cloud10Db.Exec("UPDATE user_table SET last_time = ? WHERE name = ?", time.Now(), a.Name)
}

func (va *VerifyAccount) Verify() (errorCode int) {
	rows, err := Cloud10Db.Query("SELECT password FROM user_table WHERE name = ?", va.Account)
	if err != nil {
		panic(err)
		return msg.AccountNotExist
	}
	var okMd5Password string
	for rows.Next() {
		var password string
		rows.Scan(&password)
		okMd5Password = utils.MD5(password + va.Salt)
	}
	rows.Close()
	if okMd5Password == va.Md5Password {
		return msg.OK
	}
	return msg.PasswordError
}

func (this *Account) GetDetails() interface{} {
	if len(this.Name) == 0 {
		return nil
	}
	rows, err := Cloud10Db.Query("SELECT b.nickname,b.sex,b.birthday,b.location,b.phone,b.email,b.about FROM user_detail_table b,user_table a WHERE a.detail_id = b.id AND name = ?", this.Name)
	if err != nil {
		return nil
	}
	var nickname string
	var sex int
	var birthday time.Time
	var location string
	var phone string
	var email string
	var about string
	for rows.Next() {
		rows.Scan(&nickname, &sex, &birthday, &location, &phone, &email, &about)
	}
	rows.Close()
	return &UserDetail{
		Nickname: nickname,
		Sex:      sex,
		Birthday: birthday,
		Location: location,
		Phone:    phone,
		Email:    email,
		About:    about,
	}
}

func (this *Account) CreateAccount() (errorCode int) {
	//检查账户是否存在
	rows, err := Cloud10Db.Query("SELECT COUNT(*) FROM user_table WHERE name = ?", this.Name)
	if err != nil {
		panic(err)
		return msg.AccountExist
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	rows.Close()
	if count > 0 {
		return msg.AccountExist
	}

	//创建账户
	ret, _ := Cloud10Db.Exec("INSERT INTO user_detail_table (nickname) VALUES (?)", "")
	userDetailId, err := ret.LastInsertId()

	ret, _ = Cloud10Db.Exec("INSERT INTO user_table (name,password,permission,create_time,detail_id) VALUES (?,?,?,?,?)",
		this.Name,
		this.Password,
		PermissionGeneralUser,
		time.Now(),
		userDetailId)

	return msg.OK
}

func (this *Account) UpdateDetails(ud UserDetail, fields []string) (errorCode int) {
	row := Cloud10Db.QueryRow("SELECT detail_id FROM user_table WHERE name = ?", this.Name)
	var detailId int
	row.Scan(&detailId)
	if !(detailId > 0) {
		return msg.AccountDetailsNotExist
	}

	var nickname string
	var sex int
	var birthday time.Time
	var location string
	var phone string
	var email string
	var about string
	dr := Cloud10Db.QueryRow("SELECT nickname,sex,birthday,location,phone,email,about FROM user_detail_table WHERE detail_id = ?", detailId)
	dr.Scan(&nickname, &sex, &birthday, &location, &phone, &email, &about)
	mDetail := &UserDetail{
		Nickname: nickname,
		Sex:      sex,
		Birthday: birthday,
		Location: location,
		Phone:    phone,
		Email:    email,
		About:    about,
	}

	for _, filed := range fields {
		switch filed {
		case "nickname":
			mDetail.Nickname = ud.Nickname
			break
		case "sex":
			mDetail.Sex = ud.Sex
			break
		case "birthday":
			mDetail.Birthday = ud.Birthday
			break
		case "location":
			mDetail.Location = ud.Location
			break
		case "phone":
			mDetail.Phone = ud.Phone
			break
		case "email":
			mDetail.Email = ud.Email
			break
		case "about":
			mDetail.About = ud.About
			break
		}
	}

	_, err := Cloud10Db.Exec("UPDATE user_detail_table SET nickname=?,sex=?,birthday=?,location=?,phone=?,email=?,about=? WHERE detail_id = ?",
		mDetail.Nickname,
		mDetail.Sex,
		mDetail.Birthday,
		mDetail.Location,
		mDetail.Phone,
		mDetail.Email,
		mDetail.About,
		detailId)
	if err != nil {
		panic(err)
	}
	return msg.OK
}
