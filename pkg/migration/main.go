package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	method := flag.String("method", "default", "migration method can be up or down")
	seed := flag.String("seed", "0", "seeding data, 0 for all seed")
	flag.Parse()
	// Initialize Config

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open database: %w", err))
	}

	seedVersion, _ := strconv.Atoi(*seed)
	if *method == "seed" {
		runSeed(db, seedVersion)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open database: %w", err))
	}

	// Initialize Migration
	m, err := migrate.NewWithDatabaseInstance(
		"file://query",
		os.Getenv("DB_NAME"), driver)
	if err != nil {
		log.Fatal(err)
	}

	// Check version
	version, dirty, err := m.Version()
	if err != nil {
		version = 0
		dirty = false
	}

	if dirty {
		m.Force(int(version))
	}

	var methodName string
	switch *method {
	case "up":
		if err = m.Up(); err != nil {
			m.Force(int(version))
			log.Fatal(err)
		}
		methodName = "upgrade"
	case "down":
		if err = m.Steps(-1); err != nil {
			m.Force(int(version))
			log.Fatal(err)
		}
		methodName = "down"
	case "reset":
		if err = m.Down(); err != nil {
			m.Force(int(version))
			log.Fatal(err)
		}
		methodName = "reset"
	default:
		log.Fatal("Complete with method is empty (expected?)")
		return
	}
	newVersion, _, err := m.Version()
	if err != nil {
		newVersion = 0
	}
	log.Printf("Successful database %s from version %d to %d", methodName, version, newVersion)
}

func runSeed(dbCli *sql.DB, seed int) {
	seedFile := make(map[string]string)
	// Get All files on folder seed
	root := "./pkg/migration/seed" // Current directory, can be changed to your folder path
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Open the file
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("Error opening file:", path, err)
				return nil // Continue to next file
			}
			defer file.Close() // Close when done

			// Process the file here
			name := filepath.Base(file.Name())
			index := strings.Index(name, "_")
			if _, find := seedFile[name[:index]]; find {
				return errors.New("file already exist on index " + name[:index])
			}
			seedFile[name[:index]] = path
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}

	// For each seedfile open the sql and execute using db
	for index := range seedFile {
		content, _ := os.ReadFile(seedFile[index])
		_, err = dbCli.Exec(string(content))
		if err != nil {
			fmt.Println("Error executing file:", seedFile[index], err)
			break
		}
		fmt.Println("Successfully executed file:", seedFile[index])
	}
}
