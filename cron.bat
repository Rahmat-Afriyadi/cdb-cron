@echo off

start cmd /k "cd /d D:\nextjs\mswkm && npm run start"

:loop

start cmd /k "cd /d D:\Golang\src\mswkm && go run server.go"

echo Loop 0 Executable crashed with exit code %errorlevel%. Restarting...
timeout /t 3

if %ERRORLEVEL% neq 0 (
    echo Error encountered, retrying...
    timeout /t 3 /nobreak >nul  :: Wait for 5 seconds before retrying
    goto loop
)

:loop1

start cmd /k "cd /d D:\Golang\src\cron && go run main.go"

echo Loop 1 Executable crashed with exit code %errorlevel%. Restarting...
timeout /t 3

if %ERRORLEVEL% neq 0 (
    echo Error encountered, retrying...
    timeout /t 3 /nobreak >nul  :: Wait for 5 seconds before retrying
    goto loop1
)


pause