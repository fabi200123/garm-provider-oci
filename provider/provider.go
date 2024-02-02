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

package provider

import (
	"context"
	"fmt"

	"github.com/cloudbase/garm-provider-common/execution"
	"github.com/cloudbase/garm-provider-common/params"
	"github.com/cloudbase/garm-provider-oci/config"
	"github.com/cloudbase/garm-provider-oci/internal/client"
)

var _ execution.ExternalProvider = &OciProvider{}

func NewOciProvider(ctx context.Context, configPath string, controllerID string) (*OciProvider, error) {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	ociCli, err := client.NewOciCli(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("error creating oci client: %w", err)
	}
	return &OciProvider{
		cfg:    conf,
		ociCli: ociCli,
	}, nil
}

type OciProvider struct {
	cfg    *config.Config
	ociCli *client.OciCli
}

func (o *OciProvider) CreateInstance(ctx context.Context, bootstrapParams params.BootstrapInstance) (params.ProviderInstance, error) {
	return params.ProviderInstance{}, nil
}

func (o *OciProvider) GetInstance(ctx context.Context, instanceID string) (params.ProviderInstance, error) {
	ociInstance, err := o.ociCli.GetInstance(ctx, instanceID)
	if err != nil {
		return params.ProviderInstance{}, fmt.Errorf("error getting instance: %w", err)
	}
	providerInstance := params.ProviderInstance{
		ProviderID: *ociInstance.Id,
		Name:       *ociInstance.DisplayName,
		// https://pkg.go.dev/github.com/oracle/oci-go-sdk/v49@v49.2.0/core#InstanceLifecycleStateEnum
		Status: "running", //*ociInstance.LifecycleState
		OSType: "linux",   //TODO: get from oci
		OSArch: "amd64",   //TODO: get from oci
	}
	return providerInstance, nil
}

func (o *OciProvider) DeleteInstance(ctx context.Context, instanceID string) error {
	return o.ociCli.DeleteInstance(ctx, instanceID)
}

func (o *OciProvider) ListInstances(ctx context.Context, poolID string) ([]params.ProviderInstance, error) {
	return nil, nil
}

func (o *OciProvider) RemoveAllInstances(ctx context.Context) error {
	return nil
}

func (o *OciProvider) Stop(ctx context.Context, instance string, force bool) error {
	return o.ociCli.StopInstance(ctx, instance)
}

func (o *OciProvider) Start(ctx context.Context, instance string) error {
	return nil
}
