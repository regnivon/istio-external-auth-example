# istio-external-auth-example
You need to change all images used in deployments to your own repos

Steps

1. create and label namespace

kubectl create ns test
kubectl label ns test istio-injection=enabled

2. edit the istio config to hook in the auth extension

kubectl edit configmap istio -n istio-system
```
data:
  mesh: |-  
    # Add the following content to define the external authorizers.  
    extensionProviders:  
    - name: "sample-ext-authz-grpc"  
      envoyExtAuthzGrpc:  
        service: "ext-authz.foo.svc.cluster.local" 
        port: "9000" 
    - name: "sample-ext-authz-http"  
      envoyExtAuthzHttp:  
        service: "ext-authz.foo.svc.cluster.local"  
        port: "8000"  
        includeHeadersInCheck: ["x-ext-authz"]        
```

3. apply resources (don't bother with deploy and service files if you don't want to test regular http routing with istio)

kubectl apply -f kubernetes/

4. build client and run

go build cmd/client.go
./client 

true


5. If you see true and not an error then you have succeeded. You can test it failing if you change the x-ext-authz header in the client to deny rather than allow.
