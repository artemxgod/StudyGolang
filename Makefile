include .env

bot_path=.\technosphere\lek-4\06_bot\src\main\main.go
run:
	go run main.go
executable:
	go build main.go

exe_w/o_terminal:
	go build go build -ldflags -H=windowsgui main.go

test:
	go test main_test.go

run_bot:
	go run $(bot_path)
