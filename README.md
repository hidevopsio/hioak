# hioak

<p align="center">
  <a href="https://travis-ci.org/hidevopsio/hioak?branch=master">
    <img src="https://travis-ci.org/hidevopsio/hioak.svg?branch=master" alt="Build Status"/>
  </a>
  <a href="https://codecov.io/gh/hidevopsio/hioak">
    <img src="https://codecov.io/gh/hidevopsio/hioak/branch/master/graph/badge.svg" />
  </a>
  <a href="https://opensource.org/licenses/Apache-2.0">
      <img src="https://img.shields.io/badge/License-Apache%202.0-green.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/hidevopsio/hioak">
      <img src="https://goreportcard.com/badge/github.com/hidevopsio/hioak" />
  </a>
  <a href="https://godoc.org/github.com/hidevopsio/hioak">
      <img src="https://godoc.org/github.com/golang/gddo?status.svg" />
  </a>
</p>

hioak is a library for orchestration client, it includes docker, k8s, openshift, kong, istio client

## 简介

hioak 是一个组件 结合了`hiboot`框架
分别给`docker` `k8s` `openshift` `kong` `istio`提供client
![docker](https://konghq.com/wp-content/uploads/2017/08/docker.svg)  ![k8s](https://konghq.com/wp-content/uploads/2017/08/kubernetes.svg)  ![openshift](https://ss1.baidu.com/6ONXsjip0QIZ8tyhnq/it/u=114369328,3108566132&fm=58&bpow=2000&bpoh=2000)  ![kong](https://konghq.com/wp-content/themes/konghq/assets/img/gradient-logo.svg)  ![istio](http://img2.imgtn.bdimg.com/it/u=3385702251,4239533823&fm=26&gp=0.jpg)
## 快速开始

hioak的包目录结构为：

```yml
starter
|
|____docker # docker的client
|
|____________fake #  docker mock 测试方法
|
|____kube  # k8s的client
|
|____________fake
|
|____openshift #  openshift的client
|
|_____________fake
|
|____scm  # scm的client
|
|_____________fake
|
```

### Docker

我们的docker包主要提供了一以下5种能力 分别是镜像的拉取， 打tag镜像, 推镜像, 查看所有镜像 和镜像的构建，如果你还需要docker更多的能力可以查看管方文档。

[https://docs.docker.com/engine/api/v1.24](https://docs.docker.com/engine/api/v1.24)

```go
type ClientInterface interface {
	ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error)
	ImageTag(ctx context.Context, imageID, ref string) error
	ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error)
	ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
	ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error)
}
```

### Kube

在`kube`包 我们主要是调用了k8s 提供的`kube-client`来调集群的`api-server`。
主要我们对 k8s的资源进行了增删改查的操作， 主要的资源有以下10个
`crd` `configmap` `deployment` `events` `pod` `rc` `rs` `secret` `service` `token`。
如果想知道这些资源的作用, 可查看官方文档：

[https://kubernetes.io/zh/docs](https://kubernetes.io/zh/docs)

### openshift

Openshift是一个开源容器云平台，是一个基于主流的容器技术Docker和K8s构建的云平台。Openshift底层以Docker作为容器引擎驱动，以K8s作为容器编排引擎组件，并提供了开发语言，中间件，DevOps自动化流程工具和web console用户界面等元素，提供了一套完整的基于容器的应用云平台。

我们在openshift包中主要对
`build` `dc` `image` `imagetag` `project` `rolebinding` `route` 这些资源进行操作

### SCM

scm主要是调用gitlab的Api 来对代码进行管理 主要提供了以下的接口

```go
type GroupInterface interface {
	ListGroups(token, baseUrl string, page int) ([]Group, error)
	GetGroup(token, baseUrl string, gid int) (*Group, error)
	ListGroupProjects(token, baseUrl string, gid, page int) ([]Project, error)
}

type GroupMemberInterface interface {
	ListGroupMembers(token, baseUrl string, gid, uid int) (int, error)
	GetGroupMember(token, baseUrl string, gid, uid int) (*GroupMember, error)
}

type ProjectInterface interface {
	ListProjects(baseUrl, token, search string, page int) ([]Project, error)
	GetGroupId(url, token string, pid int) (int, error)
	GetProject(baseUrl, id, token string) (int, int, error)
	Search(baseUrl, token, search string) ([]Project, error)
}

type RepositoryInterface interface {
	ListTree(baseUrl, token, ref string, pid int) ([]TreeNode, error)
}

type RepositoryFileInterface interface {
	GetRepository(baseUrl, token, filePath, ref string, pid int) (string, error)
}

type SessionInterface interface {
	GetSession(baseUrl, username, password string) error
	GetToken() string
	GetId() int
}

type UserInterface interface {
	GetUser(baseUrl, accessToken string) (*User, error)
}
```