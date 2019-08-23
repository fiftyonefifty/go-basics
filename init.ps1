Write-Host "Script:" $PSCommandPath
Write-Host "Path:" $PSScriptRoot

$Env:GOPATH = $PSScriptRoot
"GOPATH:$Env:GOPATH" 

go get -u github.com/gorilla/mux
go get -u github.com/thedevsaddam/govalidator
