package gitlab

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"os"
	"testing"
)

func TestListTree(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	file := &gitlab.TreeNode{
		ID:   "chulei",
		Name: "aaaa",
	}
	var tree []*gitlab.TreeNode
	tree = append(tree, file)
	resp := new(gitlab.Response)
	s.On("ListTree", nil, nil, nil).Return(tree, resp, nil)
	repository := NewRepository(s)
	_, err := repository.ListTree(baseUrl, os.Getenv("Token"), "pom.xml", 1)
	assert.Equal(t, nil, err)
}
