package llrb

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
)

type Int int

func (i Int) Compare(with Key) int {
	if i == with.(Int) {
		return 0
	}
	if i < with.(Int) {
		return -1
	}
	return 1
}

func TestCreateTree(t *testing.T) {
	tree := NewTree()

	if tree.root != nil {
		t.Errorf("tree.root = %T; want nil", tree.root)
	}

	if tree.size != 0 {
		t.Errorf("tree.size = %d; want 0", tree.size)
	}
}

func TestCreateNode(t *testing.T) {
	tree := NewTree()
	tree.insert(Int(1), "1")

	if isRed(tree.root) {
		t.Errorf("tree.root.color = %t; want false", tree.root.color)
	}
}

func TestSearchMissingKey(t *testing.T) {
	tree := NewTree()
	for i := 0; i < 10; i++ {
		tree.insert(Int(i), strconv.Itoa(i))
	}
	value := tree.search(Int(10))
	if value != nil {
		t.Errorf("found %T, want %T", value, nil)
	}
}

func TestInsertionSeries(t *testing.T) {
	tree := NewTree()
	for i := 0; i < 1000; i++ {
		tree.insert(Int(i), strconv.Itoa(i))
	}
}

func TestInsertAndSearchSeries(t *testing.T) {
	tree := NewTree()

	for i := 0; i < 1000; i++ {
		tree.insert(Int(i), strconv.Itoa(i))
	}

	for i := 0; i < 1000; i++ {
		got := tree.search(Int(i))
		got, _ = strconv.Atoi(got.(string))
		if i != got {
			t.Errorf("search(%d) gives %q, want %q", i, got, strconv.Itoa(i))
		}
	}
}

func TestInsertionRandomSeries(t *testing.T) {
	tree := NewTree()

	r := rand.New(rand.NewSource(42))
	var series []int
	for i := 0; i < 100; i++ {
		v := r.Int()
		tree.insert(Int(v), strconv.Itoa(v))
		series = append(series, v)
	}

	existKey := series[len(series)-1]
	tree.insert(Int(existKey), strconv.Itoa(existKey))

	for _, v := range series {
		got := tree.search(Int(v))
		got, _ = strconv.Atoi(got.(string))
		if v != got {
			t.Errorf("search(%d) gives %s, want %s", v, got, strconv.Itoa(v))
		}
	}
}

func TestHeightRandom(t *testing.T) {

	nSimilations := 10
	nSeeds := 10
	nInsertions := 50000

	upperBound := 2 * math.Log2(float64(nInsertions+1))

	var avgHeights []float64
	for n := 0; n < nSimilations; n++ {
		var seeds []int64
		for i := 0; i < nSeeds; i++ {
			seeds = append(seeds, int64(i*20+3))
		}

		var heights []int
		for _, s := range seeds {
			r := rand.New(rand.NewSource(s))

			tree := NewTree()
			for i := 0; i < nInsertions; i++ {
				v := r.Int()
				tree.insert(Int(v), strconv.Itoa(v))
			}
			h := heightOf(tree.root)
			heights = append(heights, h)
		}

		sum := 0
		for _, h := range heights {
			sum += h
		}
		avgHeights = append(avgHeights, float64(sum)/float64(nSeeds))
	}

	sum := 0.0
	for _, h := range avgHeights {
		//fmt.Println(h)
		sum += h
	}
	result := sum / float64(nSimilations)
	fmt.Println("average of average heights: ", result)
	if result > upperBound {
		t.Error("Balancing is wrong")
	}
}
