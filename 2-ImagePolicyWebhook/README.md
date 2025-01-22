** Build and deplooy Webhook APP **
```
docker build . -t ttl.sh/image-webhook:1h
```

** Generate certificat for webhook **
```

openssl genrsa -out tls.key 2048
openssl req -new -key  tls.key -out /etc/kubernetes/pki/webhook-server.csr -config webhook-server.conf
openssl x509 -req -in webhook-server.csr -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial -out tls.crt -days 365 -extfile webhook-server.conf -extensions v3_req
```

create tls secret

```
kubectl create secret tls webhook-tls-secret --cert tls.crt --key tls.key
```