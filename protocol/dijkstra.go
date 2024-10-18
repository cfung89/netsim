package protocol

import (
	"errors"
	"github.com/cfung89/sharp/netsim"
	"github.com/google/uuid"
	"math"
)

type Path struct {
	Nodes  []*Node
	Length float64
}

// Conventional Dijkstra's single-source single-destination (CDSSSD) algorithm
func Dijkstra(g Graph, source uuid.UUID) ([][]*Nodes, error) {
	paths := make(map[uuid.UUID]*Path, 0)
	for index, dest := range g.Destinations {
		length := len(g.Adjacent)

		distance := make(map[uuid.UUID]float64, length)
		used := make(map[uuid.UUID]bool, length)

		for id, _ := range g.Adjacent {
			distance[id] = math.Inf(1)
			used[id] = false
		}
		distance[source] = 0
		previous := make([]*Node, length)

		for {
			minDistance := math.Inf(1)
			var minNode uuid.UUID

			// Find closest unused node to source node (unknown nodes have infinite weight)
			for id, _ := range g.Adjacent {
				if !used[id] && distance[id] < minDistance {
					minDistance = distance[id]
					minNode := id
				}
			}
			if minDistance == math.Inf(1) {
				break
			}
			used[minNode] = true

			for idx, _ := range g.Adjacent[minNode] {
				id := g.Adjacent[minNode][idx].Node.ID
				weight := g.Adjacent[minNode][idx].Weight
				if weight <= 0 {
					return nil, errors.New("Negative or 0 weight in graph.")
				}
				shortestToMinNode := distance[minNode]
				distanceToNextNode := weight
				totalDist := shortestToMinNode + distanceToNextNode
				if totalDist < distance[id] {
					distance[id] := totalDist
					previous[id] := minNode
				}
			}
		}

		var currPath Path
		if distance[dest] == math.Inf(1) {
			// no path found to that destination
			paths[dest] = &Path{
				Nodes:  make([]*Node, 0),
				Length: math.Inf(1),
			}
		}

		var currentNode uuid.UUID = dest
		newPath := &Path{
			Nodes:  make([]*Node, 0),
			Length: distance[dest],
		}
		for currentNode != nil {
			newPath.Nodes = append(newPath.Nodes, currentNode)
			currentNode = previous[currentNode]
		}

		// reverse path
		rev := func(tempPath []*Node) ([]*Node, error) {
			if len(tempPath) == 0 {
				return []*Node, errors.New("Path of length 0")
			}
			temp := make([]*Node, len(tempPath))
			copy(temp, tempPath)
			for i, j := 0, len(tempPath)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
			return temp
		}

		newPath.Nodes, err := rev(newPath.Nodes)
		if err != nil {
			return nil, err
		}
		paths = append(path, newPath.Nodes)
	}
	return paths, nil
}
