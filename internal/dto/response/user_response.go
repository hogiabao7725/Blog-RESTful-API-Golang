package response

import "time"

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}
