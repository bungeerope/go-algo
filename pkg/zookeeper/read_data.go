package zookeeper

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"log"
	"time"
)

type dirTreeNode struct {
	name  string
	child []dirTreeNode
}

type ConnClient struct {
	Client *zk.Conn
}

type ConnectOption struct {
	Server         []string
	SessionTimeout time.Duration
}

func InitClient() *ConnClient {

	var connectOption = &ConnectOption{
		Server:         []string{"172.16.251.174:2181"},
		SessionTimeout: 3 * time.Second,
	}
	conn, _, err := zk.Connect(connectOption.Server, connectOption.SessionTimeout)
	if err != nil {
		fmt.Errorf("zk init connect err: %v", err)
	}

	return &ConnClient{
		Client: conn,
	}
}

func (conn *ConnClient) GetPath(path string) (data []byte, stat *zk.Stat, err error) {
	data, stat, err = conn.Client.Get(path)
	return
}

func printDirTree(tree dirTreeNode, prefix string) {
	fmt.Println(prefix + tree.name)
	if len(tree.child) > 0 {
		prefix += "----"
		for _, childNode := range tree.child {
			printDirTree(childNode, prefix)
		}
	}
}

func (conn *ConnClient) GetAllPath() {
	path := "/"
	tree, err := getDirTree(conn, path)
	if err != nil {
		log.Fatalln("read dir '%s' fail: %v", path, err)
	}
	fmt.Println("result is :\n")
	printDirTree(tree, "")
}

func getDirTree(conn *ConnClient, path string) (dirTreeNode, error) {
	strs, err := conn.GetChildren(path)
	if err != nil {
		log.Fatalf("Read dir '%s' failed: %v", path, err)
	}
	var tree, childNode dirTreeNode
	tree.name = path
	var fullName string
	for _, name := range strs {
		if "/" == path {
			fullName = path + name
		} else {
			fullName = path + "/" + name
		}

		strs2, err := conn.GetChildren(fullName)
		if err == nil && len(strs2) > 0 {
			childNode, err = getDirTree(conn, fullName)
			if err != nil {
				log.Fatalf("Read dir '%s' failed: %v", fullName, err)
			}
		} else {
			childNode.name = name
			childNode.child = nil
		}
		tree.child = append(tree.child, childNode)
	}
	return tree, nil
}

func (conn *ConnClient) GetChildren(path string) ([]string, error) {
	v, _, err := conn.Client.Children(path)
	if err != nil {
		fmt.Errorf("zk init connect err: %v", err)
	}
	fmt.Printf("value of path[%s]=[%s].\n", path, v)
	return v, err
}
