package deliverd

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// get key by id
func get_key_by_id(key_id string) (Keys, error) {
	db, err := get_db()
	if err != nil {
		return Keys{}, err
	}
	var key Keys
	err = db.QueryRow("SELECT * FROM deliverd.keys WHERE key_id = ?", key_id).Scan(&key.KeyID, &key.KeyType, &key.KeyOwner, &key.DateCreated, &key.Status)
	return key, err
}

// validate key in context of request
func validate_key(c fiber.Ctx) error {
	key_id := c.Get("Authorization")
	if key_id == "" {
		return errors.New("no key provided in Authorization header")
	}
	key, err := get_key_by_id(key_id)
	if err != nil {
		return errors.New("error: " + err.Error())
	}
	if key.Status != STATUS_ACTIVE {
		return errors.New("key is not active")
	}
	return nil
}

// validate key in context of request with specific type
func validate_key_with_type(c fiber.Ctx, key_type APIKeyType) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return errors.New("no key provided in Authorization header")
	}
	// Extract key_id from "Bearer {id}" format
	key_id := auth[7:] // Skip "Bearer "
	if len(auth) < 7 || auth[:7] != "Bearer " {
		return errors.New("invalid Authorization header format")
	}
	key, err := get_key_by_id(key_id)
	if err != nil {
		return errors.New("error: " + err.Error())
	}
	if key.Status != STATUS_ACTIVE {
		return errors.New("key is not active")
	}
	if key.KeyType != key_type {
		return errors.New("key is not of type " + strconv.Itoa(int(key_type)))
	}
	return nil
}

// create a new key
func (key Keys) new_key() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.keys (key_id, key_type, key_owner, date_created, status) VALUES (?, ?, ?, ?, ?)", key.KeyID, key.KeyType, key.KeyOwner, key.DateCreated, key.Status)
	return err
}

// get all API keys
func get_all_keys() ([]Keys, error) {
	db, err := get_db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM deliverd.keys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var keys []Keys
	var key Keys
	for rows.Next() {
		err = rows.Scan(&key.KeyID, &key.KeyType, &key.KeyOwner, &key.DateCreated, &key.Status)
		if err != nil {
			return keys, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}

// change key status
func change_key_status(key_id int, status APIKeyStatus) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	key, err := get_key_by_id(key_id)
	if err != nil {
		return err
	}
	if key.Status == STATUS_ACTIVE && status != STATUS_ACTIVE || key.Status != STATUS_ACTIVE && status == STATUS_ACTIVE {
		return errors.New("invalid status change")
	}
	_, err = db.Exec("UPDATE deliverd.keys SET status = ? WHERE key_id = ?", status, key.KeyID)
	return err
}

// get all keys belonging to owner
func get_keys_by_owner(owner_id int) ([]Keys, error) {
	db, err := get_db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM deliverd.keys WHERE key_owner = ?", owner_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var keys []Keys
	var key Keys

	for rows.Next() {
		err = rows.Scan(&key.KeyID, &key.KeyType, &key.KeyOwner, &key.DateCreated, &key.Status)
		if err != nil {
			return keys, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}

// change key status by owner id
func change_key_status_by_owner(owner_id int, new_status APIKeyStatus) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.keys SET status = ? WHERE key_owner = ?", new_status, owner_id)
	return err
}
