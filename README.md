# Pulumi native provider for managing a YouTube playlist
This implementation is meant to play around with the native provider concept in Pulumi. It is not meant to be used in production.  

The repository was generated using the [Pulumi boilerplate](https://github.com/pulumi/pulumi-provider-boilerplate).

## Scratchpad
- To get auth permissions set up correctly:
  - Configure OAuth consent screen on the GCP project (needs to be external if using a brand account, make it a test OAuth config and add the brand account owner's email as a test user)
  - Create credentials for for a `Desktop` application
  - Download the client JSON credentials
  - Authenticate using
     ```shell
     gcloud auth application-default login --scopes=https://www.googleapis.com/auth/youtube.force-ssl --client-id-file=<path/to/creds.json>
     ```
  - Make sure to configure the project for pulumi using `pulumi config set gcp:project <project ID>`
