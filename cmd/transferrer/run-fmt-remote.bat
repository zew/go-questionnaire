REM check admin login
REM transferrer-endpoint?survey_id=fmt&wave_id=2021-03&fetch_all=1
cls
rm ./transferrer.exe
go build && transferrer.exe -rmt=transferrer/fmt-remote.json

pause
