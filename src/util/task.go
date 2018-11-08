package src

import (
	"encoding/json"
	"fmt"
	"strconv"

	lib "github.com/cache/src/lib"
	mod "github.com/cache/src/module"
)

// RegisterTask is to register your task
func RegisterTask() []lib.Task {
	c := lib.Cache{}

	t1 := lib.NewTask(lib.T1, func(msg []byte) {
		var person mod.Person
		err := json.Unmarshal(msg, &person)
		if err != nil {
			fmt.Println(err)
		}
		key := fmt.Sprintf("key:%d", person.ID)

		c.Hmset(key, map[string]string{
			"id":      strconv.FormatInt(person.ID, 10),
			"name":    person.Name,
			"address": person.Address,
			"hp":      person.HP,
		})
	})

	t2 := lib.NewTask(lib.T2, func(msg []byte) {
		ps := struct {
			Persons []mod.Person
			mod.Rsort
		}{}
		err := json.Unmarshal(msg, &ps)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(ps.Rsort)

	})

	return []lib.Task{t1, t2}
}
