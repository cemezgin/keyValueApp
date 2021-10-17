
# Key Value In memory store app
#### Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

#### Requirements

- Go 1.16

#### Execute

    go build
    ./keyValueApp

## Api Doc

| TYPE   | URI                       | BODY                           | STATUS |
|--------|---------------------------|--------------------------------|--------|
| GET    | localhost:8090/keys/{key} | -                              | 200    |
| GET    | localhost:8090/keys       | -                              | 200    |
| POST   | localhost:8090/keys       | {"key":"foo","value":"blabla"} | 201    |
| DELETE | localhost:8090/keys       | -                              | 202    |

