#!/bin/bash
# Script to remove lustre packages from nodes and delete debug pods (may cause alerts in AKS)
nodes=$(kubectl get nodes --output=jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}')
echo "${nodes}"
for node in $nodes; do
    echo "Cleaning ${node}"
    kubectl debug node/${node} --image=ubuntu:20.04 -- /bin/bash -c "uptime;whoami; (chroot /host /bin/bash -c 'apt autoremove *lustre* -y')"
done

# # now cleanup the errant pods
debug_pods=$(kubectl get pods --no-headers=true | awk '/^node-debugger/{print $1}')
echo -e "\n\nDebug pods found: ${debug_pods}"
read -rp "Press [yY] to delete all debug pods: " _input
if [[ "$_input" =~ ^([yY][eE][sS]|[yY])+$ ]]; then
    echo "Deleting debug pods"
    kubectl delete pods ${debug_pods}
else
    echo "Exiting"
    exit 1
fi
