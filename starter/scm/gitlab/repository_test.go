package gitlab_test

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
	"os"
	"testing"
)

func TestListTree(t *testing.T) {
	fs := new(fake.RepositoriesService)
	cli := &fake.Client{
		RepositoriesService: fs,
	}
	s := gitlab.NewRepository(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	file := &gogitlab.TreeNode{
		ID:   "chulei",
		Name: "aaaa",
	}
	var tree []*gogitlab.TreeNode
	tree = append(tree, file)
	resp := new(gogitlab.Response)
	fs.On("ListTree", nil, nil, nil).Return(tree, resp, nil)
	_, err := s.ListTree("", os.Getenv("Token"), "pom.xml", 1)
	assert.Equal(t, nil, err)
}
