Getting started with OAuth2 in Go
Authentication usually is very important part in any application. You can always implement your own authentication system, but it will require a lot of work, registration, forgot password form, etc. That's why OAuth2 was created, to allow user to log in using one of the many accounts user already has.

In this video we'll create a simple web page with Google login using oauth2 Go package.

Google Project, OAuth2 keys
First of all, let's create our Google OAuth2 keys.

Go to Google Cloud Platform
Create new project or use an existing one
Go to Credentials
Click "Create credentials"
Choose "OAuth client ID"
Add authorized redirect URL, in our case it will be localhost:8080/callback
Get client id and client secret
Save it in a safe place
How OAuth2 works with Google
Obtain OAuth 2.0 credentials from the Google API Console.
Obtain an access token from the Google Authorization Server.
Send the access token to an API.
Refresh the access token, if necessary.
Structure
We'll do everything in 1 main.go file, and register 3 URL handlers:

/
/login
/welcome
Initial handlers and OAuth2 config
go get golang.org/x/oauth2
