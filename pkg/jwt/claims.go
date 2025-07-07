package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID     string   `json:"sub"`
	TenantID   string   `json:"tenant_id"`
	TenantName string   `json:"tenant_name"`
	Username   string   `json:"username"`
	Roles      []string `json:"roles"`
	Exp        int64    `json:"exp"`
	Iat        int64    `json:"iat"`
}

func GenerateToken(userID, username, tenantID, tenantName string, roles []string, secret []byte) (string, error) {
	claims := Claims{
		UserID:     userID,
		TenantID:   tenantID,
		TenantName: tenantName,
		Username:   username,
		Roles:      roles,
		Exp:        time.Now().Add(24 * time.Hour).Unix(),
		Iat:        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         claims.UserID,
		"tenant_id":   claims.TenantID,
		"tenant_name": claims.TenantName,
		"username":    claims.Username,
		"roles":       claims.Roles,
		"exp":         claims.Exp,
		"iat":         claims.Iat,
	})

	return token.SignedString(secret)
}

func ValidateToken(tokenString string, secret []byte) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	userClaims := &Claims{
		UserID:     claims["sub"].(string),
		TenantID:   claims["tenant_id"].(string),
		TenantName: claims["tenant_name"].(string),
		Username:   claims["username"].(string),
		Exp:        int64(claims["exp"].(float64)),
		Iat:        int64(claims["iat"].(float64)),
	}

	if roles, ok := claims["roles"].([]interface{}); ok {
		for _, role := range roles {
			userClaims.Roles = append(userClaims.Roles, role.(string))
		}
	}

	return userClaims, nil
}