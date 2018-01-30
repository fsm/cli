package cli

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/fsm/emitable"
)

// We add some sleep time between each type of message to make it feel more natural
// Not doing anything fancy here, and keeping it somewhat quick as this target will
// be used mostly in development.
const sleepTime = 600

type emitter struct{}

func (e *emitter) Emit(i interface{}) error {
	switch v := i.(type) {

	case string:
		e.Emit(emitable.Typing{Enabled: true})
		fmt.Println(v)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.Audio:
		e.Emit("Audio: " + v.URL)
		return nil

	case emitable.File:
		e.Emit("File: " + v.URL)
		return nil

	case emitable.Image:
		e.Emit("Image: " + v.URL)
		return nil

	case emitable.Video:
		e.Emit("Video: " + v.URL)
		return nil

	case emitable.QuickReply:
		// Message
		e.Emit(v.Message)

		// Options
		optionsBuffer := new(bytes.Buffer)
		for i, reply := range v.Replies {
			optionsBuffer.WriteString("`")
			optionsBuffer.WriteString(reply)
			optionsBuffer.WriteString("`")
			if i+2 < len(v.Replies) && len(v.Replies) > 2 {
				optionsBuffer.WriteString(", ")
			} else if i+1 < len(v.Replies) {
				if len(v.Replies) > 2 {
					optionsBuffer.WriteString(", or ")
				} else {
					optionsBuffer.WriteString(" or ")
				}
			}
		}

		// Determine format
		format := "You can %v"
		if v.RepliesFormat != "" {
			format = v.RepliesFormat
		}

		// Output
		e.Emit(fmt.Sprintf(format, optionsBuffer.String()))
		return nil

	case emitable.Sleep:
		time.Sleep(time.Duration(v.LengthMillis) * time.Millisecond)
		return nil

	case emitable.Typing:
		if v.Enabled {
			fmt.Printf("... ")
		}
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil
	}
	return errors.New("CommandLineEmitter cannot handle " + reflect.TypeOf(i).String())
}
