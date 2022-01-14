package zris

import (
	"bufio"
	"errors"
	"io"
)

// The RisObject store the ris information. The nature of it is a map of slice.
type RisObject map[string][]string

// The zris.Index() will scan the whole ris file and return a RisObject with key and value.
func Index(file io.Reader) (RisObject, error) {
	risdata := make(RisObject)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		risdata[line[0:2]] = append(risdata[line[0:2]], line[6:])

	}

	_, check := risdata["TY"]
	if !check {
		return nil, errors.New("the ris file is not valid")

	}
	return risdata, nil
}

// The method will find the value with the given key. If the value does not exist, an error will return.
func (r RisObject) Match(key string) ([]string, error) {
	result, exist := r[key]
	if !exist {
		error := errors.New("the key is not founded")
		return nil, error
	}
	return result, nil
}
