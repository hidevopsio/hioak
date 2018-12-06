package gitlab

import (
	"github.com/jinzhu/copier"
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/scm"
)

type Repository struct {
	scm.TreeNode
	newClient NewClient
}

type TreeNode struct {
	scm.TreeNode
}

// NewRepository need add test
func NewRepository(newClient NewClient) *Repository {
	return &Repository{
		newClient: newClient,
	}
}

func (r *Repository) ListTree(baseUrl, token, ref string, pid int) ([]scm.TreeNode, error) {
	log.Debug("Repository.ListTree()")
	log.Debugf("url: %v", baseUrl)
	opt := &gitlab.ListTreeOptions{
		RefName: &ref,
	}
	tree, _, err := r.newClient(baseUrl, token).Repository().ListTree(pid, opt)
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
