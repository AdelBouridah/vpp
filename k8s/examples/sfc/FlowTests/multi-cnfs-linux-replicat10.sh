#!/bin/bash
set -eu

getPodNames() {
        pods=`kubectl get pods | grep linux-cnf`
        rpcnf1=`kubectl get pods |  grep -oE "rep-linux-cnf1\S*"`
	rpcnf2=`kubectl get pods |  grep -oE "rep-linux-cnf2\S*"`
	rpcnf3=`kubectl get pods |  grep -oE "rep-linux-cnf3\S*"`
	rpcnf4=`kubectl get pods |  grep -oE "rep-linux-cnf4\S*"`
}

checkReadiness() {
        echo "Checking pod readiness...(up to 30sec)"
        duration=0
        until [[ `kubectl get pods | grep linux-cnf` != *"0/1"* ]];do
                echo "Pods not ready, waiting..."
                duration=$((duration+5))
                if [ "${duration}" -gt "25" ];then
                        echo "Timed out waiting for pods to be ready."
                        exit 1
                fi
                sleep 5
        done
        echo "Pods ready."
}

setuprpcnf1() {
	declare -i i=0
	for line in $rpcnf1; do  
		i=i+1; 
		kubectl exec $line -- ip address add 192.168.187.1/24 dev tap1
		kubectl exec $line -- ip link set dev tap1 up
		echo "Configured cnf1 with IP 192.168.187.1 for copy"$i;  
	done	
}

setuprpcnf2() {
	declare -i i=0
	for line in $rpcnf2; do  
		i=i+1; 
		kubectl exec $line -- brctl addbr br1
		kubectl exec $line -- brctl addif br1 tap1
		kubectl exec $line -- brctl addif br1 tap2
		kubectl exec $line -- ip link set dev br1 up
		kubectl exec $line -- ip link set dev tap1 up
		kubectl exec $line -- ip link set dev tap2 up
		echo "Configured cnf2 with linux bridge."
	done
}


setuprpcnf3() {
	declare -i i=0
	for line in $rpcnf3; do  
		i=i+1; 
		kubectl exec $line -- brctl addbr br1
		kubectl exec $line -- brctl addif br1 tap1
		kubectl exec $line -- brctl addif br1 tap2
		kubectl exec $line -- ip link set dev br1 up
		kubectl exec $line -- ip link set dev tap1 up
		kubectl exec $line -- ip link set dev tap2 up
		echo "Configured cnf3 with linux bridge."
	done
}

setuprpcnf4() {
	declare -i i=0
	for line in $rpcnf4; do  
		i=i+1; 
		kubectl exec $line -- ip address add 192.168.187.2/24 dev tap1
		kubectl exec $line -- ip link set dev tap1 up
		echo "Configured cnf4 with IP 192.168.187.1 for copy"$i;  
	done
}

testConnectivity() {
	declare -i i=0
	for line in $rpcnf1; do  
		i=i+1; 
		kubectl exec $line -- ping -c 1 -I tap1 192.168.187.2
		echo "Direction 1 - Connectivity verified for path";
		echo $i;
		  
	done
	i=0
	for line in $rpcnf4; do  
		i=i+1; 
		kubectl exec $line -- ping -c 1 -I tap1 192.168.187.1
		echo "Direction 2 - Connectivity verified for path";
		echo $i; 
	done

}
testFlowsOverPaths(){
	# Get First and last CNFs of each path
	declare -i i=0 j
	declare -a cnfbegintab cnfendtab
	for line in $rpcnf1; do 
		cnfbegintab[i]=$line 
		i=i+1; 		  
	done
	i=0
	for line in $rpcnf4; do  
		cnfendtab[i]=$line 
		i=i+1; 
	done
	i=0
	while (($i<100)); 	do
		j=(RANDOM%3)
		kubectl exec ${cnfbegintab[$j]} -- ping -c 20 -I tap1 192.168.187.2
		echo "Ping Path $j Direction 1 - Connectivity verified for path";
		i=$i+1;
	done 
}
getPodNames
checkReadiness
setuprpcnf1
setuprpcnf2
setuprpcnf3
setuprpcnf4
testConnectivity
: <<'END_COMMENT'
END_COMMENT

