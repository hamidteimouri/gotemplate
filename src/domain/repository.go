package domain

import "golang.org/x/net/context"

type DbRepository interface {
	FetchCoins(ctx context.Context, req *SearchCoinsEntity) ([]CoinEntity, error)
}
