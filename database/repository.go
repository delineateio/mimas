package database

import (
	"database/sql"
	"fmt"

	"github.com/avast/retry-go"
	log "github.com/delineateio/mimas/log"

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

// NewRepository returns production database access
func NewRepository() *Repository {
	info := NewInfo()
	return &Repository{
		info:     info,
		settings: NewSettings(info.Name),
	}
}

func (r *Repository) setDB() error {
	// Opens the DB using the new Gorm interface
	db, err := gorm.Open(postgres.Open(r.info.ConnectionString()), &gorm.Config{})
	if err != nil {
		log.Error("db.connection", err)
		return err
	}
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
	err := retry.Do(
		func() error {
			tries++
			err := r.setDB()
			if err != nil {
				return err
			}
			return nil
		},
		retry.Attempts(r.settings.Attempts),
		retry.Delay(r.settings.Delay),
		retry.OnRetry(func(try uint, err error) {
			log.Warn("db.open.error", "failed on attempt "+fmt.Sprint(try+1))
		}),
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
func (r *Repository) Migrate(entity interface{}) error {
	err := r.Open()
	if err != nil {
		return err
	}
	err = r.core.AutoMigrate(entity)
	if err != nil {
		// better to report the earlier error
		log.Error("db.migrate.error", err)
		_ = r.Close()
		return err
	}
	err = r.Close()
	if err != nil {
		return err
	}
	log.Info("db.migrate", "successfully migrated the db")
	return nil
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
	err := r.db.Close()
	if err != nil {
		log.Error("db.close.error", err)
		return err
	}

	return nil
}
