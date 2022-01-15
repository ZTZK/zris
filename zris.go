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

// abstract the year from the "PY" tag. An error will return if something wrong.
func (r RisObject) ConvertDateYear() (string, error) {
	var date string
	_, check := r["PY"]
	if !check {
		_, check = r["Y1"]
		if !check {
			return "", errors.New("the py tag is not found")
		}
		date = r["Y1"][0]

	} else {
		date = r["PY"][0]
	}
	if len(date) < 4 {
		return "", errors.New("error with the py tag")
	}
	return date[0:4], nil
}

// abstract the month from the "PY" tag. An error will return if something wrong.
func (r RisObject) ConvertDateMonth() (string, error) {
	var date string
	_, check := r["PY"]
	if !check {
		_, check = r["Y1"]
		if !check {
			return "", errors.New("the py tag is not found")
		}
		date = r["Y1"][0]

	} else {
		date = r["PY"][0]
	}
	if len(date) < 7 {
		return "", errors.New("error with the py tag")
	}
	return date[5:7], nil
}

// abstract the day from the "PY" tag. An error will return if something wrong.
func (r RisObject) ConvertDateDay() (string, error) {
	var date string
	_, check := r["PY"]
	if !check {
		_, check = r["Y1"]
		if !check {
			return "", errors.New("the py tag is not found")
		}
		date = r["Y1"][0]

	} else {
		date = r["PY"][0]
	}

	if len(date) < 10 {
		return "", errors.New("error with the py tag")
	}
	return date[8:10], nil
}
