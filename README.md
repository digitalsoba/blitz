# blitz

## Building the project

You can build the project using make. You can choose a Docker image or plain binaries.
* Run the following command to build the image: `make build-docker`
* Run the following command to build the binaries: `make build`

## Deploying the project
Deploy this project on any Kubernetes cluster running Nginx Ingress and kustomize. If you're using another ingress controller update the ingress.yaml manifest to fit your provider. In a perfect world this would be linked with ArgoCD or similar tools to deploy once requierments are met (merge-request, CI, etc). For now, you can use the following command: `make deploy` to deploy into any cluster running Kubernetes v1.21 or higher.
## Developer Reference

### Local development
Note: You'll need Go v1.17 and Docker to run this application locally. Run `make dev` to start the complied development server via docker-compose. You should run `make dev` again if you modify the `hooks/` or `app.go` files. 

Two containers are started:
* `blitz-webhook-1`: runs the webhook + server logic
* `blitz-client-1`: runs the client with curl installed

Test webhook requests locally with curl or within the client container. To enter the client container, run `docker exec -it blitz-client-1 sh` or `make exec`. Sample request below. If you're using curl locally or any rest client swap `blitz-client` with `localhost`.

```shell
curl --request POST \
  --url http://blitz-webhook:9000/hooks/webhook \
  --header 'Authorization: Bearer auth-token-123' \
  --header 'Content-Type: application/json' \
  --data '{
	"filename": "test.txt",
	"content": "We'\''re no strangers to love, you know the rules and so do I"
}'
```

You can also run `make test` to run go test locally.

## Blitz requirements

Use docker compose to test a webhook tool. (https://github.com/adnanh/webhook) We would like to test the above tool. Use docker compose to spin up a test. The test
should qualify the below acceptance criteria:
* Spin up a the golang code and configure the settings to bare minimum to meet rest of requirements.
* The webhook should take the name of a text file and text to append into the text file.
* It should write the file with the contents to storage that persists through docker compose up/down.
* To fully test the webhook the sender hitting the endpoint should be from outside the webhook container but not localhost, it should test and validate webhook works through docker compose networking.
* Webhook should also have a secure way of validating the source of the request and not accept just any service hitting the endpoint.
* Any commands needing to be run should be scripted with bash or something similarly simple.
* Anyone using your repo should only be invoking scripts not running cmds directly.
* The readme should provide instructions on how to spin up your solution and test it.
* The readme should contain a brief explanation of how to move this to your cloud
provider of choice, not an actual implementation just an English explanation of going from test/POC to production.
Create a public repo on github with your solution and send it in.