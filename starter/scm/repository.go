package scm

type TreeNode struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Mode string `json:"mode"`
}

type RepositoryInterface interface {
	ListTree(baseUrl, token, ref string, pid int) ([]TreeNode, error)
}

type RepositoryFileInterface interface {
	GetRepository(baseUrl, token, filePath, ref string, pid int) (string, error)
}