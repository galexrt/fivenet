package rector

import (
	"database/sql"

	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/server/audit"
	"go.uber.org/zap"
)

type Server struct {
	RectorServiceServer

	logger *zap.Logger
	db     *sql.DB
	p      perms.Permissions
	aud    audit.IAuditer
	c      *mstlystcdata.Enricher
}

func NewServer(logger *zap.Logger, db *sql.DB, p perms.Permissions, aud audit.IAuditer, c *mstlystcdata.Enricher) *Server {
	return &Server{
		logger: logger,
		db:     db,
		p:      p,
		aud:    aud,
		c:      c,
	}
}
