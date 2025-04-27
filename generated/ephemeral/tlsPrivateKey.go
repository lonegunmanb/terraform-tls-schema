package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const tlsPrivateKey = `{
  "block": {
    "attributes": {
      "algorithm": {
        "description": "Name of the algorithm to use when generating the private key. Currently-supported values are: ` + "`" + `RSA` + "`" + `, ` + "`" + `ECDSA` + "`" + `, ` + "`" + `ED25519` + "`" + `. ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ecdsa_curve": {
        "computed": true,
        "description": "When ` + "`" + `algorithm` + "`" + ` is ` + "`" + `ECDSA` + "`" + `, the name of the elliptic curve to use. Currently-supported values are: ` + "`" + `P224` + "`" + `, ` + "`" + `P256` + "`" + `, ` + "`" + `P384` + "`" + `, ` + "`" + `P521` + "`" + `. (default: ` + "`" + `P224` + "`" + `).",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "private_key_openssh": {
        "computed": true,
        "description": "Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.",
        "description_kind": "markdown",
        "sensitive": true,
        "type": "string"
      },
      "private_key_pem": {
        "computed": true,
        "description": "Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.",
        "description_kind": "markdown",
        "sensitive": true,
        "type": "string"
      },
      "private_key_pem_pkcs8": {
        "computed": true,
        "description": "Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.",
        "description_kind": "markdown",
        "sensitive": true,
        "type": "string"
      },
      "public_key_fingerprint_md5": {
        "computed": true,
        "description": "The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. ` + "`" + `aa:bb:cc:...` + "`" + `. Only available if the selected private key format is compatible, similarly to ` + "`" + `public_key_openssh` + "`" + ` and the [ECDSA P224 limitations](../../docs#limitations).",
        "description_kind": "markdown",
        "type": "string"
      },
      "public_key_fingerprint_sha256": {
        "computed": true,
        "description": "The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. ` + "`" + `SHA256:...` + "`" + `. Only available if the selected private key format is compatible, similarly to ` + "`" + `public_key_openssh` + "`" + ` and the [ECDSA P224 limitations](../../docs#limitations).",
        "description_kind": "markdown",
        "type": "string"
      },
      "public_key_openssh": {
        "computed": true,
        "description": " The public key data in [\"Authorized Keys\"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for ` + "`" + `ECDSA` + "`" + ` with curve ` + "`" + `P224` + "`" + `, as it is [not supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "markdown",
        "type": "string"
      },
      "public_key_pem": {
        "computed": true,
        "description": "Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a ` + "`" + `\\n` + "`" + ` at the end of the PEM. In case this disrupts your use case, we recommend using [` + "`" + `trimspace()` + "`" + `](https://www.terraform.io/language/functions/trimspace).",
        "description_kind": "markdown",
        "type": "string"
      },
      "rsa_bits": {
        "computed": true,
        "description": "When ` + "`" + `algorithm` + "`" + ` is ` + "`" + `RSA` + "`" + `, the size of the generated RSA key, in bits (default: ` + "`" + `2048` + "`" + `).",
        "description_kind": "markdown",
        "optional": true,
        "type": "number"
      }
    },
    "description": "-\u003e If the managed resource doesn't have a write-only argument available for the private key (first introduced in Terraform 1.11), then the private key can only be created with the managed resource variant of [` + "`" + `tls_private_key` + "`" + `](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key).\n\nCreates a PEM (and OpenSSH) formatted private key.\n\nGenerates a secure private key and encodes it in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) and [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formats. \n\nThe primary use-case for generating an ephemeral private key is to be used in combination with a write-only argument in a managed resource, which will avoid Terraform storing the private key data in the plan or state file.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func TlsPrivateKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(tlsPrivateKey), &result)
	return &result
}
