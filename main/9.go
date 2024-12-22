package main

import "fmt"

func ninthDay() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	// input := readFileAsBytes("input9example.txt")[0]
	input := readFileAsBytes("input9.txt")[0]
	filesystem := make([]int, 0)
	for _, b := range input {
		filesystem = append(filesystem, int(b-'0'))
	}
	fmt.Printf("filesystem: %v\n", filesystem)

	res := 0
	pos := filesystem[0] // jump to end of file with id 0

	// disk := writeFile(make([]int, 0), pos, 0)

	j := len(filesystem) - 1 // assuming last element is a file
	if j%2 != 0 {
		panic("last element was not a file!")
	}

	for i := 1; i < j; i += 2 {
		freeLen := filesystem[i]
		for freeLen > 0 {
			consumeLen := min(freeLen, filesystem[j])

			// move file (skip)
			// disk = writeFile(disk, consumeLen, j/2)
			// calc checkSum
			res += calcCheckSum(pos, consumeLen, j/2)

			pos += consumeLen
			freeLen -= consumeLen
			filesystem[j] -= consumeLen
			if filesystem[j] == 0 {
				j -= 2 // take next to last file
			}

		}
		// calc checkSum of existing file
		res += calcCheckSum(pos, filesystem[i+1], (i+1)/2)
		pos += filesystem[i+1]

		//disk = writeFile(disk, filesystem[i+1], (i+1)/2)
	}

	fmt.Printf("res: %v\n", res)
	//fmt.Printf("disk: %v\n", disk)
}

func writeFile(disk []int, l, fileId int) []int {
	for i := 0; i < l; i++ {
		disk = append(disk, fileId)
	}
	return disk
}
func calcCheckSum(s, l, fileId int) int {
	return fileId * l * (2*s + l - 1) / 2
}
