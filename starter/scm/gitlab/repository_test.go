package gitlab

import (
	"testing"
	"os"
	"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
)

func TestGetRepositoty(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	file := &gitlab.File{
		FileName: "chulei",
	}
	resp := new(gitlab.Response)
	s.On("GetFile", nil, nil, nil).Return(file, resp, nil)
	repository := NewRepository(s)
	_, err := repository.GetRepository(baseUrl, os.Getenv("Token"), "pom.xml", "master", 1)
	assert.Equal(t, nil, err)
}

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
