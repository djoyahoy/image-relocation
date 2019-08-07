package auth

import (
	"errors"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"os"
)

type ServiceAccountKeychain struct {
	KeyEnv string
}

func (e ServiceAccountKeychain) Resolve(name.Registry) (authn.Authenticator, error) {
	val, ok := os.LookupEnv(e.KeyEnv)
	if !ok {
		return nil, errors.New("could not read env var: " + e.KeyEnv)
	}
	return &authn.Basic{Username: "_json_key", Password: val}, nil
}
