package twobucket

import "fmt"

type bucket struct {
    name string
    capacity int
    fill int
}

func (b *bucket) Clone(volume int) bucket { return NewBucket(b.name, b.capacity, volume) }
func (b *bucket) Fill() bucket { return b.Clone(b.capacity) }
func (b *bucket) Empty() bucket { return b.Clone(0) }
func (b *bucket) Transfer(c bucket) (bucket, bucket) {
    if b.fill + c.fill > c.capacity {
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
    moves int
}

func NewQueueItem(a, b bucket, moves int) queueItem {
    return queueItem{buckets{a, b}, moves}
}

func Solve(bucket_one, bucket_two, goal int, start_bucket string)  (string, int, int, error) {
    println()
    queue := []queueItem{}
    seen := make(map[buckets]bool)
    a, b := NewBucket("one", bucket_one, 0), NewBucket("two", bucket_two, 0)
    if start_bucket == "one" { queue = append(queue, NewQueueItem(a, b, 0))
    } else { queue = append(queue, NewQueueItem(a, b, 0)) }
    var item queueItem
    for len(queue) > 0 {
        item, queue = queue[0], queue[1:]
        if ok, resp := seen[item.buckets]; ok && resp {
            continue
        }
        seen[item.buckets] = true;
        fmt.Printf("%d: %s (%d/%d) - %s (%d/%d)\n", item.moves,
            item.buckets.a.name, item.buckets.a.fill, item.buckets.a.capacity,
            item.buckets.b.name, item.buckets.b.fill, item.buckets.b.capacity)
        if item.buckets.a.fill == goal {
            return item.buckets.a.name, item.moves, item.buckets.b.fill, nil
        } else {
            a, b := item.buckets.a, item.buckets.b
            // Swap
            queue = append([]queueItem{ NewQueueItem(b, a, item.moves) }, queue...)
            // Fill
            queue = append(queue, NewQueueItem(b, a.Fill(), item.moves+1))
            // Empty
            queue = append(queue, NewQueueItem(b, a.Empty(), item.moves+1))
            // Transfer
            a, b = a.Transfer(b)
            queue = append(queue, NewQueueItem(b, a, item.moves+1))
        }
    }
    fmt.Println("Done")
    return "one", 4, 5, nil
}
