// Taken and modified from https://github.com/go-chef/chef/blob/master/authentication.go

// License: Apache License, Version 2.0
// Authors:
//   - Kraig Amador (@bigkraig)
//   - AJ Christensen (@fujin)
//   - Jesse Nelson (@spheromak)

/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package client

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type cpReader struct {
	*bytes.Buffer
}

func (r cpReader) Close() error {
	return nil
}

func LoadKey(key_file string) (*rsa.PrivateKey, error) {
	key_contents, err := ioutil.ReadFile(key_file)
	if err != nil {
		return nil, err
	}

	key_data, _ := pem.Decode(key_contents)
	key, err := x509.ParsePKCS1PrivateKey(key_data.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenerateSignature will generate a signature ( sign ) the given data
func GenerateSignature(priv *rsa.PrivateKey, data string) (enc []byte, err error) {
	sig, err := privateEncrypt(priv, []byte(data))
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// privateEncrypt implements OpenSSL's RSA_private_encrypt function
func privateEncrypt(key *rsa.PrivateKey, data []byte) (enc []byte, err error) {
	k := (key.N.BitLen() + 7) / 8
	tLen := len(data)

	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		err = errors.New("Data too long")
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(key.N) > 0 {
		err = nil
		return
	}
	var m *big.Int
	var ir *big.Int
	if key.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, key.D, key.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, key.Precomputed.Dp, key.Primes[0])
		m2 := new(big.Int).Exp(c, key.Precomputed.Dq, key.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, key.Primes[0])
		}
		m.Mul(m, key.Precomputed.Qinv)
		m.Mod(m, key.Primes[0])
		m.Mul(m, key.Primes[1])
		m.Add(m, m2)

		for i, values := range key.Precomputed.CRTValues {
			prime := key.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, key.N)
	}
	enc = m.Bytes()
	return
}

// HashStr returns the base64 encoded SHA1 sum of the toHash string
func HashStr(toHash string) string {
	h := sha1.New()
	io.WriteString(h, toHash)
	hashed := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hashed
}

// Base64BlockEncode takes a byte slice and breaks it up into a
// slice of base64 encoded strings
func Base64BlockEncode(content []byte, limit int) []string {
	resultString := base64.StdEncoding.EncodeToString(content)
	var resultSlice []string

	index := 0

	var maxLengthPerSlice int

	// No limit
	if limit == 0 {
		maxLengthPerSlice = len(resultString)
	} else {
		maxLengthPerSlice = limit
	}

	// Iterate through the encoded string storing
	// a max of <limit> per slice item
	for i := 0; i < len(resultString)/maxLengthPerSlice; i++ {
		resultSlice = append(resultSlice, resultString[index:index+maxLengthPerSlice])
		index += maxLengthPerSlice
	}

	// Add remaining chunk to the end of the slice
	if len(resultString)%maxLengthPerSlice != 0 {
		resultSlice = append(resultSlice, resultString[index:])
	}

	return resultSlice
}

func PrepRequest(req *http.Request, client_name string, priv_key *rsa.PrivateKey) error {
	body_str := ""
	if req.Body != nil {
		body_buffer, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		reader1 := cpReader{bytes.NewBuffer(body_buffer)}
		reader2 := cpReader{bytes.NewBuffer(body_buffer)}
		req.Body = reader2

		content, err := ioutil.ReadAll(reader1)
		if err != nil {
			return err
		}

		body_str = string(content)
	}

	hashed_path := HashStr(req.URL.Path)
	hashed_body := HashStr(body_str)

	req.Header.Set("Method", req.Method)
	req.Header.Set("X-Chef-Version", "12.4.0")
	req.Header.Set("X-Ops-UserId", client_name)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Hashed Path", hashed_path)
	req.Header.Set("X-Ops-Timestamp", time.Now().UTC().Format(time.RFC3339))
	req.Header.Set("X-Ops-Sign", "algorithm=sha1;version=1.0")
	req.Header.Set("X-Ops-Reporting-Protocol-Version", "0.1.0")
	req.Header.Set("X-Ops-Content-Hash", hashed_body)

	req_content := ""
	for _, key := range []string{"Method", "Hashed Path", "X-Ops-Content-Hash", "X-Ops-Timestamp", "X-Ops-UserId"} {
		req_content += fmt.Sprintf("%s:%s\n", key, req.Header.Get(key))
	}
	req_content = strings.TrimSuffix(req_content, "\n")

	sig, err := GenerateSignature(priv_key, req_content)
	if err != nil {
		return err
	}

	base64sig := Base64BlockEncode(sig, 60)
	for i, val := range base64sig {
		req.Header.Set(fmt.Sprintf("X-Ops-Authorization-%d", i+1), string(val))
	}

	return nil
}
