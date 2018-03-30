all:
	@echo
	@echo "Commands  : Description"
	@echo "--------- : -----------"
	@echo "make deps : Install the dependencies"
	@echo "make run  : Run the program"
	@echo

deps:
	@dep ensure

run:
	@go run *.go
