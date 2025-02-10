module github.com/raito-io/sdk-go

go 1.23.0

toolchain go1.23.3

require (
	github.com/Khan/genqlient v0.8.0
	github.com/agnivade/levenshtein v1.2.1
	github.com/alexflint/go-arg v1.5.1
	github.com/alexflint/go-scalar v1.2.0
	github.com/aws/aws-sdk-go-v2 v1.36.1
	github.com/aws/aws-sdk-go-v2/config v1.29.6
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.49.4
	github.com/aws/smithy-go v1.22.2
	github.com/raito-io/enumer v0.1.6
	github.com/stretchr/testify v1.10.0
	golang.org/x/tools v0.30.0
)

replace github.com/Khan/genqlient v0.7.0 => github.com/raito-io/genqlient v0.0.2

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.17.59 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.14 // indirect
	github.com/bmatcuk/doublestar/v4 v4.8.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pascaldekloe/name v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.22 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
