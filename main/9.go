package main

import "fmt"

func ninthDay() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	// input := readFileAsBytes("input9example.txt")[0]
	input := readFileAsBytes("input/input9.txt")[0]
	filesystem := make([]int, 0)
	for _, b := range input {
		filesystem = append(filesystem, int(b-'0'))
	}
	fmt.Printf("filesystem: %v\n", filesystem)

	checksum := moveFileBlocks(filesystem)
	fmt.Printf("res: %v\n", checksum)

	filesystem = make([]int, 0)
	for _, b := range input {
		filesystem = append(filesystem, int(b-'0'))
	}

	checksum = moveWholeFiles(filesystem)
	fmt.Printf("move whole files: %v\n", checksum)

}

func moveWholeFiles(filesystem []int) (checksum int) {
	checksum = 0

	filesystemCopy := make([]int, len(filesystem))
	copy(filesystemCopy, filesystem)

	disk := make([]int, 50)

	j := len(filesystem) - 1 // assuming last element is a file
	if j%2 != 0 {
		panic("last element was not a file!")
	}

	for ; j > 0; j -= 2 {
		pos := 0
		for i := 1; i < j; i += 2 {
			pos += filesystemCopy[i-1] // jump over the file
			if filesystem[i] >= filesystem[j] {
				pos += filesystemCopy[i] - filesystem[i]
				// move file into free space
				filesystem[i] -= filesystem[j]

				checksum += calcCheckSum(pos, filesystem[j], j/2)
				writeBlocks(disk, pos, filesystem[j], j/2)
				//fmt.Printf("disk: %v\n", disk)

				filesystem[j] = 0 // do not count this file in checksum on step 2

				break
			}
			pos += filesystemCopy[i] // jump over free space
		}
	}

	checksum += calcCheckSum(0, filesystem[0], 0)
	writeBlocks(disk, 0, filesystem[j], -1)
	pos := filesystem[0]
	// step 2: count checksum for files that were not moved
	for j = 2; j < len(filesystem); j += 2 {
		pos += filesystemCopy[j-1]

		checksum += calcCheckSum(pos, filesystem[j], j/2)
		writeBlocks(disk, pos, filesystem[j], j/2)
		pos += filesystemCopy[j]
	}

	fmt.Printf("disk: %v\n", disk)

	return checksum
}

func writeBlocks(disk []int, s, l int, fileId int) {
	for i := s; i < s+l && i < len(disk); i++ {
		if disk[i] != 0 {
			fmt.Println("Warning!", fileId, disk[i])
		}
		disk[i] = fileId
	}
}

func moveFileBlocks(filesystem []int) (checksum int) {

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

	//fmt.Printf("disk: %v\n", disk)

	return res

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
