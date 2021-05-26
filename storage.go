package main

import (
  "github.com/replit/database-go"
)

func ReadValue(key string) string {
  value, _ := database.Get(key)
  return value
}