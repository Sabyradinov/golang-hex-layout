# github.com/Sabyradinov/golang-hex-layout
___
Golang project layout with hexagonal architecture

### Interesting libs/frameworks can be found in the following links:
https://github.com/avelino/awesome-go

## Hexagonal architecture
- In this project implemented hexagonal architecture structure with gin web server https://github.com/gin-gonic/gin

### Pros of hexagonal architecture:

- Maintainability, low level of dependency between project layers, which gives an advantage when adding and changing functionality.
- Possibility for several developers to work in parallel in different parts of the code.
- Testability, ability to write unit tests for each component of the project without dependencies on other components.
- Adaptability to different task requirements.

### Project structure:

- In hexagonal architecture there are two main concepts - ports and adapters, conditionally for understanding: in ports(contracts) interfaces are declared, in adapters interfaces are implemented (more details here https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)).
- Also, before starting development, study the basic concepts of Domain Driven Design (DDD).
- `/common` package contains common functions, constants, errors, etc.
- `/config` package contains configuration files, environment variables, etc.
- `/internal/domain` package contains domain services, entities, value objects, etc.
- `/internal/domain/port` package contains interfaces for domain services
- `/internal/domain/service` package contains domain services
- `/internal/handler` package contains handlers for web server
- `/internal/http` package contains http server, REST API
- `/internal/grpc` package contains grpc server, gRPC API

### hexgen - generator for hexagonal architecture project layout

- to install run command `go install github.com/Sabyradinov/golang-hex-layout/hexgen`
- example of use `hexgen -path user-path -gomod new-project-gomod`, where `user-path` - absolute folder path `new-project-gomod` - name of the project
- check version `hexgen -version`

