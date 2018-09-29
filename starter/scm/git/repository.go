package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hiboot/pkg/utils/io"
	"path/filepath"
)

type Repository interface {
	Clone(url, branch, dir, username, password string) (string, error)
}

type CloneFunc func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error)

type repository struct {
	Repository
	cloneFunc CloneFunc
}

func NewRepository(cloneFunc ...CloneFunc) Repository {
	cf := git.PlainClone
	if len(cloneFunc) != 0 {
		cf = cloneFunc[0]
	}

	return &repository{
		cloneFunc: cf,
	}
}

// Clone the given repository to the given directory
func (r *repository) Clone(url, branch, dir, username, password string) (string, error) {
	projectName := io.Filename(url)
	projectName = io.Basename(projectName)
	projectPath := filepath.Join(dir, projectName)

	log.Debugf("git clone %s %s", url, projectPath)
	_, err := r.cloneFunc(projectPath, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              transport.AuthMethod(&http.BasicAuth{Username: username, Password: password}),
	})
	if err != nil {
		log.Debugf("Error", err)
		return "", err
	}
	return projectPath, nil
}