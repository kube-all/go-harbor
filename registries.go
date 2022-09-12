/*
Copyright 2022 The kubeall.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"github.com/kube-all/go-harbor/models"
	"github.com/parnurzeal/gorequest"
)

type RegistriesService struct {
	client *Client
}

func (s *RegistriesService) UpdateRegistry(registry *models.Registry) (success bool, resp gorequest.Response, err error) {
	var errs []error
	agent := s.client.
		newRequest(gorequest.POST, "/registries").
		AppendHeader("Accept", "application/json").
		Send(registry)
	if len(s.client.HeaderXRequestId) > 0 {
		agent = agent.AppendHeader("X-Request-Id", s.client.HeaderXRequestId)
	}
	resp, _, errs = agent.End()

	err = s.client.ErrsWrapper(errs)
	return
}
func (s *RegistriesService) GetRegistry(id string) (registry models.Registry, resp gorequest.Response, err error) {
	var errs []error
	agent := s.client.
		newRequest(gorequest.GET, "/registries/"+id)
	if len(s.client.HeaderXRequestId) > 0 {
		agent = agent.AppendHeader("X-Request-Id", s.client.HeaderXRequestId)
	}
	resp, _, errs = agent.EndStruct(&resp)
	err = s.client.ErrsWrapper(errs)
	return
}
func (s *RegistriesService) DeleteRegistry(id string) (success bool, resp gorequest.Response, err error) {
	var errs []error
	agent := s.client.
		newRequest(gorequest.DELETE, "/registries/"+id)
	if len(s.client.HeaderXRequestId) > 0 {
		agent = agent.AppendHeader("X-Request-Id", s.client.HeaderXRequestId)
	}
	resp, _, errs = agent.End()
	err = s.client.ErrsWrapper(errs)
	return
}
func (s *RegistriesService) CreateRegistry(registry *models.Registry) (success bool, resp gorequest.Response, err error) {
	var errs []error
	agent := s.client.
		newRequest(gorequest.POST, "/registries").
		AppendHeader("Accept", "application/json").
		Send(registry)
	if len(s.client.HeaderXRequestId) > 0 {
		agent = agent.AppendHeader("X-Request-Id", s.client.HeaderXRequestId)
	}
	resp, _, errs = agent.End()

	err = s.client.ErrsWrapper(errs)
	return
}
func (s *RegistriesService) ListRegistries(query models.Query) (success bool, resp gorequest.Response, err error) {
	var errs []error
	agent := s.client.
		newRequest(gorequest.GET, "/registries").
		AppendHeader("Accept", "application/json").
		Query(query)
	if len(s.client.HeaderXRequestId) > 0 {
		agent = agent.AppendHeader("X-Request-Id", s.client.HeaderXRequestId)
	}
	resp, _, errs = agent.End()

	err = s.client.ErrsWrapper(errs)
	return
}
