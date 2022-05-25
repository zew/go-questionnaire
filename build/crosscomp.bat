@REM from app root

SET GOOS=solaris
SET GOOS=openbsd
SET GOOS=linux

SET GOARCH=386
SET GOARCH=amd64

@REM go build -v github.com\zew\go-questionnaire -o go-questionnaire-new

go build -v  .\cmd\server\main.go
del    /s  go-questionnaire-new
rename main  go-questionnaire-new



pause