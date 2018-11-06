package factories

import (
	"errors"
	"fmt"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
	"hidevops.io/hioak/starter/scm/gitlab"
)

func (s *ScmFactory) NewGroupMember(provider int) (scm.GroupMemberInterface, error) {
	log.Debug("scm.NewSession()")
	switch provider {
	case GithubScmType:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not implemented\n", provider))
	case GitlabScmType:
		return new(gitlab.GroupMember), nil
	default:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not recognized\n", provider))
	}
	return nil, nil
}
