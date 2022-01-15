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

// The method will find the value with the given key. If the value does not exist, nil will return.
func (r RisObject) Match(key string) []string {
	result, exist := r[key]
	if !exist {
		return nil
	}
	return result
}

// Convert "PY" tag into date with 3 return values year, month and day. If the tag is not found, an error will return.
func (r RisObject) ConverDate() (string, string, string, error) {
	var year, month, day string
	_, check := r["PY"]
	if !check {
		return "", "", "", errors.New("date is not found")
	}
	date := r["PY"][0]
	year = date[0:4]
	if date[5] == '/' {
		month = ""
		if date[6] == '/' {
			day = ""
		} else {
			day = date[6:8]
		}
	} else {
		month = date[5:7]
		if date[8] == '/' {
			day = ""
		} else {
			day = date[8:10]
		}
	}
	return year, month, day, nil

}
