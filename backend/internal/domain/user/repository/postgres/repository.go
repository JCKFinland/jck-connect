package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
	userrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ userrepo.Repository = (*repository)(nil)

// New creates a PostgreSQL user repository.
func New(db *database.Database) userrepo.Repository {
	return &repository{
		db: db,
	}
}

func NewTx(
	tx database.DBTX,
) userrepo.Repository {

	return &repository{
		db: tx,
	}
}

// Create inserts a new user.
func (r *repository) Create(
	ctx context.Context,
	user *entity.User,
) error {

	const query = `
INSERT INTO users (
	id,
	pi_uid,
	pi_username,
	display_name,
	email,
	phone_number,
	role,
	status,
	created_at,
	updated_at
)
VALUES (
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.PiUID,
		user.PiUsername,
		user.DisplayName,
		user.Email,
		user.PhoneNumber,
		user.Role,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// Update updates an existing user.
func (r *repository) Update(
	ctx context.Context,
	user *entity.User,
) error {

	const query = `
UPDATE users
SET
	display_name = $2,
	email = $3,
	phone_number = $4,
	updated_at = $5
WHERE id = $1
`

	commandTag, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.DisplayName,
		user.Email,
		user.PhoneNumber,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

// FindByID returns a user by ID.
func (r *repository) FindByID(
	ctx context.Context,
	id string,
) (*entity.User, error) {

	const query = `
SELECT
	id,
	pi_uid,
	pi_username,
	display_name,
	email,
	phone_number,
	role,
	status,
	created_at,
	updated_at
FROM users
WHERE id = $1
`

	var user entity.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.PiUID,
		&user.PiUsername,
		&user.DisplayName,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sharedErrors.UserNotFound(
				sharedErrors.ErrNotFound,
			)
		}

		return nil, err
	}

	return &user, nil
}

// FindByPiUID returns a user by Pi UID.
func (r *repository) FindByPiUID(
	ctx context.Context,
	piUID string,
) (*entity.User, error) {

	const query = `
SELECT
	id,
	pi_uid,
	pi_username,
	display_name,
	email,
	phone_number,
	role,
	status,
	created_at,
	updated_at
FROM users
WHERE pi_uid = $1
`

	var user entity.User

	err := r.db.QueryRow(ctx, query, piUID).Scan(
		&user.ID,
		&user.PiUID,
		&user.PiUsername,
		&user.DisplayName,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sharedErrors.UserNotFound(
				sharedErrors.ErrNotFound,
			)
		}

		return nil, err
	}

	return &user, nil
}



// FindByPiUsername returns a user by Pi username.
func (r *repository) FindByPiUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {

	const query = `
SELECT
	id,
	pi_uid,
	pi_username,
	display_name,
	email,
	phone_number,
	role,
	status,
	created_at,
	updated_at
FROM users
WHERE pi_username = $1
`

	var user entity.User

	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.PiUID,
		&user.PiUsername,
		&user.DisplayName,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sharedErrors.UserNotFound(
				sharedErrors.ErrNotFound,
			)
		}

		return nil, err
	}

	return &user, nil
}
// List returns all users.
func (r *repository) List(
	ctx context.Context,
) ([]*entity.User, error) {

	const query = `
SELECT
	id,
	pi_uid,
	pi_username,
	display_name,
	email,
	phone_number,
	role,
	status,
	created_at,
	updated_at
FROM users
ORDER BY created_at DESC
`

rows, err := r.db.Query(
	ctx,
	query,
)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*entity.User, 0)

	for rows.Next() {
		var user entity.User

		if err := rows.Scan(
			&user.ID,
			&user.PiUID,
			&user.PiUsername,
			&user.DisplayName,
			&user.Email,
			&user.PhoneNumber,
			&user.Role,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Delete is intentionally not implemented.
// Users should be deactivated instead of permanently removed.
func (r *repository) Delete(
	ctx context.Context,
	id string,
) error {
	return errors.New("user deletion is not supported")
}
