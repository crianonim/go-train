package main

import "fmt"

// Entity is general tyoe.
type Entity struct {
	name  string
	place *Place
	money int
}

// Place you can be
type Place struct {
	name     string
	kind     string
	entities []*Entity
	places   []*Place
}

func entityString(e *Entity) string {
	return fmt.Sprintf("Entity %s is in %s and has %v money", e.name, e.place.name, e.money)
}

func placeString(place *Place) string {
	s := fmt.Sprintf("Place %v has following entities: \n", place.name)
	for i, k := range place.entities {
		s = s + fmt.Sprintf(" %v) %v \n", i, entityString(k))
	}
	return s
}
func (place *Place) put(e *Entity) bool {
	place.entities = append(place.entities, e)
	e.place = place
	fmt.Println(entityString(e))
	return true
}
