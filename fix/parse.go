package fix

import (
	"strconv"
	"strings"
)

func ParseFixMessages(message string) (map[int]string, error) {
	fields := strings.Split(message, ",")
	var result = make(map[int]string)
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		result[val] = parts[1]
	}
	return result, nil
}
