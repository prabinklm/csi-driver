// Copyright 2019 Hewlett Packard Enterprise Development LP

package vanilla

import (
	"encoding/json"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/hpe-storage/common-host-libs/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Flavor of the CSI driver
type Flavor struct {
}

// ConfigureAnnotations does nothing in the vanilla case
func (flavor *Flavor) ConfigureAnnotations(claimName string, parameters map[string]string) (map[string]string, error) {
	return parameters, nil
}

// LoadNodeInfo loads a Node object into JSON in the vanilla case.  Note this might hit the 128 character limit
// in some other CO.  A new flavor will likely need to be added in that case.
func (flavor *Flavor) LoadNodeInfo(node *model.Node) (string, error) {
	nodeBytes, err := json.Marshal(node)
	if err != nil {
		return "", fmt.Errorf("Could not marshal Node struct")
	}
	return string(nodeBytes), nil
}

// UnloadNodeInfo does nothing in the vanilla case
func (flavor *Flavor) UnloadNodeInfo() {
}

// GetNodeInfo returns the node details required to publish a volume to a node
func (flavor *Flavor) GetNodeInfo(nodeID string) (*model.Node, error) {
	var node *model.Node
	err := json.Unmarshal([]byte(nodeID), &node)

	return node, err
}

// GetCredentialsFromPodSpec :
func (flavor *Flavor) GetCredentialsFromPodSpec(volumeHandle string, podName string, namespace string) (map[string]string, error) {
	return nil, nil
}

// GetCredentialsFromSecret :
func (flavor *Flavor) GetCredentialsFromSecret(name string, namespace string) (map[string]string, error) {
	return nil, nil
}

func (flavor *Flavor) CreateNFSVolume(pvName string, reqVolSize int64, parameters map[string]string) (nfsVolume *csi.Volume, rollback bool, err error) {
	return nil, false, fmt.Errorf("NFS provisioned volume is not supported for non-k8s environments")
}

func (flavor *Flavor) DeleteNFSVolume(pvName string) error {
	return fmt.Errorf("NFS provisioned volume is not supported for non-k8s environments")
}

func (flavor *Flavor) HandleNFSNodePublish(request *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return nil, status.Error(codes.Internal, "NFS provisioned volume is not supported for non-k8s environments")
}

func (flavor *Flavor) IsNFSVolume(volumeID string) bool {
	return false
}
