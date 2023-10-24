# buuurst_dev_go

Golang middleware for Buuurst.Dev

## Installation

```
go get git.drecom.jp/diet/buuurst_dev_go
```

## Usage

### go-chi

```golang
import (
    ...

    buuurst_dev "git.drecom.jp/diet/buuurst_dev_go"
)

func main() {
    r := chi.NewRouter()

    // Add buuurst_dev middleware
	r.Use(buuurst_dev.MiddlewareFunc(
        // buuurst_dev collector configuration
		&buuurst_dev.BuuurstDevConfig{
			Enabled:      true, // Make collector enable/disable
			CollectorURL: "https://lambda-public.buuurst.dev/put-request-log",
			ProjectID:    YOUR_PROJECT_ID, // Set your project ID
			ServiceKey:   "YOUR_SERVICE_KEY", // Set your service key
			CustomHeaders: []string{ // Set headers that should be sent to buuurst.dev
				"Authorization",
			},
			IgnorePaths: []string{ // Set ignored paths
				"/ignored",
			},
		},
	))

    ....
}
```

### echo

```golang
import (
    ...

    buuurst_dev "git.drecom.jp/diet/buuurst_dev_go"
)

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(buuurst_dev.MiddlewareFunc(
		&buuurst_dev.BuuurstDevConfig{
			Enabled:      true,
			CollectorURL: "https://lambda-public.buuurst.dev/put-request-log",
			ProjectID:    YOUR_PROJECT_ID,
			ServiceKey:   "YOUR_SERVICE_KEY",
			CustomHeaders: []string{
				"Authorization",
			},
			IgnorePaths: []string{
				"/ignored",
			},
		},
	)))
    ...
}
```