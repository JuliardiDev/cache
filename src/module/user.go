package src

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	lib "github.com/cache/src/lib"
)

// Person entity
type Person struct {
	ID      int64
	Name    string
	Address string
	HP      string
}

// PersonReq entity request for person
type PersonReq struct {
	Sort Rsort
	Page int
}

// Rsort is entity for struct
type Rsort struct {
	ID      int
	Name    int
	Address int
	HP      int
}

// WriteFile to write Person on file
func WriteFile(p Person) error {
	dat, err := ioutil.ReadFile("files/data")
	if err != nil {
		return nil
	}

	f, err := os.Create("files/data")
	if err != nil {
		return err
	}
	defer f.Close()

	p.ID = rand.Int63()

	pString := fmt.Sprintf("%d:%s:%s:%s\n", p.ID, p.Name, p.Address, p.HP)
	data := fmt.Sprintf("%s%s", string(dat), pString)

	_, err = f.WriteString(data)
	if err == nil {
		go lib.SendToChannel(p, lib.T1)
	}
	return err
}

// ReadFiles to read Person on file
func ReadFiles(r PersonReq) ([]Person, error) {
	var person []Person
	dataRaw, err := ioutil.ReadFile("files/data")
	if err != nil {
		return nil, err
	}
	s := string(dataRaw)
	data := strings.Split(s, "\n")

	for _, dat := range data {
		if dat == "" {
			continue
		}
		splt := strings.Split(dat, ":")
		if len(splt) < 4 {
			for i := 0; i < len(splt); i++ {
				splt = append(splt, "")
			}
		}
		ID, _ := strconv.ParseInt(splt[0], 10, 64)
		person = append(person, Person{
			ID:      ID,
			Name:    splt[1],
			Address: splt[2],
			HP:      splt[3],
		})
	}

	// Sort
	if r.Sort.ID == 1 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].ID < person[j].ID
		})
	} else if r.Sort.ID == 2 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].ID > person[j].ID
		})
	} else if r.Sort.Name == 1 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].Name < person[j].Name
		})
	} else if r.Sort.Name == 2 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].Name > person[j].Name
		})
	} else if r.Sort.Address == 1 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].Address < person[j].Address
		})
	} else if r.Sort.Address == 2 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].Address > person[j].Address
		})
	} else if r.Sort.HP == 1 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].HP < person[j].HP
		})
	} else if r.Sort.HP == 2 {
		sort.Slice(person, func(i, j int) bool {
			return person[i].HP > person[j].HP
		})
	} else {
		sort.Slice(person, func(i, j int) bool {
			return person[i].ID < person[j].ID
		})
	}

	if r.Page != 0 {
		end := r.Page * 20
		if end > len(person) {
			end = len(person)
		}
		start := (r.Page * 20) - 20
		if start > end {
			start = end
		}
		person = person[start:end]
	}

	//Send message to channel
	ps := struct {
		Persons []Person
		Rsort
	}{
		person,
		r.Sort,
	}

	go lib.SendToChannel(ps, lib.T2)

	return person, nil
}
