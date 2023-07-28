# build the ./website binary
website : script/main.go
	go build -o website  script/main.go 

# clear out ./website binary, public/ (hugo output), and bin/ (other output)
clean :
	-rm website
	-rm -rf public
	-rm -rf bin