package floyd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_WEIGHT = math.MaxFloat64

type Floyd struct {
	Count   int
	Weights [][]float64
	Dists   [][]float64
}

func NewFloyd(count int) *Floyd {
	weights := make([][]float64, count+1)
	dists := make([][]float64, count+1)
	for i := 0; i <= count; i++ {
		weights[i] = make([]float64, count+1)
		dists[i] = make([]float64, count+1)
		for j := 0; j <= count; j++ {
			if i > 0 && j > 0 && i == j {
				weights[i][j] = 0
			} else {
				weights[i][j] = MAX_WEIGHT
			}
			dists[i][j] = weights[i][j]
		}
	}
	f := &Floyd{
		Count:   count,
		Weights: weights,
		Dists:   weights,
	}
	return f
}

func (f *Floyd) AddEdge(from, to int, weight float64) error {
	if from < 1 || to < 1 || f.Count < from || f.Count < to {
		return fmt.Errorf("out of bound error")
	}
	f.Weights[from][to] = weight
	return nil
}

func (f *Floyd) Calculate() error {
	for k := 1; k <= f.Count; k++ {
		for i := 1; i <= f.Count; i++ {
			for j := 1; j <= f.Count; j++ {
				f.Dists[i][j] = math.Min(f.Dists[i][j], f.Dists[i][k]+f.Dists[k][j])
			}
		}
	}
	return nil
}

func FloydAlgorithm(in *os.File) {
	bin := bufio.NewReader(in)
	var l []byte

	l, _, _ = bin.ReadLine()
	nodeCount, _ := strconv.Atoi(strings.TrimSpace(string(l)))

	floyd := NewFloyd(nodeCount)

	l, _, _ = bin.ReadLine()
	edgeCount, _ := strconv.Atoi(strings.TrimSpace(string(l)))

	for i := 0; i < edgeCount; i++ {
		l, _, _ = bin.ReadLine()
		infos := strings.SplitN(strings.TrimSpace(string(l)), " ", 3)
		from, _ := strconv.Atoi(infos[0])
		to, _ := strconv.Atoi(infos[1])
		weight, _ := strconv.ParseFloat(infos[2], 64)
		floyd.AddEdge(from, to, weight)
	}
	floyd.Calculate()

	l, _, _ = bin.ReadLine()
	pathCount, _ := strconv.Atoi(strings.TrimSpace(string(l)))

	for i := 0; i < pathCount; i++ {
		l, _, _ = bin.ReadLine()
		dirs := strings.SplitN(strings.TrimSpace(string(l)), " ", 2)
		from, _ := strconv.Atoi(dirs[0])
		to, _ := strconv.Atoi(dirs[1])
		fmt.Println(floyd.Dists[from][to])
	}
}
