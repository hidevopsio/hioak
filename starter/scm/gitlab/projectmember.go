package gitlab

import (
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type ProjectMember struct {
	scm.ProjectMember
	client NewClient
}

func NewProjectMember(c NewClient) *ProjectMember {
	return &ProjectMember{
		client: c,
	}
}

func (p *ProjectMember) GetProjectMember(token, baseUrl string, pid, uid, gid int) (scm.ProjectMember, error) {
	log.Debug("Product.GetProject()")
	scmProjectMember := scm.ProjectMember{}
	log.Debug("before p.project.GetProjectMember(so)", pid)
	projectMember, _, err := p.client(baseUrl, token).ProjectMember().GetProjectMember(pid, uid)
	if err != nil {
		log.Error("ProjectMembers.GetProjectMember ", err)
		return scmProjectMember, err
	}
	log.Debug("after c.Session.GetSession(so)")
	for id, permissions := range scm.Permissions {
		if projectMember.AccessLevel == id {
			scmProjectMember.ProjectPermissions = permissions
			return scmProjectMember, nil
		}
	}
	return scmProjectMember, err
}
