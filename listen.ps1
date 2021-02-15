# inotifywait for windows
#   https://mcpmag.com/articles/2019/05/01/monitor-windows-folder-for-new-files.aspx
#   but stalls on reload
# better 
#   https://github.com/cortesi/modd
$watcher = New-Object System.IO.FileSystemWatcher
$watcher.IncludeSubdirectories = $true
$watcher.Path = 'c:\Users\pbu\Documents\zew_work\git\go\go-questionnaire-v2\'
$watcher.EnableRaisingEvents = $true
$watcher.Filter = "*.go";

$action =
{
    $path       = $event.SourceEventArgs.FullPath
    $changetype = $event.SourceEventArgs.ChangeType
    Write-Host "$path was $changetype  $(get-date)"
    Get-Process | Where-Object {$_.Path -like "*go-questionnaire.exe*"} | Stop-Process -WhatIf
    Write-Host "killed prev  $(get-date)"
    go build
    Write-Host "build compl  $(get-date)"
    start go-questionnaire.exe
    Write-Host "restartet at $(get-date)"

}

Register-ObjectEvent $watcher 'Created' -Action $action
Register-ObjectEvent $watcher 'Changed' -Action $action
Register-ObjectEvent $watcher 'Renamed' -Action $action
Register-ObjectEvent $watcher 'Deleted' -Action $action

Write-Host "watcher setup complete"
# echo "watcher setup complete"
# pause