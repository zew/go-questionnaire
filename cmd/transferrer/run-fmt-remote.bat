@echo off

REM check admin login at remote host
REM transferrer-endpoint?survey_id=fmt&wave_id=2021-03&fetch_all=1
cls

@REM ENVIRONMENT is restored, whenever this batch exits; change directory is reset
setlocal

@REM for execution as scheduled task
cd c:\xampp\htdocs\go-questionnaire\cmd\transferrer\

rm ./transferrer.exe
go build

transferrer.exe -rmt=transferrer/fmt-remote.json  >>importer-1.log 2>&1

copy /Y  C:\xampp\htdocs\go-questionnaire\app-bucket\responses\downloaded\fmt-*.csv C:\xampp\htdocs\fmt\Mikrodaten-ger\


cd C:\xampp\htdocs\fmt\
php import-fmt-from-csv.php  >>c:\xampp\htdocs\go-questionnaire\cmd\transferrer\importer-1.log


@REM no pause - dont stall scheduler
@REM pause

