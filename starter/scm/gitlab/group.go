package gitlab

import (
	"github.com/jinzhu/copier"
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type Group struct {
	scm.Group
	group  GroupInterface
	newClient NewClient
}

func NewGroup(newClient NewClient) *Group {
	return &Group{
		newClient: newClient,
	}
}

func (g *Group) ListGroups(token, baseUrl string, page int) ([]scm.Group, error) {
	log.Debug("group.ListGroups()")
	var scmGroups []scm.Group
	scmGroup := &scm.Group{}
	log.Debug("before c.group.ListGroups(so)")
	opt := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page: page,
		},
	}
	groups, _, err := g.newClient(baseUrl, token).Group().ListGroups(opt)
	if err != nil {
		return nil, err
	}
	log.Debug("after c.Group.ListGroups(so)")
	for _, group := range groups {
		copier.Copy(scmGroup, group)
		scmGroups = append(scmGroups, *scmGroup)
	}
	return scmGroups, err
}

func (g *Group) GetGroup(token, baseUrl string, gid int) (*scm.Group, error) {
	log.Debug("group.GetGroup()")
	scmGroup := &scm.Group{}
	log.Debug("before c.group.ListGroups(so)")
	group, _, err := g.newClient(baseUrl, token).Group().GetGroup(gid)
	log.Debug("after c.Session.GetSession(so)")
	if err != nil {
		return nil, err
	}
	copier.Copy(scmGroup, group)
	return scmGroup, err
}

func (g *Group) ListGroupProjects(token, baseUrl string, gid, page int) ([]scm.Project, error) {
	log.Debug("group.ListGroups()")
	var scmProjects []scm.Project
	scmProject := &scm.Project{}
	log.Debug("before c.group.ListGroups")
	opt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page: page,
		},
	}
	projects, _, err := g.newClient(baseUrl, token).Group().ListGroupProjects(gid, opt)
	log.Debug("ListGroupProjects : ", len(projects))
	if err != nil {
		log.Error("Group ListGroupProjects : ", err)
		return nil, err
	}
	for _, project := range projects {
		copier.Copy(scmProject, project)
		scmProjects = append(scmProjects, *scmProject)
	}
	return scmProjects, err
}
