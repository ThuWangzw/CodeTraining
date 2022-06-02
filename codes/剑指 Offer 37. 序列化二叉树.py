# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if root==None:
            return 'null'
        ans = str(root.val)
        ans +="," + self.serialize(root.left)
        ans +=',' + self.serialize(root.right)
        return ans

        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        root, _ = self.deserialize_iter(data)
        return root

    def deserialize_iter(self, data):
        end = 0
        while True:
            if end==len(data) or data[end]==',':
                break
            end += 1
        if data[:end]=='null':
            if end<len(data) and data[end]==',':
                end += 1
            return None, end
        else:
            node = TreeNode(int(data[:end]))
            if end<len(data) and data[end]==',':
                end += 1
            left, leftend = self.deserialize_iter(data[end:])
            end += leftend
            right, rightend = self.deserialize_iter(data[end:])
            end += rightend
            node.left = left
            node.right = right
            return node, end

        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))