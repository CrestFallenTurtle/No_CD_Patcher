OUTPUT := trespasser_patcher.exe

compile:
	GOOS=windows GOARCH=386 go build -o $(OUTPUT) .

clean:
	rm $(OUTPUT)

install_dependencies:
	go get github.com/CrestFallenTurtle/ArgumentParser
	go get github.com/CrestFallenTurtle/notify_handler