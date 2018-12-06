package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type RepositoryFile struct {
	scm.TreeNode
	newClient NewClient
}

func NewRepositoryFile(newClient NewClient) *RepositoryFile {
	return &RepositoryFile{
		newClient: newClient,
	}
}

func (r *RepositoryFile) GetRepository(baseUrl, token, filePath, ref string, pid int) (string, error) {
	log.Debug("Repository.Repository()")
	log.Debugf("url: %v", baseUrl)
	opt := &gitlab.GetFileOptions{
		Ref:      &ref,
		FilePath: &filePath,
	}
	file, _, err := r.newClient(baseUrl, token).RepositoryFile().GetFile(pid, opt)
	if err != nil {
		return "", err
	}
	return file.Content, nil
}
