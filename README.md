# MFA AUTH

Used to create MFA Authenticated Credentials
By Default Credentials are valid for 12 Hours.

### Install

#### Go Installer
```
go install github.com/rajeshg007/mfa-auth
```

#### Homebrew
```
brew tap rajeshg007/tap
brew install mfa-auth
```

### Usage

Use `mfa-auth init` to configure your credentials in mfa-auth

Use `mfa-auth <token>` to Authenticate using MFA token.

You can also use `mfa-auth --help` to see all the available commands.

Note: that the Authenticated credentials are stored in the Default AWS Profile.

