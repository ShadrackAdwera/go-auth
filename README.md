# go-auth

- Login flow using Auth0.
- Auth Middleware on protected route.

## Set up
- Create .env file on the root project.
- Add the config below

AUTH0_CLIENT_ID=YOUR_CLIENT_ID
AUTH0_DOMAIN=YOUR_DOMAIN
AUTH0_CLIENT_SECRET=YOUR_CLIENT_SECRET
AUTH0_CALLBACK_URL=YOUR_CONFIGURED_CALLBACK_URLS

Create a user on your app:

Run `go mod tidy` to download the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/login](http://localhost:3000/login).

### Next steps

- Add Create Organisation config - to set up multiple tenants each with their own brand customisation.
- Add / invite users to an organisation.
- Login to an organisation

