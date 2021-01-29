package main

import "fmt"

func main() {
	personMap := map[string][]string{
		"A_HOJO": {"Programming", "IT", "Magic", "Fantasy", "Lit RPG", "video games"},
		"Jedi":   {"Light Sabers", "cool force moves", "the force"},
		"Kairi":  {"Squishies", "surprise videos", "mommy", "daddy"},
	}

	personMap["Linsday"] = []string{"Kettles", "Amazon", "Youtube"}

	for k, v := range personMap {
		fmt.Println(k)
		for i, item := range v {
			fmt.Printf("\tPosition: %d\t Likes: %s\n", i, item)
		}

	}
}
