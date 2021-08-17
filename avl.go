package avl

type AVL interface {
	Set(key int, value interface{})
	Del(key int)
	Get(key int) (value interface{}, ok bool)
	Print() (keyList []int, valueList []interface{})
}

var max = func(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

type avl struct {
	root *node
}

type node struct {
	left, right *node
	height      int
	key         int
	value       interface{}
}

func (a *avl) height(cur *node) int {
	if cur == nil {
		return 0
	}
	return cur.height

}

func (a *avl) rightSpin(cur *node) (ret *node) {

	ret = cur.left
	cur.left = ret.right
	ret.right = cur

	cur.height = max(a.height(cur.left), a.height(cur.right))
	ret.height = max(a.height(ret.left), a.height(ret.right))

	return
}

func (a *avl) leftSpin(cur *node) (ret *node) {

	ret = cur.right
	cur.right = ret.left
	ret.left = cur

	cur.height = max(a.height(cur.left), a.height(cur.right))
	ret.height = max(a.height(ret.left), a.height(ret.right))

	return
}

func (a *avl) LL_logic(cur *node) (ret *node) {

	//LL
	return a.rightSpin(cur)
}

func (a *avl) RR_logic(cur *node) (ret *node) {

	//RR
	return a.leftSpin(cur)
}

func (a *avl) LR_logic(cur *node) (ret *node) {

	//LR
	cur.left = a.leftSpin(cur.left)
	//LL
	return a.rightSpin(cur)

}

func (a *avl) RL_logic(cur *node) (ret *node) {

	//RL
	cur.right = a.rightSpin(cur.right)
	//RR
	return a.leftSpin(cur)

}

func (a *avl) insert(cur, new *node) (ret *node) {

	switch {
	case cur == nil:
		ret = new
	case cur.key > new.key:
		//left
		cur.left = a.insert(cur.left, new)
		ret = cur
		if a.height(ret.left)-a.height(ret.right) == 2 {
			if new.key < ret.left.key {
				ret = a.LL_logic(ret)
			} else {
				ret = a.LR_logic(ret)
			}
		}

		ret.height = max(a.height(ret.left), a.height(ret.right)) + 1
	case cur.key < new.key:
		//right
		cur.right = a.insert(cur.right, new)
		ret = cur
		if a.height(ret.right)-a.height(ret.left) == 2 {
			if new.key > ret.right.key {
				ret = a.RR_logic(ret)
			} else {
				//RL-spin
				ret = a.RL_logic(ret)
			}
		}

		ret.height = max(a.height(ret.left), a.height(ret.right)) + 1
	default:
		cur.value = new.value
		ret = cur
	}

	return
}

func (a *avl) minNode(cur *node) *node {

	for cur != nil {
		cur = cur.left
	}
	return cur
}

func (a *avl) maxNode(cur *node) *node {

	for cur != nil {
		cur = cur.right
	}
	return cur
}

func (a *avl) delete(cur *node, key int) (ret *node) {

	switch {
	case cur == nil:
		return nil
	case cur.key > key:
		//left
		cur.left = a.delete(cur.left, key)
		ret = cur

		if a.height(ret.right)-a.height(ret.left) == 2 {
			if a.height(ret.right.right) > a.height(ret.right.left) {
				ret = a.RR_logic(ret)
			} else {
				ret = a.RL_logic(ret)
			}
		}
	case cur.key < key:
		//right
		cur.right = a.delete(cur.right, key)
		ret = cur

		if a.height(ret.left)-a.height(ret.right) == 2 {
			if a.height(ret.left.left) > a.height(ret.left.right) {
				ret = a.LL_logic(ret)
			} else {
				ret = a.LR_logic(ret)
			}
		}
	default:
		switch {
		case cur.left != nil && cur.right == nil:
			ret = cur.left
		case cur.left == nil && cur.right != nil:
			ret = cur.right
		default:
			if a.height(cur.left) > a.height(cur.right) {
				tmpNode := a.maxNode(cur.left)
				cur.key = tmpNode.key
				cur.value = tmpNode.value
				cur.left = a.delete(cur.left, tmpNode.key)
				ret = cur
			} else {
				tmpNode := a.minNode(cur.right)
				cur.key = tmpNode.key
				cur.value = tmpNode.value
				cur.right = a.delete(cur.right, tmpNode.key)
				ret = cur
			}
		}
	}
	return
}

func (a *avl) get(cur *node, key int) (value interface{}, ok bool) {

	switch {
	case cur == nil:
		return 0, false
	case cur.key == key:
		return cur.value, true
	case key < cur.key:
		//left
		return a.get(cur.left, key)
	default:
		//right
		return a.get(cur.right, key)
	}
}

func (a *avl) print(cur *node) (keyList []int, valueList []interface{}) {

	if cur == nil {
		return
	}

	if cur.left != nil {
		tmp1, tmp2 := a.print(cur.left)
		keyList = append(keyList, tmp1...)
		valueList = append(valueList, tmp2...)
	}

	keyList = append(keyList, cur.key)
	valueList = append(valueList, cur.value)

	if cur.right != nil {
		tmp1, tmp2 := a.print(cur.right)
		keyList = append(keyList, tmp1...)
		valueList = append(valueList, tmp2...)
	}

	return
}

func GenAVL() AVL {
	return &avl{}
}

func (a *avl) Set(key int, value interface{}) {

	a.root = a.insert(a.root, &node{
		key:    key,
		value:  value,
		height: 1,
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
