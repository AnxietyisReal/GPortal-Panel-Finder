# How to use
Currently as of now, I don't have the motivation to compile the code to Windows and Linux executables and release them.

So you're gonna have to install [Go 1.21+](https://go.dev/dl/) and compile them yourself with `go build -trimpath ./` in the directory.

After that's compiled, you can run the program with two required arguments:  
(Written this code on Ubuntu, so cope.)  
`./panel-hunter -server "the server name as displayed in-game" -ip "the server's IP address that you captured through Wireshark or other tools"`