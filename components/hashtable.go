package components

type Table struct {
  lists []LinkedList
}

func (table *Table) Get(key string) (*ListNode, error) {
  list := table.lists[table.hash(key)]
  return list.Get(key)
}

func (table *Table) Add(key string, value string) {
  table.lists[table.hash(key)].Add(key, value)
}

func (table *Table) Remove(key string) {
  table.lists[table.hash(key)].Remove(key)
}

func CreateTable(size int) *Table {
  return &Table{
    lists: make([]LinkedList, size),
  }
}

func (table *Table) Print() {
  for _, list := range table.lists {
    list.Print()
  }
}

func (table *Table) hash(key string) int {
  hash := 0
  for _, char := range key {
    hash += int(char)
  }
  return (hash * 3) % len(table.lists)
}
