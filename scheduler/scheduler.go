package scheduler

import (
	"cube/node"
	"cube/task"
	"log"
	"math"
	"time"
)

const (
	// LIEB square ice constant
	// https://en.wikipedia.org/wiki/Lieb%27s_square_ice_constant
	LIEB = 1.53960071783900203869
)

// 1. Determine a set of candidate workers on which a task could run
// 2. Score the candidate workers from best to worst
// 3. Pick the worker with the best score
type Scheduler interface {
	SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node
	Score(t task.Task, nodes []*node.Node) map[string]float64
	Pick(scores map[string]float64, candidates []*node.Node) *node.Node
}

type RoundRobin struct {
	Name       string
	LastWorker int
}

func (r *RoundRobin) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	return nodes
}

func (r *RoundRobin) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	nodeScores := make(map[string]float64)

	var newWorker int
	if r.LastWorker+1 < len(nodes) {
		newWorker = r.LastWorker + 1
		r.LastWorker++
	} else {
		newWorker = 0
		r.LastWorker = 0
	}

	for idx, node := range nodes {
		if idx == newWorker {
			nodeScores[node.Name] = 0.1
		} else {
			nodeScores[node.Name] = 1.0
		}
	}

	return nodeScores
}

func (r *RoundRobin) Pick(scores map[string]float64, candidates []*node.Node) *node.Node {
	var bestNode *node.Node
	var lowestScore float64
	// If there is a node with idx = 0 among the candidates,
	// then it must be the best node with score = 0.1
	for idx, node := range candidates {
		if idx == 0 {
			bestNode = node
			lowestScore = scores[node.Name]
			continue
		}

		if scores[node.Name] < lowestScore {
			bestNode = node
			lowestScore = scores[node.Name]
		}
	}

	return bestNode
}

type Greedy struct {
	Name string
}

func (g *Greedy) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	return selectCandidateNodes(t, nodes)
}
func (g *Greedy) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	nodeScores := make(map[string]float64)

	for _, node := range nodes {
		cpuUsage, err := calculateCpuUsage(node)
		if err != nil {
			log.Printf("error calculating cpu usage for node %s, skipping: %v\n", node.Name, err)
			continue
		}
		cpuLoad := calculateLoad(float64(*cpuUsage), math.Pow(2, 0.8))
		nodeScores[node.Name] = cpuLoad
	}

	return nodeScores
}

func (g *Greedy) Pick(scores map[string]float64, candidates []*node.Node) *node.Node {
	minCpu := 0.0

	var bestNode *node.Node
	for idx, node := range candidates {
		if idx == 0 {
			minCpu = scores[node.Name]
			bestNode = node
			continue
		}

		if scores[node.Name] < minCpu {
			minCpu = scores[node.Name]
			bestNode = node
		}
	}

	return bestNode
}

// Enhanced PVM (Parallel Virtual Machine) Algorithm
//
// Implementation of the E-PVM algorithm laid out in http://www.cnds.jhu.edu/pub/papers/mosix.pdf.
// The algorithm calculates the "marginal cost" of assigning a task to a machine. In the paper and
// in this implementation, the only resources considered for calculating a task's marginal cost are
// memory and cpu.
type Epvm struct {
	Name string
}

func (e *Epvm) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	return nil
}
func (e *Epvm) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	return nil
}
func (e *Epvm) Pick(scores map[string]float64, candidates []*node.Node) *node.Node {
	return nil
}

func selectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	var candidates []*node.Node
	for node := range nodes {
		if checkDisk(t, nodes[node].Disk-nodes[node].DiskAllocated) {
			candidates = append(candidates, nodes[node])
		}
	}
	return candidates
}

func checkDisk(t task.Task, diskAvailable int64) bool {
	return t.Disk < diskAvailable
}

func calculateLoad(usage float64, capacity float64) float64 {
	return usage / capacity
}

// See discussion from this StackOverflow thread:
// https://stackoverflow.com/questions/23367857/accurate-calculation-of-cpu-usage-given-in-percentage-in-linux
func calculateCpuUsage(node *node.Node) (*float64, error) {
	stat1, err := node.GetStats()
	if err != nil {
		return nil, err
	}
	time.Sleep(3 * time.Second)

	stat2, err := node.GetStats()
	if err != nil {
		return nil, err
	}

	stat1Idle := stat1.CPUStats.Idle + stat1.CPUStats.IOWait
	stat2Idle := stat2.CPUStats.Idle + stat2.CPUStats.IOWait

	stat1NonIdle := stat1.CPUStats.User + stat1.CPUStats.Nice + stat1.CPUStats.System + stat1.CPUStats.IRQ + stat1.CPUStats.SoftIRQ + stat1.CPUStats.Steal
	stat2NonIdle := stat2.CPUStats.User + stat2.CPUStats.Nice + stat2.CPUStats.System + stat2.CPUStats.IRQ + stat2.CPUStats.SoftIRQ + stat2.CPUStats.Steal

	stat1Total := stat1Idle + stat1NonIdle
	stat2Total := stat2Idle + stat2NonIdle

	total := stat2Total - stat1Total
	idle := stat2Idle - stat1Idle

	var cpuPercentUsage float64
	if total == 0 && idle == 0 {
		cpuPercentUsage = 0.00
	} else {
		cpuPercentUsage = (float64(total) - float64(idle)) / float64(total)
	}

	return &cpuPercentUsage, nil
}
