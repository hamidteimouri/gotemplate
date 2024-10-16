package postgres

import (
	"gitlab.com/hamidteimouri/core-lens/domain"
	"time"
)

type CoinMode struct {
	Id                    uint64
	BrokerId              string
	Symbol                string
	Name                  string
	NameLocale            string
	NeedKyc               bool
	IsActive              bool
	IsSuspended           bool
	CanDeposit            bool
	CanWithdraw           bool
	Position              int
	Website               string
	SuspensionDescription string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (b *CoinMode) TableName() string {
	return "coins"
}

func (b *CoinMode) ConvertEntityToModel(e *domain.CoinEntity) {
	b.Symbol = e.Symbol
	b.Name = e.Name
	b.NameLocale = e.NameLocale
	b.NeedKyc = e.NeedKyc
	b.IsActive = e.IsActive
	b.IsSuspended = e.IsSuspended
	b.CanDeposit = e.CanDeposit
	b.CanWithdraw = e.CanWithdraw
	b.Position = e.Position
	b.Website = e.Website
	b.SuspensionDescription = e.SuspensionDescription
	b.CreatedAt = e.CreatedAt
	b.UpdatedAt = e.UpdatedAt

}

func (b *CoinMode) ConvertModelToEntity() *domain.CoinEntity {
	return &domain.CoinEntity{
		Id:                    b.Id,
		Symbol:                b.Symbol,
		Name:                  b.Name,
		NameLocale:            b.NameLocale,
		NeedKyc:               b.NeedKyc,
		IsActive:              b.IsActive,
		IsSuspended:           b.IsSuspended,
		CanDeposit:            b.CanDeposit,
		CanWithdraw:           b.CanWithdraw,
		Position:              b.Position,
		Website:               b.Website,
		SuspensionDescription: b.SuspensionDescription,
		CreatedAt:             b.CreatedAt,
		UpdatedAt:             b.UpdatedAt,
	}
}
