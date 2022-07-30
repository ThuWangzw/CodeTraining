/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
struct CmpListNode {
    bool operator()(const ListNode *a, const ListNode *b) const {
        return a->val > b->val;
    }
};

class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, std::vector<ListNode*>, CmpListNode> pq;
        for(auto node : lists) {
            if(node==nullptr) {
                continue;
            }
            pq.push(node);
        }
        ListNode* head = nullptr;
        ListNode* tail = nullptr;
        while(pq.size()>0) {
            if(head==nullptr) {
                head = tail = pq.top();
            } else {
                tail->next = pq.top();
                tail = tail->next;
            }
            pq.pop();
            if(tail->next != nullptr) {
                pq.push(tail->next);
            }
            tail->next = nullptr;
        }
        return head;
    }
};