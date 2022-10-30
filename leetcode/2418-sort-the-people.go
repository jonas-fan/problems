// #string #sort

import (
    "sort"
)

type Person struct {
    name   string
    height int
}

func sortPeople(names []string, heights []int) []string {
    people := make([]*Person, len(names))

    for i, name := range names {
        people[i] = &Person{
            name:   name,
            height: heights[i],
        }
    }

    sort.Slice(people, func(lhs int, rhs int) bool {
        return people[lhs].height > people[rhs].height
    })

    for i, person := range people {
        names[i] = person.name
    }

    return names
}
