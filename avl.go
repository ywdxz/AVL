package avl

// AVL 平衡二叉树
/*
任何一个节点的左子树和右子树高度差不超过1
*/
type AVL interface {
	Set(key int, value interface{})
	Del(key int)
	Get(key int) (value interface{}, ok bool)
	Print() (keyList []int, valueList []interface{})
}

type node struct {
	left, right *node
	key         int
	height      int
	value       interface{}
}

type avl struct {
	root *node
}

var max = func(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (a *avl) height(cur *node) int {
	if cur == nil {
		return 0
	}
	return cur.height
}

// leftSpin 左旋转
/*
  	   a              c
     /   \          /   \
    b     c   =>   a     h
   / \   / \      / \
  e   f g   h    b   g
                / \
               e   f
*/
func (a *avl) leftSpin(cur *node) (ret *node) {

	ret = cur.right
	cur.right = ret.left
	ret.left = cur

	cur.height = max(a.height(cur.left), a.height(cur.right)) + 1
	ret.height = max(a.height(ret.left), a.height(ret.right)) + 1

	return
}

// rightSpin 右旋转
/*
  	   a              b
     /   \          /   \
    b     c   =>   e     a
   / \   / \            / \
  e   f g   h          f   c
                          / \
                         g   h
*/
func (a *avl) rightSpin(cur *node) (ret *node) {

	ret = cur.left
	cur.left = ret.right
	ret.right = cur

	cur.height = max(a.height(cur.left), a.height(cur.right)) + 1
	ret.height = max(a.height(ret.left), a.height(ret.right)) + 1

	return
}

// LL_logic
// 右旋
/*
        a                b
      /   \            /   \
     b     c          d      a
    / \              / \    / \
   d   e      =>    f   g  e   c
  / \
 f   g
*/
func (a *avl) LL_logic(cur *node) *node {
	return a.rightSpin(cur)
}

//RR_logic
//左旋
/*
   a                          c
 /   \                       /  \
b     c         =>          a     e
     / \                   / \   / \
    d   e                 b   d f   g
       / \
      f   g
*/
func (a *avl) RR_logic(cur *node) *node {
	return a.leftSpin(cur)
}

// LR_logic
// 左右旋
/*
        a                  a                   e
      /   \              /   \               /   \
     b     c            e     c             b     a
    / \                / \                 / \   / \
   d   e        =>    b   g         =>    d   f g   c
      / \            / \
     f   g          d   f
*/
func (a *avl) LR_logic(cur *node) *node {
	cur.left = a.leftSpin(cur.left)
	return a.rightSpin(cur)
}

// RL_logic
// 左右旋
/*
   a
 /   \
b     c
     / \
    d   e
   / \
  f   g
*/
func (a *avl) RL_logic(cur *node) *node {
	cur.right = a.rightSpin(cur.right)
	return a.leftSpin(cur)
}

func (a *avl) checkBalance(cur *node) *node {

	if cur == nil {
		return nil
	}

	switch a.height(cur.left) - a.height(cur.right) {
	case 2:
		//L
		if a.height(cur.left.left) > a.height(cur.left.right) {
			//LL
			cur = a.LL_logic(cur)
		} else {
			//LR
			cur = a.LR_logic(cur)
		}
	case -2:
		//R
		if a.height(cur.right.left) < a.height(cur.right.right) {
			//RR
			cur = a.RR_logic(cur)
		} else {
			//RL
			cur = a.RL_logic(cur)
		}
	default:
		cur.height = max(a.height(cur.left), a.height(cur.right)) + 1
	}

	return cur
}

func (a *avl) minNode(cur *node) *node {
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (a *avl) maxNode(cur *node) *node {
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

func (a *avl) insert(cur *node, new *node) (ret *node) {

	switch {
	case cur == nil:
		cur = new
		cur.height = 1
	case cur.key > new.key:
		cur.left = a.insert(cur.left, new)
	case cur.key < new.key:
		cur.right = a.insert(cur.right, new)
	case cur.key == new.key:
		cur.value = new.value
	}

	ret = a.checkBalance(cur)
	return
}

func (a *avl) delete(cur *node, key int) (ret *node) {

	switch {
	case cur == nil:
	case cur.key > key:
		cur.left = a.delete(cur.left, key)
	case cur.key < key:
		cur.right = a.delete(cur.right, key)
	case cur.key == key:
		switch {
		case cur.left == nil && cur.right == nil:
			cur = nil
		case cur.left != nil && cur.right != nil:
			if a.height(cur.left) < a.height(cur.right) {
				tmpNode := a.minNode(cur.right)
				cur.key = tmpNode.key
				cur.value = tmpNode.value
				cur.right = a.delete(cur.right, tmpNode.key)
			} else {
				tmpNode := a.maxNode(cur.left)
				cur.key = tmpNode.key
				cur.value = tmpNode.value
				cur.left = a.delete(cur.left, tmpNode.key)
			}
		case cur.left != nil:
			cur = cur.left
		case cur.right != nil:
			cur = cur.right
		}
	}

	ret = a.checkBalance(cur)
	return
}

func (a *avl) get(cur *node, key int) (value interface{}, ok bool) {
	switch {
	case cur == nil:
		value, ok = nil, false
	case cur.key > key:
		value, ok = a.get(cur.left, key)
	case cur.key < key:
		value, ok = a.get(cur.right, key)
	case cur.key == key:
		value, ok = cur.value, true
	}
	return
}

func (a *avl) print(cur *node) (keyList []int, valueList []interface{}) {

	if cur == nil {
		return
	}

	if cur.left != nil {
		l1, l2 := a.print(cur.left)
		keyList = append(keyList, l1...)
		valueList = append(valueList, l2...)
	}

	keyList = append(keyList, cur.key)
	valueList = append(valueList, cur.value)

	if cur.right != nil {
		l1, l2 := a.print(cur.right)
		keyList = append(keyList, l1...)
		valueList = append(valueList, l2...)
	}
	return
}

func GenAVL() AVL {
	return &avl{}
}

func (a *avl) Set(key int, value interface{}) {

	a.root = a.insert(a.root, &node{
		key:   key,
		value: value,
	})
}

func (a *avl) Del(key int) {
	a.root = a.delete(a.root, key)
}

func (a *avl) Get(key int) (value interface{}, ok bool) {
	return a.get(a.root, key)
}

func (a *avl) Print() (keyList []int, valueList []interface{}) {
	return a.print(a.root)
}
