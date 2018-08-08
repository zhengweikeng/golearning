package bst

type Data struct {
	Value string
}

type BST struct {
	Key   string
	Value *Data
	Left  *BST // 左子树
	Right *BST // 右子树
	N     int  // 以该节点为根的子树中节点的总数
}

func size(tree *BST) int {
	if tree == nil {
		return 0
	}
	return tree.N
}

func (tree *BST) Size() int {
	return size(tree)
}

func get(tree *BST, key string) *Data {
	if tree == nil {
		return nil
	}

	if tree.Key > key {
		return get(tree.Left, key)
	} else if tree.Key < key {
		return get(tree.Right, key)
	} else {
		return tree.Value
	}
}

func (tree *BST) Get(key string) *Data {
	return get(tree, key)
}

func put(root *BST, key string, value *Data) *BST {
	if root == nil {
		return &BST{
			Key:   key,
			Value: value,
			N:     1,
		}
	}

	if root.Key > key {
		root.Left = put(root.Left, key, value)
	} else if root.Key < key {
		root.Right = put(root.Right, key, value)
	} else {
		root.Value = value
	}

	root.N = size(root.Left) + size(root.Right) + 1

	return root
}

func (tree *BST) Put(key string, value *Data) {
	put(tree, key, value)
}
