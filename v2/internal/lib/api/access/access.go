package access

import (
	"context"
	"github.com/s4bb4t/srest-api/internal/database"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(`b3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABDIsCk4b4SwgpWaZXbeuCXUAAAAEAAAAAEAAAGXAAAAB3NzaC1yc2EAAAADAQABAAABgQCwN27MXT2rYoNIzwqPtHxIBiJhlPLWEAakzCxQesr8W0hBHrMBWfsVvYhCF+l4vdPwcTL6Vav6FefAQICrgEpnMtzT3i25KT4vV/4Q07oqhNvNp`)

type CxtKey string

type Claims struct {
	UserId      int      `json:"id"`
	Roles       []string `json:"isAdmin"`
	UserVersion int      `json:"userVersion"`
	jwt.StandardClaims
}

type UserContext struct {
	UserId    int      `json:"id"`
	Roles     []string `json:"isAdmin"`
	IsBlocked bool     `json:"isBlocked"`
}

func NewAccessToken(id int, roles []string) (string, error) {
	expirationTime := time.Now().Add(3 * time.Minute)
	claims := &Claims{
		UserId:      id,
		Roles:       roles,
		UserVersion: database.UserVersion(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewRefreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	key := []byte(strconv.Itoa(int(time.Now().Unix())))
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if claims.UserVersion != database.UserVersion(claims.UserId) {
			http.Error(w, "Session ended", http.StatusUnauthorized)
			return
		}

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userContext := UserContext{
			UserId: claims.UserId,
			Roles:  claims.Roles,
		}
		ctx := context.WithValue(r.Context(), CxtKey("userContext"), userContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
