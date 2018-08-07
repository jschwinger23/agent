package corestore

import (
	pb "github.com/projecteru2/core/rpc/gen"
	"github.com/projecteru2/core/types"
	"golang.org/x/net/context"
)

// GetNode return a node by core
func (c *CoreStore) GetNode(nodename string) (*types.Node, error) {
	client := c.client.GetRPCClient()
	resp, err := client.GetNodeByName(context.Background(), &pb.GetNodeOptions{Nodename: nodename})
	if err != nil {
		return nil, err
	}

	cpus := types.CPUMap{}
	for k, v := range resp.Cpu {
		cpus[k] = int(v)
	}

	node := &types.Node{
		Name:      resp.Name,
		Podname:   resp.Podname,
		Endpoint:  resp.Endpoint,
		Available: resp.Available,
		CPU:       cpus,
		MemCap:    resp.Memory,
	}
	return node, nil
}

// UpdateNode update node status
func (c *CoreStore) UpdateNode(node *types.Node) error {
	client := c.client.GetRPCClient()
	opts := &pb.NodeAvailable{
		Podname:   node.Podname,
		Nodename:  node.Name,
		Available: node.Available,
	}
	_, err := client.SetNodeAvailable(context.Background(), opts)
	return err
}
