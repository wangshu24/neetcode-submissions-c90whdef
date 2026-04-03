type MinStack struct {
	stack []int
	m int
}

func Constructor() MinStack {
	return MinStack {
		stack : make([]int, 0),
		m: 0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.m = val
	} else {
		this.stack = append(this.stack, val - this.m)
		if val < this.m { 
			this.m = val
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if pop < 0 {
		this.m = this.m - pop
	}
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack)-1]
	if top > 0 {
		return top + this.m
	}
	return this.m
}

func (this *MinStack) GetMin() int {
	return this.m
}
