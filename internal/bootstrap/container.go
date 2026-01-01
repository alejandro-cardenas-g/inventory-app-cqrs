package bs

import (
	"inventory_cqrs/internal/application/commands/products"
	qproducts "inventory_cqrs/internal/application/queries/products"
	"inventory_cqrs/internal/config"
	"inventory_cqrs/internal/constants"
	"inventory_cqrs/internal/infra/elastic"
	"inventory_cqrs/internal/infra/publisher"
	"inventory_cqrs/internal/store/persistence"

	"go.uber.org/zap"
)

type Container struct {
	Uow                   *persistence.TxManager
	CreateProductHandler  *products.CreateProductHandler
	GetProductByIDHandler *qproducts.GetProductByIDHandler
}

func InjectServices(cfg config.Config) *Container {
	persistance := persistence.NewStore(cfg.Store)

	container := &Container{}
	container.Uow = persistance.TxManager

	// products
	container.CreateProductHandler = products.NewCreateProductHandler(persistance.Products, persistance.Outbox, persistance.TxManager)
	container.GetProductByIDHandler = qproducts.NewGetProductByIDHandler(persistance.Products)

	return container
}

type WorkerContainer struct {
	Uow               *persistence.TxManager
	DispatcherService *publisher.DispatcherService
}

func NewWorkerContainer(cfg config.Config, logger *zap.Logger) *WorkerContainer {
	publishersMap := make(map[constants.Event]publisher.IPublisher)

	persistance := persistence.NewStore(cfg.Store)

	publishersMap[constants.ProductCreatedEvent] = elastic.NewCreateProductPublisher(persistance.ProductView)

	dispatcherService := publisher.NewDispatcherService(publishersMap, persistance.TxManager, persistance.Outbox, logger, cfg.Worker)

	return &WorkerContainer{
		Uow:               persistance.TxManager,
		DispatcherService: dispatcherService}
}
