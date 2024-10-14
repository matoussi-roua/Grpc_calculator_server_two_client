# Grpc_calculator_server_two_client
This project implements a simple calculator service using gRPC (Google Remote Procedure Call) and Protocol Buffers for efficient communication. The service allows clients to perform addition operations on two numbers. This project demonstrates the interoperability of gRPC, showcasing its ability to work seamlessly with different programming languages and across various machines.

## Table of Contents
- [Definitions](#definitions)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [Using the Python Client](#using-the-python-client)
- [Using the Go Client](#using-the-go-client)
- [Working on Different Machine](#working_on_different_machine)
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
    └──calculater/
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

### Clone the repository:
       
       git clone https://github.com/matoussi-roua/Grpc_calculator_server_two_client.git
       cd calculator

### Install the required dependencies for the Python server:
Install the required dependencies for the Python server:
        
        pip install grpcio grpcio-tools

### For the Go client, make sure you have Go installed and set up on your machine:

        go mod tidy
        
### Generate the Protocol Buffers code:

Python:

        python -m grpc_tools.protoc -I. --python_out=./server --grpc_python_out=./server calculator.proto
        
Go:

        protoc --go_out=./client_go --go-grpc_out=./client_go calculator.proto
        
# Running the Server:
  To start the gRPC server, navigate to the server directory and run:
  
        
        python server.py
        
  The server will start running on port 50051.

# using the python client:

  To use the Python client, navigate to the client_python directory and run:
  
        
        python client.py
        
  This client will connect to the server and perform an addition operation.
  
# using the go client:

  To use the Go client, navigate to the client_go directory and run:
  
        
        go run client.go
        
  This client will also connect to the server and perform an addition operation.
  
# working on different machine

Start the server on machine 1, To allow clients on different machines to connect to the server, you need the IP address of the server machine.
On the server machine, run:
        
        ipconfig  # On Windows
        ifconfig  # On Linux/Mac
        
## Python Client:
If you're using the Python client:

Install dependencies:

        pip install grpcio
        
Edit the client.py file:

        // Update the server address to the server machine's IP address:
        with grpc.insecure_channel('192.168.1.100:50051') as channel:
        
Run the client:

        python client.py
        
## Go Client:
If you're using the Go client:

Make sure Go is installed.
Edit the client.go file:

        // Update the server address to the server machine's IP address:
        conn, err := grpc.Dial("192.168.1.100:50051", grpc.WithInsecure())
        
Run the Go client:

        go run client.go
        
Testing the Setup: When the server is running on one machine and the clients are running on different machines, the clients should successfully connect to the server and receive responses for addition operations.

## Notes on Networking
Ensure that the server machine's firewall allows inbound connections on port 50051. You may need to add a firewall rule to allow traffic through this port.
The machines must be on the same network (local network) or have proper routing if on different networks (e.g., using VPN or public IPs).
If using public IP addresses, ensure that you secure your gRPC service appropriately.
# Contributing
Contributions are welcome! If you have suggestions for improvements or want to add features, please create a pull request.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

If you need further modifications or additional sections, let me know!

