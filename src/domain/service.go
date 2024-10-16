package domain

import (
	"context"
	"github.com/sirupsen/logrus"
)

type LensService struct {
	dbRepo DbRepository
}

func NewLensService(dbRepo DbRepository) *LensService {
	ls := &LensService{
		dbRepo: dbRepo,
	}

	return ls
}

func (l *LensService) GetCoins(ctx context.Context, req *SearchCoinsEntity) ([]CoinEntity, error) {
	result, err := l.dbRepo.FetchCoins(ctx, req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": req,
		}).Error("failed to get coins")
		return nil, err
	}
	return result, nil
}
