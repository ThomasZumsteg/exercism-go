package twobucket

import "fmt"

type bucket struct {
	name     string
	capacity int
	fill     int
}

func (b *bucket) Clone(volume int) bucket { return NewBucket(b.name, b.capacity, volume) }
func (b *bucket) Fill() bucket            { return b.Clone(b.capacity) }
func (b *bucket) Empty() bucket           { return b.Clone(0) }
func (b *bucket) Transfer(c bucket) (bucket, bucket) {
	if b.fill+c.fill > c.capacity {
		return b.Clone(b.fill + c.fill - c.capacity), c.Clone(c.capacity)
	} else {
		return b.Clone(0), c.Clone(c.fill + b.fill)
	}
}

func NewBucket(name string, capacity, fill int) bucket {
	return bucket{name, capacity, fill}
}

type buckets struct {
	a bucket
	b bucket
}

type queueItem struct {
	buckets buckets
	moves   int
}

func NewQueueItem(a, b bucket, moves int) queueItem {
	return queueItem{buckets{a, b}, moves}
}

func Solve(bucket_one, bucket_two, goal int, start_bucket string) (string, int, int, error) {
	if bucket_one <= 0 || bucket_two <= 0 || goal <= 0 {
		return "", -1, -1,
			fmt.Errorf("Inalid size, bucket sizes and goal must be greater than 1")
	}
	queue := []queueItem{}
	seen := make(map[buckets]bool)
	one, two := NewBucket("one", bucket_one, 0), NewBucket("two", bucket_two, 0)
	queue = append(queue, NewQueueItem(one, two, 0))
	switch start_bucket {
	case "one":
		seen[buckets{one, two.Fill()}] = true
		seen[buckets{two.Fill(), one}] = true
	case "two":
		seen[buckets{two, one.Fill()}] = true
		seen[buckets{one.Fill(), two}] = true
	default:
		return "", -1, -1, fmt.Errorf("start bucket must be 'one' or 'two'")
	}
	var item queueItem
	var a, b bucket
	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]
		if ok, resp := seen[item.buckets]; ok && resp {
			continue
		}
		seen[item.buckets] = true
		a, b = item.buckets.a, item.buckets.b
		if item.buckets.a.fill == goal {
			return a.name, item.moves, b.fill, nil
		}
		queue = append([]queueItem{NewQueueItem(b, a, item.moves)}, queue...)
		queue = append(queue, NewQueueItem(a.Fill(), b, item.moves+1))
		queue = append(queue, NewQueueItem(a.Empty(), b, item.moves+1))
		a, b = a.Transfer(b)
		queue = append(queue, NewQueueItem(a, b, item.moves+1))
	}
	return "", -1, -1, fmt.Errorf("No solution")
}
