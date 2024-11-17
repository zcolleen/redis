package redis

import "github.com/zcolleen/redis/internal"

type Client = internal.Client

func NewClient() Client {
	return &internal.ClientEntity{}
}
