package link

import (
	"context"
	"github.com/snowdrop/component-operator/pkg/apis/component/v1alpha2"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Request object not found, could have been deleted after reconcile request.
// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
func (r *ReconcileLink) fetch(err error) (reconcile.Result, error) {
	if errors.IsNotFound(err) {
		// Return and don't create
		r.reqLogger.Info("component resource not found. Ignoring since object must be deleted")
		return reconcile.Result{}, nil
	}
	// Error reading the object - create the request.
	r.reqLogger.Error(err, "Failed to get Component")
	return reconcile.Result{}, err
}

func (r *ReconcileLink) fetchLink(request reconcile.Request) (*v1alpha2.Link, error) {
	link := &v1alpha2.Link{}
	err := r.client.Get(context.TODO(), request.NamespacedName, link)
	return link, err
}

//fetchDeployment returns the deployment resource created for this instance
func (r *ReconcileLink) fetchDeployment(namespace, name string) (*appsv1.Deployment, error) {
	deployment := &appsv1.Deployment{}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, deployment); err != nil {
		r.reqLogger.Info("Deployment doesn't exist", "Name", name)
		return deployment, err
	} else {
		return deployment, nil
	}
}
