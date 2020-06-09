// TODO: clean main file
package main

import (
  "fmt"
  "bytes"
  "math"
)

const example string = "[[\"hello world\"]]"

// baconbuffer is used to keep data about any
// expression being evaluated
type baconbuffer struct {
  position int
  and bool
  buffer bytes.Buffer
}

func (b baconbuffer) Unshift(bytes []byte) {
  newbuffer = bytes.Buffer
  newbuffer.Write([]byte"]")
  newbuffer.Write([]byte(b.buffer.Bytes())
  b.buffer = &
// create the final channel and fires the recursive method
func main() {
  final := make(chan []byte)
  go Evaluate(
    example,
    final,
  )

  fmt.Println(string(<-final))
}

// take in raw strings evaluate
// elements and create go routines
// to concurrently check conditions
//
// combines all buffers into a raw array
// of bytes to post to the channel
func evaluate(
  raw string,
  and bool,
  done chan<- []byte,
) {
  // used for keeping track of OR/AND expressions
  state := 0
  depth := 0

  // use to keep track of concurrently evaluated channels
  var buffers = make([]*baconbuffer)

  // the current buffer for the given go routine
  var current baconbuffer

  // buffers for any elements that need concurrency
  bufferchan = make(chan []*baconbuffer)
  // maximum number of expected responses on bufferchan
  nchan := 0

  // run the launch method with the appropriate arguments
  // and update the state of the current routine
  registerLaunch := func(buffer *baconbuffer) {
    go launch(
      buffer,
      bufferchan,
    )

    nchan = nchan + 1
    state = 0
  }

  // create a baconbuffer, and update
  // the current routine's state
  bufferfactory := func() *baconbuffer {
    buffer := &baconbuffer {
      channelCount,
      false,
      bytes.Buffer,
    }

    buffers = append(
      buffers,
      buffer,
    )

    current = buffers[len(buffers)-1]
  }

  for pos, char := range raw {
    // start of new expression

    if char == "[" {
      // if the previous value was a string,
      // this is a new expression
      if state == 0 {
        state = -1
        depth = 1

        buffer := *baconbuffer {
          // we use the position of the baconbuffer
          // to perserve concurrent determinism
          channelCount,
          // when we start a new
          // expression, we assume it's
          // an OR expression by default
          false,
          bytes.Buffer,
        }

        // add this buffer to the list of buffers
        buffers = append(buffers, buffer))
        // make it our current buffer
        current = *buffers[len(buffers)-1]

      // otherwise this character could
      // potentially be a part of an AND expression
      } else if state == -1 {
          state = 2
      }
    // end of an expression
    } else if char == "]" {
      if state == 2 {
        state == 3
      } else if state == 1 || state == 3 { registerlaunch(current) }
    } else if state == 3 {
      // throw it back to the front
      current.and = false
      go registerLaunch(current)
    } else if state == 0 {

      if buffers.Len() == 0 {
        buffers.append(make(
      // write the rune as a string
      current.Write([]byte(string(char)))
    }
  }


  fmt.Println("")
  done <- result.Bytes()
}

func unshift(b *baconbuffer)
// take a baconbuffer and launch a new go routine
func launch(
  buffer baconbuffer,
  done chan<- []byte,
) {
  go evaluate(
    buffer.buffer.String(),
    buffer.and,
    done,
  )
}
