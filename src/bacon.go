package main

import (
  "fmt"
  "bytes"
  "math"
)

type expression struct {
  b int
  v string
}

type baconbuffer struct {
  position int
  and bool
  buffer bytes.Buffer
}

// create the final channel and fire the recursive method
func main() {
  final := make(chan []byte)
  go exp(
    "[[\"hello world\"]]",
    final,
  )

  fmt.Println(string(<-final))
}

// take in raw strings and add the evaluated bytes to the done channel
func exp(raw string, done chan []byte) {

  // use to keep track of concurrently evaluated channels
  var buffers = make([]baconbuffer)
  channels = make(chan []byte)

  start := bytes.Buffer
  end := bytes.Buffer
  current = *start
  depth := 0
  channelcount = 0

  for pos, char := range raw {
    // start of new expression
    if char == "[" {
      if depth == 0 {
        depth = -(depth + 1)
      } else if math.Copysign(1, depth) == -1 {
        current.and = true
      }

      // we want to keep track of what's part of this buffer
      buffer := *baconbuffer {
        // we use the position of the rune to perserve concurrent determinism
        pos,
        false,
        bytes.Buffer,
      }

      buffers = append(buffers, buffer))
      current = *buffers[len(buffers)-1]

    // end of an expression
    } else if char == "]" {
      // we want to write the buffer to a new channel





    } else {
      depth = -depth

    }
    current.Write([]byte(string(char)))
  }


  fmt.Println("")
  done <- result.Bytes()
}

func launch([][]byte, done chan<- []byte) {


}
