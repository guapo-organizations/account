package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey []byte

func init() {

	if mySigningKey == nil {
		mySigningKey = []byte("zldzJwtSigninKey")
	}
}

//jwt 编码
func JwtTokenEncode(jwt_map_claims jwt.MapClaims) (string, error) {
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt_map_claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("jwt加密失败：%s", err)
	}
	return tokenString, nil
}

//jwt 解码
func JwtTokenDecode(jwt_token string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, nil
	}

	return nil, err
}
