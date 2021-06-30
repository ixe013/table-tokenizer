package tables

/*
import (
    "bufio"
    "os"
)
//*/

type TableEntry uint32
type InMemoryTable struct {
    entries []uint32
}

type SQLiteTable struct {
    decoy uint32
}

type Table interface {
    Lookup(TableEntry) TableEntry
}

func (t InMemoryTable) Lookup(e TableEntry) TableEntry {
    return TableEntry(t.entries[e])
}

func (t SQLiteTable) Lookup(e TableEntry) TableEntry {
    return TableEntry(t.decoy)
}

/*
func LoadTableFromFile(filename string) *Table {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
//*/
