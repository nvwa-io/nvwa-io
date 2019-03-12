module github.com/nvwa-io/nvwa-io/nvwa-agent

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180802221240-56440b844dfe
	golang.org/x/net => github.com/golang/net v0.0.0-20180811021610-c39426892332
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180810173357-98c5dad5d1a0
	golang.org/x/text => github.com/golang/text v0.3.0
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190110114555-6a25665e652a
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	golang.org/x/net v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/resty.v1 v1.10.3
)
