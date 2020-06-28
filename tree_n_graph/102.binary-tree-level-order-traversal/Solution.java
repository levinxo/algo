import java.util.List;
import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedList;

/**
 * https://leetcode.com/problems/binary-tree-level-order-traversal/
 * 102. 二叉树的层序遍历
 */
public class Solution {


    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) {
            val = x;
        }
    }


    /**
     * Time Complexity O(n)
     * Space Comlexity O(n)
     */
	public List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) {
            return Collections.emptyList();
        }

        LinkedList<TreeNode> nodeList = new LinkedList<TreeNode>();
        nodeList.add(root);

        List<List<Integer>> res = new ArrayList<List<Integer>>();

        while (!nodeList.isEmpty()) {
            int length = nodeList.size();
            List<Integer> r = new ArrayList<Integer>();

            for (int i = 0; i < length; i++) {
                TreeNode tmp = nodeList.poll();
                r.add(tmp.val);
                if (tmp.left != null) {
                    nodeList.add(tmp.left);
                }
                if (tmp.right != null) {
                    nodeList.add(tmp.right);
                }
            }
            res.add(r);
        }

        return res;
    }


}
