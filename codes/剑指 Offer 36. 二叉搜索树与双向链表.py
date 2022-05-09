"""
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""
class Solution:
    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if root is None:
            return None
        left = root.left
        right = root.right
        root.left = root
        root.right = root
        head = root
        if left is not None:
            lefthead = self.treeToDoublyList(left)
            head = self.joinNodes(lefthead, head)
        if right is not None:
            righthead = self.treeToDoublyList(right)
            head = self.joinNodes(head, righthead)
        return head

    def joinNodes(self, a, b):
        atail = a.left
        btail = b.left
        atail.right = b
        b.left = atail
        btail.right = a
        a.left = btail
        return a
        