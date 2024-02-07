package log

type Segment struct {
	MaxIndexBytes uint64
	MaxStoreBytes uint64
	InitialOffset uint64
}
type Config struct {
	Segment Segment
}
