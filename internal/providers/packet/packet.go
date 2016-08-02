// Copyright 2016 CoreOS, Inc.
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

// The packet provider fetches a remote configuration from the packet.net
// userdata metadata service URL.

package packet

import (
	"github.com/coreos/ignition/config"
	"github.com/coreos/ignition/config/types"
	"github.com/coreos/ignition/internal/log"
	"github.com/coreos/ignition/internal/providers"
	"github.com/coreos/ignition/internal/util"

	"github.com/packethost/packngo/metadata"
)

type Creator struct{}

func (Creator) Create(logger *log.Logger) providers.Provider {
	return &provider{
		logger: logger,
		client: util.NewHttpClient(logger),
	}
}

type provider struct {
	logger *log.Logger
	client util.HttpClient
}

func (p provider) FetchConfig() (types.Config, error) {
	p.logger.Debug("fetching config from packet metadata")
	data, err := metadata.GetUserData()
	if err != nil {
		p.logger.Err("failed to fetch config: %v", err)
		return types.Config{}, err
	}

	return config.Parse(data)
}
