package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateRatingRequest struct {
	IDUser     int    `json:"id_user"`
	IDCourse   int    `json:"id_course"`
	Opinion    int    `json:"opinion"`
	Commentary string `json:"commentary"`
}

func (r *CreateRatingRequest) Validate() error {
	if r.IDUser == 0 {
        return utils.ErrParamIsrequired("id_user", "int")
    }
    if r.IDCourse == 0 {
        return utils.ErrParamIsrequired("id_course", "int")
    }
    if r.Opinion == 0 {
        return utils.ErrParamIsrequired("opinion", "int")
    }
    if r.Commentary == "" {
        return utils.ErrParamIsrequired("commentary", "string")
    }
	return nil
}

type UpdateRatingRequest struct {
	Opinion    int    `json:"opinion"`
	Commentary string `json:"commentary"`
}

func (r *UpdateRatingRequest) Validate() error {
	if r.Opinion != 0 || r.Commentary != "" {
		return nil
	}

	return utils.ErrNoValidFields()
}
