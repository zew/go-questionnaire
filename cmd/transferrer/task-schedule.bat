
@REM https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/cc725744(v=ws.11)

@REM schtasks /QUERY /?
@REM SCHTASKS /create /?

schtasks /delete /tn "import-fmt-results" 

@REM daily - 7:00 to 20:00
@REM /f  - overwrite if already exists
@REM /RI - repetition interval in minuutes
@REM /st - start time
@REM /du - duration after start time
@REM /ET - end   time  - nut used

@REM default log location C:\Windows\System32\winevt\Logs\Microsoft-Windows-TaskScheduler*.EVTX.
@REM redirection seems difficult:  >> c:\xampp\htdocs\go-questionnaire\cmd\transferrer\task-log.csv

schtasks /create /tn "import-fmt-results" /tr "\"c:\xampp\htdocs\go-questionnaire\cmd\transferrer\run-fmt-remote.bat\"" /f /sc DAILY   /RI 60  /st 07:00  /du 13:00    

pause