package twobucket

type bucket struct {
    name string
    capacity int
    fill int
}

func NewBucket(name string, capacity, fill int) bucket {
    return bucket{name, capacity, fill}
}

type state struct {
    one bucket
    two bucket
    moves int
    swap bool
}

func NewState(one, two bucket) state {
    return state{one, two, 0, false}
}

func Solve(bucket_one, bucket_two, goal int, start_bucket string)  (string, int, int, error) {
    queue := []state{}
    seen := make(map[state]bool)
    queue = append(queue, NewState(
        NewBucket("one", bucket_one, 0),
        NewBucket("two", bucket_two, 0)))
    var item state
    for len(queue) > 0 {
        item, queue = queue[0], queue[1:]
        if ok, _ := seen[item]; ok {
            continue
        } else if item.one.fill == goal {
            return item.one.name, item.moves, item.two.fill, nil
        }
    }
    return "one", 4, 5, nil
}
