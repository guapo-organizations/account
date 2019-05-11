package register

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/account-service/model"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/help"
	"time"
)

const (
	EMAIL_TOKEN_PREFIX = "email:token:"
	PHONE_TOKEN_PREFIX = "phone:token"
)

type RegisterService struct {
}

//email在缓存中的key
func (this RegisterService) GetEmailToken(email string) string {
	return EMAIL_TOKEN_PREFIX + email
}

//phone在缓存中的key
func (this RegisterService) GetPhoneToken(phone string) string {
	return PHONE_TOKEN_PREFIX + phone
}

//email存入缓存
func (this RegisterService) StoreEmail(email string) (bool, error) {
	redis_client := cache.GetRedisClient()

	_, err := redis_client.SetNX(this.GetEmailToken(email), email, 5*60*time.Second).Result()

	if err != nil {
		return false, err
	}

	return true, nil
}

//在缓存中获取email
func (this RegisterService) GetEmailByToken(token string) (email string, err error) {
	redis_client := cache.GetRedisClient()
	email, err = redis_client.Get(token).Result()

	if err == redis.Nil {
		email = ""
		err = fmt.Errorf("token:%s无法找到email", token)
	}

	return email, err
}

//邮箱注册
func (this RegisterService) RegisterAccountByEmail(email, passwd string) (*model.Account, error) {

	if !help.VerifyEmailFormat(email) {
		return nil, fmt.Errorf("email格式不对")
	}

	db := database.GetMysqlDB()
	//执行一下new，保证指针不是nil
	account_model := new(model.Account)
	db.Where(&model.Account{
		Email: email,
	}).First(account_model)

	//找到直接返回
	if account_model.ID != 0 {
		return account_model, nil
	}

	//email注册
	account_model.Email = email
	account_model.Passwd = passwd
	db.Create(account_model)

	return account_model, nil
}

//手机号注册
func (this RegisterService) RegisterAccountByPhone(phone, passwd string) (*model.Account, error) {

	if !help.VerifyMobileFormat(phone) {
		return nil, fmt.Errorf("手机号格式不对")
	}

	db := database.GetMysqlDB()
	account_model := new(model.Account)

	db.Where(&model.Account{
		Phone: phone,
	}).First(account_model)

	//找到直接返回
	if account_model.ID != 0 {
		return account_model, nil
	}

	//手机注册
	account_model.Phone = phone
	account_model.Passwd = passwd
	db.Create(account_model)

	return account_model, nil
}
