#!/bin/sh
#
# Copyright 2018 ZTE, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

wget https://github.com/istio/istio/releases/download/1.0.0/istio-1.0.0-linux.tar.gz
tar -zxvf istio-1.0.0-linux.tar.gz
rm -rf istio-1.0.0-linux.tar.gz
sudo cp istio-1.0.0/bin/istioctl /usr/bin/
rm -rf istio-1.0.0

kubectl apply -f istio-auth.yaml
