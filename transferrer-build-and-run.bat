cd transferrer
go build -o ..\transferrer.exe
cd ..
transferrer.exe -cfg=transferrer.json  -rmt=transferrer-remote-mul.json
REM timeout /t 10 /nobreak
pause