REM https://localhost:8083/survey/
REM check admin login
REM transferrer-endpoint?survey_id=fmt&wave_id=2021-03&fetch_all=1

REM go build && transferrer.exe -rmt=transferrer/remote-fmt-localhost.json

go build && transferrer.exe -rmt=transferrer/remote-fmt-localhost.json
pause
