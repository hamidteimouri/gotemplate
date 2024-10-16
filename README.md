# Lens



### Directories structure :

```
project
│   .gitignore
│   Dockerfile    
│
└─── migrations
│   │   migrations-file.go
│   │   
└─── src
│   │
│   └─── data
│   │   │   └── Postgres
│   │   │   
│   └─── di
│   │   │   di.go (Dependency Injection)
│   │   │   
│   └─── domain
│   │   │   entity.go
│   │   │   repository.go
│   │   │   service.go
│   │   │   
│   └─── presentation
│   │   │   └── HTTP
│   │   │   └── gRPC
│   │   │   └── Kafka
│   │   │   
│   └─── pkg
│       │   ...
```

# Directory Overview
`migrations/`: Contains database migration files, helping manage schema changes over time.

`src/`:

- `data/Postgres`: Data layer where PostgreSQL-related database operations are handled.
- `di/`: Handles Dependency Injection for the application.
- `domain/`: This is the core business logic layer.
  + `entity.go`: Domain models representing business objects.
  + `repository.go`: Repository interfaces for abstracting database interactions.
  + `service.go`: Service layer for handling business logic.
- `presentation/`: Interfaces for external communication, handling:
  + `HTTP`: RESTful API endpoints.
  + `gRPC`: gRPC service implementations.
  + `Kafka`: Kafka consumer and producer services.
- `pkg/`: Utility functions and reusable components.


# Setup Instructions
To Run migration files please read these [docs](https://github.com/pressly/goose)

