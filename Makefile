compile: webami.go 
	echo "Compiling webami..."
	go build -o bin/webami webami.go 

compile_linux_amd64: webami.go
	echo "Compiling webami..."
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/ webami.go 

compile_windows_amd64: webami.go 
	echo "Compiling webami..."
	env GOOS=windows GOARCH=amd64 go build -o bin/windows/webami.exe webami.go

compile_all_systems: compile_linux_amd64 compile_windows_amd64

clean: 
	echo "Cleaning up..."
	rm -r bin/