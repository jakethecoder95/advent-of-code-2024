package day9

import (
	"advent2024/util"
	"os"
	"slices"
	"strconv"
)

type StorageType int
const (
    File = iota
    Space
)

type StorageAllocation struct {
    storageType StorageType
    blocks      int
    id          int
}

func Part1() int {

    path := os.Args[1]
    file := util.ReadFile(path)

    id := 0
    blocks := []int{}
    emptyBlocks := 0

    for i, v := range file {
        blockCount, _ := strconv.Atoi(string(v))
        for j := 1; j <= blockCount; j++ {
            if i%2 == 0 {
                blocks = append(blocks, id)
            } else {
                blocks = append(blocks, -1)
                emptyBlocks++
            }
        }
        if i%2 == 0 {
            id++
        }
    }

    emptyBlocksSeen := 0
    for i := len(blocks)-1; emptyBlocksSeen < emptyBlocks; i-- {
        if blocks[i] == -1 {
            emptyBlocksSeen++
            continue
        }
        for j := 0; j < len(blocks); j++ {
            if blocks[j] == -1 {
                blocks[j] = blocks[i]
                blocks[i] = -1
                emptyBlocksSeen++
                break
            }
        }
    }

    total := 0
    for i, block := range blocks {
        if block == -1 {
            break
        }
        total += i * block
    }

    return total
}

func Part2() int {

    path := os.Args[1]
    file := util.ReadFile(path)

    id := 0
    storage := []StorageAllocation{}

    for i := 0; i < len(file); i += 2 {
        // File
        fBlocks, _ := strconv.Atoi(string(file[i]))
        fileStorage := StorageAllocation{
            storageType: File,
            blocks: fBlocks,
            id: id,
        }
        storage = append(storage, fileStorage)
        // Storage
        if i < len(file)-1 {
            sBlocks, _ := strconv.Atoi(string(file[i+1]))
            space := StorageAllocation{
                storageType: Space,
                blocks: sBlocks,
                id: -1,
            }
            storage = append(storage, space)
        }
        id++
    }

    fileCount := (len(storage)/2)+1
    filesSeen := 0
    for i := len(storage)-1; filesSeen < fileCount; i-- {
        allocation := storage[i]
        if allocation.storageType == File {
            filesSeen++
        } else {
            continue
        }
        for j := 0; j < len(storage) && j < i; j++ {
            currAlloc := storage[j]
            if currAlloc.storageType == File || currAlloc.blocks < allocation.blocks {
                continue
            }
            storage[i] = StorageAllocation{
                storageType: Space,
                blocks: allocation.blocks,
                id: -1,
            }
            if allocation.blocks == currAlloc.blocks {
                storage[j] = allocation
            } else {
                newSpace := StorageAllocation{
                    storageType: Space,
                    blocks: currAlloc.blocks - allocation.blocks,
                    id: -1,
                }
                storage[j] = allocation
                storage = slices.Insert(storage, j+1, newSpace)
                i++
            }
            break
        }
    }

    total := 0
    index := 0
    for i := 0; i < len(storage); {
        allocation := storage[i]
        if allocation.storageType == File {
            for j := 0; j < allocation.blocks; j++ {
                total += index * allocation.id
                index++
            }
        } else {
            index += allocation.blocks
        }
        i++
    }

    return total
}
