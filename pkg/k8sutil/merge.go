// Copyright 2016 The Nho Luong DevOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sutil

import (
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

// MergePatchContainers adds patches to base using a strategic merge patch and
// iterating by container name, failing on the first error.
func MergePatchContainers(base, patches []v1.Container) ([]v1.Container, error) {
	var out []v1.Container

	containersByName := make(map[string]v1.Container)
	for _, c := range patches {
		containersByName[c.Name] = c
	}

	// Patch the containers that exist in base.
	for _, container := range base {
		patchContainer, ok := containersByName[container.Name]
		if !ok {
			// This container didn't need to be patched.
			out = append(out, container)
			continue
		}

		containerBytes, err := json.Marshal(container)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON for container %s: %w", container.Name, err)
		}

		patchBytes, err := json.Marshal(patchContainer)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON for patch container %s: %w", container.Name, err)
		}

		// Calculate the patch result.
		jsonResult, err := strategicpatch.StrategicMergePatch(containerBytes, patchBytes, v1.Container{})
		if err != nil {
			return nil, fmt.Errorf("failed to generate merge patch for container %s: %w", container.Name, err)
		}

		var patchResult v1.Container
		if err := json.Unmarshal(jsonResult, &patchResult); err != nil {
			return nil, fmt.Errorf("failed to unmarshal merged container %s: %w", container.Name, err)
		}

		// Add the patch result and remove the corresponding key from the to do list.
		out = append(out, patchResult)
		delete(containersByName, container.Name)
	}

	// Append containers that are left in containersByName.
	// Iterate over patches to preserve the order.
	for _, c := range patches {
		if container, found := containersByName[c.Name]; found {
			out = append(out, container)
		}
	}

	return out, nil
}
