package utility

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"sloan.com/service/internal/event"
)

func DumpEventsToFile(evts []*event.Event, filePath string) {
	f, _ := os.Create(filePath)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Printf("filePath %v\n", filePath)
	for _, evt := range evts {
		b, _ := json.Marshal(evt)
		fmt.Printf("writing string: %v", string(b))
		c, err := w.WriteString(string(b) + "\n")
		if err != nil {
			fmt.Printf("error writing", err, c)
		}
	}
	err := w.Flush()
	if err != nil {
		fmt.Printf("error flushing")
	}
}