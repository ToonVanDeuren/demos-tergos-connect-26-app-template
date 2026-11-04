# An app on the platform

The platform created this repository when the app was requested. Publish a release to
deploy it: **Releases > Draft a new release**, pick a tag such as `v0.1.0`, **Publish**.
The **Actions** tab shows the run, and the `production` environment links to the app.

| Yours | The platform's |
| --- | --- |
| the code (`main.go`) | the `production` environment and `DEPLOY_CONFIG` |
| the Deployment (`deploy/deployment.yaml`) | the identity the pipeline signs in as |
| | the namespace, the Service and the route |

There are no secrets in this repository. The pipeline gets a short-lived token because
it runs in this repository's `production` environment, and that token can only push
images and deploy to this app's namespace.
