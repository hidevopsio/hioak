package gitlab

import (
		"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/xanzy/go-gitlab"
	"github.com/jinzhu/copier"
	"github.com/hidevopsio/hioak/starter/scm"
)

type Repository struct {
	scm.TreeNode
	client ClientInterface
}

type TreeNode struct {
	scm.TreeNode
}


func NewRepository(c ClientInterface) scm.RepositoryInterface {
	return &Repository{
		client: c,
	}
}

func (r *Repository) GetRepository(baseUrl, token, filePath, ref string, pid int) (string, error) {
	log.Debug("Repository.Repository()")
	log.Debugf("url: %v", baseUrl)
	r.client.SetBaseURL(baseUrl + ApiVersion)
	opt := &gitlab.GetFileOptions{
		Ref: &ref,
		FilePath: &filePath,
	}
	file, _, err := r.client.GetFile(pid, opt)
	if err != nil {
		return "", err
	}
	return file.Content, nil
}

func (r *Repository) ListTree(baseUrl, token, ref string, pid int)  ([]scm.TreeNode, error){
	log.Debug("Repository.ListTree()")
	log.Debugf("url: %v", baseUrl)
	r.client.SetBaseURL(baseUrl + ApiVersion)
	opt := &gitlab.ListTreeOptions{
		RefName: &ref,
	}
	tree, _, err := r.client.ListTree(pid, opt)
	if err != nil {
		return nil, err
	}
	log.Info(tree)
	var treeNodes []scm.TreeNode
	for _, tr := range tree{
		treeNode := scm.TreeNode{}
		copier.Copy(&treeNode, tr)
		treeNodes = append(treeNodes, treeNode)
	}
	return treeNodes, nil
}
