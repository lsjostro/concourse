module github.com/concourse/concourse

require (
	code.cloudfoundry.org/clock v0.0.0-20180518195852-02e53af36e6c
	code.cloudfoundry.org/credhub-cli v0.0.0-20190415201820-e3951663d25c
	code.cloudfoundry.org/garden v0.0.0-20181108172608-62470dc86365
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/localip v0.0.0-20170223024724-b88ad0dea95c
	code.cloudfoundry.org/urljoiner v0.0.0-20170223060717-5cabba6c0a50
	github.com/DataDog/datadog-go v3.2.0+incompatible
	github.com/DataDog/zstd v1.4.0
	github.com/Masterminds/squirrel v0.0.0-20190107164353-fa735ea14f09
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Microsoft/hcsshim v0.8.6 // indirect
	github.com/NYTimes/gziphandler v1.1.1
	github.com/The-Cloud-Source/goryman v0.0.0-20150410173800-c22b6e4a7ac1
	github.com/aryann/difflib v0.0.0-20170710044230-e206f873d14a
	github.com/aws/aws-sdk-go v1.25.18
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/cloudfoundry/go-socks5 v0.0.0-20180221174514-54f73bdb8a8e // indirect
	github.com/cloudfoundry/socks5-proxy v0.0.0-20180530211953-3659db090cb2 // indirect
	github.com/concourse/baggageclaim v1.6.5
	github.com/concourse/dex v0.0.0-20191105150807-8bbc2a6df1ea
	github.com/concourse/flag v1.0.0
	github.com/concourse/go-archive v1.0.1
	github.com/concourse/retryhttp v1.0.2
	github.com/containerd/cgroups v0.0.0-20191125132625-80b32e3c75c9 // indirect
	github.com/containerd/containerd v1.3.0
	github.com/containerd/continuity v0.0.0-20190827140505-75bee3e2ccb6 // indirect
	github.com/containerd/fifo v0.0.0-20190816180239-bda0ff6ed73c // indirect
	github.com/containerd/ttrpc v0.0.0-20191028202541-4f1b8fe65a5c // indirect
	github.com/containerd/typeurl v0.0.0-20190911142611-5eb25027c9fd // indirect
	github.com/coreos/go-oidc v2.0.0+incompatible
	github.com/cppforlife/go-semi-semantic v0.0.0-20160921010311-576b6af77ae4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fatih/color v1.7.0
	github.com/felixge/httpsnoop v1.0.0
	github.com/go-sql-driver/mysql v0.0.0-20160802113842-0b58b37b664c // indirect
	github.com/gobuffalo/packr v1.13.7
	github.com/gogo/googleapis v1.3.0 // indirect
	github.com/gogo/protobuf v1.3.0
	github.com/google/jsonapi v0.0.0-20180618021926-5d047c6bc66b
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/vault/api v1.0.5-0.20191108163347-bdd38fca2cff
	github.com/hashicorp/vault/sdk v0.1.14-0.20191112033314-390e96e22eb2 // indirect
	github.com/imdario/mergo v0.3.6
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/influxdata/influxdb1-client v0.0.0-20190118215656-f8cdb5d5f175
	github.com/jessevdk/go-flags v1.4.0
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/kr/pty v1.1.8
	github.com/krishicks/yaml-patch v0.0.10
	github.com/lib/pq v0.0.0-20181016162627-9eb73efc1fcc
	github.com/mattn/go-colorable v0.1.1
	github.com/mattn/go-isatty v0.0.7
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/maxbrunsfeld/counterfeiter/v6 v6.2.2
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/miekg/dns v1.1.6
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/onsi/ginkgo v1.9.0
	github.com/onsi/gomega v1.7.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/opencontainers/runtime-spec v1.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/peterhellberg/link v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/pkg/term v0.0.0-20190109203006-aa71e9d9e942
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/prometheus/client_golang v0.9.2
	github.com/racksec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/sirupsen/logrus v1.4.2
	github.com/skratchdot/open-golang v0.0.0-20160302144031-75fb7ed4208c
	github.com/square/certstrap v1.1.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	github.com/tedsuo/ifrit v0.0.0-20180802180643-bea94bb476cc
	github.com/tedsuo/rata v1.0.1-0.20170830210128-07d200713958
	github.com/vbauerster/mpb/v4 v4.6.1-0.20190319154207-3a6acfe12ac6
	github.com/vito/go-interact v0.0.0-20171111012221-fa338ed9e9ec
	github.com/vito/go-sse v0.0.0-20160212001227-fd69d275caac
	github.com/vito/houdini v1.1.1
	github.com/vito/twentythousandtonnesofcrudeoil v0.0.0-20180305154709-3b21ad808fcb
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/tools v0.0.0-20190829051458-42f498d34c4d // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1
	gopkg.in/yaml.v2 v2.2.2
	gotest.tools v2.2.0+incompatible // indirect
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a // indirect
	k8s.io/utils v0.0.0-20190829053155-3a4a5477acf8 // indirect
	sigs.k8s.io/yaml v1.1.0
)

go 1.13

replace github.com/docker/distribution v2.7.1+incompatible => github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible
