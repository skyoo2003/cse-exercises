#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from typing import Optional

class BSTNode(object):
    def __init__(self, data: int,
                 parent: Optional['BSTNode'] = None,
                 left: Optional['BSTNode'] = None,
                 right: Optional['BSTNode'] = None) -> None:
        self.data = data
        self.parent = parent
        self.left = left
        self.right = right

    def __str__(self) -> str:
        return str(self.__dict__)

    def __repr__(self) -> str:
        return str(self)

    @staticmethod
    def deepcopy(node: 'BSTNode') -> 'BSTNode':
        parent = BSTNode.deepcopy(node.parent) if node.parent else None
        left = BSTNode.deepcopy(node.left) if node.left else None
        right = BSTNode.deepcopy(node.right) if node.right else None
        return BSTNode(data=node.data, parent=parent, left=left, right=right)

    def find(self, data: int) -> Optional['BSTNode']:
        node = self
        while node is not None:
            if data < node.data:
                node = node.left
            elif data > node.data:
                node = node.right
            else:
                return node
        return None

    def minimum(self) -> 'BSTNode':
        node = self
        while node.left is not None:
            node = node.left
        return node

    def maximum(self) -> 'BSTNode':
        node = self
        while node.right is not None:
            node = node.right
        return node

    def insert(self, data: int) -> bool:
        node = self
        while node is not None:
            if data < node.data:
                if node.left is None:
                    node.left = BSTNode(data=data, parent=node)
                    return True
                node = node.left
            elif data > node.data:
                if node.right is None:
                    node.right = BSTNode(data=data, parent=node)
                    return True
                node = node.right
            else: # NOTE: Invalid case because data already exists
                return False
        return True

    def delete(self, data: int) -> bool:
        node = self.find(data)
        if node is None:
            return False

        if node.left and node.right:
            node.data = node.left.maximum().data
            return node.left.delete(node.data)
        elif node.left:
            node.parent.right = node.left
        elif node.right:
            node.parent.left = node.right
        del node
        return True

class BST:
    def __init__(self):
        self.root: Optional['BSTNode'] = None

    def find(self, data: int) -> Optional['BSTNode']:
        if self.root is None:
            return None
        return self.root.find(data)

    def insert(self, data: int) -> bool:
        if self.root is None:
            self.root = BSTNode(data=data)
            return True
        return self.root.insert(data)

    def delete(self, data: int) -> bool:
        if self.root is None:
            return False
        return self.root.delete(data)
