// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package Openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYb2scRRj/KsfoyzV70Sq67xptMUgxpOZVOMpk77m7qbcz68xc0hAW3F0RbRKQQCwR",
	"oRFCbAuNoiL+CfbDDEnqt5CZ2dzt3e5dLuWaRkhehLndnXn+/H6/meeZNeSzIGQUqBTIW0PCb0GAzfAG",
	"54zrQchZCFwSMI8DEAI3QQ/hHg7CNiAPCRaAbBHarBBRWcJ15CC5Gpo3khPaRFHkIA6fdQiHOvIWu6vU",
	"uh+ypbvgSxQ56Ba0imZJXf+vg/A5CSVhFHlIpfsqPVTJ7yo9nP0AOT2H3jN/DmowHmCJPNQhVL5zrecW",
	"oRKawLU5Cfdkce3j+9snj56p+IGKH6r0S5U8VcmfKv0qbwWp+L6Kn6h4XcU7Kv5CxQ+LgTuoI8Dk8XUO",
	"DeSh19xeyt0s3+6C/mYwScazsgzN4Sah2Ho6mCifdWhJPNMq/c4k6w8VHxx///nR3o/H2z/no5muVsdL",
	"WBsLeaccj34T/2yo+OD5fnzy266KD14UoYGk2PjKsvIJCaBNKMyDCBkVUMbdliWThECcBYmmYdQ1gznH",
	"q/p32Jf7UQvkUBoMIreIY70qC2ghI85FKUH4HIDeoTiAoomjbzaPvt7so38TB+RMpRO9GeRXLkaqTYPf",
	"4USu3ta5s3HOAOY2/h5FHbtB6clL9nV3sZaUIYr0WoQ2mJ4miTRu3rrxYeX63Cxy0DJwkclhqjpV1TGz",
	"ECgOCfLQW+aRBlhalrgN1m6zFeH6HLC0dGJCjs6+SrZUuq2Sxyp9qn/GOypZR8YON4DP1pGHbpqVDb42",
	"WyDkDKuvWgFTCVbCOAzbxDfT3LvCMs6yq0gL6yzAEGXmXdpWyYZK1vvcLtXnu+PxRm9wY5jdGcvmi+wJ",
	"pw44fVkoIZqdaLcHk7U3q9NFt983eNd1aNeq1XNhMmo/sAeq8aHf3jwI1uE+VCiTlQbr0Mz09Ms3PXPz",
	"jetzs1Zq+uwmdBm3ibH/9kWEvkA/pWyFViD7wEESN4XG1CKJavpZV4l1aMNoJeYJZ8/uAW0+UemuHqff",
	"DtHmPARsGV6pNuN9XU1cvEIfnMPyReu0WnT+44+uJHppJBqAK7MCTJtqQpk8k2cq2TMCfGwGv6pk6/kP",
	"j072/hoixtsttnJa101MjpMu4c5B4IlAVSh0S1C70sar1UYAra4wbDk57MTq9ZYq2TptO0vFYOsS3ZdM",
	"Sgnlve8E+t3SwyYf3Ms+YIb0zlcV4P9RQhpWMV7/dbLz978bv+hKb5SGJlrdjWybtYqSn0zZeajPu3TT",
	"DHbP30if0UMPI/ZE0MuupwrgXenkEujE3C/Wovw1CvIWexcoi7Wopl/xZaOixTXU4e3sxkR4rotDMpVx",
	"ccpngYuiWvRfAAAA///mwdDuGBYAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
