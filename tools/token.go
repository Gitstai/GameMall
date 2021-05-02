package tools

import (
	"GameMall/logs"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "GameMall"
)

func GenerateToken(id int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour.Truncate(1)).Unix() //设置过期时间为1小时后
	claims["iat"] = time.Now().Unix()                            //用作和exp对比的时间
	claims["sub"] = id                                           //用于确定用户
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthCheck(tokenString string) (int64, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//验证是否是给定的加密算法
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return false, errors.New("err")
		}
		return []byte(SecretKey), nil
	})
	var id int64
	if err != nil || token == nil {
		logs.Logger.Infof("jwt.Parse err:%v, token:%v", err, token)
		return 0, false
	}
	if !token.Valid {
		logs.Logger.Infof("jwt.Parse, token:%v, not valid", err, token)
		return 0, false
	} else {
		claims := token.Claims.(jwt.MapClaims)
		id = int64(claims["sub"].(float64))
		logs.Logger.Infof("jwt.Parse success, id:%v, ", id)
		return id, true
	}
}
