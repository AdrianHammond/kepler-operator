package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type keplerSADescription struct {
	Context     context.Context
	Client      client.Client
	Scheme      *runtime.Scheme
	SA          *corev1.ServiceAccount
	Owner       metav1.Object
	role        *rbacv1.Role
	roleBinding *rbacv1.RoleBinding
}

func (d *keplerSADescription) Reconcile(l klog.Logger) (bool, error) {
	return reconcileBatch(l,
		d.ensureSA,
		d.ensureRole,
		d.ensureRoleBinding,
	)
}

func (d *keplerSADescription) ensureSA(l klog.Logger) (bool, error) {
	logger := l.WithValues("ServiceAccount", nameFor(d.SA))
	op, err := ctrlutil.CreateOrUpdate(d.Context, d.Client, d.SA, func() error {
		if err := ctrl.SetControllerReference(d.Owner, d.SA, d.Scheme); err != nil {
			logger.Error(err, "unable to set controller reference")
			return err
		}
		return nil
	})
	if err != nil {
		logger.Error(err, "ServiceAccount reconcile failed")
		return false, err
	}

	logger.V(1).Info("ServiceAccount reconciled", "operation", op)
	return true, nil
}

func (d *keplerSADescription) ensureRole(l klog.Logger) (bool, error) {
	d.role = &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind: "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.SA.Name,
			Namespace: d.SA.Namespace,
		},
	}
	logger := l.WithValues("Role", nameFor(d.role))
	op, err := ctrlutil.CreateOrUpdate(d.Context, d.Client, d.role, func() error {
		if err := ctrl.SetControllerReference(d.Owner, d.role, d.Scheme); err != nil {
			logger.Error(err, "unable to set controller reference")
			return err
		}
		d.role.Rules = []rbacv1.PolicyRule{
			{
				APIGroups: []string{"security.openshift.io"},
				Resources: []string{"securitycontextconstraints"},
				// Must match the name of the SCC that is deployed w/ the operator
				ResourceNames: []string{SCCName},
				Verbs:         []string{"use"},
			},
		}
		return nil
	})
	if err != nil {
		logger.Error(err, "Role reconcile failed")
		return false, err
	}

	logger.V(1).Info("Role reconciled", "operation", op)
	return true, nil
}

func (d *keplerSADescription) ensureRoleBinding(l klog.Logger) (bool, error) {
	d.roleBinding = &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind: "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.SA.Name,
			Namespace: d.SA.Namespace,
		},
	}
	logger := l.WithValues("RoleBinding", nameFor(d.roleBinding))
	op, err := ctrlutil.CreateOrUpdate(d.Context, d.Client, d.roleBinding, func() error {
		if err := ctrl.SetControllerReference(d.Owner, d.roleBinding, d.Scheme); err != nil {
			logger.Error(err, "unable to set controller reference")
			return err
		}
		d.roleBinding.RoleRef = rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     d.role.Name,
		}
		d.roleBinding.Subjects = []rbacv1.Subject{
			{Kind: "ServiceAccount", Name: d.SA.Name},
		}
		return nil
	})
	if err != nil {
		logger.Error(err, "RoleBinding reconcile failed")
		return false, err
	}

	logger.V(1).Info("RoleBinding reconciled", "operation", op)
	return true, nil
}
