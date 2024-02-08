// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package util

import (
	"github.com/cloudbase/garm-provider-common/params"
	"github.com/oracle/oci-go-sdk/v49/core"
)

func OciInstanceToProviderInstance(ociInstance core.Instance) params.ProviderInstance {
	details := params.ProviderInstance{
		ProviderID: *ociInstance.Id,
		Name:       *ociInstance.DisplayName,
		OSType:     "linux", //TODO: get from oci
		OSArch:     "amd64", //TODO: get from oci
	}

	switch ociInstance.LifecycleState {
	case core.InstanceLifecycleStateRunning:
		details.Status = "running"
	case core.InstanceLifecycleStateStopped,
		core.InstanceLifecycleStateTerminated:

		details.Status = "stopped"
	default:
		details.Status = "unknown"
	}

	return details
}
