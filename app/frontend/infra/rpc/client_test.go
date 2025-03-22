// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"
	"testing"

	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/user"
)

func Test_iniUserClient(t *testing.T) {
	initUserClient()
	resp, err := UserClient.Login(context.Background(), &user.LoginReq{
		Email:    "x@163.com",
		Password: "0000",
	})
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	t.Logf("resp: %v", resp)
}
