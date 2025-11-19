class Node:
    def __init__(self):
        self.children = {}
        self.end = False


class Trie:
    def __init__(self):
        self.root = Node()

    def insert(self, text):
        node = self.root
        for char in text:
            if char not in node.children:
                node.children[char] = Node()
            node = node.children[char]
        node.end = True

    def get_node(self, text):
        node = self.root
        for char in text:
            if char not in node.children:
                return None
            node = node.children[char]
        return node

    def exists(self, text):
        node = self.get_node(text)
        return node and node.end

    def search(self, prefix):
        node = self.get_node(prefix)
        if not node:
            return []
        words = []
        stack = [(node, prefix)]
        while stack:
            node, word = stack.pop()
            if node.end:
                words.append(word)
            for char, child in node.children.items():
                stack.append((child, word + char))
        return words

    def delete(self, text):
        stack = []

        node = self.root
        for char in text:
            if char not in node.children:
                return False
            stack.append((node, char))
            node = node.children[char]

        if not node.end:
            return False

        node.end = False

        while stack:
            parent, char = stack.pop()
            node = parent.children[char]
            if node.children or node.end:
                break
            del parent.children[char]

        return True


def display_trie(trie):
    def _display(node, depth=1):
        if node.children:
            print(f"{'  ' * depth}Children: {list(node.children.keys())}")
            for child in node.children.values():
                _display(child, depth + 1)

    print("Root:")
    _display(trie.root)


if __name__ == "__main__":
    trie = Trie()
    trie.insert("can")
    trie.insert("car")
    trie.insert("cart")
    print("ca 검색 결과: ", trie.exists("ca"))
    print("can 검색 결과: ", trie.exists("can"))
    print("cart 검색 결과: ", trie.exists("cart"))

    trie = Trie()
    trie.insert("can")
    trie.insert("car")
    trie.insert("cart")
    trie.insert("apple")
    trie.insert("apply")
    print(trie.search("app"))
    print(trie.search("car"))

    trie = Trie()
    trie.insert("can")
    trie.insert("car")
    trie.insert("cart")
    trie.insert("apple")
    trie.insert("apply")
    display_trie(trie)

    trie.delete("can")
    display_trie(trie)

    trie.delete("apple")
    display_trie(trie)

    print(trie.search("can"))
    print(trie.search("apple"))
