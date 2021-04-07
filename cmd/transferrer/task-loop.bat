:START
REM Execute the MS-DOS dir command ever 3600 seconds.
@REM dir
DATE /T >> importer.log
TIME /T >> importer.log
echo:  >> importer.log
run-fmt-remote.bat >> importer.log
SLEEP 3600
GOTO START
