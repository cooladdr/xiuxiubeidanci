package datasource



type Engine uint32

const (
	// Memory stands for simple memory location;
	// map[int64] datamodels.User ready to use, it's our source in this example.
	Memory Engine = iota
	// Bolt for boltdb source location.
	Bolt
	// MySQL for mysql-compatible source location.
	MySQL
)


func NewDataSource(engine Engine) {
	
}