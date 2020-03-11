package eventrepository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"sloan.com/service/internal/event"
)

var (
	defaultRepo = "default"
)

type FileRepository struct {
	absPathToDir string
	Files        map[string]*FileWrapper
}

func NewFileRepository(dirPath string) (*FileRepository, error) {
	path, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	return &FileRepository{
		absPathToDir: path,
		Files:        map[string]*FileWrapper{},
	}, nil
}

type FileWrapper struct {
	relPath string
	events  []*event.Event
}

func (r *FileRepository) Name() string {
	return defaultRepo
}

func (r *FileRepository) GetResource(name string) ([]*event.Event, error) {
	if wrapper, exists := r.Files[name]; exists {
		fmt.Printf("wrapper.events", wrapper.events)
		return wrapper.events, nil
	} else {
		wrapper, err := r.loadFile(name)
		if err != nil {
			return []*event.Event{}, err
		}
		return wrapper.events, nil
	}
}

// reads newline separated event dump
func (r *FileRepository) loadFile(filePath string) (*FileWrapper, error) {
	// if something writes to this file while we're reading..uh oh
	// todo - validate path
	absPath := filepath.Join(r.absPathToDir, filePath)
	fd, err := os.Open(absPath)
	if err != nil {
		return nil, ErrResourceNotFound
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	evts := []*event.Event{}
	var evtP *event.Event
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return nil, err
		}
		b := scanner.Bytes()
		evtP, err = parseLineToEvent(b)
		if err != nil {
			return nil, err
		}
		evts = append(evts, evtP)
	}
	return &FileWrapper{
		relPath: filePath,
		events:  evts,
	}, nil
}

func parseLineToEvent(line []byte) (*event.Event, error) {
	evt := &event.Event{}
	fmt.Printf("\nthe alleged josn\n\"%v\"EOF", string(line))
	err := json.Unmarshal(line, evt)
	if err != nil {
		return nil, err
	}
	return evt, nil
}
