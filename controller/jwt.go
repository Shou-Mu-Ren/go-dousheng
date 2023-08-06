package controller

import (
	"douyin/repository"
	"douyin/service"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Name                 string `json:"name"`
	Password             string `json:"password"`
	jwt.RegisteredClaims
}

var UserSecret = []byte("com.linxi") // 定义secret，后面会用到


func MakeToken(name string, password string) (tokenString string, err error) {
	claim := UserClaims{
		Name: name,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(UserSecret)
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("com.linxi"), nil // 我的secret
	}
}

func ParseToken(tokenss string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &UserClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token为空")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token过期")
			} else {
				return nil, errors.New("token无法解析")
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token无法解析")
}

func ValidToken(tokenss string)(*repository.User, error){
	claims, err := ParseToken(tokenss)
	if err != nil {
		return nil, err
	}
	if user, err := service.FindUserByNameAndPassword(claims.Name, claims.Password); err != nil{
		return nil, err
	}else{
		return user, nil
	}
}