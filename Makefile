install: webami.go 
	echo "Compiling webami..."
	go build -o bin/webami webami.go 

clean: 
	echo "Cleaning up..."
	rm -r bin/