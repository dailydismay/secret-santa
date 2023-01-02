package codeflowauth

import "context"

type CodeFlowAuthProvider interface {
	GetUserProfile(ctx context.Context, code string) (*UserProfile, error)
}
