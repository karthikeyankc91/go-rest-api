package showtech

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/test"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/stretchr/testify/assert"
	"gitlab.aristanetworks.com/tac-tools/show-tech-analyzer/backend/pkg/showtech/parser"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "showtechs")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	parsedData := &parser.Commands{}
	// create
	err = repo.Create(ctx, entity.STA{
		StaId:      "1",
		ParsedData: parsedData,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	showtech, err := repo.Get(ctx, "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", showtech.StaId)
	_, err = repo.Get(ctx, "2")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, entity.STA{
		StaId:      "1",
		ParsedData: parsedData,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	showtech, _ = repo.Get(ctx, "1")
	assert.Equal(t, "1", showtech.StaId)

	// query
	showtechs, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(showtechs))

	// delete
	err = repo.Delete(ctx, "1")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "1")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "1")
	assert.Equal(t, sql.ErrNoRows, err)
}
