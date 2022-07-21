package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Do(s string, i int, b bool) (string, error) {
	var res string
	switch i {
	case 1, 2, 3, 5, 8:
		res += strconv.Itoa(i)
	case 13, 21, 34:
	default:
		return "", errors.New("invalid i")
	}

	switch s {
	case "a", "b", "c", "d":
		res += s
	default:
		return "", errors.New("invalid s")
	}

	if b {
		res = strings.ToUpper(res)
	}

	if len(res) >= 2 {
		res = fmt.Sprintf("[%s]", res)
	}

	return res, nil
}
