.PHONY: run
run: main
	./$<

main: *.go go.mod
	task run

.PHONY: all
all: main
