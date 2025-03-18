# Messages gRPC APIs
The purpose of this project is to provide gRPC examples. Four of the main types of Remote Procedure Calls are implemented. 

## gRPC Definitions
The gRPC defintions are located in the following files:
- messages/emailmessages/emailmessages.proto
- messages/chat/chat.proto

## Types and Implementations

### Unary
- The client sends a single request and receives a single response, similar to a standard function call. 
- Function that implements unary is **SendEmail(EmailReq) returns (GenericResponse)**, which is defined in messages/emailmessages/emailmessages.proto.

### Server-Side Streaming
- The client sends a single request and receives a stream of responses from the server, allowing the server to send multiple messages. 
- Function that implements server-side streaming is **BroadcastLetter(LetterMessage) returns (stream Broadcast)**, which is defined in messages/emailmessages/emailmessages.proto.

### Client-Side Streaming
- The client sends a stream of requests to the server, and the server receives all the requests before sending a single response.
- Function that implements client-side sreaming is **UpdateScribers(stream SubscriberList) returns (GenericResponse)**, which is defined in messages/emailmessages/emailmessages.proto.

### Bi-Directional Streaming
- Both the client and the server send a stream of messages to each other simultaneously, enabling a full-duplex communication channel. 
- Function that implements bi-directional streaming **Chat(stream ClientMessage) returns (stream ServerResponse)**, which is defined in messages/chat/chat.proto.

## Protoc Commands
Before being able to run the Protoc commands install the necessary packages. Note that Homebrew works if you have it installed on your macOS. The protoc command is needed when working with Protocol Buffers (protobuf), a language-neutral, platform-neutral mechanism for serializing structured data. The protoc compiler takes .proto files (which define the structure of your data and services) and generates code in your desired programming language (e.g., Go, Python, Java). This generated code is then used in your application to serialize and deserialize data or implement gRPC services.

## Required Protocol Buffer Packages
```
brew install protobuf
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

## Commands to Generate Golang Code
Open a terminal and navigate to the emailmessages directory. Then run the command below.
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   	emailmessages.proto
```
Open a terminal and navigate to the chat directory. Then run the command below.
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   	chat.proto
```
## Command to Run Go Application
```
go run main.go
```
