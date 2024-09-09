@echo off


:loop1
echo Running Go program... 
go run main.go
if %ERRORLEVEL% NEQ 0 (
    echo Program crashed with exit code %ERRORLEVEL%. Restarting...
    timeout /t 2 >nul
    goto loop1
) else (
    echo Program exited normally.
)


pause