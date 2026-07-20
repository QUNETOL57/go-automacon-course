package format

import (
	"encoding/json"
	"io"
	"os"
)

type patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
type patients struct {
	List []patient
}

func get(pathFrom string) (patients, error) {
	f, err := os.Open(pathFrom)
	if err != nil {
		return patients{}, err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	list := make([]patient, 0, 3)
	for {
		var p patient
		err = d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return patients{}, err
		}
		list = append(list, p)
	}
	return patients{List: list}, nil
}

func set(pathTo string, list patients) error {
	f, err := os.Create(pathTo)
	if err != nil {
		return err
	}
	defer f.Close()
	e := json.NewEncoder(f)
	if err = e.Encode(list.List); err != nil {
		return err
	}
	return nil
}

func Do(pathFrom string, pathTo string) error {
	list, err := get(pathFrom)
	if err != nil {
		return err
	}
	err = set(pathTo, list)
	if err != nil {
		return err
	}
	return nil
}
