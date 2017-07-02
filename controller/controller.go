package controller

import (
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
	"github.com/yqf3139/rethinkdb-broker/client"
)

type errNoSuchInstance struct {
	instanceID string
}

func (e errNoSuchInstance) Error() string {
	return fmt.Sprintf("no such instance with ID %s", e.instanceID)
}

type rethinkdbController struct {
}

// CreateController creates an instance of a User Provided service broker controller.
func CreateController() controller.Controller {
	return &rethinkdbController{}
}

func (c *rethinkdbController) Catalog() (*brokerapi.Catalog, error) {
	return &brokerapi.Catalog{
		Services: []*brokerapi.Service{
			{
				Name:        "rethinkdb",
				ID:          "10faf2e0-5e61-11e7-9c17-fb844ec31790",
				Description: "rethinkdb database",
				Plans: []brokerapi.ServicePlan{
					{
						Name:        "default",
						ID:          "18413dde-5e61-11e7-93e3-179ef79e4573",
						Description: "rethinkdb database",
						Free:        true,
					},
				},
				Bindable: true,
			},
		},
	}, nil
}

func (c *rethinkdbController) CreateServiceInstance(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {
	if err := client.Install(releaseName(id), id); err != nil {
		return nil, err
	}
	glog.Infof("Created rethinkdb Service Instance:\n%v\n", id)
	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

func (c *rethinkdbController) GetServiceInstance(id string) (string, error) {
	return "", errors.New("Unimplemented")
}

func (c *rethinkdbController) RemoveServiceInstance(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	if err := client.Delete(releaseName(id)); err != nil {
		return nil, err
	}
	return &brokerapi.DeleteServiceInstanceResponse{}, nil
}

func (c *rethinkdbController) Bind(instanceID, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	host := releaseName(instanceID) + "-rethinkdb-proxy." + instanceID + ".svc.cluster.local"
	password, err := client.GetPassword(releaseName(instanceID), instanceID)
	if err != nil {
		return nil, err
	}
	return &brokerapi.CreateServiceBindingResponse{
		Credentials: brokerapi.Credential{
			"host": 	   host,
			"password.admin":  password,
			"ports.driver":    "28015",
			"ports.admin":     "8080",
		},
	}, nil
}

func (c *rethinkdbController) UnBind(instanceID string, bindingID string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}

func releaseName(id string) string {
	return "i-" + id
}
