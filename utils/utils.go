package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goblog/config"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.JWTKey),
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}
