package cli

import (
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
		fmt.Println(v)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.Audio:
		fmt.Println("Audio:", v.URL)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.File:
		fmt.Println("File:", v.URL)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.Image:
		fmt.Println("Image:", v.URL)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.Video:
		fmt.Println("Video:", v.URL)
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil

	case emitable.QuickReply:
		fmt.Print("Type one of: [ ")
		for i, reply := range v.Replies {
			fmt.Print("'" + reply + "'")
			if i+1 < len(v.Replies) {
				fmt.Print(", ")
			}
		}
		fmt.Println(" ]")
		return nil

	case emitable.Sleep:
		time.Sleep(time.Duration(v.LengthMillis) * time.Millisecond)
		return nil

	case emitable.Typing:
		if v.Enabled {
			fmt.Println("...")
		}
		e.Emit(emitable.Sleep{LengthMillis: sleepTime})
		return nil
	}
	return errors.New("CommandLineEmitter cannot handle " + reflect.TypeOf(i).String())
}
