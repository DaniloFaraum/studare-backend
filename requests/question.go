package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateQuestionRequest struct {
	Name            string `json:"name"`
	IDQuestionnaire int    `json:"id_questionnaire"`
	Text            string `json:"question"`
}

func (r *CreateQuestionRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	return nil
}

type UpdateQuestionRequest struct {
	Text string `json:"name"`
}

func (r *UpdateQuestionRequest) Validate() error {
	if r.Text != "" {
		return nil
	}

	return utils.ErrNoValidFields()
}
