package requests

import "github.com/DaniloFaraum/studere-backend/utils"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleID   int    `json:"role"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Email == "" {
		return utils.ErrParamIsrequired("email", "string")
	}
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	if r.Password != "" {
		return utils.ErrParamIsrequired("email", "string")
	}
	if r.RoleID == 0 {
		return utils.ErrParamIsrequired("role", "int")
	}
	return nil
}

type UpdateUserRequest struct{
	Email    string `json:"email"`
	Name     string `json:"name"`
	//PFP?
	Password string `json:"password"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Email != "" || r.Name != "" || r.Password != ""{
		return nil
	}

	return utils.ErrNoValidFields()
}

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
	if r.Role == 0 {
		return utils.ErrParamIsrequired("role", "int")
	}
	return nil
}
