package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}

type Context struct {
	tokens       []string
	currentToken string
	index        int
}

func NewContext(text string) *Context {
	c := &Context{
		tokens: strings.Fields(text),
		index:  0,
	}
	c.NextToken()
	return c
}

func (c *Context) NextToken() string {
	if c.index < len(c.tokens) {
		c.currentToken = c.tokens[c.index]
		c.index++
	} else {
		c.currentToken = ""
	}
	return c.currentToken
}

func (c *Context) CurrentToken() string {
	return c.currentToken
}

func (c *Context) SkipToken(token string) error {
	if c.currentToken == "" {
		return &ParseError{Message: fmt.Sprintf("Error: '%s' が期待されましたが、トークンが見つかりません。", token)}
	} else if token != c.currentToken {
		return &ParseError{Message: fmt.Sprintf("Error: '%s' が期待されましたが、'%s' が見つかりました。", token, c.currentToken)}
	}
	c.NextToken()
	return nil
}

func (c *Context) CurrentNumber() (int, error) {
	if c.currentToken == "" {
		return 0, &ParseError{Message: "Error: トークンが見つかりません。"}
	}
	number, err := strconv.Atoi(c.currentToken)
	if err != nil {
		return 0, &ParseError{Message: fmt.Sprintf("Error: %v", err)}
	}
	return number, nil
}

type Node interface {
	Parse(context *Context) error
	String() string
}

type ProgramNode struct {
	commandListNode Node
}

func NewProgramNode() *ProgramNode {
	return &ProgramNode{}
}

func (p *ProgramNode) Parse(context *Context) error {
	err := context.SkipToken("program")
	if err != nil {
		return err
	}
	p.commandListNode = NewCommandListNode()
	return p.commandListNode.Parse(context)
}

func (p *ProgramNode) String() string {
	return fmt.Sprintf("[program %s]", p.commandListNode)
}

type RepeatCommandNode struct {
	number          int
	commandListNode Node
}

func NewRepeatCommandNode() *RepeatCommandNode {
	return &RepeatCommandNode{}
}

func (r *RepeatCommandNode) Parse(context *Context) error {
	err := context.SkipToken("repeat")
	if err != nil {
		return err
	}
	num, err := context.CurrentNumber()
	if err != nil {
		return err
	}
	r.number = num
	context.NextToken()
	r.commandListNode = NewCommandListNode()
	return r.commandListNode.Parse(context)
}

func (r *RepeatCommandNode) String() string {
	return fmt.Sprintf("[repeat %d %s]", r.number, r.commandListNode)
}

type CommandListNode struct {
	list []Node
}

func NewCommandListNode() *CommandListNode {
	return &CommandListNode{
		list: []Node{},
	}
}

func (cl *CommandListNode) Parse(context *Context) error {
	for {
		if context.CurrentToken() == "" {
			return &ParseError{Message: "Error: 'end' が見つかりません。"}
		} else if context.CurrentToken() == "end" {
			err := context.SkipToken("end")
			if err != nil {
				return err
			}
			break
		} else {
			commandNode := NewCommandNode()
			err := commandNode.Parse(context)
			if err != nil {
				return err
			}
			cl.list = append(cl.list, commandNode)
		}
	}
	return nil
}

func (cl *CommandListNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, node := range cl.list {
		sb.WriteString(node.String())
		if i < len(cl.list)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

type CommandNode struct {
	node Node
}

func NewCommandNode() *CommandNode {
	return &CommandNode{}
}

func (c *CommandNode) Parse(context *Context) error {
	token := context.CurrentToken()
	if token == "repeat" {
		c.node = NewRepeatCommandNode()
		return c.node.Parse(context)
	} else {
		c.node = NewPrimitiveCommandNode()
		return c.node.Parse(context)
	}
}

func (c *CommandNode) String() string {
	return c.node.String()
}

type PrimitiveCommandNode struct {
	name string
}

func NewPrimitiveCommandNode() *PrimitiveCommandNode {
	return &PrimitiveCommandNode{}
}

func (p *PrimitiveCommandNode) Parse(context *Context) error {
	name := context.CurrentToken()
	if name == "" {
		return &ParseError{Message: "Error: <primitive command> が見つかりません。"}
	} else if name != "go" && name != "right" && name != "left" {
		return &ParseError{Message: fmt.Sprintf("Error: 未知の <primitive command>: '%s'", name)}
	}
	p.name = name
	return context.SkipToken(name)
}

func (p *PrimitiveCommandNode) String() string {
	return p.name
}

func ParseProgram(text string) (*ProgramNode, error) {
	context := NewContext(text)
	programNode := NewProgramNode()
	err := programNode.Parse(context)
	if err != nil {
		return nil, err
	}
	if context.CurrentToken() != "" {
		return nil, &ParseError{Message: fmt.Sprintf("Error: 予期しないトークン: '%s'", context.CurrentToken())}
	}
	return programNode, nil
}
