# Go REST api for serving event files

## File structure
cmd/
    create-files/
        main.go // utility to create file /tmp/foo/repos/events
    event-service-api/
        handlers/
            routes.go // set up of route handlers and dependencies
            root.go // default handler
            healthcheck.go // healthcheck handler
            stats.go // stats handler to retrieve stats
            store.go // handler for getting files
        main.go // actually starts and runs service
    internal/
        constants/
            constants.go // currently just the location of file, and the port the server runs on
        event/
            event.go // a representation of the event type
        event-repository/
            repo.go // utility structs and interfaces to retrieve repositories and provide abstraction for extension
            raw-file-repo.go // a struct which implements the Repository interface (it gets the file)
        middleware/
            stats.go // a middleware that gathers stats
        platform/
            web/
                web.go // some abstractions on top of gin to add middleware, etc
            utility/
                dumpJsonToFile.go // this is a utility which deserializes event structs into a file


# routes
- "GET", "/v1" -> returns root payload (just a test)
- "GET", "/v1/health-check" -> returns 204
- "GET", "/v1/store/:name" -> returns a file located at /tmp/foo/repos/:name
- "GET", "/v1/stats" -> returns stats

# how to use
To start server

- First install go1.14

Then, in repo root
Make sure you don't have anything important in /tmp/foo/repos...
- go run cmd/create-files/main.go 
- go run cmd/event-service-api/main.go

- hit the server at localhost:3005
