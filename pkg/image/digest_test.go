/*
 * Copyright (c) 2018-Present Pivotal Software, Inc. All rights reserved.
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

package image_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/djoyahoy/image-relocation/pkg/image"
)

var _ = Describe("Digest", func() {
	Describe("EmptyDigest", func() {
		It("should produce an empty string form", func() {
			Expect(image.EmptyDigest.String()).To(BeEmpty())
		})
	})

	Describe("NewDigest", func() {
		var (
			str    string
			digest image.Digest
			err    error
		)

		JustBeforeEach(func() {
			digest, err = image.NewDigest(str)
		})

		Context("when the input string is valid", func() {
			BeforeEach(func() {
				str = "sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
			})

			It("should produce a digest with the correct string form", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(digest.String()).To(Equal(str))
			})
		})

		Context("when the input string is invalid", func() {
			BeforeEach(func() {
				str = "sha256:adeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
			})

			It("should return a suitable error", func() {
				Expect(err).To(MatchError("invalid checksum digest length"))
			})
		})
	})
})
