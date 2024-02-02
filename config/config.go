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

package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func NewConfig(cfgFile string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(cfgFile, &config); err != nil {
		return nil, fmt.Errorf("error decoding config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}
	return &config, nil
}

type Config struct {
	TenancyID          string `yaml:"tenancy_id"`
	UserID             string `yaml:"user_id"`
	Region             string `yaml:"region"`
	Fingerprint        string `yaml:"fingerprint"`
	PrivateKeyPath     string `yaml:"private_key_path"`
	PrivateKeyPassword string `yaml:"private_key_password"`
}

func (c *Config) Validate() error {
	if c.TenancyID == "" {
		return fmt.Errorf("tenancy_id is required")
	}
	if c.UserID == "" {
		return fmt.Errorf("user_id is required")
	}
	if c.Region == "" {
		return fmt.Errorf("region is required")
	}
	if c.Fingerprint == "" {
		return fmt.Errorf("fingerprint is required")
	}
	if c.PrivateKeyPath == "" {
		return fmt.Errorf("private_key_path is required")
	}
	return nil
}
