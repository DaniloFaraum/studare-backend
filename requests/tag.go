package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateTagRequest struct {
	Name string `json:"name"`
}

func (r *CreateTagRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	return nil
}

type UpdateTagRequest struct {
	Name string `json:"name"`
}

func (r *UpdateTagRequest) Validate() error {
	if r.Name != "" {
		return nil
	}

	return utils.ErrNoValidFields()
}
