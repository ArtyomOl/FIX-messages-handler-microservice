package storage

import (
	"FIX-messages-handler-API/fix"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"user"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

func NewClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		Username:     cfg.User,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}

	return db, nil
}

func AddOrder(client *redis.Client, symbol string, price float64, quantity int, side string) error {
	key := fmt.Sprintf("%s:%s", symbol, side)
	return client.ZAdd(context.Background(), key, redis.Z{
		Score:  price,
		Member: quantity,
	}).Err()
}

func GetOrderBook(client *redis.Client, symbol string, depth int) (*fix.OrderBook, error) {
	bidsKey := fmt.Sprintf("%s:bids", symbol)
	asksKey := fmt.Sprintf("%s:asks", symbol)

	bids, err := client.ZRevRangeWithScores(context.Background(), bidsKey, 0, int64(depth-1)).Result()
	if err != nil {
		return nil, err
	}

	asks, err := client.ZRangeWithScores(context.Background(), asksKey, 0, int64(depth-1)).Result()
	if err != nil {
		return nil, err
	}
	var orderBook fix.OrderBook
	orderBook.Symbol = symbol
	for _, z := range asks {
		quantity_str, ok := z.Member.(string)
		if !ok {
			continue
		}
		quantity, err := strconv.Atoi(quantity_str)
		if err != nil {
			return nil, err
		}
		orderBook.AsksQantility = append(orderBook.AsksQantility, quantity)
		orderBook.AsksPrices = append(orderBook.AsksPrices, z.Score)
	}
	for _, z := range bids {
		quantity_str, ok := z.Member.(string)
		if !ok {
			continue
		}
		quantity, err := strconv.Atoi(quantity_str)
		if err != nil {
			return nil, err
		}
		orderBook.BidsQuantility = append(orderBook.AsksQantility, quantity)
		orderBook.BidsPrices = append(orderBook.AsksPrices, z.Score)
	}

	return &orderBook, nil
}
