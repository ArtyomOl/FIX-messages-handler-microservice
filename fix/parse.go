package fix

import (
	"FIX-messages-handler-API/storage"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ParseFixMessages(message string) (map[int]string, error) {
	fields := strings.Split(message, "|")
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

func AddFixMessage(client *redis.Client, message string) error {
	m, err := ParseFixMessages(message)
	if err != nil {
		return err
	}
	symbol := m[55]
	price, err := strconv.ParseFloat(m[44], 64)
	if err != nil {
		return err
	}
	quantity, err := strconv.Atoi(m[38])
	if err != nil {
		return err
	}
	var side string
	switch m[54] {
	case "1":
		side = "asks"
	case "2":
		side = "bids"
	default:
		return fmt.Errorf("invalid values")
	}
	storage.AddOrder(client, symbol, price, quantity, side)
	return nil
}
