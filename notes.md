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
