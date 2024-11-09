package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateCommentRequest struct {
	IDUser   int    `json:"id_user"`
	IDCourse int    `json:"id_course"`
	Content  string `json:"content"`
}

func (r *CreateCommentRequest) Validate() error {
	if r.IDUser == 0 {
		return utils.ErrParamIsrequired("id_user", "int")
	}
	if r.IDCourse == 0 {
		return utils.ErrParamIsrequired("id_course", "int")
	}
	if r.Content == "" {
		utils.ErrParamIsrequired("content", "string")
	}
	return nil
}

type UpdateCommentRequest struct {
	Content string `json:"content"`
}

func (r *UpdateCommentRequest) Validate() error {
	if r.Content != "" {
		return nil
	}
	return utils.ErrNoValidFields()
}
