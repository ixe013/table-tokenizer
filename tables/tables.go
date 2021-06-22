package tables

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

