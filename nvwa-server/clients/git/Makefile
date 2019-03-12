PATH=github.com/nvwa-io/nvwa-io/nvwa-server/clients/git

clone:
	go test -count=1 -v ${PATH}  -run=Test_Clone

pull:
	go test -count=1 -v ${PATH}  -run=Test_Pull

fetch:
	go test -count=1 -v ${PATH}  -run=Test_Fetch

checkout:
	go test -count=1 -v ${PATH}  -run=Test_Checkout

checkout-tag:
	go test -count=1 -v ${PATH}  -run=Test_CheckoutTag

branches:
	go test -count=1 -v ${PATH}  -run=Test_Branches

tags:
	go test -count=1 -v ${PATH}  -run=Test_Tags
