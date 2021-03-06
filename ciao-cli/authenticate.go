//
// Copyright (c) 2016 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"crypto/tls"
	"net/http"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

func newAuthenticatedClient(opt gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	provider, err := openstack.NewClient(opt.IdentityEndpoint)
	if err != nil {
		errorf("Could not get ProviderClient %s\n", err)
		return nil, err
	}

	if caCertPool != nil {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: caCertPool},
		}
		provider.HTTPClient.Transport = transport
	}

	err = openstack.Authenticate(provider, opt)
	if err != nil {
		errorf("Could not get AuthenticatedClient %s\n", err)
		return nil, err
	}

	return provider, nil
}
