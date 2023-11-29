package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func permute(input string) []string {
	var result []string
	chars := []rune(input)
	n := len(chars)

	swap := func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	}

	var generatePermutations func(int)
	generatePermutations = func(index int) {
		if index == n {
			result = append(result, string(chars))
			return
		}
		used := make(map[rune]bool)
		for i := index; i < n; i++ {
			if used[chars[i]] {
				continue
			}
			used[chars[i]] = true
			swap(index, i)
			generatePermutations(index + 1)
			swap(index, i)
		}
	}

	generatePermutations(0)
	return result
}

func findOdd(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}

	return 0
}

func countOccurrences(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}

func countSmileys(arr []string) int {
	pattern := `[:;][-~][)D]`

	re := regexp.MustCompile(pattern)

	count := 0

	for _, face := range arr {
		if re.MatchString(face) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("Select an option:")
	fmt.Println("2. Permutations")
	fmt.Println("3. Find the Odd Integer")
	fmt.Println("4. Count Smiley Faces")

	var choice int
	fmt.Print("Enter your choice (2/3/4): ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading choice:", err)
		return
	}

	switch choice {
	case 2:
		fmt.Print("Enter a string to permute: ")
		var input string
		fmt.Scan(&input)
		permutations := permute(input)
		sort.Strings(permutations)
		fmt.Println(permutations)
	case 3:
		fmt.Print("Enter a list of integers separated by spaces: ")
		var input string
		fmt.Scan(&input)
		re := regexp.MustCompile(`\d+`)
		numStrings := re.FindAllString(input, -1)
		var arr []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Error parsing integer: %v\n", err)
				return
			}
			arr = append(arr, num)
		}
		result := findOdd(arr)
		fmt.Printf("%v should return %d, because it occurs %d time(s) (which is odd).\n", arr, result, countOccurrences(arr, result))
	case 4:
		smileyFaces1 := []string{":)", ";(", ";}", ":-D"}
		smileyFaces2 := []string{";D", ":-(", ":-)", ";~)"}
		smileyFaces3 := []string{";]", ":[", ";*", ":$", ";-D"}
		result1 := countSmileys(smileyFaces1)
		result2 := countSmileys(smileyFaces2)
		result3 := countSmileys(smileyFaces3)
		fmt.Printf("Test case 1 result: %d\n", result1)
		fmt.Printf("Test case 2 result: %d\n", result2)
		fmt.Printf("Test case 3 result: %d\n", result3)
	default:
		fmt.Println("Invalid choice. Please enter 2, 3, or 4.")
	}
}
