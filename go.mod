module hidevops.io/hioak

go 1.12

require (
	github.com/Microsoft/go-winio v0.4.11
	github.com/Shopify/goreferrer v0.0.0-20180807163728-b9777dc9f9cc
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/docker/distribution v2.6.2+incompatible
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.3.3
	github.com/emirpasic/gods v1.12.0
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.1.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0
	github.com/google/go-querystring v1.0.0
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/googleapis/gnostic v0.2.0
	github.com/hashicorp/hcl v1.0.0
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c
	github.com/imdario/mergo v0.3.6
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/iris-contrib/blackfriday v2.0.0+incompatible
	github.com/iris-contrib/formBinder v0.0.0-20171010160137-ad9fb86c356f
	github.com/iris-contrib/go.uuid v2.0.0+incompatible
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/json-iterator/go v1.1.5
	github.com/kataras/golog v0.0.0-20180321173939-03be10146386
	github.com/kataras/iris v11.0.3+incompatible
	github.com/kataras/pio v0.0.0-20180511174041-a9733b5b6b83
	github.com/kevinburke/ssh_config v0.0.0-20180830205328-81db2a75821e
	github.com/klauspost/compress v1.4.0
	github.com/klauspost/cpuid v0.0.0-20180405133222-e7e905edc00e
	github.com/konsorten/go-windows-terminal-sequences v1.0.1
	github.com/magiconair/properties v1.8.0
	github.com/microcosm-cc/bluemonday v1.0.1
	github.com/mitchellh/go-homedir v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/openshift/api v3.9.0+incompatible
	github.com/openshift/client-go v3.9.0+incompatible
	github.com/pelletier/go-buffruneio v0.2.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/common v0.0.0-20181020173914-7e9e6cabbd39
	github.com/sergi/go-diff v1.0.0
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/src-d/gcfg v1.4.0
	github.com/stretchr/objx v0.1.1
	github.com/stretchr/testify v1.2.2
	github.com/xanzy/go-gitlab v0.0.0-20170825130035-896163fa8f7a
	github.com/xanzy/ssh-agent v0.2.0
	golang.org/x/crypto v0.0.0-20190211182817-74369b46fc67
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/sys v0.0.0-20181031143558-9b800f95dbbc
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/inf.v0 v0.9.1
	gopkg.in/src-d/go-billy.v4 v4.3.0
	gopkg.in/src-d/go-git.v4 v4.7.1
	gopkg.in/warnings.v0 v0.1.2
	gopkg.in/yaml.v2 v2.2.1
	hidevops.io/hiboot v1.0.4
	hidevops.io/viper v1.3.2
	k8s.io/api v0.0.0-20180601181742-8b7507fac302
	k8s.io/apiextensions-apiserver v0.0.0-20180601203502-8e7f43002fec
	k8s.io/apimachinery v0.0.0-20180601181227-17529ec7eadb
	k8s.io/client-go v7.0.0+incompatible
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.36.0
	golang.org/x/build => github.com/golang/build v0.0.0-20190215225244-0261b66eb045
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181030022821-bc7917b19d8f
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190212162250-21964bba6549
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181217174547-8f45f776aaf1
	golang.org/x/net => github.com/golang/net v0.0.0-20181029044818-c44066c5c816
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20181017192945-9dcd33a902f4
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190124201629-844a5f5b46f4
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181029174526-d69651ed3497
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190214204934-8dcb7bc8c7fe
	golang.org/x/vgo => github.com/golang/vgo v0.0.0-20180912184537-9d567625acf4
	google.golang.org/api => github.com/googleapis/googleapis v0.0.0-20190215163516-1a4f0f12777d
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190215211957-bd968387e4aa
	google.golang.org/grpc => github.com/grpc/grpc-go v1.14.0
)
