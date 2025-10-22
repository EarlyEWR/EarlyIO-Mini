Param()
$ErrorActionPreference = 'Stop'
Set-Location "$PSScriptRoot/.."
if (-not (Get-Command protoc -ErrorAction SilentlyContinue)) {
  Write-Error "protoc not found. Install from https://github.com/protocolbuffers/protobuf/releases"
  exit 1
}
if (-not (Get-Command protoc-gen-go -ErrorAction SilentlyContinue)) {
  Write-Error "protoc-gen-go not found. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
  exit 1
}
if (-not (Get-Command protoc-gen-go-grpc -ErrorAction SilentlyContinue)) {
  Write-Error "protoc-gen-go-grpc not found. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
  exit 1
}
$env:PATH = "$env:GOPATH\bin;$env:PATH"
protoc --go_out=. --go-grpc_out=. protos/codeconvert.proto
Write-Host "Proto generation complete." -ForegroundColor Green
