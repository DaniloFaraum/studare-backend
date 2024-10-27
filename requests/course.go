package requests

import (
	"errors"
	"time"

	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateCourseRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Duration    time.Time `json:"duration"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	IDImage     int       `json:"id_image"`
}

func (r *CreateCourseRequest) Validate() error {
	switch {
	case r.Name == "":
		return errors.New("name is required")
	case r.Description == "":
		return errors.New("description is required")
	case r.Link == "":
		return errors.New("link is required")
	case r.Duration.IsZero():
		return errors.New("duration is required")
	case r.Author == "":
		return errors.New("author is required")
	case r.Institution == "":
		return errors.New("institution is required")
	case r.IDImage == 0:
		return errors.New("id_image is required")
	}

	return nil
}

type UpdateCourseRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Duration    time.Time `json:"duration"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
}

func (r *UpdateCourseRequest) Validate() error {
	if r.Name != "" || r.Description != "" || r.Link != "" || !r.Duration.IsZero() || r.Author != "" || r.Institution != ""{
		return nil
	}

	return utils.ErrNoValidFields()
}
