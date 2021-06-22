package servers

import (
	"context"
	"fmt"
	"k8s-web/src/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretService struct {
	SecretMaps *SecretMap `inject:"-"`
	K8sClient  *kubernetes.Clientset `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

func (se *SecretService) Getall(ns string) []*models.Secrets {
	secretsobj,err := se.SecretMaps.GetALL(ns)
	if err != nil{
		fmt.Println(err)
	}
	SecretsModels := make([]*models.Secrets,len(secretsobj))
	for key,va := range secretsobj {
		SecretsModels[key] = &models.Secrets{
			Name:       va.Name,
			NameSpace:  va.Namespace,
			Type:       models.Secret_Type_Select[string(va.Type)],
			CreateTime: va.CreationTimestamp.String(),
		}
	}
	return SecretsModels
}

func (se *SecretService) DeleteSecret(ns,name string) error {
	return se.K8sClient.CoreV1().Secrets(ns).Delete(context.TODO(),name,metav1.DeleteOptions{})
}