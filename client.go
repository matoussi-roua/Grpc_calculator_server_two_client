package main

import (
    "context"
    "fmt"
    "log"
    "time"

    pb "calculator/calculator" // Modifiez ceci en fonction du chemin d'importation
    "google.golang.org/grpc"
)

func main() {
    // Connexion au serveur
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewCalculatorClient(conn)

    // Définir les variables à additionner
    a := 10.0
    b := 20.0

    // Appeler la méthode Add
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    response, err := client.Add(ctx, &pb.AddRequest{A: a, B: b})
    if err != nil {
        log.Fatalf("could not add: %v", err)
    }

    fmt.Printf("Result: %v + %v = %v\n", a, b, response.Result)
}
