package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1920853199/passwd/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var Jwt = JwtConfig{
	Key:       "passwd",
	TokenName: "token",
}

type JwtConfig struct {
	Key       string
	TokenName string
}

func CheckJwt(req *http.Request) error {
	token, err := Jwt.ParseToken(req)
	if err != nil {
		return err
	}

	macAddr, err := service.Db.Get([]byte(service.TOKENKEY))
	if err != nil {
		return err
	}

	if token.Id != string(macAddr) {
		return fmt.Errorf("%s", "invalid token.")
	}

	return nil
}

func (conf JwtConfig) SetToken(jwtClaims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	t, err := token.SignedString([]byte(conf.Key))
	if err != nil {
		return t, err
	}

	// 保存 Token
	service.Db.Set([]byte(service.TOKENKEY), []byte(jwtClaims.Id))

	return t, err

}

func (conf JwtConfig) ParseToken(req *http.Request) (*jwt.StandardClaims, error) {
	token2, err := request.ParseFromRequest(req, request.MultiExtractor{request.ArgumentExtractor{conf.TokenName}, request.HeaderExtractor{conf.TokenName}}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Key), nil
	})
	if token2 == nil || token2.Claims == nil {
		return nil, err
	}
	mClaims := token2.Claims.(jwt.MapClaims)
	claims := &jwt.StandardClaims{
		Audience:  fmt.Sprintf("%v", mClaims["aud"]),
		ExpiresAt: emp(mClaims),
		Id:        fmt.Sprintf("%v", mClaims["jti"]),
		Subject:   fmt.Sprintf("%v", mClaims["sub"]),
		Issuer:    fmt.Sprintf("%v", mClaims["iss"]),
	}
	return claims, err
}

func emp(m map[string]interface{}) int64 {
	switch exp := m["exp"].(type) {
	case float64:
		return int64(exp)
	case json.Number:
		v, _ := exp.Int64()
		return v
	}
	return 0
}
