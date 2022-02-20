FROM envoyproxy/envoy-dev:3ce38d696c2f4b080b14a40512ab571f42ac6c85
COPY ./envoy/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml