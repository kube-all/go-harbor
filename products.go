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
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ProductsService struct {
	client *Client
}

//RemoveLabel Remove label from the specified chart version.
func (s *ProductsService) RemoveLabel(repo, name, version string, id int64) (resp gorequest.Response, err error) {
	var errs []error
	resp, _, errs = s.client.

		newRequest(gorequest.DELETE,
			fmt.Sprintf("/chartrepo/%s/charts/%s/%s/labels/%d", repo, name, version, id)).
		End()
	err = s.client.ErrsWrapper(errs)
	return
}
