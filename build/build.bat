@echo off
setlocal


@REM call from app root via
@REM .\build\build.bat



@REM argument -o only works from within local directory
@REM cd ..\cmd\server\
@REM go build -v -race   -o "go-questionnaire.exe"


@REM this always creates main.exe


IF "%1"=="race-check" (
    REM since go 1.22 
    echo race-check doubles executable size, needs CGO
    SET CGO_ENABLED=1
    go build -v -race .\cmd\server\main.go
) ELSE (
    echo building without race detector
    echo    no CGO needed - executable not bloated
    echo    
    SET CGO_ENABLED=0
    go build -v       .\cmd\server\main.go
)




del    /s  go-questionnaire.exe
rename main.exe  go-questionnaire.exe
@REM upx -1 go-questionnaire.exe
go-questionnaire.exe


endlocal