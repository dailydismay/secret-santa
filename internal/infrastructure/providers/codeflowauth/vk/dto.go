package vk

import (
	"fmt"
	"secretsanta/internal/domain/codeflowauth"
)

type ProfileResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo     string `json:"photo_400_orig"`
}

func (pr *ProfileResponse) ToDomain() *codeflowauth.UserProfile {
	return codeflowauth.NewUserProfile(
		codeflowauth.NewUserProfileID(fmt.Sprint(pr.ID)),
		pr.FirstName,
		pr.LastName,
		pr.Photo,
	)
}
