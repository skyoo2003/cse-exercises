"""
https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem
[NOTE] This answer only fits this question. The best solution is here. ( https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/editorial )
"""

"""
Node is defined as
self.left (the left child of the node)
self.right (the right child of the node)
self.data (the value of the node)
"""

def parents(node, value):
    if not node:
        return []
    if node.data == value:
        return [node]
    elif node.data > value:
        return parents(node.left, value) + [node]
    else:
        return parents(node.right, value) + [node]

def lca(root, v1, v2):
    p1 = parents(root, v1)
    p2 = parents(root, v2)
    l, r = 0, 0
    while True:
        if p1[l].data == p2[r].data:
            return p1[l]
        if l+1 < len(p1) and len(p1)-l > len(p2)-r:
            l += 1
        elif r+1 < len(p2) and len(p2)-r > len(p1)-l:
            r += 1
        else: break
    return root
