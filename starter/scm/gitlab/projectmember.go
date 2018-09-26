package gitlab

import (
	"github.com/xanzy/go-gitlab"
		"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/scm"
)
type ProjectMember struct {
	scm.ProjectMember
	client ClientInterface
}


func NewProjectMember(c ClientInterface) scm.ProjectMemberInterface {
	return &ProjectMember{
		client: c,
	}
}

func (p *ProjectMember) GetProjectMember(token, baseUrl string, pid, uid, gid int) (scm.ProjectMember, error) {
	log.Debug("Product.GetProject()")
	scmProjectMember := scm.ProjectMember{}
	p.client.SetBaseURL(baseUrl + ApiVersion)
	log.Debug("before p.project.GetProjectMember(so)", pid)
	opt := &gitlab.ListGroupMembersOptions{
		ListOptions: gitlab.ListOptions{
			Page: 50,
		},
	}
	groupMembers, _, err := p.client.ListGroupMembers(gid, opt)
	if err != nil {
		log.Error("Groups.ListGroupMembers err:", err)
		return scmProjectMember, err
	}
	log.Debug("Groups.ListGroupMembers ")
	for _, groupMember := range groupMembers {
		if groupMember.ID == uid {
			for id, permissions := range scm.Permissions  {
				if groupMember.AccessLevel == id {
					scmProjectMember.ProjectPermissions = permissions
					return scmProjectMember, nil
				}
			}
		}
	}
	projectMember, _, err := p.client.GetProjectMember(pid, uid)
	if err != nil {
		log.Error("ProjectMembers.GetProjectMember ", err)
		return scmProjectMember, err
	}
	log.Debug("after c.Session.GetSession(so)")
	for id, permissions := range scm.Permissions  {
		if projectMember.AccessLevel == id {
			scmProjectMember.ProjectPermissions = permissions
			return scmProjectMember, nil
		}
	}
	return scmProjectMember, err
}
