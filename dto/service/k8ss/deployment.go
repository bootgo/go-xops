package k8ss

import (
	"context"
	"fmt"
	"go-xops/models/k8sm"
	"go-xops/pkg/common"
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetDeployments(ctx context.Context, namespace, label string) ([]k8sm.Deployment, error) {
	result := make([]k8sm.Deployment, 0)
	data, err := common.ClientSet.AppsV1().Deployments(namespace).List(ctx, v1.ListOptions{
		LabelSelector: label,
	})

	if err != nil {
		return result, err
	}

	for _, deployment := range data.Items {
		result = append(result, k8sm.Deployment{
			Name: deployment.ObjectMeta.Name,
			UID:  string(deployment.ObjectMeta.UID),
		})
	}

	return result, nil
}

func GetDeployment(ctx context.Context, namespace, name string) (k8sm.Deployment, error) {
	var deployment k8sm.Deployment
	data, err := common.ClientSet.AppsV1().Deployments(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return deployment, err
	}
	deployment.Name = data.ObjectMeta.Name
	deployment.UID = string(data.ObjectMeta.UID)
	return deployment, nil
}

func UpdateDeployment(ctx context.Context, namespace, name, data string) (bool, error) {
	_, err := common.ClientSet.AppsV1().Deployments(namespace).Patch(
		ctx,
		name,
		types.JSONPatchType,
		[]byte(data),
		v1.PatchOptions{},
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetDeploymentStatus(ctx context.Context, namespace, name string, limit int) (bool, error) {
	// 等待直到k8s选择deployment
	time.Sleep(10 * time.Second)

	deployment, err := common.ClientSet.AppsV1().Deployments(namespace).Get(ctx, name, v1.GetOptions{})

	if err != nil {
		return false, err
	}

	status := true

	for i := 0; i < limit; i++ {
		status = true

		deployment, err = common.ClientSet.AppsV1().Deployments(namespace).Get(ctx, name, v1.GetOptions{})

		if err != nil {
			return false, err
		}

		if int(deployment.Generation) != int(deployment.Status.ObservedGeneration) {
			status = false
		}

		if int(deployment.Status.UnavailableReplicas) > 0 {
			status = false
		}

		if int(int32(*deployment.Spec.Replicas)) != int(deployment.Status.AvailableReplicas) {
			status = false
		}

		if !status {
			log.Println("Deployment Success")
			time.Sleep(2 * time.Second)
		} else {
			log.Println("Deployment Success")
			return true, nil
		}
	}

	log.Println("Deployment success")

	return false, fmt.Errorf(fmt.Sprintf(
		"Deployment %s failed: namespace %s, Generation %d, ObservedGeneration %d,"+
			" UnavailableReplicas %d, Replicas %d, AvailableReplicas %d",
		name,
		namespace,
		int(deployment.Generation),
		int(deployment.Status.ObservedGeneration),
		int(deployment.Status.UnavailableReplicas),
		int(int32(*deployment.Spec.Replicas)),
		int(deployment.Status.AvailableReplicas),
	))
}