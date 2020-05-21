import java.util.HashMap;

/**
 * 141. Linked List Cycle
 * https://leetcode.com/problems/linked-list-cycle/
 */
public class Solution {


    class ListNode {
        int val;
        ListNode next;
        ListNode(int x) {
            val = x;
            next = null;
        }
    }


    public boolean hasCycle1(ListNode head) {
        if (head == null || head.next == null) {
            return false;
        }

        ListNode step1 = head, step2 = head.next;
        while (true) {
            if (step1 == null || step2 == null || step2.next == null) {
                break;
            } else if (step1 == step2) {
                return true;
            }
            step1 = step1.next;
            step2 = step2.next.next;
        }
        return false;
    }


    public boolean hasCycle2(ListNode head) {
        ListNode node = head;
        HashMap<ListNode, Boolean> map = new HashMap<>();

        while (node != null) {
            if (map.containsKey(node)) {
                return true;
            }
            map.put(node, true);
            node = node.next;
        }
        return false;
    }


    public static void main(String[] args) {
    }

}
