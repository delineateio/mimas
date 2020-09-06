package db

import (
	"database/sql"
	"fmt"

	"github.com/avast/retry-go"
	"github.com/delineateio/mimas/log"

	// required for the postgres driver
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// IRepository for all repositories
type IRepository interface {
	Migrate() error
}

// Repository that reprents the access to the underlying database
type Repository struct {
	info     *Info
	settings *Settings
	core     *gorm.DB
	db       *sql.DB
}

// NewDefaultRepository returns production database access
func NewDefaultRepository() (*Repository, error) {
	info, err := NewDefaultInfo()
	if err != nil {
		return nil, err
	}
	return NewRepository(info)
}

// NewRepository returns production database access
func NewRepository(info *Info) (*Repository, error) {
	return &Repository{
		info:     info,
		settings: NewSettings(info.Name),
	}, nil
}

func (r *Repository) setDB() error {
	// Opens the DB using the new Gorm interface
	db, err := gorm.Open(postgres.Open(r.info.ConnectionString()), &gorm.Config{})
	if err != nil {
		log.Error("db.connection", err)
		return err
	}
	r.core = db
	r.db, err = db.DB()
	if err != nil {
		log.Error("db.connection", err)
		return err
	}
	return nil
}

// Ping pings the underlying database to ensure it's contactable
func (r *Repository) Ping() error {
	err := r.setDB()
	if err != nil {
		log.Error("db.ping", err)
		return err
	}

	return r.db.Ping()
}

// Open the database and sets the underlying configuration
func (r *Repository) Open() error {
	tries := 0

	opts := []retry.Option{
		retry.Attempts(r.settings.Attempts),
		retry.Delay(r.settings.Delay),
		retry.OnRetry(func(try uint, err error) {
			log.Warn("db.open.error", "failed on attempt "+fmt.Sprint(try+1))
		}),
	}

	err := retry.Do(
		func() error {
			tries++
			err := r.setDB()
			if err != nil {
				return err
			}
			return nil
		},
		opts...,
	)
	if err != nil {
		log.Error("db.open.error", err)
		return err
	}

	// Sets the more advanced settings
	r.db.SetMaxOpenConns(r.settings.MaxOpen)
	r.db.SetMaxIdleConns(r.settings.MaxIdle)
	r.db.SetConnMaxLifetime(r.settings.MaxLifetime)

	return nil
}

// Migrate placeholder for service specific migration
func (r *Repository) Migrate(entities []interface{}) error {
	if entities == nil {
		log.Warn("db.migrate.entities", "no db entities were provided to be migrated")
		return nil
	}
	err := r.Open()
	if err != nil {
		return err
	}
	err = r.core.AutoMigrate(entities...)
	if err != nil {
		// better to report the earlier error
		log.Error("db.migrate.error", err)
		_ = r.Close()
		return err
	}
	log.Info("db.migrate", "successfully migrated the db")
	return r.Close()
}

// Create the entity in the database
func (r *Repository) Create(entity interface{}) error {
	err := r.Open()
	if err != nil {
		return err
	}
	err = r.core.Create(entity).Error
	if err != nil {
		log.Error("db.create", err)
	}
	return r.Close()
}

// Close the DB connection
func (r *Repository) Close() error {
	if r.db == nil {
		return nil
	}
	return r.db.Close()
}
