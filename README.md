# Messages APIs
The purpose of this project is to provide gRPC examples. Four of the main types of Remote Procedure Calls are implemented. 

## Unary
The client sends a single request and receives a single response, similar to a standard function call. 

## Server-Side Streaming
The client sends a single request and receives a stream of responses from the server, allowing the server to send multiple messages. 

## Client-Side Streaming
The client sends a stream of requests to the server, and the server receives all the requests before sending a single response.

## Bi-Directional Streaming
Both the client and the server send a stream of messages to each other simultaneously, enabling a full-duplex communication channel. 
