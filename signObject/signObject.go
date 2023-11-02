// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// [START storage_generate_signed_url_v4]
import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// generateV4GetObjectSignedURL generates object signed URL with GET method.
func main() {
	bucket := os.Args[1]
	object := os.Args[2]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		println(err.Error())
		return
	}
	defer client.Close()

	// Signing a URL requires credentials authorized to sign a URL. You can pass
	// these in through SignedURLOptions with one of the following options:
	//    a. a Google service account private key, obtainable from the Google Developers Console
	//    b. a Google Access ID with iam.serviceAccounts.signBlob permissions
	//    c. a SignBytes function implementing custom signing.
	// In this example, none of these options are used, which means the SignedURL
	// function attempts to use the same authentication that was used to instantiate
	// the Storage client. This authentication must include a private key or have
	// iam.serviceAccounts.signBlob permissions.
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	u, err := client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		println(err.Error())
	}

	f, err := os.Create("/workspace/source/signedURL.txt")

	if err != nil {
		println(err.Error())
	}

	defer f.Close()

	_, err2 := f.WriteString(u)

	if err2 != nil {
		println(err.Error())
	}

}
