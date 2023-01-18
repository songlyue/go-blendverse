package utils

import (
	"blendverse/global"
	"blendverse/model/system/request"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
}

func (j *JWT) CreateClaims(claims request.BaseClaims) request.CustomClaims {
	return request.CustomClaims{
		BaseClaims: claims,
		BufferTime: int64(3600 * 24),
		StandardClaims: jwt.StandardClaims{
			Issuer: global.GVA_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
}

func (j *JWT) createToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, TokenInvalid
		}
	} else {
		return nil, TokenInvalid
	}
}
