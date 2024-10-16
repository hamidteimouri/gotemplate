package postgres

import (
	"context"
	"gitlab.com/hamidteimouri/core-lens/domain"
)

func (p *Postgres) FetchCoins(ctx context.Context, req *domain.SearchCoinsEntity) ([]domain.CoinEntity, error) {
	var models []CoinMode

	query := p.db.WithContext(ctx).Model(CoinMode{})

	if req.Symbol != "" {
		query.Where("symbol = ?", req.Symbol)
	}

	if req.IsActive != "" {
		if req.IsActive == "true" {
			query.Where("is_active = ?", true)
		}
		if req.IsActive == "false" {
			query.Where("is_active = ?", false)
		}
	}

	query.Order("position ASC")

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	objects := make([]domain.CoinEntity, len(models))
	for i, cm := range models {
		objects[i] = *cm.ConvertModelToEntity()
	}

	return objects, nil
}
