package middleware

import (
    "context"
    "net/http"
    "strings"
    "github.com/Hanumath1006/skillsync/utils"
)

type contextKey string

const userCtxKey = contextKey("user")

// AuthMiddleware verifies JWT in Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

        claims, err := utils.ValidateJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Add user info to context
        ctx := context.WithValue(r.Context(), userCtxKey, claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetUserFromContext pulls user info from request context
func GetUserFromContext(r *http.Request) *utils.Claims {
    user, ok := r.Context().Value(userCtxKey).(*utils.Claims)
    if !ok {
        return nil
    }
    return user
}
