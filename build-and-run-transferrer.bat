cd transferrer
go build -o ..\transferrer.exe
cd ..
transferrer.exe -cfg=remote.json
timeout /t 10 /nobreak