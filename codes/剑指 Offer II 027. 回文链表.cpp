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
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if(head==nullptr || head->next==nullptr) {
            return true;
        }
        auto fast = head->next;
        auto slow = head;
        while(fast != nullptr) {
            if(fast->next == nullptr) {
                break;
            }
            fast = fast->next->next;
            slow = slow->next;
        }
        auto tmp = slow;
        slow = slow->next;
        tmp->next = nullptr;
        auto rtail = slow;
        auto last = slow;
        slow = slow->next;
        while(slow!=nullptr) {
            auto tmp = slow;
            slow = slow->next;
            tmp->next = last;
            last = tmp;
        }
        rtail->next = nullptr;
        auto p = head, q = last;
        while(p!=nullptr && q!=nullptr) {
            if(p->val != q->val) {
                return false;
            }
            p = p->next;
            q = q->next;
        }
        return true;
    }
};