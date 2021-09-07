// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main KCC network.
var MainnetBootnodes = []string{
	"enode://d16cf20214ae987c0b200252df331633a680866e75fec511ae25573c06585a8ce2d7379526c9ca186501fbcf1af92443a178f7a800049947a178d5a7ba130566@13.230.226.55:30303",  // kcc-mainnet-node-boot-01
	"enode://0a3e227c2babb0f63e7fbb942f6c1abf00e47dc2e8b970b8c10b6c99dba7ac5eb223a11259eae50d195f88d27d3779a9b7b97618ca2afa5932c7ffdfabc91c31@35.74.148.173:30303",  //kcc-mainnet-node-boot-02
	"enode://17cdac7246184e3c406e62d6dbd61f8ffcfbf7c4e652695b19f184b6d2d9c0ddc8afa745c80371fbc486462c9eb0ac897f044dbcdcff629839b3621cabc7a265@34.97.13.219:30303",   //kcc-mainnet-node-boot-03
	"enode://58c66d5d9bdd369fecb8f4d3a081e02f7506e2b950aeeeed65e03ff4cbfd8177eac31c80c9f730b150f92e482828879d44e862ace9013d2f9bd318b42edf4989@34.97.171.225:30303",  //kcc-mainnet-node-boot-04
	"enode://b36a0aabf8e4e01aeb953d8b87991f164451b4b6ec810e477f633ca7982c5dc34ec025fc4b18e3b89536801388b3950f362bd675d3c20f89c39f43e22d44bdac@35.194.173.113:30303", //kcc-mainnet-node-boot-05
	"enode://848b74725913fa96041226a34a44c4ab9e4903ab25f91533ecb2c525fea2b28154d0ea126d0b30dd8a227abbc1941a7a8fe1dd50c71bda452e8a8f1c12ef688c@35.198.210.190:30303", //kcc-mainnet-node-boot-06
	"enode://fcad93335b50f0a73068879d8b51df4f86ca7249294d45ed12c4fff17285a01b6d979547d58915c4a6e7de8eb653ca29e3b367bd2550e5f4c90a9713d5ad521b@35.226.11.40:30303",   //kcc-mainnet-node-boot-07
	"enode://f5bdccf125fb971f3ae8a0d18691a219892a140cfbbf78b886ce631724212902cd8b4143575faa48811d06ce9fe3cd888c9b4d3c5c61a01e9d94c3d705bb7366@35.243.247.243:30303", //kcc-mainnet-node-boot-08
	"enode://446d3ab031578915d90b07b49fbc26f7b99a48f670ced4985d53d413bed49e0d332585dd5ee27ad3fe782aa2ccfe0bbd5973d70272fe75e8a75642133de8280c@34.82.155.210:30303",  //kcc-mainnet-node-boot-09
	"enode://c3b18858f03546f130bb6f0ba677f33f4087acd23bb1da05bc76738c3545afa6c4d69686a16c74f78437c26733140ff984a597532301d0b2752321d00ca70971@35.187.66.121:30303",  //kcc-mainnet-node-boot-10
	"enode://ef65959234a42fa02572d4ce65805ca58963289df172d54e2b078240bbdfef0cc5040b614b5aa6c72bd6486ad43f392f59b7367fa58ade1f1b5f88b4c3c95b8b@35.234.99.7:30303",    //kcc-mainnet-node-boot-11
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://8c5f9c393d84c671608ce8b259a53810629019059944e7ca1ca6fcd2a5842e167bc0115241fb1e6336158f11ce8ac2169338211d3c9a4c9751788642faf7d091@161.35.11.58:30303", // kcc-testnet-node-boot-01
	"enode://8c5f9c393d84c671608ce8b259a53810629019059944e7ca1ca6fcd2a5842e167bc0115241fb1e6336158f11ce8ac2169338211d3c9a4c9751788642faf7d091@127.0.0.1:30303",
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
