package model

import (
  "context"
  "fmt"
  "time"

  "github.com/google/uuid"
)

type Base struct {
  ID        uuid.UUID `gorm:"column:id;primary_key" json:"id"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
}

type Pagination struct {
  PageNumber int `json:"pageNumber"`
  PageSize   int `json:"pageSize"`
}

func MustParse(s string) uuid.UUID {
  u, err := uuid.Parse(s)
  if err != nil {
    fmt.Errorf("error generating uuid from string %s: %w", s, err)
    return uuid.UUID{}
  }
  return u
}


