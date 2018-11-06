package gitlab

import (
	"github.com/jinzhu/copier"
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type Repository struct {
	scm.TreeNode
	client NewClient
}

type TreeNode struct {
	scm.TreeNode
}

func NewRepository(c NewClient) *Repository {
	return &Repository{
		client: c,
	}
}

func (r *Repository) ListTree(baseUrl, token, ref string, pid int) ([]scm.TreeNode, error) {
	log.Debug("Repository.ListTree()")
	log.Debugf("url: %v", baseUrl)
	opt := &gitlab.ListTreeOptions{
		RefName: &ref,
	}
	tree, _, err := r.client(baseUrl, token).Repository().ListTree(pid, opt)
	if err != nil {
		return nil, err
	}
	log.Info(tree)
	var treeNodes []scm.TreeNode
	for _, tr := range tree {
		treeNode := scm.TreeNode{}
		copier.Copy(&treeNode, tr)
		treeNodes = append(treeNodes, treeNode)
	}
	return treeNodes, nil
}
