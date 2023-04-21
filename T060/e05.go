package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	// the sorting is done in sequence first by age and last by last name
	// also afther this the sorting is applied on saying for each user (also by last remains as well)

	users := []user{u1, u2, u3}

	fmt.Println("")
	fmt.Println(users)
	fmt.Println("")
	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println("")
	sort.Sort(ByLast(users))
	fmt.Println(users)
	fmt.Println("Before sort sayings")
	fmt.Println("")
	for _, v := range users {
		fmt.Printf("%v %v %v\n\n", v.First, v.Last, v.Age)
		for _, e := range v.Sayings {
			fmt.Printf("\t%v\n", e)
		}
		fmt.Println("")
	}
	fmt.Println("After sort sayings")
	fmt.Println("")
	sort.Strings(u1.Sayings)
	sort.Strings(u2.Sayings)
	sort.Strings(u3.Sayings)
	for _, v := range users {
		fmt.Printf("%v %v %v\n\n", v.First, v.Last, v.Age)
		for _, e := range v.Sayings {
			fmt.Printf("\t%v\n", e)
		}
		fmt.Println("")
	}
}
