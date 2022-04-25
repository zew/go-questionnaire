@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=fmt&wave_id=2021-03&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
@REM CD c:\xampp\htdocs\go-questionnaire\cmd\transferrer\

SET JOBTIME=%date:~6,4%-%date:~3,2%-%date:~0,2%-%time:~0,5%
SET LOGFILE=import-%date:~6,4%-%date:~3,2%-%date:~0,2%.log
@REM ECHO "JOBTIME %JOBTIME%"
ECHO "LOGFILE %LOGFILE%"


@REM quotes will be in log - but I dont care
ECHO "  "             >>%LOGFILE%
ECHO %JOBTIME%        >>%LOGFILE%
ECHO "============="  >>%LOGFILE%


rm ./transferrer.exe
go build
ECHO "built finished"

@REM standalone execution...
@REM config dir is app-bucket
transferrer.exe -rmt=transferrer/fmt-remote.json  >>%LOGFILE% 2>&1
ECHO "transfer finished; see log file"

@REM COPY /Y  C:\xampp\htdocs\go-questionnaire\app-bucket\responses\downloaded\fmt-*.csv C:\xampp\htdocs\fmt\Mikrodaten-ger\


@REM CD C:\xampp\htdocs\fmt\
@REM php import-fmt-from-csv.php  >>c:\xampp\htdocs\go-questionnaire\cmd\transferrer\%LOGFILE%


@REM no pause - dont stall scheduler
@REM pause

