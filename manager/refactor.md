# Refactor Oct 22

More modular, testable design pattern
More idiomatic error handling and server behavior

## Todo

### DataModel

router > Handlers > DataModel > Models

#### Handlers

Signature as is (they're ServeHTTP methods)
Call `DataModel` methods with `[]byte()` in, `[]byte()` out

#### New Package `DataModel`:

Wraps database (and Redis?) and contains all DB method handlers?

Dependencies:
- Models package - as is currently
- Database
- Redis?

[]bytes in, []bytes out?

`type DataModel struct {db *gorm.DB}`

hang all CRUD operations off DataModel

#### Package `Models`:

Struct/GORM definitions
Request/Response struct definitions
Methods for translating to/from Models
`toJSON()` and `fromJSON()` methods

#### Package `Database`:

GORM database
accepts connection string?
- `nil` = get environment variables?

### Server Graceful Shutdown

wrap `err = ListenAndServe` in goroutine
- if err != nil, log fatal

```go
sigChan := make(chan, os.Signal)
signal.Notify(sigChan, os.Kill)
signal.Notify(sigChan, os.Interrupt)

sig := <-sigChan // blocks
log.Pringln("Shutting down server", sig)

ctx, _ := contest.WithTimeout(context.Background(), 30 *time.Seconds)
os.Shutdown(ctx)
```

### Repo Structure

fanarest
|
|- /manager
|
|- /dash
|
|- Dockerfile

