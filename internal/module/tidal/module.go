package tidal

import (
	authgateway "poseidon/internal/module/tidal/gateway/auth"
	mediagateway "poseidon/internal/module/tidal/gateway/media"
	authservice "poseidon/internal/module/tidal/service/auth"
	mediaservice "poseidon/internal/module/tidal/service/media"

	"go.uber.org/fx"
)

// Module wires the Tidal integration layer into an Fx application.
var Module = fx.Module("tidal",
	fx.Provide(
		fx.Private,
		fx.Annotate(
			authgateway.NewTidalAuthAPI,
			fx.As(new(authgateway.TidalAuthAPI)),
		),
	),
	fx.Provide(
		fx.Private,
		fx.Annotate(
			mediagateway.NewTidalMediaAPI,
			fx.As(new(mediagateway.TidalMediaAPI)),
		),
	),
	fx.Provide(
		fx.Annotate(
			authservice.NewTidalAuthService,
			fx.As(new(authservice.TidalAuthService)),
		),
	),
	fx.Provide(
		fx.Annotate(
			mediaservice.NewTidalMediaService,
			fx.As(new(mediaservice.TidalMediaService)),
		),
	),
)
