/**
 * 146. LRU Cache
 * https://leetcode.com/problems/lru-cache/
 */

import java.util.HashMap;

class LRUCache {

    class Node {
        Node prev;
        Node next;
        int key;
        int value;

        public Node (int key, int value) {
            this.key = key;
            this.value = value;
            prev = null;
            next = null;
        }
    }

    private int capacity;
    private HashMap<Integer, Node> map;

    private Node head;
    private Node tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<Integer, Node>(capacity);
        head = new Node(Integer.MAX_VALUE, -1);
        tail = new Node(Integer.MAX_VALUE, -1);
        head.next = tail;
        tail.prev = head;
    }

    private void addNode(Node node){
        node.next = head.next;
        node.next.prev = node;
        head.next = node;
        node.prev = head;
        map.put(node.key, node);
    }

    private void removeNode(Node node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    private void removeTail() {
        Node current = tail.prev;
        current.prev.next = tail;
        tail.prev = current.prev;
        map.remove(current.key);
    }

    public int get(int key) {
        if (!map.containsKey(key)) {
            return -1;
        }

        Node current = map.get(key);

        removeNode(current);
        addNode(current);
        return current.value;
    }

    public void put(int key, int value) {
        int val = get(key);
        if (val != -1) {
            map.get(key).value = value;
            return;
        }

        if (map.size() == capacity) {
            removeTail();
        }

        Node current = new Node(key, value);
        addNode(current);
    }

    public static void main(String[] args) {
        LRUCache cache = new LRUCache(2);

        cache.put(1, 1);
        cache.put(2, 2);
        System.out.println(cache.get(1));       // returns 1
        cache.put(3, 3);    // evicts key 2
        System.out.println(cache.get(2));       // returns -1 (not found)
        cache.put(4, 4);    // evicts key 1
        System.out.println(cache.get(1));       // returns -1 (not found)
        System.out.println(cache.get(3));       // returns 3
        System.out.println(cache.get(4));       // returns 4
    }

}
