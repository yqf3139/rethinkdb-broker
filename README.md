# RethinkDB Broker

This is an implementation of a Service Broker that uses Helm to provision
instances of [RethinkDB](https://kubeapps.com/charts/stable/rethinkdb). This is a
**proof-of-concept** for the [Kubernetes Service
Catalog](https://github.com/kubernetes-incubator/service-catalog), and should not
be used in production. Thanks to the [mariadb broker repo](https://github.com/prydonius/mariadb-broker).

## Prerequisites

1. Kubernetes cluster
2. [Helm 2.x](https://github.com/kubernetes/helm)
3. [Service Catalog API](https://github.com/kubernetes-incubator/service-catalog) - follow the [walkthrough](https://github.com/kubernetes-incubator/service-catalog/blob/master/docs/walkthrough.md)

## Installing the Broker

The RethinkDB Service Broker can be installed using the Helm chart in this
repository.

```
$ git clone https://github.com/yqf3139/rethinkdb-broker.git
$ cd rethinkdb-broker
$ helm install --name rethinkdb-broker --namespace rethinkdb-broker charts/rethinkdb-broker
```

To register the Broker with the Service Catalog, create the Broker object:

```
$ kubectl --context service-catalog create -f examples/rethinkdb-broker.yaml
```

If the Broker was successfully registered, the `rethinkdb` ServiceClass will now
be available in the catalog:

```
$ kubectl --context service-catalog get serviceclasses
NAME      KIND
rethinkdb   ServiceClass.v1alpha1.servicecatalog.k8s.io
```

## Usage

### Create the Instance object

```
$ kubectl --context service-catalog create -f examples/rethinkdb-instance.yaml
```

This will result in the installation of a new RethinkDB chart:

```
$ helm list
NAME                                  	REVISION	UPDATED                 	STATUS  	CHART               	NAMESPACE
i-3e0e9973-a072-49ba-8308-19568e7f4669	1       	Sat May 13 17:28:35 2017	DEPLOYED	rethinkdb-0.6.1       	3e0e9973-a072-49ba-8308-19568e7f4669
```

### Create a Binding to fetch credentials

```
$ kubectl --context service-catalog create -f examples/rethinkdb-binding.yaml
```

A secret called `rethinkdb-instance-credentials` will be created containing the
connection details for this RethinkDB instance.

```
$ kubectl get secret rethinkdb-instance-credentials -o yaml
```
