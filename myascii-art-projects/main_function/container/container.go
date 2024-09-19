package container

var (
	Format        [][]string
	Format1d      []string
	Element_width []int
	Max           int64
	Result        string
)

var Hash_table = make([]byte, 2e9) // It is the hash table for each element of current banner elements
