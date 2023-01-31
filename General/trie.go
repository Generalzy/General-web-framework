package General

type node struct {
	// 待匹配路由
	pattern  string
	// 路由中的一部分
	part     string
	// 子节点列表
	children []*node
	// 是否模糊匹配
	// 路由: "/:name/hello
	// url: "/Generalzy/hello
	// 不添加该字段会导致路由与url无法匹配
	isWild bool
}

// matchChild 寻找匹配到的第一个子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// 精准匹配到路径或本次match是模糊查找的情况下返回节点
		if child.part == part|| child.isWild{
			return child
		}
	}
	return nil
}

// insert 插入节点
// height 路由层数
//
// 以 /:name/hello,Get前缀树 为例:
// 1. pattern="/:name/hello" parts=[":name","hello"] height=0 part=":name" node1={"",":name",[]}
// 结果:root{"","",[node1]}->node1{"",":name",[]}
// 2. pattern="/:name/hello" parts=[":name","hello"] height=1 part="hello" node2={"",":hello",[]}
// 结果:root{"","",[node1]}->node1{"",":name",[node2]}->node2{"","hello",[]}
func (n *node)insert(pattern string,parts []string,height int){
	// 1. 初始化根节点 pattern = / height = 0 parts = []
	// 2. 当height==len(parts)即,到了最后一个节点,将路由整体放入
	if len(parts)==height{
		n.pattern=pattern
		return
	}

	// 获取当前层级的部分url
	part:=parts[height]
	// 遍历根节点寻找part节点
	child := n.matchChild(part)
	// 未找到part节点
	if child == nil{
		// 新建node并且赋值给child
		// 并判断是否有一段路由需要模糊匹配
		child = &node{part: part,isWild: part[0] == ':'}
		n.children = append(n.children, child)
	}

	// 递归下一层路由
	child.insert(pattern, parts, height+1)
}

// matchChildren 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0,buf)
	for _, child := range n.children {
		// 当本次需要模糊匹配时,直接将node添加到nodes
		if child.part == part || child.isWild{
			nodes = append(nodes, child)
		}
	}
	return nodes
}


// search 寻找节点
//
// 以 /:name/hello,Get前缀树 为例:
// 1. parts=[":name","hello"] height=0 part=":name" n=root children=[node1{"",":name",[node1]}] child=node1
// 2. parts=[":name","hello"] height=1 part="hello" n=node1 children=[node2{"","hello",[]}] child=node2 返回node2
func (n *node)search(parts []string, height int) *node {
	// 当遍历到最后一个节点时,返回node
	if len(parts) == height{
		if n.pattern==""{
			return nil
		}
		// 直接将根节点返回
		return n
	}

	part:=parts[height]
	children := n.matchChildren(part)

	// 递归下一层级
	for _, child := range children {
		// 接受返回的node
		result := child.search(parts, height+1)
		// 返回找到的第一个节点
		if result != nil {
			return result
		}
	}
	return nil
}
