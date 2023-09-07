# build the ./website binary
website : website-upload-tool/main.go
	cd website-upload-tool; go build -o ../website  main.go 

# clear out ./website binary, public/ (hugo output), and bin/ (other output)
clean :
	-rm website
	-rm -rf public/*
	-rm -rf bin/*