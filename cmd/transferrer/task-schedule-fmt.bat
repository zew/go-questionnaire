
@REM https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/cc725744(v=ws.11)

@REM schtasks /QUERY /?
@REM SCHTASKS /create /?

@REM daily - 7:15 to 20:00
@REM /f  - overwrite if already exists
@REM /RI - repetition interval in minuutes
@REM /st - start time
@REM /du - duration after start time
@REM /ET - end   time  - nut used

@REM default log location C:\Windows\System32\winevt\Logs\Microsoft-Windows-TaskScheduler*.EVTX.


schtasks /delete  /tn "import-fmt-results" /F
schtasks /create  /tn "import-fmt-results" /tr "c:\xampp\htdocs\go-questionnaire\cmd\transferrer\run-fmt-remote.bat" /f /sc DAILY   /RI 60  /st 07:15  /du 13:00    


@REM Above task yields "permission denied"
@REM 
@REM We dont need remote or host params to solve this
@REM    schtasks /create /S domain /RU login /RP password  
@REM 
@REM Cause is some windows problem: social.technet.microsoft.com/Forums/windowsserver/en-US/639fd44b-e874-4c64-92d2-7a6c2388cbd2/schtasks-access-is-denied
@REM Solution: goto UI - export as XML - reimport as XML
@REM 
@REM Upon re-import change 
@REM    Execute indepenently of user login
@REM         LogonType S4U - service for user ; no active session ; InteractiveToken - User must be logged on
@REM    ExecutionTimeLimit to one hour
@REM    You may remove the idle setting: removed



pause