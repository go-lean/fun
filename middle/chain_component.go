/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package middle

type ChainComponent struct {
	chain *Chain
}

func NewChainComponent(steps ...Step) *ChainComponent {

	return &ChainComponent{
		chain: New(steps...),
	}
}

func (c *ChainComponent) Use(steps ...Step) *ChainComponent {

	c.chain.Add(steps...)
	return c
}

func (c *ChainComponent) UseChain(chain *Chain) *ChainComponent {

	c.chain.Merge(chain)
	return c
}

func (c *ChainComponent) CloneChain() *Chain {
	return c.chain.Clone()
}

func (c *ChainComponent) Chain() *Chain {
	return c.chain
}
