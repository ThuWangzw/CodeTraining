/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        if(root==nullptr) return "n";
        string ans = to_string(root->val)+",";
        ans += serialize(root->left)+",";
        ans += serialize(root->right);
        return ans;
    }

    TreeNode* deserializeStr(string &data) {
        if(data.size()==0) return nullptr;
        auto it = data.find(',');
        string token(data, 0, it);
        data = string(data.begin()+it+1, data.end());
        if(token=="n") {
            return nullptr;
        } else {
            auto node = new TreeNode(stoi(token));
            node->left = deserializeStr(data);
            node->right = deserializeStr(data);
            return node;
        }
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        return deserializeStr(data);
    }
};

// Your Codec object will be instantiated and called as such:
// Codec ser, deser;
// TreeNode* ans = deser.deserialize(ser.serialize(root));