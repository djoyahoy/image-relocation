package auth

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"io/ioutil"
)

type ServiceAccountKeychain struct {
	KeyFile string
}

func (e ServiceAccountKeychain) Resolve(name.Registry) (authn.Authenticator, error) {
	buf, err := ioutil.ReadFile(e.KeyFile)
	if err != nil {
		return nil, err
	}
	return &authn.Basic{Username: "_json_key", Password: string(buf)}, nil
}
