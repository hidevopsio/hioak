package gitlab_test

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"os"
	"testing"
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
)

func TestGetRepositoty(t *testing.T) {
	fs := new(fake.RepositoryFilesService)
	cli := &fake.Client{
		RepositoryFilesService: fs,
	}
	s := gitlab.NewRepositoryFile(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	file := &gogitlab.File{
		FileName: "chulei",
	}
	resp := new(gogitlab.Response)
	fs.On("GetFile", nil, nil, nil).Return(file, resp, nil)
	_, err := s.GetRepository("", os.Getenv("Token"), "pom.xml", "master", 1)
	assert.Equal(t, nil, err)
}
