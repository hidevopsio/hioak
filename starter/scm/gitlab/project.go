package gitlab

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/scm"
	"github.com/jinzhu/copier"
	"github.com/xanzy/go-gitlab"
)

type Project struct {
	scm.Project
	client ProjectInterface
}

func NewProject(c ProjectInterface) *Project {
	return &Project{
		client: c,
	}
}

func (p *Project) GetProject(baseUrl, id, token string) (int, int, error) {
	log.Debug("project.GetProject()")
	log.Debug("before c.project.GetProject(so)")
	project, _, err := p.client.GetProject(id)
	if err != nil {
		log.Error("Projects.GetProject err:", err)
		return 0, 0, err
	}
	return project.ID, project.Namespace.ID, err
}

func (p *Project) GetGroupId(url, token string, pid int) (int, error) {
	log.Debug("project.GetProject()")
	log.Debug("before c.Session.GetSession(so)")
	project, _, err := p.client.GetProject(pid)
	log.Debug("after c.project.GetProject(so)", project)
	return project.Namespace.ID, err
}

func (p *Project) ListProjects(baseUrl, token, search string, page int) ([]scm.Project, error) {
	log.Debug("project.ListProjects()")
	log.Debugf("url: %v", baseUrl)
	listProjectsOptions := &gitlab.ListProjectsOptions{}
	if search != "" {
		listProjectsOptions = &gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 200,
			},
			Search: &search,
		}
	} else {
		listProjectsOptions = &gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page: page,
			},
		}
	}
	ps, _, err := p.client.ListProjects(listProjectsOptions)
	log.Debugf("after project: %v", len(ps))
	var projects []scm.Project
	project := &scm.Project{}
	for _, pro := range ps {
		copier.Copy(project, pro)
		projects = append(projects, *project)
	}
	return projects, err
}

func (p *Project) Search(baseUrl, token, search string) ([]scm.Project, error) {
	log.Debug("Search.GetProjects()")
	log.Debug("before Search.project(so)", search)
	listProjectsOptions := &gitlab.ListProjectsOptions{
		Search: &search,
	}
	ps, _, err := p.client.ListProjects(listProjectsOptions)
	if err != nil {
		return nil, err
	}
	log.Debugf("after Search.project: %v", len(ps))
	var projects []scm.Project
	project := &scm.Project{}
	for _, pro := range ps {
		copier.Copy(project, pro)
		projects = append(projects, *project)
	}
	return projects, err
}
