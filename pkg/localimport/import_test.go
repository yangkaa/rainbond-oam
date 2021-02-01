// RAINBOND, Application Management Platform
// Copyright (C) 2020-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package localimport

import (
	"os"
	"testing"

	"github.com/docker/docker/client"
	"github.com/goodrain/rainbond-oam/pkg/ram/v1alpha1"
	"github.com/sirupsen/logrus"
)

func TestImport(t *testing.T) {
	c, _ := client.NewEnvClient()
	im := New(logrus.StandardLogger(), c, "/tmp/ram/springboot")
	info, err := im.Import("/Users/barnett/Downloads/若依SpringBoot-3.2.zip", v1alpha1.ImageInfo{
		HubPassword: os.Getenv("PASS"),
		Namespace:   "test",
		HubURL:      "image.goodrain.com",
		HubUser:     "root",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)
}
