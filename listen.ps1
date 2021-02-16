$fulldir    = "c:\Users\pbu\Documents\zew_work\git\go\go-questionnaire-v2\"

# inotifywait for windows
#   https://mcpmag.com/articles/2019/05/01/monitor-windows-folder-for-new-files.aspx
#   but stalls on reload
# better
#   https://github.com/cortesi/modd
$watcher = New-Object System.IO.FileSystemWatcher
$watcher.IncludeSubdirectories = $true
$watcher.Path = $fulldir
$watcher.EnableRaisingEvents = $true
$watcher.Filter = "*.ps1";
$watcher.Filter = "*.go";

$action =
{
    $path       = $event.SourceEventArgs.FullPath
    $changetype = $event.SourceEventArgs.ChangeType
 
    # again !
    $fulldir    = "c:\Users\pbu\Documents\zew_work\git\go\go-questionnaire-v2\"
    $exe        = "go-questionnaire.exe"
 
    Write-Host "$path was $changetype  $(get-date)"
    Get-Process | Where-Object {$_.Path -like "*go-questionnaire.exe*"} | Stop-Process -WhatIf
    Write-Host "killed prev  $(get-date)"

    # https://ss64.com/nt/start.html
    # to escape quotation marks -> prefix with backtick `"
    # the command - last argument needs *NO* quotes
    # Write-Host "start `"must-have-title--build-window`" /D $fulldir /WAIT go build"
    # start "must-have-title--build-window" /D "c:\Users\pbu\Documents\zew_work\git\go\go-questionnaire-v2\" /WAIT go build

    # https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/start-process
    Start-Process -Wait -FilePath "rm"  -ArgumentList $exe      -WorkingDirectory "c:\Users\pbu\Go\bin\"
    Write-Host "exe from bin $(get-date)"  

    # Start-Process       -FilePath "notepad"                    -WorkingDirectory "c:\WINDOWS" 
    # Start-Process -Wait -FilePath "go"  -ArgumentList "build " -WorkingDirectory "c:\Users\pbu\Documents\zew_work\git\go\go-questionnaire-v2\"
    Start-Process -Wait -FilePath "go"  -ArgumentList "build -x" -WorkingDirectory $fulldir

    Write-Host "build compl  $(get-date)"  
    Start-Process       -FilePath $exe  -ArgumentList "      " -WorkingDirectory $fulldir
    Write-Host "restartet at $(get-date)"

}

Register-ObjectEvent $watcher 'Created' -Action $action
Register-ObjectEvent $watcher 'Changed' -Action $action
Register-ObjectEvent $watcher 'Renamed' -Action $action
Register-ObjectEvent $watcher 'Deleted' -Action $action

Write-Host "watcher setup complete"
# echo "watcher setup complete"
# pause

