package main

import "errors"

var (
	ErrCampaignServiceUnreachable = errors.New("cannot reach campaign cache")
	ErrCampaignCacheTooOld        = errors.New("campaign cache is too old")
	ErrLoadCache                  = errors.New("error loading cache")
)

type BookCache struct {
	books map[string]Book
}
