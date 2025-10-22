Param()
Set-Location "$PSScriptRoot/..\infra\terraform"
terraform init
terraform apply -auto-approve
