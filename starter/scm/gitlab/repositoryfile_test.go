package gitlab

import (
	"testing"
	"os"
	"github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
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
	repository := NewRepositoryFile(s)
	_, err := repository.GetRepository(baseUrl, os.Getenv("Token"), "pom.xml", "master", 1)
	assert.Equal(t, nil, err)
}
