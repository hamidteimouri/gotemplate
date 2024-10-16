package domain

import (
	"time"
)

type CoinEntity struct {
	Id                    uint64
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

type SearchCoinsEntity struct {
	Symbol   string
	IsActive string
}
