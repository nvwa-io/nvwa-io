PATH=github.com/nvwa-io/nvwa-io/nvwa-server/clients/jenkins

csrf:
	go test -count=1 -v ${PATH}  -run=Test_Csrf

create:
	go test -count=1 -v ${PATH}  -run=Test_Create

update:
	go test -count=1 -v ${PATH}  -run=Test_Update

is-exist:
	go test -count=1 -v ${PATH}  -run=Test_IsExist

delete:
	go test -count=1 -v ${PATH}  -run=Test_Delete

get-text-log:
	go test -count=1 -v ${PATH}  -run=Test_GetTextLog

get-status:
	go test -count=1 -v ${PATH}  -run=Test_GetStatus
