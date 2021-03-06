package main

//*
import (
    tables "tables/tables"
    "strings"
)
//*/

func tokenize(sensitive string, table tables.Table) string {
    left := tables.TableEntry(2)
    right := tables.TableEntry(0)

    right += table.Lookup(left)

    return strings.ToUpper(sensitive)
}

func main() {
    forward := tables.GenerateInMemoryTable(10)

    tokenize("asdf", forward)

    return
}
