package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateQuestionnaireRequest struct {
	IDUser int    `json:"id_user"`
	Name   string `json:"title"`
	Ready  int8   `json:"ready"`
}

func (r *CreateQuestionnaireRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	if r.IDUser == 0 {
		return utils.ErrParamIsrequired("id_user", "int")
	}
	return nil
}

type UpdateQuestionnaireRequest struct {
	Name  string `json:"title"`
	Ready int8 `json:"ready"`
}

func (r *UpdateQuestionnaireRequest) Validate() error {
	if r.Name != ""{
		return nil
	}

	return utils.ErrNoValidFields()
}
