package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const emptyID = ""

func init() {
	rand.Seed(time.Now().UnixNano())
}

type NotFoundError error

func IsStateNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(NotFoundError)
	return ok
}

func NewDriver(typeName string) *Driver {
	return &Driver{typeName: typeName}
}

type Driver struct {
	typeName string
}

func (d *Driver) Create(value interface{}) (string, error) {
	id := d.generateID()
	statePath := d.stateFilePath(id)

	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return emptyID, fmt.Errorf("marshaling to JSON is failed: %s", err)
	}

	err = ioutil.WriteFile(statePath, data, 0644)
	if err != nil {
		return emptyID, fmt.Errorf("writing to file is failed: %s", err)
	}

	return id, nil
}

func (d *Driver) Read(id string) ([]byte, error) {
	if !d.stateExists(id) {
		return nil, NotFoundError(errors.New("state not found"))
	}
	statePath := d.stateFilePath(id)
	return ioutil.ReadFile(statePath)
}

func (d *Driver) Update(id string, value interface{}) error {
	if !d.stateExists(id) {
		return NotFoundError(errors.New("state not found"))
	}
	statePath := d.stateFilePath(id)

	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling to JSON is failed: %s", err)
	}

	err = ioutil.WriteFile(statePath, data, 0644)
	if err != nil {
		return fmt.Errorf("writing to file is failed: %s", err)
	}

	return nil
}

func (d *Driver) Delete(id string) error {
	if !d.stateExists(id) {
		return NotFoundError(errors.New("state not found"))
	}
	statePath := d.stateFilePath(id)

	if err := os.Remove(statePath); err != nil {
		return fmt.Errorf("removing file is failed: %s", err)
	}
	return nil
}

func (d *Driver) generateID() string {
	return fmt.Sprintf("%d", rand.Int31())
}

func (d *Driver) stateFilePath(id string) string {
	return fmt.Sprintf("%s-%s.json", d.typeName, id)
}

func (d *Driver) stateExists(id string) bool {
	exists := true
	if _, err := os.Stat(d.stateFilePath(id)); os.IsNotExist(err) {
		exists = false
	}
	return exists
}
