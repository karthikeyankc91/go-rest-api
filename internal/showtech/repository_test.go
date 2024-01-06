package showtech

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/showtech/parser"
	"github.com/qiangxue/go-rest-api/internal/test"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/stretchr/testify/assert"
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

	showtechData := entity.STA{
		Id: "1",
		Data: &parser.Commands{
			ShowArp: parser.ShowArp{
				Output: parser.Output{
					Meta: parser.Meta{
						Command: "show arp",
					},
				},
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	showtechs, err := showtechData.Marshal()
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, showtechs)
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	showtechs, err = repo.Get(ctx, "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", showtechs.Id)
	_, err = repo.Get(ctx, "2")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, showtechs)
	assert.Nil(t, err)
	showtechs, _ = repo.Get(ctx, "1")
	assert.Equal(t, "1", showtechs.Id)

	// query
	showtechsList, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(showtechsList))

	// delete
	err = repo.Delete(ctx, "1")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "1")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "1")
	assert.Equal(t, sql.ErrNoRows, err)
}
