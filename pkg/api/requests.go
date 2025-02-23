package api

import (
	o7tapiauth "github.com/openshift/api/authorization/v1"
	o7tapiquota "github.com/openshift/api/quota/v1"
	o7tapiroute "github.com/openshift/api/route/v1"
	o7tapisecurity "github.com/openshift/api/security/v1"
	o7tapiuser "github.com/openshift/api/user/v1"

	k8sapiapps "k8s.io/api/apps/v1"
	k8sapicore "k8s.io/api/core/v1"
	k8sapistorage "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resources represent api resources used in report
type Resources struct {
	QuotaList            *o7tapiquota.ClusterResourceQuotaList
	NodeList             *k8sapicore.NodeList
	PersistentVolumeList *k8sapicore.PersistentVolumeList
	StorageClassList     *k8sapistorage.StorageClassList
	NamespaceList        []NamespaceResources
	RBACResources        RBACResources
}

// RBACResources contains all resources related to RBAC report
type RBACResources struct {
	UsersList                      *o7tapiuser.UserList
	GroupList                      *o7tapiuser.GroupList
	ClusterRolesList               *o7tapiauth.ClusterRoleList
	ClusterRolesBindingsList       *o7tapiauth.ClusterRoleBindingList
	SecurityContextConstraintsList *o7tapisecurity.SecurityContextConstraintsList
}

// NamespaceResources holds all resources that belong to a namespace
type NamespaceResources struct {
	NamespaceName     string
	DaemonSetList     *k8sapiapps.DaemonSetList
	DeploymentList    *k8sapiapps.DeploymentList
	PodList           *k8sapicore.PodList
	ResourceQuotaList *k8sapicore.ResourceQuotaList
	RolesList         *o7tapiauth.RoleList
	RouteList         *o7tapiroute.RouteList
	PVCList           *k8sapicore.PersistentVolumeClaimList
}

var listOptions metav1.ListOptions

// ListNamespaces list all namespaces, wrapper around client-go
func ListNamespaces() (*k8sapicore.NamespaceList, error) {
	return K8sClient.CoreV1().Namespaces().List(listOptions)
}

// ListPods list all pods in namespace, wrapper around client-go
func ListPods(namespace string) (*k8sapicore.PodList, error) {
	return K8sClient.CoreV1().Pods(namespace).List(listOptions)
}

// ListPVs list all PVs, wrapper around client-go
func ListPVs() (*k8sapicore.PersistentVolumeList, error) {
	return K8sClient.CoreV1().PersistentVolumes().List(listOptions)
}

// ListNodes list all nodes, wrapper around client-go
func ListNodes() (*k8sapicore.NodeList, error) {
	return K8sClient.CoreV1().Nodes().List(listOptions)
}

// ListQuotas list all cluster quotas classes, wrapper around client-go
func ListQuotas() (*o7tapiquota.ClusterResourceQuotaList, error) {
	return O7tClient.quotaClient.ClusterResourceQuotas().List(listOptions)
}

// ListResourceQuotas list all quotas classes, wrapper around client-go
func ListResourceQuotas(namespace string) (*k8sapicore.ResourceQuotaList, error) {
	return K8sClient.CoreV1().ResourceQuotas(namespace).List(listOptions)
}

// ListRoutes list all routes classes, wrapper around client-go
func ListRoutes(namespace string) (*o7tapiroute.RouteList, error) {
	return O7tClient.routeClient.Routes(namespace).List(listOptions)
}

// ListStorageClasses list all storage classes, wrapper around client-go
func ListStorageClasses() (*k8sapistorage.StorageClassList, error) {
	return K8sClient.StorageV1().StorageClasses().List(listOptions)
}

// ListDeployments will list all deployments seeding in the selected namespace
func ListDeployments(namespace string) (*k8sapiapps.DeploymentList, error) {
	return K8sClient.AppsV1().Deployments(namespace).List(listOptions)
}

// ListDaemonSets will collect all DS from specific namespace
func ListDaemonSets(namespace string) (*k8sapiapps.DaemonSetList, error) {
	return K8sClient.AppsV1().DaemonSets(namespace).List(listOptions)
}

// ListUsers list all users, wrapper around client-go
func ListUsers() (*o7tapiuser.UserList, error) {
	return O7tClient.userClient.Users().List(listOptions)
}

// ListGroups list all users, wrapper around client-go
func ListGroups() (*o7tapiuser.GroupList, error) {
	return O7tClient.userClient.Groups().List(listOptions)
}

// ListRoles list all storage classes, wrapper around client-go
func ListRoles(namespace string) (*o7tapiauth.RoleList, error) {
	return O7tClient.authClient.Roles(namespace).List(listOptions)
}

// ListClusterRoles list all storage classes, wrapper around client-go
func ListClusterRoles() (*o7tapiauth.ClusterRoleList, error) {
	return O7tClient.authClient.ClusterRoles().List(listOptions)
}

// ListClusterRolesBindings list all storage classes, wrapper around client-go
func ListClusterRolesBindings() (*o7tapiauth.ClusterRoleBindingList, error) {
	return O7tClient.authClient.ClusterRoleBindings().List(listOptions)
}

// ListSCC list all security context constraints, wrapper around client-go
func ListSCC() (*o7tapisecurity.SecurityContextConstraintsList, error) {
	return O7tClient.securityClient.SecurityContextConstraints().List(listOptions)
}

// ListPVCs list all PVs, wrapper around client-go
func ListPVCs(namespace string) (*k8sapicore.PersistentVolumeClaimList, error) {
	return K8sClient.CoreV1().PersistentVolumeClaims(namespace).List(listOptions)
}
