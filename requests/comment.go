package requests

import (
    "errors"
    "github.com/DaniloFaraum/studere-backend/utils"
)

type CreateCommentRequest struct {
    IDUser   int    `json:"id_user"`
    IDCourse int    `json:"id_course"`
    Content  string `json:"content"`
}

func (r *CreateCommentRequest) Validate() error {
    if r.IDUser == 0 {
        return errors.New("user ID is required")
    }
    if r.IDCourse == 0 {
        return errors.New("course ID is required")
    }
    if r.Content == "" {
        return errors.New("content is required")
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