# Generator Pattern

## What is this pattern?
The **Generator Pattern** is a function that produces a stream of values and sends them through a channel.  
It allows consumers to process values as they are generated, without needing the entire dataset in memory.

---

## Key Benefits of Generator Pattern
- **Non-Blocking Execution** → generation and processing happen simultaneously.  
- **Memory Efficiency** → values are produced and consumed one at a time, no need to store all in memory.  
- **Infinite Sequences** → supports generating infinite sequences without exhausting memory.  
- **Backpressure Handling** → if the consumer is slow, the generator naturally slows down due to channel blocking, preventing overload.  

