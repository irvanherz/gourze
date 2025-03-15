package order

import "go.uber.org/fx"

// Module exports dependencies for the order module
var Module = fx.Module("order",
	fx.Provide(NewOrderService),
	fx.Provide(NewOrderController),
)
