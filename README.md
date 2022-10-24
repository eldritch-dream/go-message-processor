# go-message-processor
This golang message processor aims to fulfill the requirements listed in this technical writeup https://saladtech.notion.site/Salad-s-Technical-Interview-029c555d7ea5447fabd4f3d167007899


## How to run

In order to run this simple processor application you will need to run the command `go run .` in the root directory as well as in the test_server directory. You can also change the `server_addr` property in processor.go to a real tcp server if one is available to send messages (see the comment on line 12).

 Once you run the program on the command line it will try to establish a connection to the server and read the binary message sent along the wire. The application then passes the binary data off to be parsed into a coherent message as defined by the spec in the techinical writeup. Finally, the processor prints out what it decoded from the message.

 ## How to test

Run `go test ./...` from the root directory, this will the unit tests.

## Potential improvements

* Implement a makefile build process to make formatting/building/testing/running a little more flexible
* Use github actions to make build actions happen on commit/push
