//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

type JsonCommand struct {
	Id               string
	RmrMessageType   string
	SendCommandId    string
	ReceiveCommandId string
	TransactionId    string
	RanName          string
	Meid             string
	RanIp            string
	RanPort          int
	PayloadHeader    string
	PackedPayload    string
	Payload          string
	Action           string
	RepeatCount      int
	RepeatDelayInMs  int
}