package v1

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

type Generator interface {
	Generate(option ReturnOption) (*rsa.PrivateKey, *rsa.PublicKey,  error)
	generatePrivateKey(bitSize int)  (*rsa.PrivateKey, error)
	encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte
	encodePublicKeyToPEM(publickey *rsa.PublicKey) []byte
}

type Key struct {
	 Bit_size int
}

func New() Key {
	return Key{4096}
}
type ReturnOption struct {
	ConvertToString bool
	ReturnByteArray bool
}

func(key Key) Generate(option ReturnOption) (private interface{}, public interface{},  error error){
	privateKey, err := key.generatePrivateKey()
	if err != nil {
		log.Println("Failed to generate private key: [ERROR]: "+err.Error())
		return nil,nil,err
	}
	privateKeyBytes := key.encodePrivateKeyToPEM(privateKey)
	publicKey:=key.encodePublicKeyToPEM(&privateKey.PublicKey)

	if (option.ConvertToString==false && option.ReturnByteArray==false) || option.ConvertToString==true{
		return string(privateKeyBytes),string(publicKey),nil
	}
		return privateKeyBytes,publicKey,nil

}
func(k Key)  generatePrivateKey()(*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, k.Bit_size)
	if err != nil {
		return nil, err
	}
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	log.Println("Private Key generated")
	return privateKey, nil
}


func (k Key) encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

func(k Key) encodePublicKeyToPEM(publickey *rsa.PublicKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PublicKey(publickey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}
