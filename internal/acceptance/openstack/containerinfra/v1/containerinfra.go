package v1

import (
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/containerinfra/v1/clustertemplates"
	"github.com/gophercloud/gophercloud/v2/openstack/containerinfra/v1/quotas"
)

// CreateClusterTemplateCOE will create a random cluster template for the specified orchestration engine.
// An error will be returned if the cluster template could not be created.
func CreateClusterTemplateCOE(t *testing.T, client *gophercloud.ServiceClient, coe string) (*clustertemplates.ClusterTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// workaround for https://bugs.launchpad.net/magnum/+bug/2109685

// CreateClusterTemplate will create a random swarm cluster template.
// An error will be returned if the cluster template could not be created.
func CreateClusterTemplate(t *testing.T, client *gophercloud.ServiceClient) (*clustertemplates.ClusterTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateKubernetesClusterTemplate will create a random kubernetes cluster template.
// An error will be returned if the cluster template could not be created.
func CreateKubernetesClusterTemplate(t *testing.T, client *gophercloud.ServiceClient) (*clustertemplates.ClusterTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteClusterTemplate will delete a given cluster-template. A fatal error will occur if the
// cluster-template could not be deleted. This works best as a deferred function.
func DeleteClusterTemplate(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

// CreateClusterTimeout will create a random cluster and wait for it to reach CREATE_COMPLETE status
// within the given timeout duration. An error will be returned if the cluster could not be created.
func CreateClusterTimeout(t *testing.T, client *gophercloud.ServiceClient, clusterTemplateID string, timeout time.Duration) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// createTimeout is the creation timeout on the magnum side in minutes

// CreateCluster will create a random cluster. An error will be returned if the
// cluster could not be created. Has a timeout of 300 seconds.
func CreateCluster(t *testing.T, client *gophercloud.ServiceClient, clusterTemplateID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CreateKubernetesCluster is the same as CreateCluster with a longer timeout necessary for creating a kubernetes cluster
func CreateKubernetesCluster(t *testing.T, client *gophercloud.ServiceClient, clusterTemplateID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func DeleteCluster(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

func WaitForCluster(client *gophercloud.ServiceClient, clusterID string, status string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateQuota will create a random quota. An error will be returned if the
// quota could not be created.
func CreateQuota(t *testing.T, client *gophercloud.ServiceClient) (*quotas.Quotas, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
