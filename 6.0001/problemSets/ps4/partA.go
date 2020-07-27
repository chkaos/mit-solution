package main

import "fmt"

/*
Return all permutations of a string
Example,
Find all perm of abc:
We will find all perms of the given string based on all perms of its substring
To find all perms of abc, we need to find all perms of bc, and then add a to those perms
To find all perms of bc, we need to find all perms of c, and add b to those perms
To find all perms of c, well we know that is only c
Now we can find all perms of bc, insert c at every available space in b
bc --> bc, cb
Now we need to add a to all perms of bc
abc, bac, bca (insert a in bc) -- acb, cab, cba (insert a in cb)
*/

func main() {
	fmt.Printf("All perms are - %v", getPerms("abc"))
}

func getPerms(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}

	current := str[0:1] // current char
	remStr := str[1:]   // remaining string

	perms := getPerms(remStr) // get perms for remaining string

	allPerms := make([]string, 0) // array to hold all perms of the string based on perms of substring

	// for every perm in the perms of substring
	for _, perm := range perms {
		// add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}

	return allPerms
}

// Insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}

//todo 可拓展文章
