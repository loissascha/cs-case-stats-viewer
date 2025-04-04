package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ContainerUnlock struct {
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
}

type Item struct {
	PlusMinus string `json:"plusMinus"`
	Name      string `json:"name"`
}

type kv struct {
	key   string
	value int
}

func main() {
	data, err := os.ReadFile("unlocked_container.json")
	if err != nil {
		panic(err)
	}

	var unlocks []ContainerUnlock

	unlocks_count := 0

	json.Unmarshal(data, &unlocks)
	unlocks_count = len(unlocks)

	fmt.Println("You unlocked", unlocks_count, "cases!")

	analyseCaseTypes(&unlocks)
}

func analyseCaseTypes(unlocks *[]ContainerUnlock) {
	result := make(map[string]int)
	for _, unlock := range *unlocks {
		for _, item := range unlock.Items {
			if item.PlusMinus != "-" {
				continue
			}
			itemName := item.Name
			if strings.Contains(itemName, "Key") {
				itemName = strings.Replace(itemName, "Key", "", -1)
			}
			itemName = strings.TrimSpace(itemName)
			v, ok := result[itemName]
			if ok {
				result[itemName] = v + 1
			} else {
				result[itemName] = 1
			}
		}
	}

	var items []kv
	for key, value := range result {
		items = append(items, kv{key: key, value: value})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value < items[j].value
	})

	for _, item := range items {
		fmt.Println("You opened the case", item.key, "for", item.value, "times")
	}
}

// example data
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
