EXECUTABLE=mytool

build:
	go build -o $(EXECUTABLE)

run:
	./$(EXECUTABLE) $(ARGS)
