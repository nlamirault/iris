// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"

	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
)

// CAdvisor represents a cAdvisor client
type CAdvisor struct {

	// cAdvisor client instance
	*client.Client
	NumStats int
	Stdout   bool
	Query    *info.ContainerInfoRequest
}

// Metrics contains cAdvisor metrics
type Metrics struct {
	NodeSpec       *info.MachineInfo
	NodeInfo       *info.ContainerInfo
	ContainersInfo []info.ContainerInfo
}

// NewCAdvisor creates an instance of CAdvisor
func NewCAdvisor(url string, numstats int) (*CAdvisor, error) {
	client, err := client.NewClient(url)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] [cadvisor] Create cAdvisor client using: %s", url)
	return &CAdvisor{
		Client:   client,
		NumStats: numstats,
		Query:    &info.ContainerInfoRequest{NumStats: numstats},
	}, nil
}

// GetMetrics retrieve metrics from cAdvisor
func (c *CAdvisor) GetMetrics() (*Metrics, error) {
	log.Printf("[DEBUG] [cadvisor] Get metrics")
	nodeSpec, err := c.Client.MachineInfo()
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] [cadvisor] Receive machine info : %#v",
		nodeSpec)
	nodeInfo, err := c.Client.ContainerInfo("/", c.Query)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] [cadvisor] Receive container info : %#v",
		nodeInfo)
	containersInfo, err := c.Client.AllDockerContainers(c.Query)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] [cadvisor] Receive all docker containers : %#v",
		containersInfo)
	metrics := &Metrics{
		NodeSpec:       nodeSpec,
		NodeInfo:       nodeInfo,
		ContainersInfo: containersInfo,
	}
	//log.Printf("[INFO] [cadvisor] Metrics: %#v", metrics)
	return metrics, nil
}
