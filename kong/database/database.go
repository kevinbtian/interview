package database

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"

    _ "github.com/lib/pq"
)

// Database connection configuration struct
type Config struct {
    ProjectID  string
    InstanceID string
    Database   string
    Username   string
    Password   string
}

type Database struct {
    cfg  Config
    conn *sql.DB
}

func NewDatabase(cfg Config) *Database {
    return &Database{cfg: cfg}
}

// Function to create a new database connection
func (d *Database) NewConnection() error {
    dbSocketDir := os.Getenv("DB_SOCKET_DIR")
    if dbSocketDir == "" {
        dbSocketDir = filepath.Join(os.TempDir(), "cloudsql")
    }

    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s/%s sslmode=disable", d.cfg.Username, d.cfg.Password, d.cfg.Database, dbSocketDir, d.cfg.InstanceID)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("failed to open connection to db: %w", err)
    }

    if err = db.Ping(); err != nil {
        return fmt.Errorf("failed to ping db: %w", err)
    }

    d.conn = db
    return nil
}

func (d *Database) Close() {
    if d.conn != nil {
        d.conn.Close()
    }
}

func (d *Database) GetInstance() (*sql.DB, error) {
    if d.conn == nil {
        if err := d.NewConnection(); err != nil {
            return nil, fmt.Errorf("failed to create connection to db: %w", err)
        }
    }

    return d.conn, nil
}

func (d *Database) GetServices(q string, pageSize int, pageNum int) (string, error) {
    // conn, err := d.GetInstance()
    // if err != nil {
    //     return "", fmt.Errorf("failed to get db instance: %w", err)
    // }

    // perform database query and return results
    return fmt.Sprintf("q: %v, pageSize: %v, pageNum: %v", q, pageSize, pageNum), nil
}
