# Worker Pool Pattern

## What is this pattern?
The **Worker Pool Pattern** is a concurrency pattern that distributes work across a fixed number of workers (goroutines).  
Jobs are sent into a channel, and workers consume them concurrently, producing results through another channel.  
This ensures efficient use of system resources and prevents spawning an unbounded number of goroutines.

---

## Key Benefits of Worker Pool Pattern
- **Concurrency Control** → limits the number of goroutines running at the same time.  
- **Resource Efficiency** → workers are reused, avoiding overhead of creating new goroutines for each job.  
- **Scalability** → easy to scale up or down by adjusting the number of workers.  
- **Synchronization** → workers process jobs concurrently, while results are collected in an ordered and controlled way.  

## Some examples are
- Handling incoming HTTP requests in a web server.
- Processing images concurrently.