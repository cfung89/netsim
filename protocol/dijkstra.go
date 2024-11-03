package protocol

import (
	"errors"
	"math"

	"github.com/cfung89/sharp/netsim"
	"github.com/google/uuid"
)

// Conventional Dijkstra's single-source single-destination (CDSSSD) algorithm
func Dijkstra(g *netsim.Graph, source uuid.UUID) (map[uuid.UUID]*netsim.Path, error) {
	paths := make(map[uuid.UUID]*netsim.Path, 0)
	for _, dest := range g.Destinations {
		previous := make(map[uuid.UUID]*netsim.Node)
		distance := make(map[uuid.UUID]float64)
		used := make(map[uuid.UUID]bool)

		for id := range g.Adjacent {
			distance[id] = math.Inf(1)
			used[id] = false
		}
		distance[source] = 0

		first := true
		for {
			minDistance := math.Inf(1)
			var minNode *netsim.Node

			// Find closest unused node to source node (unknown nodes have infinite weight)
			if first {
				minDistance = distance[source]
				minNode = g.Network[source]
				first = false
			}
			for id := range g.Adjacent {
				if id != source && !used[id] && distance[id] < minDistance {
					minDistance = distance[id]
					minNode = g.Network[id]
				}
			}

			// Cannot get to the node
			if minDistance == math.Inf(1) {
				break
			}
			used[minNode.ID] = true

			for idx := range g.Adjacent[minNode.ID] {
				id := g.Adjacent[minNode.ID][idx].Node.ID
				weight := g.Adjacent[minNode.ID][idx].Weight
				if weight <= 0 {
					return nil, errors.New("Negative or 0 weight in graph.")
				}
				shortestToMinNode := distance[minNode.ID]
				distanceToNextNode := weight
				totalDist := shortestToMinNode + distanceToNextNode
				if totalDist < distance[id] {
					distance[id] = totalDist
					previous[id] = minNode
				}
			}
		}

		if distance[dest] == math.Inf(1) {
			// no path found to that destination
			paths[dest] = &netsim.Path{
				Nodes:  make([]*netsim.Node, 0),
				Length: math.Inf(1),
			}
			continue
		}

		var currentID uuid.UUID = dest
		newPath := &netsim.Path{
			Nodes:  make([]*netsim.Node, 0),
			Length: distance[dest],
		}

		for {
			newPath.Nodes = append(newPath.Nodes, g.Network[currentID])
			_, ok := previous[currentID]
			if !ok {
				break
			}
			currentID = previous[currentID].ID
		}

		// reverse path
		rev := func(tempPath []*netsim.Node) ([]*netsim.Node, error) {
			if len(tempPath) == 0 {
				return nil, errors.New("Path of length 0")
			}
			temp := make([]*netsim.Node, len(tempPath))
			copy(temp, tempPath)
			for i, j := 0, len(tempPath)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
			return temp, nil
		}

		newPath.Nodes, _ = rev(newPath.Nodes)
		paths[dest] = newPath
	}
	return paths, nil
}
