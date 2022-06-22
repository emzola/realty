package data

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/emzola/realty/internal/validator"
	"github.com/lib/pq"
)

// Property contains information about a property
type Property struct {
	ID          int64             `json:"id"`
	CreatedAt   time.Time         `json:"created_at,omitempty"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	City        string            `json:"city"`
	Location    string            `json:"location"`
	Latitude    float64           `json:"latitude,omitempty"`
	Longitude   float64           `json:"longitude,omitempty"`
	Type        []string          `json:"type,omitempty"`
	Category    []string          `json:"category,omitempty"`
	Features    Features				  `json:"features,omitempty"`
	Price       float64           `json:"price"`
	Currency    []string          `json:"currency"`
	Nearby      Nearby 						`json:"nearby,omitempty"`
	Amenities   []string          `json:"amenities,omitempty"`
	Version     int32             `json:"version"`
}

// Features contains features of a property
type Features map[string]interface{}

func (a Features) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Features) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
			return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// Nearby contains information about a nearby facilities
type Nearby map[string]interface{}

func (a Nearby) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Nearby) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
			return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}


// ValidateProperty validates a property based on set validation criteria.
func ValidateProperty(v *validator.Validator, property *Property) {
	v.Check(property.Title != "", "title", "must be provided")
	v.Check(len(property.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(property.Description != "", "description", "must be provided")
	v.Check(property.City != "", "city", "must be provided")
	v.Check(property.Location != "", "location", "must be provided")
	v.Check(len(property.Type) <= 1, "type", "must not contain more than 1 type")
	// v.Check(property.Type[0] != "", "type", "must be a provided")
	v.Check(len(property.Category) <= 1, "category", "must not contain more than 1 category")
	// v.Check(property.Category[0] != "", "category", "must be provided")
	v.Check(property.Features != nil, "features", "must be provided")
	v.Check(len(property.Features) >= 1, "features", "must contain at least 1 feature")
	v.Check(len(property.Features) <= 20, "features", "must not contain more than 20 features")
	v.Check(property.Price != 0, "price", "must be provided")
	v.Check(property.Price > 0, "price", "must be a positive number")
	v.Check(len(property.Currency) <= 1, "currency", "must not contain more than 1 currency")
	// v.Check(property.Currency[0] == "USD" || property.Currency[0] == "GBP" || property.Currency[0] == "EUR", "currency", "must be USD, GBP or EUR")
	v.Check(property.Nearby != nil, "nearby", "must be provided")
	v.Check(len(property.Nearby) >= 1, "nearby", "must contain at least 1 facility")
	v.Check(len(property.Nearby) <= 10, "nearby", "must not contain more than 10 facilities")
	v.Check(validator.Unique(property.Amenities), "amenities", "must not contain duplicate values")
}


// PropertyModel struct wraps a sql.DB connection pool.
type PropertyModel struct {
	DB *sql.DB
}

// Insert inserts a new record into the property table.
func (p PropertyModel) Insert(property *Property) error {
	query := `
	INSERT INTO properties(title, description, city, location, latitude, longitude, type, category, features, price, currency, nearby, amenities)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id, created_at, version`


	args := []interface{}{property.Title, property.Description, property.City, property.Location, property.Latitude, property.Longitude, pq.Array(property.Type), pq.Array(property.Category), property.Features, property.Price, pq.Array(property.Currency), property.Nearby, pq.Array(property.Amenities)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return p.DB.QueryRowContext(ctx, query, args...).Scan(&property.ID, &property.CreatedAt, &property.Version)
}

// Get fetches a specific record from the properties table.
func (p PropertyModel) Get(id int64) (*Property, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT id, created_at, title, description, city, location, latitude, longitude, type, category, features, price, currency, nearby, amenities, version
	FROM properties
	WHERE id = $1`

	var property Property

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&property.ID, 
		&property.CreatedAt, 
		&property.Title, 
		&property.Description, 
		&property.City, 
		&property.Location, 
		&property.Latitude, 
		&property.Longitude, 
		pq.Array(&property.Type), 
		pq.Array(&property.Category), 
		&property.Features, 
		&property.Price, 
		pq.Array(&property.Currency), 
		&property.Nearby, 
		pq.Array(&property.Amenities),
		&property.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err			
		}
	}

	return &property, nil
}

// Update updates a specific record in the properties table.
func (p PropertyModel) Update(property *Property) error {
	query := `UPDATE properties
	SET title = $1, description = $2, city = $3, location = $4, latitude = $5, longitude = $6, type = $7, category = $8, features = $9, price = $10, currency = $11, nearby = $12, amenities = $13, version = version + 1
	WHERE id = $14 AND version = $15
	RETURNING version`

	args := []interface{}{property.Title, property.Description, property.City, property.Location, property.Latitude, property.Longitude, pq.Array(property.Type), pq.Array(property.Category), property.Features, property.Price, pq.Array(property.Currency), property.Nearby, pq.Array(property.Amenities), property.ID, property.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() 
	
	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&property.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// Delete deletes a specific record from the peoperties table
func (p PropertyModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE FROM properties
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() 

	result, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
