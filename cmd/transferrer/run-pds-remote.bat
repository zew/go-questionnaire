@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=pds&wave_id=2021-03&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
CD c:\xampp\htdocs\go-questionnaire\cmd\transferrer\

SET JOBTIME=%date:~6,4%-%date:~3,2%-%date:~0,2%-%time:~0,5%
SET LOGFILE=pds-import-%date:~6,4%-%date:~3,2%-%date:~0,2%.log


@REM quotes will be in log - but I dont care
ECHO "  "             >>%LOGFILE%
ECHO %JOBTIME%        >>%LOGFILE%
ECHO "============="  >>%LOGFILE%


@REM rm ./transferrer.exe
@REM go build

transferrer.exe -rmt=transferrer/pds-remote.json  >>%LOGFILE% 2>&1


COPY /Y  C:\xampp\htdocs\go-questionnaire\app-bucket\responses\downloaded\pds-*.csv C:\xampp\htdocs\fmt\export\pcag\


