package main

import (
    "flag"
    "github.com/benrowe/demo-web-services/src/go-swagger/gen/restapi"
    "github.com/benrowe/demo-web-services/src/go-swagger/gen/restapi/operations"
    "github.com/go-openapi/loads"
    "log"
)

func main() {
    portFlag := flag.Int("port", 80, "Port the service runs on")
    swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
    if err != nil {
        log.Fatalf("an error has occurred analysing the swagger spec: %v", err)
    }

    api := operations.NewGoSwaggerDemoAPI(swaggerSpec)
    server := restapi.NewServer(api)

    defer server.Shutdown()

    flag.Parse()
    server.Port = *portFlag

    if err := server.Serve(); err != nil {
        log.Fatalf("error occured serving: %v", err)
    }
}
