package entity

import (
	"time"

	"gitlab.aristanetworks.com/tac-tools/show-tech-analyzer/backend/pkg/showtech/parser"
)

// STA represents an STA record.
type STA struct {
	StaId string `json:"id"`

	ParsedData *parser.Commands `json:"parsed_data"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}
