# Demo Web Services

## Objective

Demonstrate how the same functionality can be implemented by server-server and server-client models with different implementations/protocols.

## Implementations

- [ ] RESTful/swagger w/ go
- [ ] Graphql w/ go
- [ ] GRPC w/ go

### Problem

Using the sample database, solve the problem

- Employee
  - List
  - Read
    - Current Salary
    - All salaries
    - Current departments
    - All departments
    - Current Title
  - Create
  - Update
  - Fire
  - Associate to department
- Departments
  - List
  - Read
  - Create
  - Update
  - Archive
  


## Criteria

1. Must be schema-first implementation
2. Service must have
   1. Validation
   2. Error handling
   3. Logging
3. Share the same data source

## Setup

`docker-compose up`

## TODO

- [x] start all implementations via docker-compose + common dependencies
- [ ] run migrations automatically on `up`
- [x] Get all impementations responding to http requests
- [ ] scaffold basic employees list implementation
