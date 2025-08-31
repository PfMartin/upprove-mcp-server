# upprove-mcp-server

MCP Server for the Upprove data. Upprove provides data for performance tracking in order to improve your performance over time.

## Developer Information

### TL;DR

#### Run dev server

```bash
make run-server-dev
```

#### Build server

```bash
make build
```

#### Run compiled binary

```bash
make run-server
```

### Layered folder structure

[Organize Like a pro](https://medium.com/@smart_byte_labs/organize-like-a-pro-a-simple-guide-to-go-project-folder-structures-e85e9c1769c2)

```bash
project
├── cmd                      # Command-related files
│   └── app                  # Application entry point
│       └── main.go          # Main application logic
├── internal                 # Internal codebase
│   ├── handlers             # HTTP request handlers (controllers)
│   │   └── user_handler.go  # User-specific handler
│   ├── services             # Business logic (service layer)
│   │   └── user_service.go  # User-specific service
│   ├── repositories         # Data access (repository layer)
│   │   └── user_repo.go     # User-specific repository
│   └── models               # Data models (entities)
│       └── user.go          # User model
├── pkg                      # Shared utilities or helpers
├── configs                  # Configuration files
├── go.mod                   # Go module definition
└── go.sum                   # Go module checksum file
```
