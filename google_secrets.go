package google_secrets

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"

	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

// Main function
func Get(project string, key string) (secret string, err error) {
	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + project + "/secrets/" + key + "/versions/latest",
	}

	resp, err := c.AccessSecretVersion(ctx, req)
	secret = string(resp.Payload.Data)

	return
}
