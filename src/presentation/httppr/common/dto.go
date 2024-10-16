package common

import (
	"gitlab.com/hamidteimouri/core-lens/domain"
	"time"
)

type CoinsDto struct {
	Id                    uint64 `json:"id"`
	Symbol                string `json:"symbol"`
	Name                  string `json:"name"`
	NameLocale            string `json:"name_locale"`
	NeedKyc               bool   `json:"need_kyc"`
	IsActive              bool   `json:"is_active"`
	IsSuspended           bool   `json:"is_suspended"`
	CanDeposit            bool   `json:"can_deposit"`
	CanWithdraw           bool   `json:"can_withdraw"`
	Position              int    `json:"position"`
	Website               string `json:"website"`
	SuspensionDescription string `json:"suspension_description"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}

func ConvertCoinEntityToDto(item *domain.CoinEntity) CoinsDto {

	return CoinsDto{
		Id:                    item.Id,
		Symbol:                item.Symbol,
		Name:                  item.Name,
		NameLocale:            item.NameLocale,
		NeedKyc:               item.NeedKyc,
		IsActive:              item.IsActive,
		IsSuspended:           item.IsSuspended,
		CanDeposit:            item.CanDeposit,
		CanWithdraw:           item.CanWithdraw,
		Position:              item.Position,
		Website:               item.Website,
		SuspensionDescription: item.SuspensionDescription,
		CreatedAt:             item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             item.UpdatedAt.Format(time.RFC3339),
	}
}
