# Advance Concepts
## Error Coded
- https://grpc.io/docs/guides/error/
- http://avi.im/grpc-errors/
- example: ```status.Errorf(codes.Internal, "message")```

## Deadline
- it is recommended to have deadlines for every api call
- https://grpc.io/blog/deadlines/
- Code snippet:
* clientDeadline := time.Now().Add(time.Duration(*deadlineMs) * time.Millisecond)
* ctx, cancel := context.WithDeadline(ctx, clientDeadline) // Pass in this context
* defer cancel() // always cancel the context in the end
- Identifying a timeout err: 
* statusErr, ok := status.FromErr(err) // if it is ok, then it is a grpc error
*   if ok {
      if statusErr.Code() == codes.DeadlineExceeded {
        "Send your timeout err message here"
      } else {
        ...
      }
    } else {
      ...
    }

## SSL/TLS Encryption
- In production gRPC calls should be running with encryption enable to prevent "man in the middle attack"
- This is done by generating SSL certificates
- SSL allows communication to be secure end-to-end
- There are two ways of using SSL, 
  1. 1-way verification, example, browser verifying the webserver. This is Encryption.
  2. 2-way verification, example, SSL Authentication 
- https://grpc.io/docs/guides/auth/
- Use the instructions.sh file to generate the necessary files

## gRPC reflections & Evans CLI
- Reflections help the server to expose the API endpoints (basically the client can ask the server what APIs it has)
- https://godoc.org/google.golang.org/grpc/reflection
- https://github.com/ktr0731/evans