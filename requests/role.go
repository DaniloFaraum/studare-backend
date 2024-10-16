package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func (r *CreateRoleRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	return nil
}

type UpdateRoleRequest struct {
	Name string `json:"name"`
}

func (r *UpdateRoleRequest) Validate() error {
	if r.Name != "" {
		return nil
	}

	return utils.ErrNoValidFields()
}
