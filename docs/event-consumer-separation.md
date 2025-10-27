# Separating Event Consumer from API Service

## The Question

Should we split the Alpaca SSE event consumer into a separate service from the TUI/API application?

```
Option 1: Monolith              Option 2: Separated
┌─────────────────────┐         ┌──────────┐  ┌─────────────┐
│   Pony TUI App      │         │ TUI App  │  │ Event Worker│
│  - TUI              │         │ - TUI    │  │ - SSE       │
│  - API calls        │         │ - API    │  │ - Updates   │
│  - SSE consumer     │         └────┬─────┘  └──────┬──────┘
│  - DB updates       │              │               │
└──────────┬──────────┘              └───────┬───────┘
           │                                 │
           ▼                                 ▼
        Database                         Database
```

## PROS: Keeping Them Together (Monolith)

### 1. **Simplicity**

- One process to manage
- One deployment
- One place to look for logs
- Easier to debug (everything in one stack trace)
- Less operational overhead

### 2. **Lower Latency**

- Direct in-memory communication
- No network hop between components
- Events can update TUI state immediately via channels
- Real-time UI updates without polling

### 3. **Consistency Guarantees**

- Single transaction context
- Easier to maintain data consistency
- No distributed system problems
- Immediate consistency (no eventual consistency issues)

### 4. **Development Speed**

- Faster iteration
- No need to coordinate deployments
- Easier testing (run one thing)
- Simpler local development setup

### 5. **Resource Efficiency**

- One database connection pool
- Shared memory for caching
- Less overhead (no extra containers/processes)

### 6. **For a TUI App Specifically**

- **TUI apps are typically single-user desktop tools**
- Only ONE instance running at a time per user
- No scaling concerns (not serving multiple users)
- The TUI needs real-time updates anyway

## CONS: Keeping Them Together

### 1. **Coupling**

- TUI crashes → Event consumer stops
- Event consumer bugs → TUI crashes
- Harder to isolate issues

### 2. **Resource Contention**

- Event processing competes with UI rendering
- Blocking event handling can freeze UI
- Memory leaks in one affect both

### 3. **Restart Impact**

- Restarting TUI disconnects from event stream
- Might miss events during restart
- Need to resync state after restart

### 4. **Limited Scalability** (not relevant for TUI)

- Can't scale event consumer independently
- But... you're building a TUI, not a web service!

## PROS: Separating Them

### 1. **Isolation**

- Event consumer crashes don't affect TUI
- Can restart/update consumer without TUI downtime
- Easier to debug issues in isolation

### 2. **Independent Scaling** (mostly irrelevant for TUI)

- Scale event processing independently
- But... TUI is 1 user = 1 instance

### 3. **Technology Flexibility**

- Could rewrite consumer in different language
- Different deployment schedules
- Different monitoring/alerting

### 4. **Failure Isolation**

- TUI can survive consumer crashes
- Consumer can survive TUI crashes
- Better fault tolerance

### 5. **Resource Allocation**

- Dedicated resources for event processing
- Won't starve UI of resources
- Can prioritize critical path

### 6. **Multiple Clients**

- One consumer can feed multiple TUI instances
- Useful if you want web + TUI + mobile later
- Centralized event handling

## CONS: Separating Them

### 1. **Complexity**

- Two services to deploy/manage
- Need inter-process communication (IPC/network)
- More moving parts = more failure modes
- Harder to debug distributed issues

### 2. **Latency**

- Network/IPC overhead
- Can't update TUI in real-time as easily
- Need polling or WebSocket to TUI

### 3. **Consistency Issues**

- Race conditions between services
- Eventual consistency problems
- Need distributed transaction handling

### 4. **Development Overhead**

- Two repos or complex monorepo
- Coordinated deployments
- More integration tests needed
- Local dev requires running multiple services

### 5. **Operational Complexity**

- Need orchestration (Docker Compose, K8s, etc.)
- More logs to aggregate
- More metrics to monitor
- Higher infrastructure cost

### 6. **State Synchronization**

- How does TUI get consumer state?
- Need message queue or shared DB
- Cache invalidation problems

## Recommendation for Your TUI App

### Keep Them Together Because:

1. **It's a TUI** - Single-user desktop application
   - No multi-user scaling needed
   - User launches one instance
   - Direct memory communication is perfect

2. **Real-time Updates**

   ```go
   // Beautiful in-process pattern
   eventCh, _ := broker.StreamEvents(ctx, accountID)

   go func() {
       for event := range eventCh {
           model.Update(event)  // Direct update!
       }
   }()
   ```

3. **Simpler Operations**
   - User runs: `pony`
   - Not: `docker-compose up` with 3 services

4. **YAGNI** (You Aren't Gonna Need It)
   - You don't have scaling requirements
   - Premature optimization is evil
   - Start simple, split later if needed

### When to Split Later?

Split them when you:

1. **Add a Web UI**
   - Multiple TUI/web clients need same events
   - Then centralized consumer makes sense

2. **Event Processing Gets Heavy**
   - Complex transformations
   - ML/analytics on event stream
   - Blocking operations

3. **Need High Availability**
   - Event consumer must never stop
   - TUI can restart but consumer stays up

4. **Multiple Data Sources**
   - Consuming from Alpaca + other brokers
   - Complex event routing logic

## Hybrid Approach (Best of Both Worlds)

For your TUI, consider this pattern:

```go
// pkg/events/consumer.go
type Consumer struct {
    broker domain.BrokerClient
    store  Store
    eventCh chan domain.Event
}

// Can run in-process or out-of-process
func (c *Consumer) Start(ctx context.Context) {
    sseEvents, errs := c.broker.StreamEvents(ctx, accountID)

    for {
        select {
        case event := <-sseEvents:
            // Update DB
            c.store.UpdateFromEvent(event)

            // Forward to TUI (if in-process)
            if c.eventCh != nil {
                c.eventCh <- event
            }
        case err := <-errs:
            log.Error(err)
        case <-ctx.Done():
            return
        }
    }
}
```

**Benefits:**

- Clean abstraction
- Can easily extract later
- Testable in isolation
- But still runs in-process for simplicity

## The Verdict

**For your TUI trading app: Keep them together.**

You can always split later if you:

- Add a web interface
- Need to scale (you won't)
- Have heavy event processing (you don't)

Remember: **Duplication is far cheaper than the wrong abstraction.** Same applies to services - one service is far cheaper than premature microservices.

## Further Reading

- [Monolith First](https://martinfowler.com/bliki/MonolithFirst.html) - Martin Fowler
- [The Majestic Monolith](https://m.signalvnoise.com/the-majestic-monolith/) - DHH
- [Microservices Prerequisites](https://martinfowler.com/bliki/MicroservicePrerequisites.html)
