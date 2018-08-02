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

echo "************install docker************"
sudo apt-get update
sudo apt-get install -y docker.io

echo "*************set up kubernetes apt-get source************"
sudo apt-get update && sudo apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update

read -p "Install kubelet (y/n)?" choice
case "$choice" in
  y|Y ) sudo apt-get install -y kubelet=1.11.1-00;;
esac
printf "\n"

read -p "Install kubeadm (y/n)?" choice
case "$choice" in
  y|Y ) sudo apt-get install -y kubeadm=1.11.1-00;;
esac
printf "\n"

read -p "Install kubectl (y/n)?" choice
case "$choice" in
  y|Y ) sudo apt-get install -y kubectli=1.11.1-00;;
esac
printf "\n"

echo "*************dry run to test kubeadm.conf************"
sudo kubeadm init --config kubeadm.conf --dry-run

read -p "Create kubernetees master(y/n)?" choice
case "$choice" in
  y|Y )
    sudo kubeadm init --config kubeadm.conf
    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config
    sudo chmod o+wr  $HOME/.kube/config
    ;;
esac
printf "\n"

read -p "Install calico network plugin (y/n)?" choice
case "$choice" in
  y|Y ) kubectl apply -f https://docs.projectcalico.org/v2.6/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml;;
esac
printf "\n"

read -p "Install helm (y/n)?" choice
case "$choice" in
  y|Y )
    wget https://storage.googleapis.com/kubernetes-helm/helm-v2.8.2-linux-amd64.tar.gz
    tar -zxvf helm-v2.8.2-linux-amd64.tar.gz
    chmod o+x linux-amd64/helm
    sudo mv linux-amd64/helm /usr/local/bin/helm
    rm -rf linux-amd64
    rm -rf helm-v2.8.2-linux-amd64.tar.gz

    kubectl create -f helm_service_account.yaml
    helm init --service-account tiller
    ;;
esac
printf "\n"
