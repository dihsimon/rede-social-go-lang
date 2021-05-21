package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	persmissoes := jwt.MapClaims{}
	persmissoes["autorized"] = true
	persmissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	persmissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, persmissoes)
	return token.SignedString([]byte(config.SecretKey))
}
