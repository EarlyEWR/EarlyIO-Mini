Param()
Write-Host "Building all Go binaries..." -ForegroundColor Cyan
$ErrorActionPreference = 'Stop'
Set-Location "$PSScriptRoot/.."
# TCP server
go build -o bin/tcp-echo-server ./examples/network/tcp
# TCP client
go build -o bin/tcp-echo-client ./examples/network/tcpclient
# HTTP server
go build -o bin/http-server ./examples/network/http
# gRPC server (will fail until proto generated)
(go build -o bin/codeconvert-server ./services/codeconvert/server) 2>$null
Write-Host "Done." -ForegroundColor Green
