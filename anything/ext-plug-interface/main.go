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

func runScenario(title string, nodes []Node, pod Pod, plugins []WeightedPlugin) {
	println("\n---", title, "---")

	bestScore := -1
	var bestNode Node

	for _, node := range nodes {
		totalScore := 0
		println("Node:", node.name)

		for _, wp := range plugins {
			score := wp.plugin.Score(node, pod)
			weighted := score * wp.weight

			println("  Plugin:", wp.plugin.Name(),
				"Raw:", score,
				"Weight:", wp.weight,
				"Weighted:", weighted)

			totalScore += weighted
		}

		println("  Total Score:", totalScore)

		if totalScore > bestScore {
			bestScore = totalScore
			bestNode = node
		}
	}

	println("=> Selected Node:", bestNode.name)
}

func main() {
	nodes := []Node{
		{"A", 8, 2},
		{"B", 8, 6},
		{"C", 8, 4},
	}

	pod := Pod{2}

	runScenario("Equal Weights",
		nodes,
		pod,
		[]WeightedPlugin{
			{&CPU{}, 1},
			{&BalancedUsage{}, 1},
		},
	)

	runScenario("CPU Dominates",
		nodes,
		pod,
		[]WeightedPlugin{
			{&CPU{}, 14},
			{&BalancedUsage{}, 1},
		},
	)

	runScenario("Balanced Usage Dominates",
		nodes,
		pod,
		[]WeightedPlugin{
			{&CPU{}, 1},
			{&BalancedUsage{}, 10},
		},
	)

	runScenario("Extreme Weight Bias",
		nodes,
		pod,
		[]WeightedPlugin{
			{&CPU{}, 100},
			{&BalancedUsage{}, 1},
		},
	)

	nodesWithEdge := append(nodes, Node{"D", 8, 8})

	runScenario("Includes Fully Utilized Node", nodesWithEdge, pod, []WeightedPlugin{
		{&CPU{}, 1},
		{&BalancedUsage{}, 1},
	})
}
