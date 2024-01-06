package entity

import (
	"encoding/json"
	"time"

	"github.com/qiangxue/go-rest-api/internal/showtech/parser"
)

type Showtechs struct {
	Id string `json:"id"`

	ParsedData json.RawMessage `json:"parsed_data,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type STA struct {
	Id        string
	Data      *parser.Commands
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *STA) Validate() error {
	return nil
}

func (s *Showtechs) Unmarshal() (*STA, error) {
	showtechs := &STA{}
	showtechs.Data = &parser.Commands{}

	err := json.Unmarshal(s.ParsedData, showtechs.Data)
	if err != nil {
		return nil, err
	}
	showtechs.Id = s.Id
	showtechs.CreatedAt = s.CreatedAt
	showtechs.UpdatedAt = s.UpdatedAt
	return showtechs, nil
}

func (s *STA) Marshal() (Showtechs, error) {
	showtech := Showtechs{}
	data, err := json.Marshal(s.Data)
	if err != nil {
		return showtech, err
	}
	showtech.ParsedData = data

	showtech.Id = s.Id
	showtech.CreatedAt = s.CreatedAt
	showtech.UpdatedAt = s.UpdatedAt

	return showtech, nil
}
