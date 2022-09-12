# README

This is a template for running multiple Go AWS Lambdas with SAM that require common components.

Ensure that you correct the Makefile replace statement in each Lambda directory.

You can run this SAM project either with the high level SAM template in the root directory, or with SAM templates in each lambda directory.

The Go Lambdas use the AL2 runtime and each Lambda has a Makefile for the build process.

```bash
# From the root
sam build
sam local invoke lambda1
```

You might be wondering why all the trouble with this. I wanted a project with common Go dependencies that aren't imported separately. I've had 'too much fun' with that approach and opted for a local `pkg` import instead. Here are the files included:

```bash
├── LICENSE
├── README.md
├── cmd
│   ├── lambda1
│   │   ├── README.md
│   │   ├── cmd
│   │   │   ├── Makefile
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   └── template.yaml
│   ├── lambda2
│   │   ├── README.md
│   │   ├── cmd
│   │   │   ├── Makefile
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   └── template.yaml
│   └── lambda3
│       ├── README.md
│       ├── cmd
│       │   ├── Makefile
│       │   ├── go.mod
│       │   ├── go.sum
│       │   ├── main.go
│       │   └── main_test.go
│       └── template.yaml
├── go.mod
├── go.sum
├── pkg
│   ├── common
│   │   └── common.go
│   └── go.mod
└── template.yaml
```



