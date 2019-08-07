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

package irel

import (
	"fmt"
	"github.com/djoyahoy/image-relocation/pkg/image"
	"github.com/djoyahoy/image-relocation/pkg/registry"
	"github.com/spf13/cobra"
	"log"
)

func init() { Root.AddCommand(newCmdCopy()) }

func newCmdCopy() *cobra.Command {
	return &cobra.Command{
		Use:     "copy SRC_REF SRC_JSON_KEY DST_REF DST_JSON_KEY",
		Aliases: []string{"cp"},
		Short:   "Efficiently copy a remote image from one repository to another",
		Args:    cobra.ExactArgs(4),
		Run:     copy,
	}
}

func copy(cmd *cobra.Command, args []string) {
	srcStr, srcKeyFile, dstStr, dstKeyFile := args[0], args[1], args[2], args[3]
	src, err := image.NewName(srcStr)
	if err != nil {
		log.Fatalf("invalid reference %q: %v", srcStr, err)
	}
	dst, err := image.NewName(dstStr)
	if err != nil {
		log.Fatalf("invalid reference %q: %v", dstStr, err)
	}

	regClient := registry.NewRegistryClient()
	dig, _, err := regClient.Copy(src, srcKeyFile, dst, dstKeyFile)
	if err != nil {
		log.Fatalf("copy failed: %v", err)
	}
	fmt.Printf("copied %s to %s with content digest %s\n", src, dst, dig)
}
