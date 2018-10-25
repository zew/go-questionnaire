cd transferrer
go build -o ..\transferrer.exe
cd ..
transferrer.exe -cfg=remote.json
REM timeout /t 10 /nobreak
pause