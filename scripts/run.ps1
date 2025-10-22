Param(
  [ValidateSet('tcp-server','tcp-client','http','grpc-server','grpc-client')]
  [string]$Target = 'tcp-server'
)
Set-Location "$PSScriptRoot/.."
switch ($Target) {
  'tcp-server' { go run ./examples/network/tcp }
  'tcp-client' { go run ./examples/network/tcpclient }
  'http' { go run ./examples/network/http }
  'grpc-server' { go run ./services/codeconvert/server }
  'grpc-client' { go run ./services/codeconvert/client }
}
