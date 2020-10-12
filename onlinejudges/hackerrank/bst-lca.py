"""
https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem

This answer only fits this question. The best solution is here. (https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/editorial)
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
    left, right = 0, 0
    while True:
        if p1[left].data == p2[right].data:
            return p1[left]
        if left + 1 < len(p1) and len(p1) - left > len(p2) - right:
            left += 1
        elif right + 1 < len(p2) and len(p2) - right > len(p1) - left:
            right += 1
        else:
            break
    return root
