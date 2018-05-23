FROM bitnami/minideb:latest
COPY ./rethinkdb-broker /rethinkdb-broker
ADD https://kubernetes-charts.storage.googleapis.com/rethinkdb-0.0.2.tgz /rethinkdb-chart.tgz
CMD ["/rethinkdb-broker", "-logtostderr"]
