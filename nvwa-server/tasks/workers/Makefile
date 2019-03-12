PATH=github.com/nvwa-io/nvwa-io/nvwa-server/tasks/workers

deal-once:
	go test -count=1 -v ${PATH}  -run=Test_BuildDealOnce

jenkins-temp:
	go test -count=1 -v ${PATH}  -run=Test_JenkinsTemp

deploy-once:
	go test -count=1 -v ${PATH}  -run=Test_DeployDealOnce

