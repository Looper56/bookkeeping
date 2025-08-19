package util

import (
	"errors"
	"fmt"
	"net"

	"github.com/bwmarrin/snowflake"
)

var maker *labelMaker

func init() {
	maker = NewLabelMaker()
}

// UniqueID 生成唯一ID
func UniqueID() string {
	return maker.uniqueID()
}

// labelMaker 服务标识生成
type labelMaker struct {
	node *snowflake.Node
}

// NewLabelMaker new
func NewLabelMaker() *labelMaker {
	maker := &labelMaker{}
	_ = maker.initNode()
	return maker
}

// initNode 初始化节点
func (l *labelMaker) initNode() error {
	ipInt, err := l.getLocalIP()
	if err != nil {
		return err
	}
	nodeNum := l.getNodeNum(ipInt)
	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		return err
	}
	l.node = node
	return nil
}

// UniqueID 生成唯一id,雪花算法
func (l *labelMaker) uniqueID() string {
	id := l.node.Generate()
	return id.Base58()
}

// getNodeNum 获取节点id，node节点占10bit，共1024个
func (l *labelMaker) getNodeNum(seed uint) int64 {
	return int64(seed % 1024)
}

// getLocalIP 获取本地ip
func (l *labelMaker) getLocalIP() (uint, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return 0, err
	}
	for _, address := range addr {
		// 检查ip地址判断是否回环地址
		ipNet, ok := address.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}
		ip4 := ipNet.IP.To4()
		if ip4 == nil {
			continue
		}
		fmt.Printf("[info] local ip: %s\n", ipNet.IP.String())
		return uint(ip4[3]) |
			uint(ip4[2])<<8 |
			uint(ip4[1])<<16 |
			uint(ip4[0])<<24, nil
	}
	return 0, errors.New("do not have local ip")
}
