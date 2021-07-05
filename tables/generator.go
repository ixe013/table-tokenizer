package tables

import (
	"math/rand"
)

func makeRange(max uint32) []uint32 {
    r := make([]uint32, max+1)
    for i := range r {
        r[i] = uint32(i)
    }
    return r
}

func GenerateInMemoryTable(size uint32) *InMemoryTable {
    var table InMemoryTable
    table.entries = makeRange(size-1)
    rand.Shuffle(len(table.entries), func(i, j int) {
        table.entries[i], table.entries[j] = table.entries[j], table.entries[i]
    })

    // We will have to replace Go's internal random with an HSM source
    // See https://stackoverflow.com/a/35208651/591064
    return &table
}

