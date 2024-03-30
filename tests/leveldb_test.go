package tests

import (
	"github.com/OpenDevNexus/GoLeveldb"
	"testing"
)

// TestLevelDBPutAndGet tests the Put and Get functions of the LevelDB package.
func TestLevelDBPutAndGet(t *testing.T) {
	// Initialize test database name.
	dbName := "testdb"

	// Open the database.
	options := GoLevelDb.NewOptions()
	options.SetCreateIfMissing(true)
	db, err := GoLevelDb.Open(dbName, options)
	if err != nil {
		t.Errorf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	// Put a key-value pair into the database.
	writeOptions := GoLevelDb.NewWriteOptions()
	err = db.Put(writeOptions, []byte("key1"), []byte("value1"))
	if err != nil {
		t.Errorf("Error putting key-value pair: %v", err)
		return
	}

	// Retrieve the value for the key from the database.
	readOptions := GoLevelDb.NewReadOptions()
	value, err := db.Get(readOptions, []byte("key1"))
	if err != nil {
		t.Errorf("Error getting value for key: %v", err)
		return
	}

	// Verify the retrieved value matches the expected value.
	expectedValue := []byte("value1")
	if string(value) != string(expectedValue) {
		t.Errorf("Got unexpected value. Expected: %s, Got: %s", expectedValue, value)
	}
}

// TestLevelDBDelete tests the Delete function of the LevelDB package.
func TestLevelDBDelete(t *testing.T) {
	// Initialize test database name.
	dbName := "testdb"

	// Open the database.
	options := GoLevelDb.NewOptions()
	options.SetCreateIfMissing(true)
	db, err := GoLevelDb.Open(dbName, options)
	if err != nil {
		t.Errorf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	// Put a key-value pair into the database.
	writeOptions := GoLevelDb.NewWriteOptions()
	err = db.Put(writeOptions, []byte("key2"), []byte("value2"))
	if err != nil {
		t.Errorf("Error putting key-value pair: %v", err)
		return
	}

	// Delete the key-value pair from the database.
	err = db.Delete(writeOptions, []byte("key2"))
	if err != nil {
		t.Errorf("Error deleting key-value pair: %v", err)
		return
	}
}
