package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"net/http"
	"strings"
)

type ClientInterface interface {
	SetBaseURL(baseUrl string) error
	GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error)
	ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
	GetGroup(gid interface{}, options ...gitlab.OptionFunc) (*gitlab.Group, *gitlab.Response, error)
	ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
	ListGroupMembers(gid interface{}, opt *gitlab.ListGroupMembersOptions, options ...gitlab.OptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error)
	GetProject(pid interface{}, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error)
	ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
	GetProjectMember(pid interface{}, user int, options ...gitlab.OptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)
	GetFile(pid interface{}, opt *gitlab.GetFileOptions, options ...gitlab.OptionFunc) (*gitlab.File, *gitlab.Response, error)
	ListTree(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.OptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error)
	CurrentUser(options ...gitlab.OptionFunc) (*gitlab.User, *gitlab.Response, error)
}

func NewClient(token string) *gitlab.Client {
	len := strings.Count(token, "") - 1
	if len <= 20 {
		return gitlab.NewClient(&http.Client{}, token)
	}
	return gitlab.NewOAuthClient(&http.Client{}, token)
}
