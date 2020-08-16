package common

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"github.com/jinzhu/gorm"

	// Used internally by gorm to load the postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// DefaultDBType is the default db type
const DefaultDBType = "postgres"
const defaultRetries = 3
const defaultDelayMilliseconds = 500
const defaultMaxIdle = 500
const defaultMaxOpen = 50
const defaultMaxLifetime = 60

// IRepository for all repositories
type IRepository interface {
	Migrate() error
}

// Repository that reprents the access to the underlying database
type Repository struct {
	Name           string
	Database       *gorm.DB
	Username       string
	Password       string
	DBName         string
	DBTypeKey      string
	AllowedDBTypes []string
	DefaultDBType  string
	Attempts       uint // Number of attempts
	Delay          time.Duration
	MaxIdle        int
	MaxOpen        int
	MaxLifetime    time.Duration
	SetDBFunc      func() (*gorm.DB, error)
	Info           DBInfo
}

// DBInfo represents the connection details
type DBInfo struct {
	Type             string
	ConnectionString string
	Tries            int // Actual attempts
}

// NewRepository returns production database access
func NewRepository(name string) *Repository {
	return &Repository{
		Name:           name,
		Username:       os.Getenv("DB_USERNAME"),
		Password:       os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBTypeKey:      "db." + name + ".type",
		AllowedDBTypes: []string{"postgres"},
		DefaultDBType:  DefaultDBType,
		Attempts:       GetUint("db."+name+".retries.attempts", defaultRetries),
		Delay:          GetDuration("db."+name+".retries.delay", defaultDelayMilliseconds*time.Millisecond),
		MaxIdle:        GetInt("db."+name+".limits.maxIdle", defaultMaxIdle),
		MaxOpen:        GetInt("db."+name+".limits.maxOpen", defaultMaxOpen),
		MaxLifetime:    GetDuration("db."+name+".limits.maxLifetime", defaultMaxLifetime*time.Minute),
	}
}

// Default to postgres
// Get the connection string
// Get a value not in approved list (error)
func (r *Repository) dbTypeAllowed(expect string, list []string) bool {
	for _, current := range list {
		if strings.EqualFold(expect, current) {
			return true
		}
	}
	return false
}

func (r *Repository) getDatabaseType() (string, error) {
	var err error
	dbType := GetString(r.DBTypeKey, DefaultDBType)

	if !r.dbTypeAllowed(dbType, r.AllowedDBTypes) {
		err = errors.New("no db type was provided")
		Error("db.connection.error", err)
		dbType = DefaultDBType
	}

	return dbType, err
}

func (r *Repository) getConnectionString() (string, error) {
	// Retrieves from env variables
	var err error

	if r.Username == "" || r.Password == "" || r.DBName == "" {
		err = errors.New("no connection string was provided")
		Error("db.connection.error", err)
		return "", err
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", r.Username, r.Password, r.DBName)
	return connectionString, nil
}

func (r *Repository) getInfo() (DBInfo, error) {
	dbType, err := r.getDatabaseType()
	if err != nil {
		return DBInfo{}, err
	}

	dbConnectionString, err := r.getConnectionString()
	if err != nil {
		return DBInfo{}, err
	}

	Debug("db.connection", dbType+" - "+dbConnectionString)

	info := DBInfo{
		Type:             dbType,
		ConnectionString: dbConnectionString,
		Tries:            0,
	}

	r.Info = info
	return info, nil
}

// Ping pings the underlying database to ensure it's contactable
func (r *Repository) Ping() error {
	info, err := r.getInfo()
	if err != nil {
		return err
	}

	// Ensures func set to open DB
	r.setDB(info)

	db, err := r.SetDBFunc()
	if err != nil {
		Error("db.connection", err)
		return err
	}

	return db.DB().Ping()
}

// Ensures that a func is set to open the DB
func (r *Repository) setDB(info DBInfo) {
	// Enables the replacing of the underlying DB connection
	if r.SetDBFunc == nil {
		r.SetDBFunc = func() (*gorm.DB, error) {
			return gorm.Open(info.Type, info.ConnectionString)
		}
	}
}

// Open the database and sets the underlying configuration
func (r *Repository) Open() error {
	info, err := r.getInfo()
	if err != nil {
		return err
	}

	// Ensures func set to open DB
	r.setDB(info)

	err = retry.Do(
		func() error {
			r.Info.Tries++
			r.Database, err = r.SetDBFunc()
			if err != nil {
				return err
			}
			return nil
		},
		retry.Attempts(r.Attempts),
		retry.Delay(r.Delay),
		retry.OnRetry(func(n uint, err error) {
			Warn("db.open.error", "failed on attempt "+fmt.Sprint(n+1))
		}),
	)
	if err != nil {
		Error("db.open.error", err)
		return err
	}

	// Sets the more advanced settings
	r.Database.DB().SetMaxOpenConns(r.MaxOpen)
	r.Database.DB().SetMaxIdleConns(r.MaxIdle)
	r.Database.DB().SetConnMaxLifetime(r.MaxLifetime)

	return nil
}

// Migrate placeholder for service specific migration
func (r *Repository) Migrate(entity interface{}) error {
	err := r.Open()
	if err != nil {
		return err
	}
	err = r.Database.AutoMigrate(entity).Error
	if err != nil {
		// better to report the earlier error
		Error("db.migrate.error", err)
		_ = r.Close()
		return err
	}
	err = r.Close()
	if err != nil {
		return err
	}
	Info("db.migrate", "successfully migrated the db")
	return nil
}

// Create the entity in the database
func (r *Repository) Create(entity interface{}) error {
	err := r.Open()
	if err != nil {
		return err
	}

	err = r.Database.Create(entity).Error
	if err != nil {
		Error("db.create", err)
	}

	err = r.Close()
	if err != nil {
		return err
	}

	return nil
}

// Close the DB connection
func (r *Repository) Close() error {
	err := r.Database.Close()
	if err != nil {
		Error("db.close.error", err)
		return err
	}

	return nil
}
