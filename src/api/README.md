The `api.proto` file is not usable like this: it needs to get compiled.

The compilation will generate code for the chosen language that your app will actually call.

In our case we are going to compile it to `Go`:

    $ protoc -I api\ -I %GOPATH%/src --go_out=plugins=grpc:api api\api.proto   

This command generates the file `api/api.pb.go`, a `Go` source file that implements the gRPC code your app will use.