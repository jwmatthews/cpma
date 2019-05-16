package oauth

import (
	"encoding/base64"

	"github.com/fusor/cpma/pkg/transform/secrets"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"

	configv1 "github.com/openshift/api/legacyconfig/v1"
)

// IdentityProviderHTPasswd is a htpasswd specific identity provider
type IdentityProviderHTPasswd struct {
	identityProviderCommon `yaml:",inline"`
	HTPasswd               `yaml:"htpasswd"`
}

// HTPasswd contains htpasswd FileData
type HTPasswd struct {
	FileData FileData `yaml:"fileData"`
}

// FileData from htpasswd file
type FileData struct {
	Name string `yaml:"name"`
}

func buildHTPasswdIP(serializer *json.Serializer, p IdentityProvider) (IdentityProviderHTPasswd, secrets.Secret, error) {
	var err error
	var idP IdentityProviderHTPasswd
	var secret *secrets.Secret

	var htpasswd configv1.HTPasswdPasswordIdentityProvider

	_, _, err = serializer.Decode(p.Provider.Raw, nil, &htpasswd)
	if err != nil {
		return idP, *secret, err
	}

	idP.Name = p.Name
	idP.Type = "HTPasswd"
	idP.Challenge = p.UseAsChallenger
	idP.Login = p.UseAsLogin
	idP.MappingMethod = p.MappingMethod
	idP.HTPasswd.FileData.Name = htpasswd.File

	secretName := p.Name + "-secret"
	idP.HTPasswd.FileData.Name = secretName

	encoded := base64.StdEncoding.EncodeToString(p.HTFileData)

	secret, err = secrets.GenSecret(secretName, encoded, "openshift-config", "htpasswd")
	if err != nil {
		return idP, *secret, err
	}

	return idP, *secret, nil
}
