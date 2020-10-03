#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from unittest import TestCase
from cs.trees.bst import BSTNode, BST

class BSTNodeTestCase(TestCase):
    def setUp(self):
        self.left = BSTNode(data=5)
        self.right = BSTNode(data=15)
        self.root = BSTNode(data=10, left=self.left, right=self.right)

    def test_assigned_variables(self):
        self.assertEqual(self.left.data, 5)
        self.assertEqual(self.left.left, None)
        self.assertEqual(self.left.right, None)

        self.assertEqual(self.right.data, 15)
        self.assertEqual(self.right.left, None)
        self.assertEqual(self.right.right, None)

        self.assertEqual(self.root.data, 10)
        self.assertEqual(self.root.left, self.left)
        self.assertEqual(self.root.right, self.right)

class BSTTestCase(TestCase):
    def setUp(self):
        self.tree = BST()

    def test_insert_data(self):
        success = self.tree.insert(10)
        self.assertEqual(success, True)
        self.assertEqual(self.tree.root.data, 10)

        success = self.tree.insert(5)
        self.assertEqual(success, True)
        self.assertEqual(self.tree.root.left.data, 5)

        success = self.tree.insert(15)
        self.assertEqual(success, True)
        self.assertEqual(self.tree.root.right.data, 15)

    def test_delete_data(self):
        self.tree.insert(10)
        self.tree.insert(8)
        self.tree.insert(4)
        self.tree.insert(9)
        self.tree.insert(3)
        self.tree.insert(5)

        # NOTE: Expected tree structure below
        #         10
        #       8
        #     4   9
        #   3   5
        self.assertEqual(self.tree.root.data, 10)
        self.assertEqual(self.tree.root.left.data, 8)
        self.assertEqual(self.tree.root.left.left.data, 4)
        self.assertEqual(self.tree.root.left.right.data, 9)
        self.assertEqual(self.tree.root.left.left.left.data, 3)
        self.assertEqual(self.tree.root.left.left.right.data, 5)
        
        # NOTE: Expected tree structure after deletion
        #         10
        #       5
        #     4   9
        #   3   
        success = self.tree.delete(8)
        self.assertEqual(success, True)
        self.assertEqual(self.tree.root.left.data, 5)
        self.assertEqual(self.tree.root.left.left.data, 4)
        self.assertEqual(self.tree.root.left.left.left.data, 3)
        self.assertEqual(self.tree.root.left.right.data, 9)

        node = self.tree.find(8)
        self.assertIsNone(node)
