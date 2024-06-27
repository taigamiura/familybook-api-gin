package util

import "os"

func GetGinMode() string {
  if ginMode := os.Getenv("GIN_MODE"); ginMode != "" {
    return ginMode
  }
  return "development"
}
