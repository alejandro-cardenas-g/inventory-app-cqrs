package api

import (
	"inventory_cqrs/internal/api/dtos"
	"inventory_cqrs/internal/application/commands/products"
	qproducts "inventory_cqrs/internal/application/queries/products"
	dproducts "inventory_cqrs/internal/domain/products"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *App) BindProductsRoutes(router chi.Router) {
	router.Route("/products", func(r chi.Router) {	
		r.Post("/", s.CreateProduct)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetProductByID)
		})
	})
}

func (s *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateProductDTO
	if err := readJSON(w, r, &req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := Validate.Struct(req); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	result, err := s.container.CreateProductHandler.Handler(r.Context(), products.CreateProductCommand{
		SKU: req.SKU,
		Name: req.Name,
		Description: req.Description,
		CategoryID: req.CategoryID,
		Brand: req.Brand,
		Price: req.Price,
		Currency: req.Currency,
		Stock: req.Stock,
		Attributes: req.Attributes,
	})

	if err != nil {
		switch err {
		case dproducts.ErrSKUAlreadyExists:
			writeJSONError(w, http.StatusConflict, err.Error())
			return
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	writeJSON(w, http.StatusCreated, dtos.CreateProductResultDTO{
		ID: result.ID,
	})
}

func (s *App) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	productID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := s.container.GetProductByIDHandler.Handler(r.Context(), qproducts.GetProductByIDQuery{
		ID: productID,
	})

	if err != nil {
		switch err {
		case dproducts.ErrProductNotFound:
			writeJSONError(w, http.StatusNotFound, err.Error())
			return
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	writeJSON(w, http.StatusOK, dtos.GetProductByIDResultDTO{
		ID: product.Product.GetID(),
		SKU: product.Product.GetSKU(),
		Name: product.Product.GetName(),
		Description: product.Product.GetDescription(),
		CategoryID: product.Product.GetCategoryID(),
		Brand: product.Product.GetBrand(),
		Price: product.Product.GetPriceCents(),
		Currency: product.Product.GetCurrency(),
		Stock: product.Product.GetStock(),
		Attributes: product.Product.GetAttributes(),
		CreatedAt: product.Product.GetCreatedAt(),
		UpdatedAt: product.Product.GetUpdatedAt(),
	})
}