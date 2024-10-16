package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.com/hamidteimouri/core-lens/domain"
	"gitlab.com/hamidteimouri/core-lens/presentation/httppr/common"
)

type LensHandler struct {
	service *domain.LensService
}

func NewLensHandler(service *domain.LensService) *LensHandler {
	return &LensHandler{
		service: service,
	}
}

// GetCoins This function handle and returns list of coins
func (l *LensHandler) GetCoins(c echo.Context) error {
	req := common.GetCoinsRequest{}

	err := c.Bind(&req)
	if err != nil {
		return common.RespBadRequest(c, errors.New("bad_request"))
	}

	coins, err := l.service.GetCoins(c.Request().Context(), &domain.SearchCoinsEntity{
		Symbol: req.Symbol,
	})
	if err != nil {
		return common.RespBadRequest(c, err)
	}

	objects := make(map[string]common.CoinsDto, len(coins))
	for _, item := range coins {
		objects[item.Symbol] = common.ConvertCoinEntityToDto(&item)
	}

	return common.RespOK(c, objects)
}
