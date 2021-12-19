package main

import (
   "fmt"
   "time"
)

type program struct{}

func (p program) run() {
   for {
      fmt.Println("Service is running")
      time.Sleep(1 * time.Second)
   }
}

func main() {
   err = run()
   if err != nil {
      fmt.Println("Cannot start the service: " + err.Error())
   }
}
