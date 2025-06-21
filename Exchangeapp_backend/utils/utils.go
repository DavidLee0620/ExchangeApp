package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/*Bcrypt加密算法，加盐后哈希密码

$2a$12$R9h/cIPz0gi.URNNX3kh2OPST9/PgBkqquzi.Ss7KIUgO2t0jWMUW
\__/\/ \____________________/\_____________________________/
Alg Cost      Salt                        Hash

*/

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

func GenerateJWT(name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"username": name,
	})
	sightoken, err := token.SignedString([]byte("secret"))
	return "Bearer " + sightoken, err
}
func ParseJWT(token string) (string, error) {
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected Signing Method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username is not string")
		}
		return username, nil
	}
	return "", err

}
func CheckPwd(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil

}
