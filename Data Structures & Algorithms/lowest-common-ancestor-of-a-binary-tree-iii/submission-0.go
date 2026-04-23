/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	seen := make(map[*Node]bool)
	for p != nil {
		seen[p] = true
		p = p.Parent
	}
	for q != nil {
		if seen[q] {
			return q
		}
		q = q.Parent
	}
	return nil
}
