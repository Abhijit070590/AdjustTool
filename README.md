# AdjustTool
Assignment for Adjust

## Steps to run 
- In the Makefile set variable `EXECUTABLE` to the name of the executable wanted 
- To build run `make build`
- run the executable e.g if the `EXECUTABLE` is mytool then run `./mytool google.com yahoo.com`
- unit test cases are in Tool_test.go and can be run by `go test`

## Cases that can handle
- urls without http 
- multiple urls 
- in  parallel mode with flag `parallel` as asked in the assignment sheet. 
- In case of the wrong urls like `http://google` not `http://google.com` the programs gives appropriate message in the log