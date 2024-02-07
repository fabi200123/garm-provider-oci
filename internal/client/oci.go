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

package client

import (
	"context"
	"fmt"

	"github.com/cloudbase/garm-provider-oci/config"
	"github.com/oracle/oci-go-sdk/v49/common"
	"github.com/oracle/oci-go-sdk/v49/core"
)

func NewOciCli(ctx context.Context, cfg *config.Config) (*OciCli, error) {
	confProvider := common.NewRawConfigurationProvider(
		cfg.TenancyID,
		cfg.UserID,
		cfg.Region,
		cfg.Fingerprint,
		cfg.PrivateKeyPath,
		common.String(cfg.PrivateKeyPassword),
	)
	computeClient, err := core.NewComputeClientWithConfigurationProvider(confProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating compute client: %w", err)
	}
	return &OciCli{
		computeClient: computeClient,
	}, nil
}

type OciCli struct {
	computeClient core.ComputeClient
}

func (o *OciCli) CreateInstance(ctx context.Context, instanceID string) error {
	return nil
}

func (o *OciCli) GetInstance(ctx context.Context, instanceID string) (core.Instance, error) {
	req := core.GetInstanceRequest{
		InstanceId: &instanceID,
	}
	resp, err := o.computeClient.GetInstance(ctx, req)
	if err != nil {
		return core.Instance{}, fmt.Errorf("error getting instance: %w", err)
	}

	return resp.Instance, nil
}

func (o *OciCli) DeleteInstance(ctx context.Context, instanceID string) error {
	request := core.TerminateInstanceRequest{
		InstanceId: &instanceID,
	}

	_, err := o.computeClient.TerminateInstance(ctx, request)
	if err != nil {
		return fmt.Errorf("error terminating instance: %w", err)
	}
	return nil
}

func (o *OciCli) ListInstances(ctx context.Context) error {
	return nil
}

func (o *OciCli) StopInstance(ctx context.Context, instanceID string) error {
	return nil
}

func (o *OciCli) StartInstance(ctx context.Context, instanceID string) error {
	return nil
}
