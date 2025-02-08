@REM call from app root via
@REM .\build\build.bat


@REM argument -o only works from within local directory
@REM cd ..\cmd\server\
@REM go build -v -race   -o "go-questionnaire.exe"


@REM this always creates main.exe

@REM -race needs CGO since 1.22 
SET CGO_ENABLED=1
@REM SET CGO_ENABLED=0

go build -v -race .\cmd\server\main.go
@REM go build -v       .\cmd\server\main.go

del    /s  go-questionnaire.exe
rename main.exe  go-questionnaire.exe
@REM upx -1 go-questionnaire.exe
go-questionnaire.exe
