run:
	go run main.go
executable:
	go build main.go

exe_w/o_terminal:
	go build go build -ldflags -H=windowsgui main.go