package org.vatsuvaksi;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;

import java.util.logging.Logger;

/*
 * Huffman Encoding Algorithm John Cricket Challenge No. 4
 * */
public class Main {

    private static final Logger logger = Logger.getLogger(Main.class.getName());

    private interface HuffBaseNode {
        boolean isLeaf();

        int weight();
    }

    /**
     * Huffman tree node: Leaf class
     */
    static class HuffLeafNode implements HuffBaseNode {
        private final char element;      // Element for this node
        private final int weight;        // Weight for this node

        /**
         * Constructor
         */
        HuffLeafNode(char el, int wt) {
            element = el;
            weight = wt;
        }

        /**
         * @return The element value
         */
        char value() {
            return element;
        }

        /**
         * @return The weight
         */
        public int weight() {
            return weight;
        }

        /**
         * Return true
         */
        public boolean isLeaf() {
            return true;
        }
    }

    /**
     * Huffman tree node: Internal class
     */
    static class HuffInternalNode implements HuffBaseNode {
        private final int weight;
        private final HuffBaseNode left;
        private final HuffBaseNode right;

        /**
         * Constructor
         */
        HuffInternalNode(HuffBaseNode l,
                         HuffBaseNode r, int wt) {
            left = l;
            right = r;
            weight = wt;
        }

        /**
         * @return The left child
         */
        HuffBaseNode left() {
            return left;
        }

        /**
         * @return The right child
         */
        HuffBaseNode right() {
            return right;
        }

        /**
         * @return The weight
         */
        public int weight() {
            return weight;
        }

        /**
         * Return false
         */
        public boolean isLeaf() {
            return false;
        }
    }

    /**
     * A Huffman coding tree
     */
    static class HuffTree implements Comparable<HuffTree> {
        private final HuffBaseNode root;

        /**
         * Constructors
         */
        HuffTree(char el, int wt) {
            root = new HuffLeafNode(el, wt);
        }

        HuffTree(HuffBaseNode l, HuffBaseNode r, int wt) {
            root = new HuffInternalNode(l, r, wt);
        }

        HuffBaseNode root() {
            return root;
        }

        int weight() // Weight of tree is weight of root
        {
            return root.weight();
        }

        @Override
        public int compareTo(HuffTree that) {
            return Integer.compare(root.weight(), that.weight());
        }

        // Overriding because sonar lint said when we override compare to we should also override equals
        @Override
        public boolean equals(Object obj) {
            if (null == obj) return false;

            if (this.getClass() != obj.getClass()) {
                return false;
            }

            HuffTree that = (HuffTree) obj;

            return that.root.weight() == this.root.weight();
        }

        // Not Overriding hashcode because equals will not use
        @Override
        public int hashCode() {
            return super.hashCode();
        }
    }

    public static void main(String[] args) {

        logger.info("Welcome to my compression tool - Vatsal [@vatsuvaksi]!");

        // Reading a file from input
        String importedFile = readFile(args[0]);

        //Calculate frequency of chars in the imported file
        Map<Character, Integer> frequenceMap = calculateFrequency(importedFile == null ? "" : importedFile);

        // Building a binary tree for this frequencyMap
        HuffTree huffTree = buildBinaryTreeForFrequencyMap(frequenceMap);
        System.out.println(huffTree.root.weight());

    }



    // Builds a binary tree for the frequency map
    private static HuffTree buildBinaryTreeForFrequencyMap(Map<Character, Integer> frequenceMap) {

        // 01 Building Priority Queue
        Queue<HuffTree> queue = new PriorityQueue<>();
        for (Map.Entry<Character, Integer> entry : frequenceMap.entrySet()) {
            queue.add(new HuffTree(entry.getKey(), entry.getValue()));
        }

        return buildHuffTreeHelper(queue);
    }

    private static HuffTree buildHuffTreeHelper(Queue<HuffTree> queue) {

        HuffTree mergedHuffTree = null;
        while (queue.size() > 1) {
            HuffTree first = queue.remove();
            HuffTree second = queue.remove();

            mergedHuffTree = new HuffTree(first.root, second.root, first.root().weight() + second.root.weight());
            queue.add(mergedHuffTree);
        }
        return mergedHuffTree;
    }

    // Calculates the frequency of each occurring character
    private static Map<Character, Integer> calculateFrequency(String importedFile) {
        Map<Character, Integer> frequenceMap = new HashMap<>();
        for (char c : importedFile.toCharArray()) {
            frequenceMap.put(c, frequenceMap.getOrDefault(c, 0) + 1);
        }
        return frequenceMap;
    }

    // Reading the content of the file
    private static String readFile(String filePath) {
        try {
            return new String(Files.readAllBytes(Paths.get(filePath)));
        } catch (IOException e) {
            logger.warning("Oops.. Could not read file: " + filePath);
            return null;
        }
    }
}