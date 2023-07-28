# build the ./website binary
website : script/main.go
	go build -o website  script/main.go 

# clear out ./website binary, piping output to /dev/null (prevent outputting errors)
clean :
	-rm website 2> /dev/null