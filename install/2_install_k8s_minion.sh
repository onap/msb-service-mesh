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
  y|Y ) sudo apt-get install -y kubectl=1.11.1-00;;
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
    ;;
esac
printf "\n"

cat << EOF
########################################################################################################################
1. You can now join this machines by running "kubeadmin join" command as root:
Please note that this is just an example, please refer to the output of the "kubeamin init" when cteating the k8s master for the exact comand to use in your k8s cluter!!!
  kubeadm join 10.12.6.108:6443 --token 43utwe.inl7h8dxn26p26iv --discovery-token-ca-cert-hash sha256:54cc1bcf72218de70c6b98edf4d486f79fb6d921a92ac5b7d10c76dbf96d006f

2. If you would like to get kubectl talk to your k8s master, you need to copy the dministrator kubeconfig file from your master to your workstation like this:

scp root@<master ip>:/etc/kubernetes/admin.conf .
kubectl --kubeconfig ./admin.conf get nodes

or you can manually copy the content of this file to ~/.kube/conf if scp can't be used due to security reason.
########################################################################################################################

EOF
