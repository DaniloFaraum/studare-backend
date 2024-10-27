package requests

import (
    "errors"
    "github.com/DaniloFaraum/studere-backend/utils"
)

type CreateImageRequest struct {
    Name           string `json:"name"`
    EncryptedImage []byte `json:"encrypted_image"`
}

func (r *CreateImageRequest) Validate() error {
    if r.Name == "" {
        return errors.New("name is required")
    }
    if len(r.EncryptedImage) == 0 {
        return errors.New("encrypted image is required")
    }
    return nil
}

type UpdateImageRequest struct {
    Name           string `json:"name"`
    EncryptedImage []byte `json:"encrypted_image"`
}

func (r *UpdateImageRequest) Validate() error {
    if r.Name != "" || len(r.EncryptedImage) != 0 {
        return nil
    }

    return utils.ErrNoValidFields()
}