package handlers

import (
	"errors"
	"github.com/Aegon95/mytheresa-product-api/internal/db/entity"
	"github.com/Aegon95/mytheresa-product-api/internal/service"
	"github.com/Aegon95/mytheresa-product-api/internal/util"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ProductHandler interface {
	GetProducts() http.HandlerFunc
}

type productHandler struct {
	log *zap.SugaredLogger
	productService service.ProductService
}

func NewProductHandler(log *zap.SugaredLogger, productService service.ProductService) ProductHandler {

	return  &productHandler{
		log,
		productService,
	}

}

func (p *productHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		category := r.URL.Query().Get("category")

		priceLessThanParam := r.URL.Query().Get("priceLessThan")

		priceLessThan, err := strconv.ParseInt(priceLessThanParam, 10, 64)

		if !entity.IsValid(category) {
			p.log.Info("Error occurred while parsing query params")
			err = errors.New("bad request.")
			util.RenderErrInvalidRequest(w, err)
			return
		}
		products, err := p.productService.GetProducts(ctx, entity.Category(category), priceLessThan)

		if err != nil {
			err = errors.New("internal server error.")
			util.RenderErrInternal(w, err)
			return
		}

		if products == nil {
			util.RenderErrNotFound(w)
			return
		}

		util.RenderJSON(w, http.StatusOK, products)

	}
}
