package main

import "time"

type DimissionContext struct {
	Name           string
	WorkingYears   int
	DepartmentName string
	LastDay        time.Time
}

type Handler interface {
	process(ctx *DimissionContext) error
}

type DirectorHandler struct {
}

func (d *DirectorHandler) process(ctx *DimissionContext) error {
	// 确认开发任务完成
	// 确认测试任务完成
	// 确认上线任务完成
	// 确认代码合并完成
	// 确认上线文档编写完成
	// 确认上线文档审核完成
	// 确认交接完成
	return nil
}

type ManagerHandler struct {
}

func (m *ManagerHandler) process(ctx *DimissionContext) error {
	// 确认离职人员工作交接
	// 确认离职人员离职交接
	// 确认离职人员离职面谈
	return nil
}

type HRHandler struct {
}

func (h *HRHandler) process(ctx *DimissionContext) error {
	// 确认离职人员离职面谈
	// 确认离职人员离职交接
	// 确认权限关闭
	// 确认离职人员工作交接
	return nil
}

type CEOHandler struct {
}

func (c *CEOHandler) process(ctx *DimissionContext) error {
	// do something
	return nil
}

type HandlerNode struct {
	handler Handler
	next    *HandlerNode
}

func (h *HandlerNode) handle(ctx *DimissionContext) error {
	err := h.handler.process(ctx)
	if err != nil {
		return err
	}
	if h.next != nil {
		return h.next.handle(ctx)
	}
	return nil
}

func NewDimissionHandlerChain() *HandlerNode {
	directorNode := &HandlerNode{handler: &DirectorHandler{}}
	managerNode := &HandlerNode{handler: &ManagerHandler{}}
	hrNode := &HandlerNode{handler: &HRHandler{}}
	ceoNode := &HandlerNode{handler: &CEOHandler{}}

	directorNode.next = managerNode
	managerNode.next = hrNode
	hrNode.next = ceoNode

	return directorNode
}
