package requests

import (
	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateImageRequest struct {
	Name           string `json:"name"`
	EncryptedImage []byte `json:"encrypted_image"`
}

func (r *CreateImageRequest) Validate() error {
	if r.Name == "" {
		return utils.ErrParamIsrequired("name", "string")
	}
	if len(r.EncryptedImage) == 0 {
		return utils.ErrParamIsrequired("encrypted image", "string") //need to change type later
	}
	return nil
}
