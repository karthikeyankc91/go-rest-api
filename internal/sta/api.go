package sta

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Use(authHandler)

	r.Post("/analyze", res.analyze)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) analyze(c *routing.Context) error {
	var input CreateSTARequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	analysis, err := r.service.Analyze(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(analysis, http.StatusCreated)
}
