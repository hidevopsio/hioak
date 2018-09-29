// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package gitlab_test

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	gg "github.com/xanzy/go-gitlab"
	"testing"
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestSession(t *testing.T) {
	gs := &gg.Session{
		Username: "chulei",
	}
	gr := new(gg.Response)
	fs := new(fake.SessionService)
	cli := &fake.Client{
		SessionService: fs,
	}
	s := gitlab.NewSession(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	cli.On("Session", nil, nil).Return(fs)
	fs.On("GetSession", nil, nil).Return(gs, gr, nil)
	e := s.GetSession("", "", "")
	assert.Equal(t, nil, e)
}
