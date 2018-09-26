package gitlab

import (
			"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/jinzhu/copier"
		"github.com/hidevopsio/hioak/starter/scm"
)

type User struct {
	scm.User
	client ClientInterface
}

func NewUser(c ClientInterface) scm.UserInterface {
	return &User{
		client: c,
	}
}

func (s *User) GetUser(baseUrl, accessToken string) (*scm.User, error) {
	log.Debug("Session get user")
	s.client.SetBaseURL(baseUrl + ApiVersion)
	user, _, err := s.client.CurrentUser()
	if err != nil {
		return nil, err
	}
	scmUser := &scm.User{}
	copier.Copy(scmUser, user)
	log.Debugf("get User : %v", user)
	return scmUser, nil

}
