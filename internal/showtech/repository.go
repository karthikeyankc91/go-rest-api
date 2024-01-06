package showtech

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access showtech from the data source.
type Repository interface {
	// Get returns the showtech with the specified showtech ID.
	Get(ctx context.Context, id string) (entity.Showtechs, error)
	// Count returns the number of showtech.
	Count(ctx context.Context) (int, error)
	// Query returns the list of showtech with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Showtechs, error)
	// Create saves a new showtech in the storage.
	Create(ctx context.Context, showtech entity.Showtechs) error
	// Update updates the showtech with given ID in the storage.
	Update(ctx context.Context, showtech entity.Showtechs) error
	// Delete removes the showtech with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists showtech in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new showtech repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the showtech with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Showtechs, error) {
	var showtech entity.Showtechs
	err := r.db.With(ctx).Select().Model(id, &showtech)
	return showtech, err
}

// Create saves a new showtech record in the database.
// It returns the ID of the newly inserted showtech record.
func (r repository) Create(ctx context.Context, showtech entity.Showtechs) error {
	return r.db.With(ctx).Model(&showtech).Insert()
}

// Update saves the changes to an showtech in the database.
func (r repository) Update(ctx context.Context, showtech entity.Showtechs) error {
	return r.db.With(ctx).Model(&showtech).Update()
}

// Delete deletes an showtech with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	showtech, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&showtech).Delete()
}

// Count returns the number of the showtech records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("showtechs").Row(&count)
	return count, err
}

// Query retrieves the showtech records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Showtechs, error) {
	var showtech []entity.Showtechs
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&showtech)
	return showtech, err
}
