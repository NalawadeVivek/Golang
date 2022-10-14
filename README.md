# Golang
Steps 

To run go

    go run main.go

To build and generate a binary or exe. file

    go build main.go -> this will generate the exe file in windows // binary in linux

To build & generate desired output exe with another name

    go build -o app.exe main.go

To build for other operating systems
    To know list of OS & ARCH go to the terminal & use 
        
        go env
        go tool dist list
    
   IMP parameter here are [GOOS, GOARCH] 
       
       export GOOS=linux 
       export GOARCH=amd64 

   for windows
       
       export GOOS=windows
       export GOARCH=amd64

   check the go env parameter

    GOOS=linux && GOARCH=amd64 go build -o app main.go  

To instll the exe
    
    go install main.go

   this file will install in GOBIN path by default it look for GOPATH or GOROOT
