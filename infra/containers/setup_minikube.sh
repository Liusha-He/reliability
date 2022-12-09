#!/bin/bash

# step 1 - Update system
sudo apt update
sudo apt install apt-transport-https
sudo apt upgrade -y

[ -f /var/run/reboot-required ] && sudo reboot -f

# Step 2: Install KVM or VirtualBox Hypervisor
sudo apt install virtualbox virtualbox-ext-pack

# Step 3: Download minikube on Ubuntu 22.04|20.04|18.04
wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
chmod +x minikube-linux-amd64
sudo mv minikube-linux-amd64 /usr/local/bin/minikube

minikube version

# Step 4: Install kubectl on Ubuntu
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl

chmod +x ./kubectl

sudo mv ./kubectl /usr/local/bin/kubectl

kubectl version -o json  --client

