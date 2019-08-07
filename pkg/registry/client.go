/*
 * Copyright (c) 2019-Present Pivotal Software, Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package registry

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/djoyahoy/image-relocation/pkg/auth"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/djoyahoy/image-relocation/pkg/image"
)

type Client interface {
	Copy(source image.Name, sourceKeyFile string, target image.Name, targetKeyFile string) (image.Digest, int64, error)
}

type client struct {
	readRemoteImage  func(n image.Name, keychain authn.Keychain) (v1.Image, error)
	writeRemoteImage func(i v1.Image, n image.Name, keychain authn.Keychain) error
}

// NewRegistryClient returns a new Client.
func NewRegistryClient() Client {
	return &client{
		readRemoteImage:  readRemoteImage,
		writeRemoteImage: writeRemoteImage,
	}
}

func (r *client) Copy(source image.Name, sourceKeyFile string, target image.Name, targetKeyFile string) (image.Digest, int64, error) {
	img, err := r.readRemoteImage(source, auth.ServiceAccountKeychain{KeyFile: sourceKeyFile})
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to read image %v: %v", source, err)
	}

	hash, err := img.Digest()
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to read digest of image %v: %v", source, err)
	}

	err = r.writeRemoteImage(img, target, auth.ServiceAccountKeychain{KeyFile: targetKeyFile})
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to write image %v: %v", target, err)
	}

	dig, err := image.NewDigest(hash.String())
	if err != nil {
		return image.EmptyDigest, 0, err
	}

	rawManifest, err := img.RawManifest()
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to get raw manifest of image %v: %v", source, err)
	}

	return dig, int64(len(rawManifest)), nil
}
