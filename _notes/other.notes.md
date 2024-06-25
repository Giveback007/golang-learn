## Start chrome from terminal
`start chrome`

---
https://superuser.com/questions/807883/dynamically-find-chrome-on-end-users-windows-machine
## Is there a way to dynamically determine where Chrome is installed?
The following command will determine where chrome is installed and set the CHROMEPATH environment variable to this value:
```
for /f "usebackq tokens=1,2,3,4,5" %a in (`reg query HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\ /s /f \chrome.exe ^| findstr Application`) do set CHROMEPATH=%c%d%e
```

Example output:
```
echo %CHROMEPATH%
C:\ProgramFiles(x86)\Google\Chrome\Application\chrome.exe
```

To use in a batch file you need to double up the percents as follows:
```
for /f "usebackq tokens=1,2,3,4,5" %%a in (`reg query HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\ /s /f \chrome.exe ^| findstr Application`) do set CHROMEPATH=%%c%%d%%e
```

### Other option:
where.exe chrome.exe