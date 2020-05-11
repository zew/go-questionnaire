certutil –addstore -enterprise –f "Root" server.pem


REM https://superuser.com/questions/1031444/importing-pem-certificates-on-windows-7-on-the-command-line
REM certutil.exe
REM certutil –addstore -enterprise –f “Root” <pathtocertificatefile>
REM 
REM mmc
REM 	snap-in-hinzufügen - Zertifikate
REM 		lokaler computer
REM 			dieser computer


REM 
REM http://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate
REM 	chrome://flags/#allow-insecure-localhost
REM 
REM http://stackoverflow.com/questions/25692084/force-chrome-to-accept-any-ssl-certificate-regardless-of-who-it-was-signed-by	
REM 	chrome  --ignore-certificate-errors
REM 
REM 