/*
Copyright © 2019 Thorsten Klein <iwilltry42@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cluster

import (
	"fmt"

	k3drt "github.com/rancher/k3d/pkg/runtimes"
	k3d "github.com/rancher/k3d/pkg/types"
	log "github.com/sirupsen/logrus"
)

// CreateCluster creates a new cluster consisting of
// - some containerized k3s nodes
// - a docker network
func CreateCluster(cluster *k3d.Cluster, runtime k3drt.Runtime) error {

	// used for node suffices
	masterCount := 0
	workerCount := 0

	for _, node := range cluster.Nodes {

		// cluster specific settings
		node.Labels = make(map[string]string)
		node.Labels["cluster"] = cluster.Name

		// node role specific settings
		suffix := 0
		if node.Role == "master" {
			// name suffix
			suffix = masterCount
			masterCount++
		} else if node.Role == "worker" {
			// name suffix
			suffix = workerCount
			workerCount++
		}

		node.Name = fmt.Sprintf("%s-%s-%s-%d", k3d.DefaultObjectNamePrefix, cluster.Name, node.Role, suffix)

		// create node
		log.Infoln("Creating node", node.Name)
		if err := CreateNode(&node, runtime); err != nil {
			log.Errorln("...failed")
		}
	}

	log.Debugln("...success")
	return nil
}

// DeleteCluster deletes an existing cluster
func DeleteCluster(cluster *k3d.Cluster, runtime k3drt.Runtime) error {
	return nil
}