module github.com/isan-rivkin/kubefigure

go 1.16

require (
	github.com/aws/aws-sdk-go v1.42.23
	github.com/hashicorp/consul/api v1.12.0
	github.com/hashicorp/go-hclog v1.1.0 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/hashicorp/terraform v0.15.2
	github.com/hashicorp/vault/api v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0
	github.com/zclconf/go-cty v1.8.3
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

//replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.2.0
//replace github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.12.0
