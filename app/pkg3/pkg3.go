package pkg3

import (
    "encoding/json"

    "github.com/google/uuid"
    "github.com/pkg/errors"
)

type Payload struct {
    ID string `json:"id"`
}

func Work() string {
    p := Payload{ID: uuid.New().String()}
    _, _ = json.Marshal(p)
    _ = errors.New("dummy")
    return "pkg3 "
}
