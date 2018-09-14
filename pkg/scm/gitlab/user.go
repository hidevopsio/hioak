package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"net/http"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/jinzhu/copier"
	"strings"
	"github.com/hidevopsio/hioak/pkg/scm"
)

type User struct {
	scm.User
}


func (s *User) GetUser(baseUrl, accessToken string) (*scm.User, error) {
	log.Debug("Session get user")
	c := CheckToken(accessToken)
	c.SetBaseURL(baseUrl + ApiVersion)
	user, _, err := c.Users.CurrentUser()
	if err != nil {
		return nil, err
	}
	scmUser := &scm.User{}
	copier.Copy(scmUser, user)
	log.Debugf("get User : %v", user)
	return scmUser, nil

}

func CheckToken(token string) *gitlab.Client {
	len := strings.Count(token,"") - 1
	if len <= 20 {
		return gitlab.NewClient(&http.Client{}, token)
	}
	return gitlab.NewOAuthClient(&http.Client{}, token)
}
