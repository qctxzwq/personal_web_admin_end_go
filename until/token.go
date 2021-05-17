package until

import (
	"admin/models"
	"errors"
	"github.com/astaxie/beego/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 创建token
func CreateToken(user models.Users) (string, error) {
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return "", err
	}
	expireTime, err := BConfig.Int("token::expire_time")
	if err != nil {
		expireTime = 7 * 24
	}
	tokenKey := BConfig.String("token::token_key")

	claims := make(jwt.MapClaims)
	claims["id"] = user.Id
	claims["name"] = user.Name
	claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenKey string) (jwt.MapClaims, error) {
	token, err := ValToken(tokenKey)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("get ParseToken claims error")
	}
	return claims, nil
}

// token 验证
func ValToken(tokenString string) (t *jwt.Token, err error) {
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return nil, err
	}
	tokenKey := BConfig.String("token::token_key")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("this is not a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errors.New("token 过期")
			} else {
				// Couldn't handle this token
				return nil, errors.New("无法处理的 token")
			}
		} else {
			// Couldn't handle this token
			return nil, errors.New("the token can't be handle")
		}
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
