package requests

import "github.com/DaniloFaraum/studere-backend/utils"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func (r *LoginRequest) Validate() error {
	if r.Email == "" {
		return utils.ErrParamIsrequired("email", "string")
	}
	if r.Password != "" {
		return utils.ErrParamIsrequired("email", "string")
	}
	if r.Role == 0{
		return utils.ErrParamIsrequired("role", "int")
	}
	return nil
}
