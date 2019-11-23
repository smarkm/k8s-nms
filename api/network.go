package api

import (
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//PodNetwork test
func PodNetwork(c *kubernetes.Clientset, ns string) (rs interface{}, err error) {
	data := struct {
		Nodes []Node `json:"nodes"`
		Edges []Edge `json:"edges"`
	}{}
	data.Nodes = make([]Node, 0)
	data.Edges = make([]Edge, 0)
	nodes, err := c.CoreV1().Nodes().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, n := range nodes.Items {
		ip := n.Status.Addresses[0].Address
		data.Nodes = append(data.Nodes, Node{ID: ip, Label: n.Name, IP: ip, Type: "node"})
		fmt.Println()
	}
	pods, err := c.CoreV1().Pods(ns).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, p := range pods.Items {
		data.Nodes = append(data.Nodes, Node{ID: p.Name, Label: p.Name, IP: p.Status.PodIP, Status: p.Status.HostIP, Type: "pod"})
		data.Edges = append(data.Edges, Edge{p.Status.HostIP, p.Name})
	}
	return data, err
}
