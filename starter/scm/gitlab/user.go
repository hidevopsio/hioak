package gitlab

import (
	"github.com/jinzhu/copier"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type User struct {
	scm.User
	newClient NewClient
}

func NewUser(newClient NewClient) *User {
	return &User{
		newClient: newClient,
	}
}

func (s *User) GetUser(baseUrl, accessToken string) (*scm.User, error) {
	log.Debug("Session get user")
	user, _, err := s.newClient(baseUrl, accessToken).User().CurrentUser()
	if err != nil {
		return nil, err
	}
	scmUser := &scm.User{}
	copier.Copy(scmUser, user)
	log.Debugf("get User : %v", user)
	return scmUser, nil

}
