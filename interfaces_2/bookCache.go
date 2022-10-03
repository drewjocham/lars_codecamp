package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrCacheUnreachable = errors.New("cannot reach book cache")
	ErrLoadCache        = errors.New("error loading cache")
)

type BookCache struct {
	cache     map[string]*Book
	repo      *PostgresRepository
	cacheLock *sync.RWMutex
	updatedAt time.Time
}

func (c *BookCache) load() error {
	bookCache := make(map[string]*Book)
	books, err := c.repo.getAllBooks()

	if err != nil {
		return ErrLoadCache
	}

	// Talk about why I did it like this
	for i, _ := range books {
		id := books[i].Id.String()
		bookCache[id] = &books[i]
	}

	c.cache = bookCache

	return nil
}

// TODO: why is retires not working?
func (c *BookCache) Init(maxRetries int, retryPeriod time.Duration) error {
	var retries int

	for retries = 0; retries < maxRetries; retries++ {
		err := c.load()

		if err == nil {
			break
		}

		retries++
		return fmt.Errorf("%s: %w", ErrLoadCache, err)
	}

	if retries == maxRetries {
		return ErrCacheUnreachable
	}

	return nil
}

func (c *BookCache) StartRefresh(ctx context.Context) error {

	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			err := c.load()
			if err != nil {
				log.Fatal(ErrLoadCache)
			}
		}
	}

	return nil
}

func (c *BookCache) GetBookById(id string) (*Book, bool) {
	log.Println("Getting book from cache with id ", id)

	book, ok := c.cache[id]

	return book, ok
}

func (c *BookCache) GetAllBooks() []Book {
	var books []Book
	for _, book := range c.cache {
		books = append(books, *book)
	}
	return books

}

func NewBookCache(r *PostgresRepository) *BookCache {
	return &BookCache{repo: r}
}
