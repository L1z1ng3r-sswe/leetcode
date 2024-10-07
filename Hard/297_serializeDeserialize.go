	type Codec struct {
	}

	func Constructor() Codec {
		return Codec{}
	}

	func (this *Codec) serialize(root *TreeNode) string {
		if root == nil {
			return "null"
		}

		var res []string

		queue := []*TreeNode{root}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left, node.Right)
			}
		}

		for len(res) > 0 && res[len(res)-1] == "null" {
			res = res[:len(res)-1]
		}

		return strings.Join(res, ",")
	}

	func (this *Codec) deserialize(data string) *TreeNode {
		if data == "null" {
			return nil
		}

		dataList := strings.Split(data, ",") // slice of strings

		rootVal, _ := strconv.Atoi(dataList[0])
		root := &TreeNode{Val: rootVal}

		queue := []*TreeNode{root}

		i := 1
		for i < len(dataList) && len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			// left value
			if dataList[i] != "null" {
				leftVal, _ := strconv.Atoi(dataList[i])
				node.Left = &TreeNode{Val: leftVal}
				queue = append(queue, node.Left)
			}

			// shift to the right
			i++

			// right value
			if i < len(dataList) && dataList[i] != "null" {
				rightVal, _ := strconv.Atoi(dataList[i])
				node.Right = &TreeNode{Val: rightVal}
				queue = append(queue, node.Right)
			}
			i++
		}

		return root
	}

	// serizize
	// n - number of nodes
	// b - number of nodes on the last level
	// time:  O(n)
	// space: O(b)

	// deserialize
	// n - number of nodes
	// b - number of nodes on the last level
	// time: O(n)
	// space: O(b)