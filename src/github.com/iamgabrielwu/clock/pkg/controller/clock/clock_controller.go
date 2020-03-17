package clock

import (
	"context"
	"strconv"

	"fmt"

	clockv1 "github.com/iamgabrielwu/clock/pkg/apis/clock/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_clock")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Clock Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileClock{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("clock-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Clock
	err = c.Watch(&source.Kind{Type: &clockv1.Clock{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Clock
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &clockv1.Clock{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileClock implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileClock{}

// ReconcileClock reconciles a Clock object
type ReconcileClock struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Clock object and makes changes based on the state read
// and what is in the Clock.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileClock) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Clock")

	// Fetch the Clock instance
	instance := &clockv1.Clock{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("object not found or garbage collected ")

			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		fmt.Println(err)
		reqLogger.Info("object found with err ")
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	ck := newClockForCR(instance)

	// Set Clock instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, ck, r.scheme); err != nil {
		fmt.Println("Set Clock instance as the owner and controller")
		reqLogger.Info("Set Clock instance as the owner and controller ")
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &clockv1.Clock{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: ck.Name, Namespace: ck.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Clock", "Clock.Namespace", ck.Namespace, "Clock.Name", ck.Name)
		err = r.client.Create(context.TODO(), ck)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		fmt.Println("Set Clock created successfully - don't requeue")
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Clock already exists", "Clock.Namespace", found.Namespace, "Clock.Name", found.Name)
	return reconcile.Result{}, nil
}

// newClockForCR returns a busybox pod with the same name/namespace as the cr
func newClockForCR(cr *clockv1.Clock) *clockv1.Clock {
	labels := map[string]string{
		"app":   cr.Name,
		"reset": strconv.FormatBool(cr.Spec.Reset),
	}
	fmt.Println("cr Reset", cr.Spec.Reset)
	if cr.Spec.Reset == true {
		fmt.Println("Resetting Clock, not matter what current date/time is")
		cr.Spec.CurrentDay = "1970-01-01"
		cr.Spec.CurrentTime = "00:00"
	}
	return &clockv1.Clock{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: clockv1.ClockSpec{
			Reset:       cr.Spec.Reset,
			CurrentDay:  cr.Spec.CurrentDay,
			CurrentTime: cr.Spec.CurrentTime,
		},
	}
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
// func newPodForCR(cr *clockv1.Clock) *corev1.Pod {
// 	labels := map[string]string{
// 		"app": cr.Name,
// 	}
// 	return &corev1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      cr.Name + "-pod",
// 			Namespace: cr.Namespace,
// 			Labels:    labels,
// 		},
// 		Spec: corev1.PodSpec{
// 			Containers: []corev1.Container{
// 				{
// 					Name:    "busybox",
// 					Image:   "busybox",
// 					Command: []string{"sleep", "3600"},
// 				},
// 			},
// 		},
// 	}
// }
