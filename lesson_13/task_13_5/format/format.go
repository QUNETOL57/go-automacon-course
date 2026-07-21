package format

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"sort"
)

type patient struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Email string `xml:"Email"`
}
type patients struct {
	List []patient `xml:"Patient"`
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

	e := xml.NewEncoder(f)
	e.Indent("", " ")
	_, err = f.WriteString(xml.Header)
	if err != nil {
		return err
	}
	if err = e.Encode(list); err != nil {
		return err
	}
	return nil
}

func (p patients) sort() patients {
	sort.Slice(p.List, func(i, j int) bool {
		return p.List[i].Age < p.List[j].Age
	})
	return p
}

func Do(pathFrom string, pathTo string) error {
	list, err := get(pathFrom)
	if err != nil {
		return err
	}
	err = set(pathTo, list.sort())
	if err != nil {
		return err
	}
	return nil
}
