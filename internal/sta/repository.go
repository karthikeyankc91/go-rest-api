package sta

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access analysis from the data source.
type Repository interface {
	// Get returns the analysis with the specified analysis ID.
	Get(ctx context.Context, id string) (entity.Analysis, error)
	// Get returns the analysis with the specified analysis ID.
	GetByStaid(ctx context.Context, id string) (entity.Analysis, error)
	// Count returns the number of analysis.
	Count(ctx context.Context) (int, error)
	// Query returns the list of analysis with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Analysis, error)
	// Create saves a new analysis in the storage.
	Create(ctx context.Context, analysis entity.Analysis) error
	// Update updates the analysis with given ID in the storage.
	Update(ctx context.Context, analysis entity.Analysis) error
	// Delete removes the analysis with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists analysis in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new analysis repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the analysis with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Analysis, error) {
	var analysis entity.Analysis
	err := r.db.With(ctx).Select().Model(id, &analysis)
	return analysis, err
}

// Get reads the analysis with the specified ID from the database.
func (r repository) GetByStaid(ctx context.Context, id string) (entity.Analysis, error) {
	var analysis entity.Analysis
	err := r.db.With(ctx).
		Select("id", "staid", "knowledgemap", "created_at", "updated_at").
		From("analysis").
		Where(dbx.HashExp{"staid": id}).
		One(&analysis)
	return analysis, err
}

// Create saves a new analysis record in the database.
// It returns the ID of the newly inserted analysis record.
func (r repository) Create(ctx context.Context, analysis entity.Analysis) error {
	return r.db.With(ctx).Model(&analysis).Insert()
}

// Update saves the changes to an analysis in the database.
func (r repository) Update(ctx context.Context, analysis entity.Analysis) error {
	return r.db.With(ctx).Model(&analysis).Update()
}

// Delete deletes an analysis with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	analysis, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&analysis).Delete()
}

// Count returns the number of the analysis records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("analysis").Row(&count)
	return count, err
}

// Query retrieves the analysis records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Analysis, error) {
	var analysis []entity.Analysis
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&analysis)
	return analysis, err
}
