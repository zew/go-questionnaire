REM run from app root
go test -v ./... 
REM go test -v ./... -covermode=count -coverprofile=count.log github.com/zew/util
REM go test -v ./... -covermode=count -coverprofile=count.log github.com/zew/exceldb
REM go tool cover -html=count.log

