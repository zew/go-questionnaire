@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=kneb1&wave_id=2021-03&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
CD c:\xampp\htdocs\go-questionnaire\cmd\transferrer\

SET JOBTIME=%date:~6,4%-%date:~3,2%-%date:~0,2%-%time:~0,5%
mkdir "logs-kneb"
SET LOGFILE=.\logs-kneb\import-%date:~6,4%-%date:~3,2%-%date:~0,2%.log


@REM quotes will be in log - but I dont care
ECHO "  "             >>%LOGFILE%
ECHO %JOBTIME%        >>%LOGFILE%
ECHO "============="  >>%LOGFILE%


@REM built newly or not
@REM rm ./transferrer.exe
@REM go build

transferrer.exe -rmt=transferrer/kneb-remote.json  >>%LOGFILE% 2>&1

COPY /Y  C:\xampp\htdocs\go-questionnaire\app-bucket\responses\downloaded\kneb1-*.csv C:\xampp\htdocs\kneb\

