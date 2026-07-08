@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=lix&wave_id=2026-06&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
@REM CD c:\xampp\htdocs\go-questionnaire\cmd\transferrer\
CD C:\goprojects\go-questionnaire\cmd\transferrer\

SET JOBTIME=%date:~6,4%-%date:~3,2%-%date:~0,2%-%time:~0,5%
mkdir "logs-lix"
SET LOGFILE=.\logs-lix\lix-import-%date:~6,4%-%date:~3,2%-%date:~0,2%.log

@REM quotes will be in log - but I dont care
ECHO "  "             >>%LOGFILE%
ECHO %JOBTIME%        >>%LOGFILE%
ECHO "============="  >>%LOGFILE%


@REM built newly or not
@REM rm ./transferrer.exe
@REM go build

@REM pwd
@REM /c/goprojects/go-questionnaire/cmd/transferrer
@REM /c/goprojects/go-questionnaire/app-bucket/transferrer

@REM transferrer.exe -rmt=transferrer/lix-remote.json    
transferrer.exe -rmt=transferrer/lix-remote.json    >>%LOGFILE% 2>&1


@REM COPY /Y  C:\goprojects\go-questionnaire\app-bucket\responses\downloaded\lix-*.csv C:\xampp\htdocs\fmt\export\lix\


echo "finished"