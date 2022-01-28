package main

type RoundRobinBalancer struct{
	stat []int
	node_num chan int
	do chan bool 
}

func (balancer* RoundRobinBalancer) Init(n int) {
	balancer.stat = make([]int, n)
	balancer.node_num = make(chan int)
	balancer.do = make(chan bool)

	node := 0

	go func(){
		do := balancer.do
		for {
			if <-do{
				balancer.stat[node]++
				node = (node+1) % n
				balancer.node_num <- node
			}
		}
	}()
}

func (balancer RoundRobinBalancer) GiveStat() []int{
	return balancer.stat
}

func (balancer* RoundRobinBalancer) GiveNode() int{
	balancer.do <- true
	return <-balancer.node_num 
}