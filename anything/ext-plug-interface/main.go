package main

type Node struct {
	name     string
	noOfCPUs int
	usedCPUs int
}

type WeightedPlugin struct {
	plugin ScorePlugin
	weight int
}

type Pod struct {
	requiredCPUs int
}

type ScorePlugin interface {
	Name() string
	Score(node Node, pod Pod) int
}

type CPU struct{}
type BalancedUsage struct{}

func (c *CPU) Name() string {
	return "CPUScore"
}

func (c *CPU) Score(node Node, pod Pod) int {
	free := node.noOfCPUs - node.usedCPUs
	return free
}

func (b *BalancedUsage) Name() string {
	return "BalancedUsage"
}

func (b *BalancedUsage) Score(node Node, pod Pod) int {
	usageRatio := float64(node.usedCPUs) / float64(node.noOfCPUs)

	diff := usageRatio - 0.5
	if diff < 0 {
		diff = -diff
	}

	return int((1 - diff) * 100)
}

func schedule(nodes []Node, pod Pod, plugins []WeightedPlugin) Node {
	bestScore := -1
	var bestNode Node
	for _, node := range nodes {
		totalScore := 0
		for _, wp := range plugins {
			score := wp.plugin.Score(node, pod)
			totalScore += score * wp.weight
		}

		if totalScore > bestScore {
			bestScore = totalScore
			bestNode = node
		}

	}
	return bestNode
}

func main() {
	nodes := []Node{
		{"A", 8, 2},
		{"B", 8, 6},
		{"C", 8, 4},
	}

	pod := Pod{2}

	weightedPlugins := []WeightedPlugin{
		{&CPU{}, 14},
		{&BalancedUsage{}, 1},
	}

	best := schedule(nodes, pod, weightedPlugins)

	println("Scheduled on node:", best.name)
}
