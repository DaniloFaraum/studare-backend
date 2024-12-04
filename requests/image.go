package requests

import (
	"os"
	//	"github.com/DaniloFaraum/studere-backend/utils"
)

type CreateImageRequest struct {
	Image *os.File `json:"encrypted_image"`
}

// func (r *CreateImageRequest) Validate() error {
// 	if len(r.EncryptedImage) == 0 {
// 		return utils.ErrParamIsrequired("encrypted image", "string") //need to change type later
// 	}
// 	return nil
// }
