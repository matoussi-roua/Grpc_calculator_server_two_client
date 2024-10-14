import grpc
import calculator_pb2
import calculator_pb2_grpc

def run():
    # Create a channel to the gRPC server
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = calculator_pb2_grpc.CalculatorStub(channel)
        
        # Create a request
        request = calculator_pb2.AddRequest(a=5.0, b=3.0)
        
        # Make the RPC call
        response = stub.Add(request)
        
        # Print the response
        print(f'The result of adding {request.a} and {request.b} is: {response.result}')

if __name__ == '__main__':
    run()
