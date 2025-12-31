package bs

import (
	"inventory_cqrs/internal/application/commands/products"
	qproducts "inventory_cqrs/internal/application/queries/products"
	"inventory_cqrs/internal/config"
	"inventory_cqrs/internal/store/persistance"
)

type Container struct {
	Uow *persistance.TxManager
	CreateProductHandler *products.CreateProductHandler
	GetProductByIDHandler *qproducts.GetProductByIDHandler
}


func InjectServices(cfg config.Config) *Container {
	persistance := persistance.NewStore(cfg.Store)

	container := &Container{}
	container.Uow = persistance.TxManager

	// products
	container.CreateProductHandler = products.NewCreateProductHandler(persistance.Products)
	container.GetProductByIDHandler = qproducts.NewGetProductByIDHandler(persistance.Products)

	return container
}