package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/utils"
)

type ctxKey string

const (
	ctxUserIDKey ctxKey = "auth_user_id"
	ctxRoleIDKey ctxKey = "auth_role_id"
)

const (
	RoleAdminID int64 = 1
	RoleUserID  int64 = 2
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{jwtSecret: jwtSecret}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			errorx.WriteError(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		parts := strings.SplitN(authorizationHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			errorx.WriteError(w, http.StatusUnauthorized, "invalid authorization header")
			return
		}

		claims, err := utils.ParseAccessToken(m.jwtSecret, parts[1])
		if err != nil {
			errorx.WriteError(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, ctxRoleIDKey, claims.RoleID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) RequireRoles(roles ...int64) func(http.Handler) http.Handler {
	allowed := make(map[int64]struct{}, len(roles))
	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roleID, ok := RoleIDFromContext(r.Context())
			if !ok {
				errorx.WriteError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if _, exists := allowed[roleID]; !exists {
				errorx.WriteError(w, http.StatusForbidden, "forbidden")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(ctxUserIDKey).(int64)
	return userID, ok
}

func RoleIDFromContext(ctx context.Context) (int64, bool) {
	roleID, ok := ctx.Value(ctxRoleIDKey).(int64)
	return roleID, ok
}
