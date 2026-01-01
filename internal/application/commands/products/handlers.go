package products

import (
	"context"
	"encoding/json"
	"inventory_cqrs/internal/domain/outbox"
	"inventory_cqrs/internal/domain/products"
	p "inventory_cqrs/internal/store/persistence"
	"time"
)

type CreateProductHandler struct {
	productsRepository *p.ProductRepository
	outboxRepository   *p.OutboxRepository
	uow                *p.TxManager
}

func NewCreateProductHandler(productsRepository *p.ProductRepository, outboxRepository *p.OutboxRepository, uow *p.TxManager) *CreateProductHandler {
	return &CreateProductHandler{productsRepository: productsRepository, outboxRepository: outboxRepository, uow: uow}
}

func (h *CreateProductHandler) Handler(ctx context.Context, command CreateProductCommand) (*CreateProductResult, error) {

	var r *CreateProductResult

	err := h.uow.WithTx(ctx, func(tx p.DBTX) error {

		pr := h.productsRepository.UseTX(tx)
		ob := h.outboxRepository.UseTX(tx)

		product, err := command.ToProduct()

		if err != nil {
			return err
		}

		result, err := pr.Create(ctx, product)

		if err != nil {
			return err
		}

		payload, err := json.Marshal(products.ProductCreated{
			ProductID:  result.ID,
			OccurredAt: time.Now(),
		})

		if err != nil {
			return products.ErrCreatingProduct
		}

		event := outbox.New(
			products.ProductCreatedEventType,
			"Product",
			result.ID,
			payload,
			nil,
			nil,
		)

		err = ob.Save(ctx, event)

		if err != nil {
			return err
		}

		r = &CreateProductResult{
			ID: result.ID,
		}

		return nil
	})

	return r, err
}
