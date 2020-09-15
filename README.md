# AppEngine Proxy Server

A simple proxy service which adds its service account ID token to the request.

Designed for use in conjunction with IAP and Cloud Run.

## Usage

Depending on your previous actions, some steps can be skipped.

1. Create an Oauth2 Client ID in the console ([link](https://console.cloud.google.com/apis/credentials)).
1. Deploy the real frontend to Cloud Run with flag `no-allow-unauthenticated`
2. Write down the reported URL - this is the proxy target scheme.
3. Open `app.yaml`
4. Fill `TARGET` with proxy target scheme (i.e. `https://www.example.com`)
5. Fill `AUDIENCE` with your Oauth2 Client ID
6. Deploy the AppEngine: `gcloud app deploy`
7. Activate and configure [Identity Aware Proxy](https://cloud.google.com/iap/docs)


## License
Copyright 2020 Hayo van Loon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
