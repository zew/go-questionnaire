SET GOOS=solaris
SET GOOS=openbsd
SET GOOS=linux

SET GOARCH=386
SET GOARCH=amd64

REM extension 'exe' for linux too, so that the *.sh scripts work for both compilates
go build -v github.com\zew\go-questionnaire\updater.exe

pause