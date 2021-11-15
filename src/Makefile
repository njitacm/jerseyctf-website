SHELL=/bin/bash
 
all = compile

# compile -> Build Go Executable
compile: clean
	@go build main.go

clean:
# Linux / MacOS
ifneq (,$(wildcard ./main))
	@rm ./main 
endif
# Windows
ifneq (,$(wildcard ./main.exe))
	@rm ./main.exe 
endif
# Idk when it will show up but
ifneq (, $(wildcard ./a.out))
	@rm ./a.out
endif