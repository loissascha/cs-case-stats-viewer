package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContainerUnlock struct {
	Date string `json:"date"`
}

func main() {
	data, err := os.ReadFile("unlocked_container.json")
	if err != nil {
		panic(err)
	}

	var unlocks []ContainerUnlock

	json.Unmarshal(data, &unlocks)
	fmt.Println("You unlocked", len(unlocks), "cases!")
}

// [
//   {
//     "date": "3 Apr, 2025",
//     "time": "9:01pm",
//     "description": "Unlocked a container",
//     "items": [
//       {
//         "plusMinus": "+",
//         "name": "StatTrak™ USP-S | PC-GRN",
//         "image": "https://community.fastly.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpoo6m1FBRp3_bGcjhQ08mlhJO0leXhJ77XmXlS7fp2mOzE-7P5gVO8v109a2n0ItSRJFc7MArX_AO2w-3thJW7tJqdwXNgvClx53bamhe2hB0dcKUx0gbwPWYB/120x40"
//       },
//       {
//         "plusMinus": "-",
//         "name": "Fever Case Key",
//         "image": "https://community.fastly.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXX7gNTPcUxuxpJSXPbQv2S1MDeXkh6LBBOienwZV9mgPaaJGhGvIXjldaIw6DwMOyDwmkAsMMoi73Crd_z2QXjqURsYWHtZNjCri8QUkI/120x40"
//       }
//     ]
//   },
