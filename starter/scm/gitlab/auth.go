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

package gitlab

import (
	"github.com/jinzhu/copier"
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type Session struct {
	scm.Session
	newClient NewClient
}

const (
	ApiVersion = "/api/v3"
)

func NewSession(c NewClient) *Session {
	return &Session{
		newClient: c,
	}
}

func (s *Session) GetSession(baseUrl, username, password string) error {
	log.Debug("Session.GetSession()")
	log.Debugf("url: %v, username: %v", baseUrl, username)
	so := &gitlab.GetSessionOptions{
		Login:    &username,
		Password: &password,
	}
	session, _, err := s.newClient(baseUrl, "").Session().GetSession(so)
	log.Debug("after c.Session.GetSession(so)", err)

	copier.Copy(s, session)

	return err
}

func (s *Session) GetToken() string {
	return s.PrivateToken
}

func (s *Session) GetId() int {
	return s.ID
}
