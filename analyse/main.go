package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
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
	unlocks := readUserData("unlocked_container.json")
	skins := readSkinsJson()
	unlocks_count := 0
	unlocks_count = len(unlocks)

	fmt.Println("You unlocked", unlocks_count, "cases!")

	analyseCaseTypes(&unlocks)
	analyseSkinRarities(&unlocks, &skins)

	// debug: print out all the possible rarities
	rarities := []string{}
	for _, skin := range skins {
		if slices.Contains(rarities, skin.Rarity.Name) {
			continue
		}
		rarities = append(rarities, skin.Rarity.Name)
	}
	fmt.Println(rarities)
	////

}

func analyseSkinRarities(unlocks *[]ContainerUnlock, skins *[]Skin) {
	gotRarities := make(map[string]int)
	for _, unlock := range *unlocks {
		for _, item := range unlock.Items {
			if item.PlusMinus != "+" {
				continue
			}

			// this is what the user got out of it!
			r := getSkinRarity(item.Name, skins)
			ra, ok := gotRarities[r]
			if ok {
				gotRarities[r] = ra + 1
			} else {
				gotRarities[r] = 1
			}
		}
	}
	fmt.Println(gotRarities)

}

func getSkinRarity(name string, skins *[]Skin) string {
	if strings.Contains(name, "StatTrak™") {
		name = strings.Replace(name, "StatTrak™", "", -1)
		name = strings.TrimSpace(name)
	}
	if strings.Contains(name, "Souvenir") {
		name = strings.Replace(name, "Souvenir", "", -1)
		name = strings.TrimSpace(name)
	}
	if strings.Contains(name, "Sticker") {
		return "Sticker"
	}
	if strings.Contains(name, "Patch") {
		return "Patch"
	}
	if strings.Contains(name, "Graffiti") {
		return "Graffiti"
	}
	for _, skin := range *skins {
		if skin.Name == name {
			return skin.Rarity.Name
		}
	}
	return "Unknown"
}

func readUserData(filepath string) []ContainerUnlock {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var unlocks []ContainerUnlock
	json.Unmarshal(data, &unlocks)
	cleanedUnlocks := []ContainerUnlock{}
	for _, v := range unlocks {
		skippedUnlocke := false
		for _, i := range v.Items {
			if i.PlusMinus != "+" {
				continue
			}
			if strings.Contains(i.Name, "Sticker") {
				skippedUnlocke = true
			}
			if strings.Contains(i.Name, "Graffiti") {
				skippedUnlocke = true
			}
			if strings.Contains(i.Name, "Patch") {
				skippedUnlocke = true
			}
		}
		if skippedUnlocke {
			continue
		}
		cleanedUnlocks = append(cleanedUnlocks, v)
	}
	return cleanedUnlocks
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
		fmt.Println("You opened", item.key, item.value, "times")
	}
}

type Skin struct {
	Name   string     `json:"name"`
	Rarity SkinRarity `json:"rarity"`
}

type SkinRarity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func readSkinsJson() []Skin {
	data, err := os.ReadFile("skins.json")
	if err != nil {
		panic(err)
	}

	var skins []Skin
	json.Unmarshal(data, &skins)
	return skins
}

// example data skins.`json:""`
// [
//   {
//     "id": "skin-e757fd7191f9",
//     "name": "★ Hand Wraps | Spruce DDPAT",
//     "description": "Preferred by hand-to-hand fighters, these wraps protect the knuckles and stabilize the wrist when punching. The outer wrap is fabric screen printed with grey digital camo.\\n\\n<i>Some people say they're tough... others show it</i>",
//     "weapon": {
//       "id": "leather_handwraps",
//       "weapon_id": 5032,
//       "name": "Hand Wraps"
//     },
//     "category": {
//       "id": "sfui_invpanel_filter_gloves",
//       "name": "Gloves"
//     },
//     "pattern": {
//       "id": "handwrap_camo_grey",
//       "name": "Spruce DDPAT"
//     },
//     "min_float": 0.06,
//     "max_float": 0.8,
//     "rarity": {
//       "id": "rarity_ancient",
//       "name": "Extraordinary",
//       "color": "#eb4b4b"
//     },
//     "stattrak": false,
//     "souvenir": false,
//     "paint_index": "10010",
//     "wears": [
//       {
//         "id": "SFUI_InvTooltip_Wear_Amount_0",
//         "name": "Factory New"
//       },
//       {
//         "id": "SFUI_InvTooltip_Wear_Amount_1",
//         "name": "Minimal Wear"
//       },
//       {
//         "id": "SFUI_InvTooltip_Wear_Amount_2",
//         "name": "Field-Tested"
//       },
//       {
//         "id": "SFUI_InvTooltip_Wear_Amount_3",
//         "name": "Well-Worn"
//       },
//       {
//         "id": "SFUI_InvTooltip_Wear_Amount_4",
//         "name": "Battle-Scarred"
//       }
//     ],
//     "collections": [],
//     "crates": [
//       {
//         "id": "crate-4288",
//         "name": "Glove Case",
//         "image": "https://raw.githubusercontent.com/ByMykel/counter-strike-image-tracker/main/static/panorama/images/econ/weapon_cases/crate_community_15_png.png"
//       },
//       {
//         "id": "crate-4352",
//         "name": "Operation Hydra Case",
//         "image": "https://raw.githubusercontent.com/ByMykel/counter-strike-image-tracker/main/static/panorama/images/econ/weapon_cases/crate_community_17_png.png"
//       }
//     ],
//     "team": {
//       "id": "both",
//       "name": "Both Teams"
//     },
//     "legacy_model": false,
//     "image": "https://raw.githubusercontent.com/ByMykel/counter-strike-image-tracker/main/static/panorama/images/econ/default_generated/leather_handwraps_handwrap_camo_grey_light_png.png"
//   },

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
