package eventrepository

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type RawFileRepository struct {
	absPathToDir string
}

func NewRawFileRepository(dirPath string) (*RawFileRepository, error) {
	path, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	return &RawFileRepository{
		absPathToDir: path,
	}, nil
}

func (rw *RawFileRepository) GetResource(name string) (Resource, error) {
	wrapper, err := rw.loadFile(name)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}

func (rw *RawFileRepository) Name() string {
	return defaultRepo
}

func (r *RawFileRepository) loadFile(filePath string) (*RawFileWrapper, error) {
	// if something writes to this file while we're reading..uh oh
	// todo - validate path
	absPath := filepath.Join(r.absPathToDir, filePath)
	fd, err := os.Open(absPath)
	if err != nil {
		return nil, ErrResourceNotFound
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	contentType := http.DetectContentType(b)
	return &RawFileWrapper{
		relPath:     filePath,
		file:        b,
		contentType: contentType,
	}, nil
}

type RawFileWrapper struct {
	relPath     string
	file        []byte
	contentType string
}

func (w *RawFileWrapper) GetContentType() string {
	return w.contentType
}

func (w *RawFileWrapper) GetData() []byte {
	return w.file
}
