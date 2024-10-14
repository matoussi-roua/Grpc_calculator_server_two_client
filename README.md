# Grpc_calculator_server_two_client
This project implements a simple calculator service using gRPC (Google Remote Procedure Call) and Protocol Buffers for efficient communication. The service allows clients to perform addition operations on two numbers.

## Table of Contents
- [Definitions](#definitions)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [Using the Python Client](#using-the-python-client)
- [Using the Go Client](#using-the-go-client)
- [Contributing](#contributing)
- [License](#license)

## Definitions

### gRPC
gRPC is an open-source remote procedure call (RPC) framework that enables efficient communication between client and server applications. It allows for the definition of services and the methods they expose, along with message types, using a simple interface definition language (IDL).

### Protocol Buffers (protobuf)
Protocol Buffers is a language-agnostic binary serialization format developed by Google. It allows you to define data structures and services in a `.proto` file, which can then be compiled to generate source code in multiple programming languages.

### Stubs
In the context of gRPC, stubs are the client-side representations of the server's methods. When a client calls a method on the stub, it automatically handles the necessary network communication to send the request to the server and receive the response.

## Project Structure
calculator/ 
│ ├── server/ 
│ ├── server.py # Main server implementation in Python 
  calculater/
  └── calculator_pb2.py # Generated Protocol Buffers code for messages 
  └── calculator_pb2_grpc.py # Generated gRPC code for the service 
│ ├── client.py # Python client implementation
│ ├── client.go # Go client implementation 
│ ├── calculator.proto # Protocol Buffers definition file 
└── README.md # Project documentation

### File Descriptions

- **calculator.proto**: This file contains the service definition and message types used in the gRPC service. It defines the `Calculator` service with the `Add` method and the corresponding request and response messages.

- **server.py**: This is the main implementation of the gRPC server in Python. It starts the server and handles incoming requests from clients.

- **calculator/calculator_pb2.py**: This file is generated from the `calculator.proto` file. It contains the definitions of the message types used in the gRPC service.

- **calculator/calculator_pb2_grpc.py**: This file is also generated from the `calculator.proto` file. It contains the definition of the `Calculator` service and the client stub for making requests.

- **client.py**: This file contains the implementation of a Python client that connects to the gRPC server, sends requests to the `Add` method, and prints the results.

- **client.go**: This file contains the implementation of a Go client that performs the same operations as the Python client.

## Installation

To run the server and clients, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/matoussi-roua/Grpc_calculator_server_two_client.git
   cd calculator

2. Install the required dependencies for the Python server:

  pip install grpcio grpcio-tools

3. For the Go client, make sure you have Go installed and set up on your machine.

4. Running the Server
  To start the gRPC server, navigate to the server directory and run:
  
  python server.py
  The server will start running on port 50051.

5. Running the client
  Using the Python Client
  To use the Python client, navigate to the client_python directory and run:
  
  python client.py
  This client will connect to the server and perform an addition operation.

6. Using the Go Client
  To use the Go client, navigate to the client_go directory and run:

  go run client.go
  This client will also connect to the server and perform an addition operation.
