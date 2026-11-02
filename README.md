# demos-tergos-connect-26-app-template

Template for applications on the demo platform. Do not use it directly: request
an `App` and the platform creates a repository from this template.

Publish a release to build the image and deploy it:

```bash
gh release create v0.1.0 --generate-notes
```

The application lives in `main.go`, its Deployment in `deploy/deployment.yaml`.
