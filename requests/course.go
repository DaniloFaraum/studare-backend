package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateCourseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Duration    int64  `json:"duration"`
	Author      string `json:"author"`
	Institution string `json:"institution"`
	IDImage     int    `json:"id_image"`
}

func (r *CreateCourseRequest) Validate() error {
	switch {
	case r.Name == "":
		return utils.ErrParamIsrequired("name", "string")
	case r.Description == "":
		return utils.ErrParamIsrequired("desciption", "string")
	case r.Link == "":
		return utils.ErrParamIsrequired("link", "string")
	case r.Duration == 0:
		return utils.ErrParamIsrequired("duration", "time")
	case r.Author == "":
		return utils.ErrParamIsrequired("author", "string")
	case r.Institution == "":
		return utils.ErrParamIsrequired("institution", "string")
	case r.IDImage == 0:
		return utils.ErrParamIsrequired("id_image", "int")
	}

	return nil
}

type UpdateCourseRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Duration    int64  `json:"duration"`
	Author      string `json:"author"`
	Institution string `json:"institution"`
}

func (r *UpdateCourseRequest) Validate() error {
	if r.Name != "" || r.Description != "" || r.Link != "" || r.Duration != 0 || r.Author != "" || r.Institution != "" {
		return nil
	}

	return utils.ErrNoValidFields()
}
