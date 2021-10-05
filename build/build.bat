@REM call from app root via
@REM .\build\build.bat


@REM argument -o only works from within local directory
@REM cd ..\cmd\server\
@REM go build -v -race   -o "go-questionnaire.exe"


@REM this always creates main.exe
go build -v -race .\cmd\server\main.go
del    /s  go-questionnaire.exe
rename main.exe  go-questionnaire.exe
@REM upx -1 go-questionnaire.exe
go-questionnaire.exe
