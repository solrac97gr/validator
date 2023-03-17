# Validator

I create this packages for use it as part of my future template with multi packages separeted of the principal template.

Its a simple validator for use in your projects who contains two principal interfaces and structs who implements it.

## Interfaces

### Validator

```go
type Validator interface {
    Struct(s EvaluableStruct) error
}
```

### ValidatorFunc

```go
type ValidatorFunc func(args ...interface{}) error
```

## EvaluableStruct

```go
type EvaluableStruct interface {
    Validate(args ...interface{}) error
}
```

## Example

```go
package main

import (
    "fmt"
    "github.com/solrac97gr/validator"
)

type User struct {
    Name string 
    Age  int    
}

func (u *User) Validate(args ...interface{}) error {
    err := validator.StringRequired(u)
    if err != nil {
        return err
    }

   err = validator.IntGratherThan(u.Age,18)
    if err != nil {
        return err
    }

    return nil
}


func main() {
    user := &User{
        Name: "Solrac",
        Age:  17,
    }

    val := validator.New()

    err := val.Struct(user)
    if err != nil {
        fmt.Println(err)
    }
}
```
## Idea

You must create a struct who implements the EvaluableStruct interface and add the Validate method who contains the logic for validate the struct.

The validator must be part of your dependencies in your depencency injection container.

## Considerations

This package is in development and is not ready for production.


## License

MIT