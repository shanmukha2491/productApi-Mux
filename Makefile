check_install:

	which swagger || GO111MODULE=off go get -u github.com/go-swagger/cmd/swagger
swagger: check_install
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models


