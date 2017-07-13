// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package hdkeychain_test

// References:
//   [BIP32]: BIP0032 - Hierarchical Deterministic Wallets
//   https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki

import (
	"bytes"
	"encoding/hex"
	"errors"
	"reflect"
	"testing"

	"github.com/abcsuite/abcd/chaincfg"
	"github.com/abcsuite/abcutil/hdkeychain"
)

// TestBIP0032Vectors tests the vectors provided by [BIP32] to ensure the
// derivation works as intended.
func TestBIP0032Vectors(t *testing.T) {
	// The master seeds for each of the two test vectors in [BIP32].
	testVec1MasterHex := "000102030405060708090a0b0c0d0e0f"
	testVec2MasterHex := "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542"
	hkStart := uint32(0x80000000)

	tests := []struct {
		name     string
		master   string
		path     []uint32
		wantPub  string
		wantPriv string
		net      *chaincfg.Params
	}{
		// Test vector 1
		{
			name:     "test vector 1 chain m",
			master:   testVec1MasterHex,
			path:     []uint32{},
			wantPub:  "apub7FQF1bgMrRnREMGQsphPkVpA4xd9Nyr9TcdxrG6sx26RdNxa7eacLqs2RLyTzQvVmBdJ8ShczqahfJJL6yc1dGcq7c7YdMB6Wzxoyrq6R2e",
			wantPriv: "aprv2iiCsBe7i46kLtEczB2oM6A3Vw7DDefEtK3BZk4aXkJ2nthoVFm2AMFAFDR2xPwWRBjE21nMynz9dTAKY6G4vrS3nwrKSejK62X5dDhfxig",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 1 chain m/0H",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart},
			wantPub:  "apub7JfeSc2PLJnZFXF1id1P6hC4AdB7NB7x73fCwso9XSKM7jkK3hVRd3rGWdkxj4WCQZnUr8cUuXbdYuCGSQXStQ1CPLcpn6aqms57wcUBWH7",
			wantPriv: "aprv2mycJBz9Bw6tN4DDpyLnhHXwbbfBCqw3Xk4RfMkr7AWxHFVYRJfqSZEQLVzAMSzdmQVYEmUfYCmsGR7qa5jJvASCvvEUv3wjXy8giR2Kiqr",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 1 chain m/0H/1",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1},
			wantPub:  "apub7Ld8Ry6ys8LjPzwNoFjJSak2SbHrCF66oW986R1SqsKFM6nR1rvTxLW4iUP4DPRvg7bbZc7fgtyn4p71vqJdBQSR9q9LvdUfpcD3DxH837w",
			wantPriv: "aprv2ow6HZ4jikf4WXuauc4i3B5usZmv2uuCECYLoty9RbWrWcXePU6smqtCYKLk2kz4XKSHye8BHgcxKxoziX8aJgXvuyssehc9cn4CEm7eP2F",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2},
			wantPub:  "apub7Nk8yyLVnTyvC3Va574CDsf15TMQZxCoydZyMU1ukfjuE7XXfBRNSLBBX4qaxmtKZRVRW3vQGwSiQ5MDEnoxdfU13o2PibPuGUV7LgUQKsK",
			wantPriv: "aprv2r46qZJFe6JFJaTnBTPbpTztWRqUQd1uQKyC4wycLPwWPdGm2nbnFqZKLvq4JWWdZaFmuAuDvCLpS6eeuQ1ZparLafMLWehdxEALtsWtBXA",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H/2",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2, 2},
			wantPub:  "apub7QKF4xHE9NRwia6EKrFzbh4hV8XtYHEPz4oZAbwUcRy1TNzDw2kG8RBzSwmrC8iWjZrGsk1pWyb4Cyd4YKptM3Qpm6MzZJzUnUK4fbjVVY5",
			wantPriv: "aprv2sdCvYEyzzkGq74SSCbQCHQav71xNx3VQmCmt5uBCAAcctjTJdvfwva8GoBPajbXLCd9kmkG2K3D9mUxK48e1kuwnRDr5GE6cWYcahmBJVP",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H/2/1000000000",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2, 2, 1000000000},
			wantPub:  "apub7SVn1sEwHNxTRpNEV6EeUPhGEUTGJNXHuqoxuj9vX4UXUNAVr7shijxkGnNL4ytmtNEVBfFt6s2FTLwAVWhRC6p6ibT2UxjBuRJk5rTfsLA",
			wantPriv: "aprv2uojsTCh91GnYMLSbSa44z39fSwL93LPLYDBdD7d6ng8dsujDj47YFLt6fe8XEp9ci85Li3Tv3JvotQQhTX48WoeCsf2KM8tCqUjVqMQARC",
			net:      &chaincfg.MainNetParams,
		},

		// Test vector 2
		{
			name:     "test vector 2 chain m",
			master:   testVec2MasterHex,
			path:     []uint32{},
			wantPub:  "apub7FQF1bgMrRnRDxmjyymgChXvK9ma4c6y9CRh9CqhvQDHvMymHoMYmk2cmT4ShfHTfjs637MnPv5xP62ZEAXAmBcmyYK1TqfkqNdvFp4kdhn",
			wantPriv: "aprv2iiCsBe7i46kLVjx6L75oHsok8FduGv4ZtpurgoQW8Qu5sizfQXxbFQkbHC68uVbfEUemJEpFu2LiytRaWemYDkpFMvZaze3fH9JzCcKCbC",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 2 chain m/0",
			master:   testVec2MasterHex,
			path:     []uint32{0},
			wantPub:  "apub7HZDHVDQbYHh1YAWXxS6Eb6tSyT6AY9oFdNoMxRkRGi5htcxbVhNZ9MneukPqiJb54ztYwmLfS5GcCdLHBubSSUM2cM1vCKc1UpVZGuir5n",
			wantPriv: "aprv2ksB95BATAc2858ieJmVqBSmswwA1CxtgKn25SPSzzugsQNBy6snNejvUnBx7wM8FuN7Jv4MZVap9ShLJfqpp77W2ZdxXVKFqK7nDjNQbTg",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 2 chain m/0/2147483647H",
			master:   testVec2MasterHex,
			path:     []uint32{0, hkStart + 2147483647},
			wantPub:  "apub7KsXcqhopNBzftYtQPQERLAurRy7VV2gHFmxCFKvw1H6nfGz6qTZbFHsYRFAaymc3g1vkB4fkU8xYtT4L9hstErtwaPj9fKyRDyE2t4Qbb2",
			wantPriv: "aprv2oBVURfZfzWKnRX6Wjje1vWoHQTBL9qmhxBAujHdWjUhxB2DUSdyQkg1NFvaqiwXMDnVdKhiQKj1rhPVK6m874UyqBHxAjnsATLScx4HbLm",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 2 chain m/0/2147483647H/1",
			master:   testVec2MasterHex,
			path:     []uint32{0, hkStart + 2147483647, 1},
			wantPub:  "apub7MT6AyaeHdfi7Gd5hZHbV5E4a1vmCkt2FWSiBKjMzXkXx7ABoun86QbFWnRUYHP93DXR9evqKmiE3DfBXsmLFpZE6SBjUZd4wDkqxCY28Th",
			wantPriv: "aprv2pm42ZYQ9Fz3DobHoud15fZwzzQq3Rh7gCqvtoh4aFx97cuRBWxXuuyPLd7P2s4U4YC8wTZ2k3ffaw7keov1DnXhXXQv8Cd9EUsqmScAZTU",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 2 chain m/0/2147483647H/1/2147483646H",
			master:   testVec2MasterHex,
			path:     []uint32{0, hkStart + 2147483647, 1, hkStart + 2147483646},
			wantPub:  "apub7PSHEX9CpSzb4S8YGRswhzGaC6eGRvWEMDRZUEhiShMwPuHVUFTA1kCfPFPDG1v7awzzKQsbBBv1WN24ZNbiLBswEErZSWC115rpmZ1TG7L",
			wantPriv: "aprv2rkF676xg5JvAy6kNnDMJacTd58LGbKKmupnBifR2RZYZR2iqrdZqFaoD8g1BeTHCQJmcYD4wE8BDZQcWNYNGdq8DNJZg48Hmhum7dUxbuL",
			net:      &chaincfg.MainNetParams,
		},
		{
			name:     "test vector 2 chain m/0/2147483647H/1/2147483646H/2",
			master:   testVec2MasterHex,
			path:     []uint32{0, hkStart + 2147483647, 1, hkStart + 2147483646, 2},
			wantPub:  "apub7RCL85sVf2zMfMFYokYZ8oMXFEyrzDZCYQp3YCQEu1Vi5HZwsjejz4vTAYYQkFA3c8frsEBUPhfJC3odUTvEwDVh33agbdCnpowMNWfi3Bm",
			wantPriv: "aprv2tWHyfqFWfJgmtDkv6sxjPhQgDTvptNHy7DGFgMwUjhKEoKBFLq9oaJazSRvhKRajqd9BBeMi8hATN2CKgYDcPL2Bri6xDz4cW8f8WCRENd",
			net:      &chaincfg.MainNetParams,
		},

		// Test vector 1 - Testnet
		{
			name:     "test vector 1 chain m - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{},
			wantPub:  "tpubVhnMyQmZAhoosedBTX7oacwyCNc5qtdEMoNHudUCW1R6WZTvqCZQoNJHSn4H11puwdk4qyDv2ET637EDap4r8HH3odjBC5nEjmnPcsDfLwm",
			wantPriv: "tprvZUo1ZuEfLLFWfAYiMVaoDV1EeLmbSRuNzaSh7F4awft7dm8nHfFAFZyobWQyV8Qr26r8M2CmNw6nEb35HaECWFGy1vzx2ZGdyfBeaaHudoi",
			net:      &chaincfg.TestNet2Params,
		},
		{
			name:     "test vector 1 chain m/0H - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart},
			wantPub:  "tpubVm3mQR7aeaowtpbnJKRnvpKsJ3A3q5u31EPY1FAU5Re1zvFfmFUE5aHXY4qmjfQcb1uFZf8mvvU1vi89vEzHPQfR5NETLqByzdthaYfQGja",
			wantPriv: "tprvZY4QzuagpDFegLXKCHtnZgP8k1KZRdBBe1TwCrkrX67387vXDi9yXmy3gnz6tBTyNKcSZmu4wLtVsYzbKZhSVZH89uP7VxV4RboFfozTBMQ",
			net:      &chaincfg.TestNet2Params,
		},
		{
			name:     "test vector 1 chain m/0H/1 - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1},
			wantPub:  "tpubVo1FPnCBBQN83JJ9Nx9iGhsqa1Gnf9sBhgsT9nNmPrdvEHHmjQuGQrwKjuTsDzLLrZiNH8dxiHrASd2uQfmTgR6dqrkyVN5p3P2crvfgpEQ",
			wantPriv: "tprvZa1tzGfHM2opppDgGvchuZw71ySJFh9LLTwrMPy9qX6wMUxdBsb1s4cqtcLgZVTQ8EZCJeYagpjaw6gkU16ht5Nr8y2WEc9UWQimC4Y6MFs",
			net:      &chaincfg.TestNet2Params,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2},
			wantPub:  "tpubVq8FwnRh6k1JqLrLeoUc3znpCsLM2rytspJJQqPEJf4a7J2tNjQAtrcSYVvPyNnjjscCDaShJLK6mtH6idGo8g8Djpe2HL13VFJgygvRmc9",
			wantPriv: "tprvZc8uYGtoGNT1crmsYmwbgrr5eqVrdQG3WbNhcSyckKXbEVhjqC5vM4HxhDpzqEyyAVNgEBKdKLTT3EXQesyhPyhFoeVy6ZExqrpurCCRvrF",
			net:      &chaincfg.TestNet2Params,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H/2 - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2, 2},
			wantPub:  "tpubVrhN2mNRTeTLMsSzuYgQRpCWcYWq1C1UtFXtDyJoARHgLZVaeaj4awdFUNrfCjcvv1y3bGY7YNTSanYx2AHir453T7yd83bd1F8eJgtdq3F",
			wantPriv: "tprvZdi1dFqXdGu39PNXoX9Q4gFn4WgLbjHdX2cHRauBc5khTmAS73Qp39Jmd6BL7U4rw7k45nAfRT9qkuMi4Y6mb9ks1QNUfAmRW9DBY5g5ELT",
			net:      &chaincfg.TestNet2Params,
		},
		{
			name:     "test vector 1 chain m/0H/1/2H/2/1000000000 - testnet",
			master:   testVec1MasterHex,
			path:     []uint32{hkStart, 1, hkStart + 2, 2, 1000000000},
			wantPub:  "tpubVtstygL8beyr57j14nf4JWq5MtSCmHJNp2YHy6XF53oCMYfrZfrWBGQ1JDT95aoC4pMFuBnB8Ftdq9s3yMAFh7UKQd4f3hLL8C8KipwSek6",
			wantPriv: "tprvZftYaAoEmHRYrdeXxm83wNtLorbiMpaXSochAi7dWiGDUkLi28YFdU5XSxe53yHVDdEyfiTsKBRZR2HASwVBhueZRroeuFgD6U9JTC1mUyU",
			net:      &chaincfg.TestNet2Params,
		},
	}

tests:
	for i, test := range tests {
		masterSeed, err := hex.DecodeString(test.master)
		if err != nil {
			t.Errorf("DecodeString #%d (%s): unexpected error: %v",
				i, test.name, err)
			continue
		}

		extKey, err := hdkeychain.NewMaster(masterSeed, test.net)
		if err != nil {
			t.Errorf("NewMaster #%d (%s): unexpected error when "+
				"creating new master key: %v", i, test.name,
				err)
			continue
		}

		for _, childNum := range test.path {
			var err error
			extKey, err = extKey.Child(childNum)
			if err != nil {
				t.Errorf("err: %v", err)
				continue tests
			}
		}

		privStr, _ := extKey.String()
		if privStr != test.wantPriv {
			t.Errorf("Serialize #%d (%s): mismatched serialized "+
				"private extended key -- got: %s, want: %s", i,
				test.name, privStr, test.wantPriv)
			continue
		}

		pubKey, err := extKey.Neuter()
		if err != nil {
			t.Errorf("Neuter #%d (%s): unexpected error: %v ", i,
				test.name, err)
			continue
		}

		// Neutering a second time should have no effect.
		pubKey, err = pubKey.Neuter()
		if err != nil {
			t.Errorf("Neuter #%d (%s): unexpected error: %v", i,
				test.name, err)
			return
		}

		pubStr, _ := pubKey.String()
		if pubStr != test.wantPub {
			t.Errorf("Neuter #%d (%s): mismatched serialized "+
				"public extended key -- got: %s, want: %s", i,
				test.name, pubStr, test.wantPub)
			continue
		}
	}
}

// TestPrivateDerivation tests several vectors which derive private keys from
// other private keys works as intended.
func TestPrivateDerivation(t *testing.T) {
	// The private extended keys for test vectors in [BIP32].
	testVec1MasterPrivKey := "aprv2iiCsBe7i46kLsQXfeN52TWFHNmaNT6UFfnNBCqYm9htp7nKLW91URmxbDw6S3mzhTPVykqVUZ7mN1csZZDCab7947x7HwvxfzyF2Ti1dDu"
	testVec2MasterPrivKey := "aprv2iiCsBe7i46kMeRMtyWtxeeSdmr6ixXbra64JMxpG3n9UVXHyvuCdGRc2fpDgtYrFoDZSpWD9nMkLWLJYcZemBXgu4sJxhCDxfWb1wyp9tL"

	tests := []struct {
		name     string
		master   string
		path     []uint32
		wantPriv string
	}{
		// Test vector 1
		{
			name:     "test vector 1 chain m",
			master:   testVec1MasterPrivKey,
			path:     []uint32{},
			wantPriv: "aprv2iiCsBe7i46kLsQXfeN52TWFHNmaNT6UFfnNBCqYm9htp7nKLW91URmxbDw6S3mzhTPVykqVUZ7mN1csZZDCab7947x7HwvxfzyF2Ti1dDu",
		},
		{
			name:     "test vector 1 chain m/0",
			master:   testVec1MasterPrivKey,
			path:     []uint32{0},
			wantPriv: "aprv2m3vK2wiVn2hbscDb2FUSvozZWPRrr5bDYxS2R23E5T6yc7275LBE8QhbiiK6YroJ2nHUWysgbdk4Bnase8QhRGzvZNJaLCn7LnvUVYoQs9",
		},
		{
			name:     "test vector 1 chain m/0/1",
			master:   testVec1MasterPrivKey,
			path:     []uint32{0, 1},
			wantPriv: "aprv2pFMrVGBbob3qyzShNmTA2hkvefgWyX3b6tRcqvFSpzS1hMCXxyHRxZKx7K4hjmrCD9qAW7PjLUE1kSYWVhZrG68bo7GyP5FjapvGPtuxdG",
		},
		{
			name:     "test vector 1 chain m/0/1/2",
			master:   testVec1MasterPrivKey,
			path:     []uint32{0, 1, 2},
			wantPriv: "aprv2r4fERpBM9CHbeMy5kjqFnxEwyQzrJwQka931EXcBX6jHmggFA5y3QhkmkL9HrvZyNLxPNNeNka5ae5n8jhqDLF8PhHu9b67fhTZ2n8KJKr",
		},
		{
			name:     "test vector 1 chain m/0/1/2/2",
			master:   testVec1MasterPrivKey,
			path:     []uint32{0, 1, 2, 2},
			wantPriv: "aprv2t71d6VzQ26AdaCg8hgGbpo7msGyW5KA4291T1TcCXAtkFvD9bNeQxNe1WZkhEFGDxRqN2kRiRr1ixHyvNFDv8KnrGXMx7G7kAtzeJ3FQS2",
		},
		{
			name:     "test vector 1 chain m/0/1/2/2/1000000000",
			master:   testVec1MasterPrivKey,
			path:     []uint32{0, 1, 2, 2, 1000000000},
			wantPriv: "aprv2tpd2gDCwafXbzA6A4sh2yNuXV3eeNKes6z3GSGUynCFoJQGEa4vh4ADj9b3NK8FQshqyvJTChuzE3TnUxNkSdjjdXRo2vkgwDGyYXMPhT6",
		},

		// Test vector 2
		{
			name:     "test vector 2 chain m",
			master:   testVec2MasterPrivKey,
			path:     []uint32{},
			wantPriv: "aprv2iiCsBe7i46kMeRMtyWtxeeSdmr6ixXbra64JMxpG3n9UVXHyvuCdGRc2fpDgtYrFoDZSpWD9nMkLWLJYcZemBXgu4sJxhCDxfWb1wyp9tL",
		},
		{
			name:     "test vector 2 chain m/0",
			master:   testVec2MasterPrivKey,
			path:     []uint32{0},
			wantPriv: "aprv2koykvFadbxizFmsz8C2bBTjcLQUvc7hwsqMfasMgt76s8Jv3y4PboAURXGXf5BL9SVsWuDPC2WJFhHtrw6bHzS9nUx6SQK7cAkc9QM4dR8",
		},
		{
			name:     "test vector 2 chain m/0/2147483647",
			master:   testVec2MasterPrivKey,
			path:     []uint32{0, 2147483647},
			wantPriv: "aprv2pEWQMaDBraisAtpXen5zkBcyNYdRB4Z7STM6bh1QjKaEoX9KwyTzZdY1eY9vnDcGGqevat1jihfH6Y2W8p2ah2NiyivDxQxwz3pVn5CyjB",
		},
		{
			name:     "test vector 2 chain m/0/2147483647/1",
			master:   testVec2MasterPrivKey,
			path:     []uint32{0, 2147483647, 1},
			wantPriv: "aprv2r8ya2FxNtAeMeMAk858Gr8s5PPHJQ7eKooXfBtaYT9WGV4VGE3nasm2ZKkCZZPvY3LwceDd9eHoV7NVkRjQhnB9DU5tGwvuLjKjDPpHTLS",
		},
		{
			name:     "test vector 2 chain m/0/2147483647/1/2147483646",
			master:   testVec2MasterPrivKey,
			path:     []uint32{0, 2147483647, 1, 2147483646},
			wantPriv: "aprv2rMjygXxBz8ierVzHkksgQiKtWUuzrgED1tiYFwgcafD2bLsC9ELKhdbGR6BDHhM7eAEoborpvV11t4hvcidqvRFg373BPpS5i89nTeJPzw",
		},
		{
			name:     "test vector 2 chain m/0/2147483647/1/2147483646/2",
			master:   testVec2MasterPrivKey,
			path:     []uint32{0, 2147483647, 1, 2147483646, 2},
			wantPriv: "aprv2tB9HPMC3CoLYjJmNXj91HExXaYhfm4kj6V1ySDGkhGw9V7kL6tdF5wq8AQNxve63KWzMvhKL1nYLSSwJEHcXUebjjMkk6u9rNjEC12xkdS",
		},

		// Custom tests to trigger specific conditions.
		{
			// Seed 000000000000000000000000000000da.
			name:     "Derived privkey with zero high byte m/0",
			master:   "dprv3jFfEhxvVxy6NJWopujhfg7syQL71xCRgNoGUpQTtjTpCwzigwtCwssQGbRQsby7PBs1Yp8Wu7isu396qeNof13EZuxbCTJVF1xkoFAQHWj",
			path:     []uint32{0},
			wantPriv: "dprv3mWLns1v1fdLfeu5DKTA6NWQHLF6pFsPSwKCS6q4h4nkjm2DfuH5X2iDnW15jhHTGa3rzxSpvskuXugcbBcUUVWCETKKzjW7ja4V2jL4aw4",
		},
	}

tests:
	for i, test := range tests {
		extKey, err := hdkeychain.NewKeyFromString(test.master)
		if err != nil {
			t.Errorf("NewKeyFromString #%d (%s): unexpected error "+
				"creating extended key: %v", i, test.name,
				err)
			continue
		}

		for _, childNum := range test.path {
			var err error
			extKey, err = extKey.Child(childNum)
			if err != nil {
				t.Errorf("err: %v", err)
				continue tests
			}
		}

		privStr, _ := extKey.String()
		if privStr != test.wantPriv {
			t.Errorf("Child #%d (%s): mismatched serialized "+
				"private extended key -- got: %s, want: %s", i,
				test.name, privStr, test.wantPriv)
			continue
		}
	}
}

// TestPublicDerivation tests several vectors which derive public keys from
// other public keys works as intended.
func TestPublicDerivation(t *testing.T) {
	// The public extended keys for test vectors in [BIP32].
	testVec1MasterPubKey := "apub7FQF1bgMrRnRDCM5zy2bM49V6nNqCjaYM7xZkVW8Pm6Qpf7eWX7xzM2eh3UMEvqxncdBcTdrbbxGQWhPRu58N2kEDqut15fazbuYL3Fex3K"
	testVec2MasterPubKey := "apub7FQF1bgMrRnRExjaGctxCpBRbk8nYfWsdw8UnqZNGELuiZ6ZbacgCG2jGaaDU7JvidAoP99N6ssKUV5FMLXS2DnTvwYBYMTsCeMFWd1SesM"

	tests := []struct {
		name    string
		master  string
		path    []uint32
		wantPub string
	}{
		// Test vector 1
		{
			name:    "test vector 1 chain m",
			master:  testVec1MasterPubKey,
			path:    []uint32{},
			wantPub: "apub7FQF1bgMrRnRDCM5zy2bM49V6nNqCjaYM7xZkVW8Pm6Qpf7eWX7xzM2eh3UMEvqxncdBcTdrbbxGQWhPRu58N2kEDqut15fazbuYL3Fex3K",
		},
		{
			name:    "test vector 1 chain m/0",
			master:  testVec1MasterPubKey,
			path:    []uint32{0},
			wantPub: "apub7HocgkAEf7SLAwp3N3ehe6V7eNgMWSaPsBesb3cHMiXxeYhupURzvtwj1MX36GNbWvzsjPmwcdwkAJqvysQAFkMZLoFKXEvXkFGU9TJ1rym",
		},
		{
			name:    "test vector 1 chain m/0/1",
			master:  testVec1MasterPubKey,
			path:    []uint32{0, 1},
			wantPub: "apub7LXeprSM8rHLEL5zym7kU8ZGAefbjgLM67CJ5HNvsnFxZdvME4LCjxEptpGh4QY7bpWsnfeUaQ2AHr79WRVxR8LrEagrgsqpogdPVBDifA2",
		},
		{
			name:    "test vector 1 chain m/0/1/2",
			master:  testVec1MasterPubKey,
			path:    []uint32{0, 1, 2},
			wantPub: "apub7M56p4mMMhFvQyoE4Nm3FKyoHQ5SR6no2AXxZTgWAt147iEspnvVs3zuspYeNTsnAfsuqgyKkpBZs1LRcNbnKak5XCLhULzVeSm4HEKSnHq",
		},
		{
			name:    "test vector 1 chain m/0/1/2/2",
			master:  testVec1MasterPubKey,
			path:    []uint32{0, 1, 2, 2},
			wantPub: "apub7PfsasXZ5wSXZpwxA2UGzPKVkHchDTSGVs3KbPqU4s9cyQJrw148cfUt8ND5K8W3HYVrqPUaeTRucBmX2XzTJg2p8VPuiM4sAvKEXhcdH4v",
		},
		{
			name:    "test vector 1 chain m/0/1/2/2/1000000000",
			master:  testVec1MasterPubKey,
			path:    []uint32{0, 1, 2, 2, 1000000000},
			wantPub: "apub7SMn3QY8GnnyT884SeadFgNhLafen4p6FphJW6wyWBkuYWEFM2M6jCkkwHhoeWyq9HsDDHmJRvTmPySPUnjx7c7gXkJTyNJfWcuRDu9R1V3",
		},

		// Test vector 2
		{
			name:    "test vector 2 chain m",
			master:  testVec2MasterPubKey,
			path:    []uint32{},
			wantPub: "apub7FQF1bgMrRnRExjaGctxCpBRbk8nYfWsdw8UnqZNGELuiZ6ZbacgCG2jGaaDU7JvidAoP99N6ssKUV5FMLXS2DnTvwYBYMTsCeMFWd1SesM",
		},
		{
			name:    "test vector 2 chain m/0",
			master:  testVec2MasterPubKey,
			path:    []uint32{0},
			wantPub: "apub7HdtsSnL5JNbdhF12kEZJRUFm4eoX7PSnfHYimjGcxMH4FVFJKWTZh4vD311cYKe2NXMT8nK7fsvuLgTbQrxXWkAnBvCqUeTypHDQNUcbWn",
		},
		{
			name:    "test vector 2 chain m/0/2147483647",
			master:  testVec2MasterPubKey,
			path:    []uint32{0, 2147483647},
			wantPub: "apub7LfXCPae8PVa2C8cAKyoTbR7HmH8sNQMSqNPXowRcEw7D5dAyyKd3adb4muzfJhLv2Bud9pMneZNH8nZgCp3gkxU1Y5a1gtW9oW1QPvamLE",
		},
		{
			name:    "test vector 2 chain m/0/2147483647/1",
			master:  testVec2MasterPubKey,
			path:    []uint32{0, 2147483647, 1},
			wantPub: "apub7M5KWdVKeXq2spiWw33PbEv2DgV4t2WdordLxwQGJtYhptivTruufLcNnzfKKtWxG4RiKjxKGs9WMAG2e4V2k24bkQoMMbTPL1TmGtLrBWg",
		},
		{
			name:    "test vector 2 chain m/0/2147483647/1/2147483646",
			master:  testVec2MasterPubKey,
			path:    []uint32{0, 2147483647, 1, 2147483646},
			wantPub: "apub7QLT53oWFrU9ZdzopAAdgfF9c6kUsbwiBpETBWpmeBboJyyv3kb5bZsqM4kJRgACo9u8VHhVYDwLY2zzUgAbn2aXENLbvpZjiPoyuqgPJzT",
		},
		{
			name:    "test vector 2 chain m/0/2147483647/1/2147483646/2",
			master:  testVec2MasterPubKey,
			path:    []uint32{0, 2147483647, 1, 2147483646, 2},
			wantPub: "apub7RZuDz3AwuupTZax3Jt2vxQ5dB1KqFGYoYJ4spwFmMBhmuwWxdY6vBNoCUrDz6zi7EWfQuXEyqqL6S1thaQfKbDU7iXJQjrqAmTf2eTATwj",
		},
	}

tests:
	for i, test := range tests {
		extKey, err := hdkeychain.NewKeyFromString(test.master)
		if err != nil {
			t.Errorf("NewKeyFromString #%d (%s): unexpected error "+
				"creating extended key: %v", i, test.name,
				err)
			continue
		}

		for _, childNum := range test.path {
			var err error
			extKey, err = extKey.Child(childNum)
			if err != nil {
				t.Errorf("err: %v", err)
				continue tests
			}
		}

		pubStr, _ := extKey.String()
		if pubStr != test.wantPub {
			t.Errorf("Child #%d (%s): mismatched serialized "+
				"public extended key -- got: %s, want: %s", i,
				test.name, pubStr, test.wantPub)
			continue
		}
	}
}

// TestGenenerateSeed ensures the GenerateSeed function works as intended.
func TestGenenerateSeed(t *testing.T) {
	wantErr := errors.New("seed length must be between 128 and 512 bits")

	tests := []struct {
		name   string
		length uint8
		err    error
	}{
		// Test various valid lengths.
		{name: "16 bytes", length: 16},
		{name: "17 bytes", length: 17},
		{name: "20 bytes", length: 20},
		{name: "32 bytes", length: 32},
		{name: "64 bytes", length: 64},

		// Test invalid lengths.
		{name: "15 bytes", length: 15, err: wantErr},
		{name: "65 bytes", length: 65, err: wantErr},
	}

	for i, test := range tests {
		seed, err := hdkeychain.GenerateSeed(test.length)
		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("GenerateSeed #%d (%s): unexpected error -- "+
				"want %v, got %v", i, test.name, test.err, err)
			continue
		}

		if test.err == nil && len(seed) != int(test.length) {
			t.Errorf("GenerateSeed #%d (%s): length mismatch -- "+
				"got %d, want %d", i, test.name, len(seed),
				test.length)
			continue
		}
	}
}

// TestExtendedKeyAPI ensures the API on the ExtendedKey type works as intended.
func TestExtendedKeyAPI(t *testing.T) {
	tests := []struct {
		name       string
		extKey     string
		isPrivate  bool
		parentFP   uint32
		privKey    string
		privKeyErr error
		pubKey     string
		address    string
	}{
		{
			name:      "test vector 1 master node private",
			extKey:    "aprv2iiCsBe7i46kN2N2SdHH92zjBphvYTAgdRrYgtcaWruEtCkprjUNX4N2UTnPY9LRty6GW5pYuH9ksGenvLTxbji8TxKxkXKTJrw6EJfJKKj",
			isPrivate: true,
			parentFP:  0,
			privKey:   "ce065119c36c1b61a5042d12a4337176fbd2b41fe2acb2cadb3d417ef68ffe19",
			pubKey:    "02cda9fd97bd1b9ebcf3040d43702ccd7e9ec892f250d5f19c88adcbe92beb3631",
			address:   "AbGz6ZmYnFUuxVpxp5sSRrjw3PiE3qXWnpU",
		},
		{
			name:       "test vector 2 chain m/0/2147483647/1/2147483646/2",
			extKey:     "apub7RZuDz3AwuupTZax3Jt2vxQ5dB1KqFGYoYJ4spwFmMBhmuwWxdY6vBNoCUrDz6zi7EWfQuXEyqqL6S1thaQfKbDU7iXJQjrqAmTf2eTATwj",
			isPrivate:  false,
			parentFP:   1730864243,
			privKeyErr: hdkeychain.ErrNotPrivExtKey,
			pubKey:     "03a44a50f1dd4ca367aef365654e956480a74bab0bc08e1f0160c8ae829fe3fa3f",
			address:    "Ab7Xp2BPXfBGcuytkkHVB1EDmFU9XjUZpdv",
		},
	}

	for i, test := range tests {
		key, err := hdkeychain.NewKeyFromString(test.extKey)
		if err != nil {
			t.Errorf("NewKeyFromString #%d (%s): unexpected "+
				"error: %v", i, test.name, err)
			continue
		}

		if key.IsPrivate() != test.isPrivate {
			t.Errorf("IsPrivate #%d (%s): mismatched key type -- "+
				"want private %v, got private %v", i, test.name,
				test.isPrivate, key.IsPrivate())
			continue
		}

		parentFP := key.ParentFingerprint()
		if parentFP != test.parentFP {
			t.Errorf("ParentFingerprint #%d (%s): mismatched "+
				"parent fingerprint -- want %d, got %d", i,
				test.name, test.parentFP, parentFP)
			continue
		}

		serializedKey, _ := key.String()
		if serializedKey != test.extKey {
			t.Errorf("String #%d (%s): mismatched serialized key "+
				"-- want %s, got %s", i, test.name, test.extKey,
				serializedKey)
			continue
		}

		privKey, err := key.ECPrivKey()
		if !reflect.DeepEqual(err, test.privKeyErr) {
			t.Errorf("ECPrivKey #%d (%s): mismatched error: want "+
				"%v, got %v", i, test.name, test.privKeyErr, err)
			continue
		}
		if test.privKeyErr == nil {
			privKeyStr := hex.EncodeToString(privKey.Serialize())
			if privKeyStr != test.privKey {
				t.Errorf("ECPrivKey #%d (%s): mismatched "+
					"private key -- want %s, got %s", i,
					test.name, test.privKey, privKeyStr)
				continue
			}
		}

		pubKey, err := key.ECPubKey()
		if err != nil {
			t.Errorf("ECPubKey #%d (%s): unexpected error: %v", i,
				test.name, err)
			continue
		}
		pubKeyStr := hex.EncodeToString(pubKey.SerializeCompressed())
		if pubKeyStr != test.pubKey {
			t.Errorf("ECPubKey #%d (%s): mismatched public key -- "+
				"want %s, got %s", i, test.name, test.pubKey,
				pubKeyStr)
			continue
		}

		addr, err := key.Address(&chaincfg.MainNetParams)
		if err != nil {
			t.Errorf("Address #%d (%s): unexpected error: %v", i,
				test.name, err)
			continue
		}
		if addr.EncodeAddress() != test.address {
			t.Errorf("Address #%d (%s): mismatched address -- want "+
				"%s, got %s", i, test.name, test.address,
				addr.EncodeAddress())
			continue
		}
	}
}

// TestNet ensures the network related APIs work as intended.
func TestNet(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		origNet   *chaincfg.Params
		newNet    *chaincfg.Params
		newPriv   string
		newPub    string
		isPrivate bool
	}{
		// Private extended keys.
		{
			name:      "mainnet -> simnet",
			key:       "aprv2iiCsBe7i46kMZBBFnwKDi7VpnwS3wpXZS6Qw6o97udGuhoiEeghEB6UH9MeRnrFgqX62uqsfwPTVHUJ9x5uyqHcjDvE4adCKhcVu578zyq",
			origNet:   &chaincfg.MainNetParams,
			newNet:    &chaincfg.SimNetParams,
			newPriv:   "sprvZ9xkGEZkBei2pYJ9nkhgt1sBoETNL75AkayhhCuG9pqUhURTnQ1GhLAycEPzxcfft7XBcSficR9YD8jVBje2kzSw5sNS4aM2wcDRTsDondA",
			newPub:    "spubVNx6fk6e22GL32NctnEhF9ovMGHrjZo27ouJVbJsiANTaGkcKwKXF8VTTWuZKitQLJ2mZCrGFTKbUFCtSjN2rTEKwp2wU8Pgmm8HMGauz8q",
			isPrivate: true,
		},
		{
			name:      "simnet -> mainnet",
			key:       "sprvZ9xkGEZkBei2p9e1uBZRQMGtfGEQNGApP1W19PyNRqg9nuEs2X4ynkvAXWaBiGb5WKiaqcbiKgmyB1HYgcX3mnxiUs7UWeWEfe4tnSpbXLv",
			origNet:   &chaincfg.SimNetParams,
			newNet:    &chaincfg.MainNetParams,
			newPriv:   "aprv2iiCsBe7i46kMAX3NDo3k3XCgpiU66vBBrciPHsFPvTx18d7UmkQKbqfCRXqBSmfK3iVG5msPD1tTA2MepxvzdoQ8DfGWenQ3jTyDhHH3vN",
			newPub:    "apub7FQF1bgMrRnREdYqFsTe9TBKFrEQFS75mADVfouYpCGLqcst7AZzW6TXNaMSN7R9CVS4ZxE9gwmiuZiaEBtQdwCYEGaSaz9gu1TN9fuUs69",
			isPrivate: true,
		},

		// Public extended keys.
		{
			name:      "mainnet -> simnet",
			key:       "apub7FQF1bgMrRnREdYqFsTe9TBKFrEQFS75mADVfouYpCGLqcst7AZzW6TXNaMSN7R9CVS4ZxE9gwmiuZiaEBtQdwCYEGaSaz9gu1TN9fuUs69",
			origNet:   &chaincfg.MainNetParams,
			newNet:    &chaincfg.SimNetParams,
			newPub:    "spubVNx6fk6e22GL2diV1D6RmVDdDJ4tmitfkERbwnNyzBD8fha1a4PELZEeNoUfNofdyJS2Y19tFgHZQ62tzKwELiBA3xVeZowLr4DJQ7xGuao",
			isPrivate: false,
		},
		{
			name:      "simnet -> mainnet",
			key:       "spubVNx6fk6e22GL2diV1D6RmVDdDJ4tmitfkERbwnNyzBD8fha1a4PELZEeNoUfNofdyJS2Y19tFgHZQ62tzKwELiBA3xVeZowLr4DJQ7xGuao",
			origNet:   &chaincfg.SimNetParams,
			newNet:    &chaincfg.MainNetParams,
			newPub:    "apub7FQF1bgMrRnREdYqFsTe9TBKFrEQFS75mADVfouYpCGLqcst7AZzW6TXNaMSN7R9CVS4ZxE9gwmiuZiaEBtQdwCYEGaSaz9gu1TN9fuUs69",
			isPrivate: false,
		},
	}

	for i, test := range tests {
		extKey, err := hdkeychain.NewKeyFromString(test.key)
		if err != nil {
			t.Errorf("NewKeyFromString #%d (%s): unexpected error "+
				"creating extended key: %v", i, test.name,
				err)
			continue
		}

		if !extKey.IsForNet(test.origNet) {
			t.Errorf("IsForNet #%d (%s): key is not for expected "+
				"network %v", i, test.name, test.origNet.Name)
			continue
		}

		extKey.SetNet(test.newNet)
		if !extKey.IsForNet(test.newNet) {
			t.Errorf("SetNet/IsForNet #%d (%s): key is not for "+
				"expected network %v", i, test.name,
				test.newNet.Name)
			continue
		}

		if test.isPrivate {
			privStr, _ := extKey.String()
			if privStr != test.newPriv {
				t.Errorf("Serialize #%d (%s): mismatched serialized "+
					"private extended key -- got: %s, want: %s", i,
					test.name, privStr, test.newPriv)
				continue
			}

			extKey, err = extKey.Neuter()
			if err != nil {
				t.Errorf("Neuter #%d (%s): unexpected error: %v ", i,
					test.name, err)
				continue
			}
		}

		pubStr, _ := extKey.String()
		if pubStr != test.newPub {
			t.Errorf("Neuter #%d (%s): mismatched serialized "+
				"public extended key -- got: %s, want: %s", i,
				test.name, pubStr, test.newPub)
			continue
		}
	}
}

// TestErrors performs some negative tests for various invalid cases to ensure
// the errors are handled properly.
func TestErrors(t *testing.T) {
	// Should get an error when seed has too few bytes.
	net := &chaincfg.MainNetParams
	_, err := hdkeychain.NewMaster(bytes.Repeat([]byte{0x00}, 15), net)
	if err != hdkeychain.ErrInvalidSeedLen {
		t.Errorf("NewMaster: mismatched error -- got: %v, want: %v",
			err, hdkeychain.ErrInvalidSeedLen)
	}

	// Should get an error when seed has too many bytes.
	_, err = hdkeychain.NewMaster(bytes.Repeat([]byte{0x00}, 65), net)
	if err != hdkeychain.ErrInvalidSeedLen {
		t.Errorf("NewMaster: mismatched error -- got: %v, want: %v",
			err, hdkeychain.ErrInvalidSeedLen)
	}

	// Generate a new key and neuter it to a public extended key.
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		t.Errorf("GenerateSeed: unexpected error: %v", err)
		return
	}
	extKey, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		t.Errorf("NewMaster: unexpected error: %v", err)
		return
	}
	pubKey, err := extKey.Neuter()
	if err != nil {
		t.Errorf("Neuter: unexpected error: %v", err)
		return
	}

	// Deriving a hardened child extended key should fail from a public key.
	_, err = pubKey.Child(hdkeychain.HardenedKeyStart)
	if err != hdkeychain.ErrDeriveHardFromPublic {
		t.Errorf("Child: mismatched error -- got: %v, want: %v",
			err, hdkeychain.ErrDeriveHardFromPublic)
	}

	// NewKeyFromString failure tests.
	tests := []struct {
		name      string
		key       string
		err       error
		neuter    bool
		neuterErr error
	}{
		{
			name: "invalid key length",
			key:  "dpub1234",
			err:  hdkeychain.ErrInvalidKeyLen,
		},
		{
			name: "bad checksum",
			key:  "dpubZF6AWaFizAuUcbkZSs8cP8Gxzr6Sg5tLYYM7gEjZMC5GDaSHB4rW4F51zkWyo9U19BnXhc99kkEiPg248bYin8m9b8mGss9nxV6N2QpU8vj",
			err:  hdkeychain.ErrBadChecksum,
		},
		{
			name: "pubkey not on curve",
			key:  "dpubZ9169KDAEUnyoTzA7pDGtXbxpji5LuUk8johUPVGY2CDsz6S7hahGNL6QkeYrUeAPnaJD1MBmrsUnErXScGZdjL6b2gjCRX1Z1GNhLdVCjv",
			err:  errors.New("pubkey [0,50963827496501355358210603252497135226159332537351223778668747140855667399507] isn't on secp256k1 curve"),
		},
		{
			name:      "unsupported version",
			key:       "4s9bfpYH9CkJboPNLFC4BhTENPrjfmKwUxesnqxHBjv585bCLzVdQKuKQ5TouA57FkdDskrR695Z5U2wWwDUUVWXPg7V57sLpc9dMgx74LsVZGEB",
			err:       nil,
			neuter:    true,
			neuterErr: chaincfg.ErrUnknownHDKeyID,
		},
	}

	for i, test := range tests {
		extKey, err := hdkeychain.NewKeyFromString(test.key)
		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("NewKeyFromString #%d (%s): mismatched error "+
				"-- got: %v, want: %v", i, test.name, err,
				test.err)
			continue
		}

		if test.neuter {
			_, err := extKey.Neuter()
			if !reflect.DeepEqual(err, test.neuterErr) {
				t.Errorf("Neuter #%d (%s): mismatched error "+
					"-- got: %v, want: %v", i, test.name,
					err, test.neuterErr)
				continue
			}
		}
	}
}

// TestZero ensures that zeroing an extended key works as intended.
func TestZero(t *testing.T) {
	tests := []struct {
		name   string
		master string
		extKey string
		net    *chaincfg.Params
	}{
		// Test vector 1
		{
			name:   "test vector 1 chain m",
			master: "000102030405060708090a0b0c0d0e0f",
			extKey: "aprv2iiCsBe7i46kLtEczB2oM6A3Vw7DDefEtK3BZk4aXkJ2nthoVFm2AMFAFDR2xPwWRBjE21nMynz9dTAKY6G4vrS3nwrKSejK62X5dDhfxig",
			net:    &chaincfg.MainNetParams,
		},

		// Test vector 2
		{
			name:   "test vector 2 chain m",
			master: "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
			extKey: "aprv2iiCsBe7i46kLVjx6L75oHsok8FduGv4ZtpurgoQW8Qu5sizfQXxbFQkbHC68uVbfEUemJEpFu2LiytRaWemYDkpFMvZaze3fH9JzCcKCbC",
			net:    &chaincfg.MainNetParams,
		},
	}

	// Use a closure to test that a key is zeroed since the tests create
	// keys in different ways and need to test the same things multiple
	// times.
	testZeroed := func(i int, testName string, key *hdkeychain.ExtendedKey) bool {
		// Zeroing a key should result in it no longer being private
		if key.IsPrivate() != false {
			t.Errorf("IsPrivate #%d (%s): mismatched key type -- "+
				"want private %v, got private %v", i, testName,
				false, key.IsPrivate())
			return false
		}

		parentFP := key.ParentFingerprint()
		if parentFP != 0 {
			t.Errorf("ParentFingerprint #%d (%s): mismatched "+
				"parent fingerprint -- want %d, got %d", i,
				testName, 0, parentFP)
			return false
		}

		wantKey := "zeroed extended key"
		_, errZeroed := key.String()
		if errZeroed.Error() != wantKey {
			t.Errorf("String #%d (%s): mismatched serialized key "+
				"-- want %s, got %s", i, testName, wantKey,
				errZeroed)
			return false
		}

		wantErr := hdkeychain.ErrNotPrivExtKey
		_, err := key.ECPrivKey()
		if !reflect.DeepEqual(err, wantErr) {
			t.Errorf("ECPrivKey #%d (%s): mismatched error: want "+
				"%v, got %v", i, testName, wantErr, err)
			return false
		}

		wantErr = errors.New("pubkey string is empty")
		_, err = key.ECPubKey()
		if !reflect.DeepEqual(err, wantErr) {
			t.Errorf("ECPubKey #%d (%s): mismatched error: want "+
				"%v, got %v", i, testName, wantErr, err)
			return false
		}

		wantAddr := "AbC3Lvy49rPuwhnQHGPeFbEHg56Fw8Pf7oL"
		addr, err := key.Address(&chaincfg.MainNetParams)
		if err != nil {
			t.Errorf("Addres s #%d (%s): unexpected error: %v", i,
				testName, err)
			return false
		}
		if addr.EncodeAddress() != wantAddr {
			t.Errorf("Address #%d (%s): mismatched address -- want "+
				"%s, got %s", i, testName, wantAddr,
				addr.EncodeAddress())
			return false
		}

		return true
	}

	for i, test := range tests {
		// Create new key from seed and get the neutered version.
		masterSeed, err := hex.DecodeString(test.master)
		if err != nil {
			t.Errorf("DecodeString #%d (%s): unexpected error: %v",
				i, test.name, err)
			continue
		}
		key, err := hdkeychain.NewMaster(masterSeed, test.net)
		if err != nil {
			t.Errorf("NewMaster #%d (%s): unexpected error when "+
				"creating new master key: %v", i, test.name,
				err)
			continue
		}
		neuteredKey, err := key.Neuter()
		if err != nil {
			t.Errorf("Neuter #%d (%s): unexpected error: %v", i,
				test.name, err)
			continue
		}

		// Ensure both non-neutered and neutered keys are zeroed
		// properly.
		key.Zero()
		if !testZeroed(i, test.name+" from seed not neutered", key) {
			continue
		}
		neuteredKey.Zero()
		if !testZeroed(i, test.name+" from seed neutered", key) {
			continue
		}

		// Deserialize key and get the neutered version.
		key, err = hdkeychain.NewKeyFromString(test.extKey)
		if err != nil {
			t.Errorf("NewKeyFromString #%d (%s): unexpected "+
				"error: %v", i, test.name, err)
			continue
		}
		neuteredKey, err = key.Neuter()
		if err != nil {
			t.Errorf("Neuter #%d (%s): unexpected error: %v", i,
				test.name, err)
			continue
		}

		// Ensure both non-neutered and neutered keys are zeroed
		// properly.
		key.Zero()
		if !testZeroed(i, test.name+" deserialized not neutered", key) {
			continue
		}
		neuteredKey.Zero()
		if !testZeroed(i, test.name+" deserialized neutered", key) {
			continue
		}
	}
}
