package factories

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"errors"
	"github.com/hidevopsio/hioak/starter/scm"
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
)

func (s *ScmFactory) NewGroup(provider int) (scm.GroupInterface, error)  {
	log.Debug("scm.NewSession()")
	switch provider {
	case GithubScmType:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not implemented\n", provider))
	case GitlabScmType:
		return new(gitlab.Group), nil
	default:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not recognized\n", provider))
	}
	return nil, nil
}

