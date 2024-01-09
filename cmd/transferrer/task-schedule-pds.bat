schtasks /delete  /tn "import-pds-results" /F
schtasks /create  /tn "import-pds-results" /tr "c:\xampp\htdocs\go-questionnaire\cmd\transferrer\run-pds-remote.bat" /f /sc DAILY   /RI 60  /st 07:15  /du 13:00    

pause