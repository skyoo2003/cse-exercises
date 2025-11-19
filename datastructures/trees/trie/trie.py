class Node:
    def __init__(
        self,
        key,
        data=None,
    ) -> None:
        self.key = key
        self.data = data
        self.children = {}
