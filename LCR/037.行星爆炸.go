func asteroidCollision(asteroids []int) []int {
	stack := []int{} // 定义一个空栈，用于存放幸存的小行星
	for _, asteroid := range asteroids {
		collision := false // 标记是否发生了碰撞
        // 如果栈空，先引入一颗行星
        // 如果栈不为空，当前小行星向右移动（正值），且栈顶小行星向左移动（负值），不会发生碰撞
		// 如果栈不为空，当前小行星向左移动（负值），且栈顶小行星向右移动（正值），则可能发生碰撞
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1] // 取栈顶的小行星
			if top < -asteroid { // 如果栈顶小行星的绝对值小于当前小行星，栈顶小行星爆炸
				stack = stack[:len(stack)-1] // 栈顶小行星被移除，继续比较
				continue // 继续检查栈中是否有其他小行星需要处理
			} else if top == -asteroid { // 如果栈顶小行星和当前小行星大小相等，两颗行星同时爆炸
				stack = stack[:len(stack)-1] // 移除栈顶小行星
			}
			collision = true // 发生碰撞，跳出循环
			break
		}
		if !collision { // 如果没有发生碰撞，将当前小行星加入栈中
			stack = append(stack, asteroid)
		}
	}
	return stack // 返回栈中的小行星（剩余没有爆炸的小行星）
}
