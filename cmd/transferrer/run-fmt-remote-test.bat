@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=fmt&wave_id=2021-03&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
CD c:\goprojects\go-questionnaire\cmd\transferrer\

SET JOBTIME=%date:~6,4%-%date:~3,2%-%date:~0,2%-%time:~0,5%
mkdir "logs-fmt"
SET LOGFILE=.\logs-fmt\import-%date:~6,4%-%date:~3,2%-%date:~0,2%.log


@REM quotes will be in log - but I dont care
ECHO "  "             >>%LOGFILE%
ECHO %JOBTIME%        >>%LOGFILE%
ECHO "============="  >>%LOGFILE%


@REM built newly or not
@REM rm ./transferrer.exe
@REM go build

transferrer.exe -rmt=transferrer/fmt-remote.json  >>%LOGFILE% 2>&1

exit /b

@REM COPY /Y  C:\goprojects\go-questionnaire\app-bucket\responses\downloaded\fmt-*.csv C:\xampp\htdocs\fmt\Mikrodaten-ger\


@REM CD C:\xampp\htdocs\fmt\
@REM php import-fmt-from-csv.php  >>c:\goprojects\go-questionnaire\cmd\transferrer\%LOGFILE%


@REM no pause - dont stall scheduler
@REM pause


@REM call FILENAME
@REM call "run-kneb-remote.bat"
