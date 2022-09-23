package devfile

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/devfile/registry-support/index/generator/schema"
)

func TestGetRegistrySamples(t *testing.T) {
	tests := []struct {
		name        string
		registry    string
		wantSamples []schema.Schema
		wantErr     bool
	}{
		{
			name:     "Fetch the sample",
			registry: DEVFILE_STAGING_REGISTRY_URL,
			wantSamples: []schema.Schema{
				{
					Name:        "nodejs-basic",
					DisplayName: "Basic Node.js",
					Description: "A simple Hello World Node.js application",
					Tags:        []string{"NodeJS", "Express"},
					Icon:        "https://nodejs.org/static/images/logos/nodejs-new-pantone-black.svg",
					Type:        schema.SampleDevfileType,
					ProjectType: "nodejs",
					Language:    "nodejs",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/nodeshift-starters/devfile-sample.git",
						},
					},
				},
				{
					Name:        "code-with-quarkus",
					DisplayName: "Basic Quarkus",
					Description: "A simple Hello World Java application using Quarkus",
					Tags:        []string{"Java", "Quarkus"},
					Icon:        "https://design.jboss.org/quarkus/logo/final/SVG/quarkus_icon_rgb_default.svg",
					Type:        schema.SampleDevfileType,
					ProjectType: "quarkus",
					Language:    "java",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/devfile-samples/devfile-sample-code-with-quarkus.git",
						},
					},
				},
				{
					Name:        "java-springboot-basic",
					DisplayName: "Basic Spring Boot",
					Description: "A simple Hello World Java Spring Boot application using Maven",
					Tags:        []string{"Java", "Spring"},
					Icon:        "https://spring.io/images/projects/spring-edf462fec682b9d48cf628eaf9e19521.svg",
					Type:        schema.SampleDevfileType,
					ProjectType: "springboot",
					Language:    "java",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/devfile-samples/devfile-sample-java-springboot-basic.git",
						},
					},
				},
				{
					Name:        "python-basic",
					DisplayName: "Basic Python",
					Description: "A simple Hello World application using Python",
					Tags:        []string{"Python"},
					Icon:        "https://raw.githubusercontent.com/devfile-samples/devfile-stack-icons/main/python.svg",
					Type:        schema.SampleDevfileType,
					ProjectType: "python",
					Language:    "python",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/devfile-samples/devfile-sample-python-basic.git",
						},
					},
				},
				{
					Name:        "go-basic",
					DisplayName: "Basic Go",
					Description: "A simple Hello World application using Go",
					Tags:        []string{"Go"},
					Icon:        "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
					Type:        schema.SampleDevfileType,
					ProjectType: "go",
					Language:    "go",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/devfile-samples/devfile-sample-go-basic.git",
						},
					},
				},
				{
					Name:        "dotnet60-basic",
					DisplayName: "Basic .NET 6.0",
					Description: "A simple application using .NET 6.0",
					Tags:        []string{"dotnet"},
					Icon:        "https://github.com/dotnet/brand/raw/main/logo/dotnet-logo.png",
					Type:        schema.SampleDevfileType,
					ProjectType: "dotnet",
					Language:    "dotnet",
					Provider:    "Red Hat",
					Git: &schema.Git{
						Remotes: map[string]string{
							"origin": "https://github.com/devfile-samples/devfile-sample-dotnet60-basic.git",
						},
					},
				},
			},
		},
		{
			name:     "Invalid registry",
			registry: "invalid",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := GetRegistrySamples(tt.registry)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error from test but got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("Got unexpected error: %s", err)
			} else if !tt.wantErr {
				var parsedRegistryIndex []schema.Schema
				err = json.Unmarshal(bytes, &parsedRegistryIndex)
				actualRegistryIndex, _ := json.MarshalIndent(parsedRegistryIndex, "", "  ")
				expectedRegistryIndex, _ := json.MarshalIndent(tt.wantSamples, "", "  ")
				if err != nil {
					t.Errorf("Got unexpected error: %s", err)
					return
				}
				if strings.Compare(string(expectedRegistryIndex), string(actualRegistryIndex)) != 0 {
					t.Errorf("expected %s does not match actual %s", expectedRegistryIndex, actualRegistryIndex)
				}

			}
		})
	}
}
