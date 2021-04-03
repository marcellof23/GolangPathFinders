package main

func DFS(G *Graph, source *Vertex, path func(int), vis []bool) {

	if source == nil {
		return
	}
	vis[source.Key] = true
	path(source.Key)

	for _, v := range source.Nodes {
		if !vis[v.Key] {
			DFS(G, v, path, vis)
		}

	}
}
