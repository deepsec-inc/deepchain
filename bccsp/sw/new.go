/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package sw

import (
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
	"reflect"

	"deepchain/bccsp"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

// NewDefaultSecurityLevel returns a new instance of the software-based BCCSP
// at security level 256, hash family SHA2 and using FolderBasedKeyStore as KeyStore.
func NewDefaultSecurityLevel(keyStorePath string) (bccsp.BCCSP, error) {
	ks := &fileBasedKeyStore{}
	if err := ks.Init(nil, keyStorePath, false); err != nil {
		return nil, errors.Wrapf(err, "Failed initializing key store at [%v]", keyStorePath)
	}

	return NewWithParams(256, "SHA2", ks)
}

// NewDefaultSecurityLevel returns a new instance of the software-based BCCSP
// at security level 256, hash family SHA2 and using the passed KeyStore.
func NewDefaultSecurityLevelWithKeystore(keyStore bccsp.KeyStore) (bccsp.BCCSP, error) {
	return NewWithParams(256, "SHA2", keyStore)
}

// NewWithParams returns a new instance of the software-based BCCSP
// set at the passed security level, hash family and KeyStore.
func NewWithParams(securityLevel int, hashFamily string, keyStore bccsp.KeyStore) (bccsp.BCCSP, error) {
	// Init config
	conf := &config{}
	err := conf.setSecurityLevel(securityLevel, hashFamily)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed initializing configuration at [%v,%v]", securityLevel, hashFamily)
	}

	swbccsp, err := New(keyStore)
	if err != nil {
		return nil, err
	}

	// Notice that errors are ignored here because some test will fail if one
	// of the following call fails.

	// Set the Encryptors
	swbccsp.AddWrapper(reflect.TypeOf(&aesPrivateKey{}), &aescbcpkcs7Encryptor{})

	// Set the Decryptors
	swbccsp.AddWrapper(reflect.TypeOf(&aesPrivateKey{}), &aescbcpkcs7Decryptor{})

	// Set the Signers
	swbccsp.AddWrapper(reflect.TypeOf(&ecdsaPrivateKey{}), &ecdsaSigner{})
	swbccsp.AddWrapper(reflect.TypeOf(&rsaPrivateKey{}), &rsaSigner{})

	// Set the Verifiers
	swbccsp.AddWrapper(reflect.TypeOf(&ecdsaPrivateKey{}), &ecdsaPrivateKeyVerifier{})
	swbccsp.AddWrapper(reflect.TypeOf(&ecdsaPublicKey{}), &ecdsaPublicKeyKeyVerifier{})
	swbccsp.AddWrapper(reflect.TypeOf(&rsaPrivateKey{}), &rsaPrivateKeyVerifier{})
	swbccsp.AddWrapper(reflect.TypeOf(&rsaPublicKey{}), &rsaPublicKeyKeyVerifier{})

	// Set the Hashers
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.SHAOpts{}), &hasher{hash: conf.hashFunction})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.SHA256Opts{}), &hasher{hash: sha256.New})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.SHA384Opts{}), &hasher{hash: sha512.New384})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.SHA3_256Opts{}), &hasher{hash: sha3.New256})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.SHA3_384Opts{}), &hasher{hash: sha3.New384})

	// Set the key generators
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAKeyGenOpts{}), &ecdsaKeyGenerator{curve: conf.ellipticCurve})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAP256KeyGenOpts{}), &ecdsaKeyGenerator{curve: elliptic.P256()})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAP384KeyGenOpts{}), &ecdsaKeyGenerator{curve: elliptic.P384()})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.AESKeyGenOpts{}), &aesKeyGenerator{length: conf.aesBitLength})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.AES256KeyGenOpts{}), &aesKeyGenerator{length: 32})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.AES192KeyGenOpts{}), &aesKeyGenerator{length: 24})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.AES128KeyGenOpts{}), &aesKeyGenerator{length: 16})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSAKeyGenOpts{}), &rsaKeyGenerator{length: conf.rsaBitLength})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSA1024KeyGenOpts{}), &rsaKeyGenerator{length: 1024})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSA2048KeyGenOpts{}), &rsaKeyGenerator{length: 2048})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSA3072KeyGenOpts{}), &rsaKeyGenerator{length: 3072})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSA4096KeyGenOpts{}), &rsaKeyGenerator{length: 4096})

	// Set the key generators
	swbccsp.AddWrapper(reflect.TypeOf(&ecdsaPrivateKey{}), &ecdsaPrivateKeyKeyDeriver{})
	swbccsp.AddWrapper(reflect.TypeOf(&ecdsaPublicKey{}), &ecdsaPublicKeyKeyDeriver{})
	swbccsp.AddWrapper(reflect.TypeOf(&aesPrivateKey{}), &aesPrivateKeyKeyDeriver{conf: conf})

	// Set the key importers
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.AES256ImportKeyOpts{}), &aes256ImportKeyOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.HMACImportKeyOpts{}), &hmacImportKeyOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAPKIXPublicKeyImportOpts{}), &ecdsaPKIXPublicKeyImportOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAPrivateKeyImportOpts{}), &ecdsaPrivateKeyImportOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.ECDSAGoPublicKeyImportOpts{}), &ecdsaGoPublicKeyImportOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.RSAGoPublicKeyImportOpts{}), &rsaGoPublicKeyImportOptsKeyImporter{})
	swbccsp.AddWrapper(reflect.TypeOf(&bccsp.X509PublicKeyImportOpts{}), &x509PublicKeyImportOptsKeyImporter{bccsp: swbccsp})

	return swbccsp, nil
}

// func New(securityLevel int, hashFamily string, keyStore bccsp.KeyStore) (bccsp.BCCSP, error) {

// 	// Init config
// 	conf := &config{}
// 	err := conf.setSecurityLevel(securityLevel, hashFamily)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "Failed initializing configuration at [%v,%v]", securityLevel, hashFamily)
// 	}

// 	// Check KeyStore
// 	if keyStore == nil {
// 		return nil, errors.Errorf("Invalid bccsp.KeyStore instance. It must be different from nil.")
// 	}

// 	// Set the encryptors
// 	encryptors := make(map[reflect.Type]Encryptor)
// 	encryptors[reflect.TypeOf(&gmsm4PrivateKey{})] = &gmsm4Encryptor{} //sm4 加密选项

// 	// Set the decryptors
// 	decryptors := make(map[reflect.Type]Decryptor)
// 	decryptors[reflect.TypeOf(&gmsm4PrivateKey{})] = &gmsm4Decryptor{} //sm4 解密选项

// 	// Set the signers
// 	signers := make(map[reflect.Type]Signer)
// 	signers[reflect.TypeOf(&gmsm2PrivateKey{})] = &gmsm2Signer{} //sm2 国密签名
// 	signers[reflect.TypeOf(&ecdsaPrivateKey{})] = &ecdsaPrivateKeySigner{}

// 	// Set the verifiers
// 	verifiers := make(map[reflect.Type]Verifier)
// 	verifiers[reflect.TypeOf(&gmsm2PrivateKey{})] = &gmsm2PrivateKeyVerifier{}  //sm2 私钥验签
// 	verifiers[reflect.TypeOf(&gmsm2PublicKey{})] = &gmsm2PublicKeyKeyVerifier{} //sm2 公钥验签
// 	verifiers[reflect.TypeOf(&ecdsaPrivateKey{})] = &ecdsaPrivateKeyVerifier{}
// 	verifiers[reflect.TypeOf(&ecdsaPublicKey{})] = &ecdsaPublicKeyKeyVerifier{}

// 	// Set the hashers
// 	hashers := make(map[reflect.Type]Hasher)
// 	hashers[reflect.TypeOf(&bccsp.SHAOpts{})] = &hasher{hash: conf.hashFunction}
// 	hashers[reflect.TypeOf(&bccsp.GMSM3Opts{})] = &hasher{hash: sm3.New} //sm3 Hash选项
// 	hashers[reflect.TypeOf(&bccsp.SHA256Opts{})] = &hasher{hash: sha256.New}
// 	hashers[reflect.TypeOf(&bccsp.SHA384Opts{})] = &hasher{hash: sha512.New384}
// 	hashers[reflect.TypeOf(&bccsp.SHA3_256Opts{})] = &hasher{hash: sha3.New256}
// 	hashers[reflect.TypeOf(&bccsp.SHA3_384Opts{})] = &hasher{hash: sha3.New384}

// 	impl := &impl{
// 		conf:       conf,
// 		ks:         keyStore,
// 		encryptors: encryptors,
// 		decryptors: decryptors,
// 		signers:    signers,
// 		verifiers:  verifiers,
// 		hashers:    hashers}

// 	// Set the key generators
// 	keyGenerators := make(map[reflect.Type]KeyGenerator)
// 	keyGenerators[reflect.TypeOf(&bccsp.GMSM2KeyGenOpts{})] = &gmsm2KeyGenerator{}
// 	keyGenerators[reflect.TypeOf(&bccsp.GMSM4KeyGenOpts{})] = &gmsm4KeyGenerator{length: 32}
// 	impl.keyGenerators = keyGenerators

// 	// Set the key derivers
// 	keyDerivers := make(map[reflect.Type]KeyDeriver)
// 	impl.keyDerivers = keyDerivers

// 	// Set the key importers
// 	keyImporters := make(map[reflect.Type]KeyImporter)
// 	keyImporters[reflect.TypeOf(&bccsp.GMSM4ImportKeyOpts{})] = &gmsm4ImportKeyOptsKeyImporter{}
// 	keyImporters[reflect.TypeOf(&bccsp.GMSM2PrivateKeyImportOpts{})] = &gmsm2PrivateKeyImportOptsKeyImporter{}
// 	keyImporters[reflect.TypeOf(&bccsp.GMSM2PublicKeyImportOpts{})] = &gmsm2PublicKeyImportOptsKeyImporter{}
// 	keyImporters[reflect.TypeOf(&bccsp.X509PublicKeyImportOpts{})] = &x509PublicKeyImportOptsKeyImporter{bccsp: impl}
// 	keyImporters[reflect.TypeOf(&bccsp.ECDSAGoPublicKeyImportOpts{})] = &ecdsaGoPublicKeyImportOptsKeyImporter{}
// 	keyImporters[reflect.TypeOf(&bccsp.ECDSAPrivateKeyImportOpts{})] = &ecdsaPrivateKeyImportOptsKeyImporter{}
// 	keyImporters[reflect.TypeOf(&bccsp.ECDSAPKIXPublicKeyImportOpts{})] = &ecdsaPKIXPublicKeyImportOptsKeyImporter{}

// 	impl.keyImporters = keyImporters
// 	return impl, nil
// }
