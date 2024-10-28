package requests

import (
    "github.com/DaniloFaraum/studere-backend/utils"
)

type CreateAnswerRequest struct {
    IDQuestion int    `json:"id_question"`
    Text       string `json:"text"`
}

func (r *CreateAnswerRequest) Validate() error {
    if r.Text == "" {
        return utils.ErrParamIsrequired("text", "string")
    }
    if r.IDQuestion == 0 {
        return utils.ErrParamIsrequired("id_question", "int")
    }
    return nil
}

type UpdateAnswerRequest struct {
    Text string `json:"text"`
}

func (r *UpdateAnswerRequest) Validate() error {
    if r.Text != "" {
        return nil
    }
    return utils.ErrNoValidFields()
}