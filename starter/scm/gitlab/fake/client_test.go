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

package fake_test

import (
	"testing"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
)

func init()  {
	log.SetLevel(log.DebugLevel)
}

func TestFakeClient(t *testing.T) {
	ss := fake.NewClient("")
	gs := new(gitlab.Session)
	gr := new(gitlab.Response)
	ss.On("GetSession", nil, nil).Return(gs, gr, nil)
	s, r, e := ss.GetSession(nil, nil)
	assert.Equal(t, nil, e)
	assert.Equal(t, s, gs)
	assert.Equal(t, r, gr)
}
