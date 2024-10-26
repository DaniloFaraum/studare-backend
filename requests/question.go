package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateQuestionRequest struct {
	Name            string `json:"name"`
	IDQuestionnaire int    `json:"id_questionnaire"`
	Question        string `json:"question"`
}

func (r *CreateQuestionRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	return nil
}

type UpdateQuestionRequest struct {
	Name string `json:"name"`
}

func (r *UpdateQuestionRequest) Validate() error {
	if r.IDQuestionnaire != nil {
		return nil
	}

	return utils.ErrNoValidFields()
}
