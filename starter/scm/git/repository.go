package git

import (
	"gopkg.in/src-d/go-git.v4"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/utils/io"
	"path/filepath"
)

type Repository interface {
	Clone(cloneOptions *git.CloneOptions, destDir string) (string, error)
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
func (r *repository) Clone(cloneOptions *git.CloneOptions, destDir string) (string, error) {
	projectName := io.Filename(cloneOptions.URL)
	projectName = io.Basename(projectName)
	projectPath := filepath.Join(destDir, projectName)

	log.Debugf("git clone %s %s", cloneOptions.URL, projectPath)
	_, err := r.cloneFunc(projectPath, false, cloneOptions)
	if err != nil {
		log.Debugf("Error", err)
		return "", err
	}
	return projectPath, nil
}
