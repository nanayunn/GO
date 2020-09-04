# 쿠버네티스



[TOC]



## 1. 클러스터 생성하기

* 쿠버네티스?
  * **클러스터**를 상호조정 => 
    * 서로 연결되었지만 단일 유닛처럼 동작
    * 고가용성( HA )
* 쿠버네티스의 배포 모델을 사용하기 위해서는 애플리케이션의 **컨테이너화**가 필요
  * **컨테이너화된 애플리 케이션**
    * 호스트에 매우 깊이 통합된 패키지
    * 특정 머신에 직접 설치되는 예전의 배포 모델보다 유연하고 가용성이 높다.
* **즉**, 쿠버네티스는?
  * 애플리케이션 컨테이너를 클러스터에 분산, 스케줄링
    * 이를 효율적으로 자동화.



* 클러스터의 구성

  * **마스터**

    * 클러스터 관리를 담당
      * 애플리케이션을 스케줄링, 항상성 유지, 스케일링, 변경사항 반영 등등..
  * 물리, 가상 머신 모두에 설치 가능
  
* **노드**( 작업자 )
  
    * VM 또는 물리적인 컴퓨터
    * 애플리케이션을 구동
    * Kubelet 에이전트 보유
      * 마스터와 통신
        * 마스터가 제공하는 Kubernetes API로 통신
    * Docker또는 rkt와 같은 툴 보유
      * 컨테이너 운영을 위한 툴
  * 
  
  
  
  * <img src="https://d33wubrfki0l68.cloudfront.net/99d9808dcbf2880a996ed50d308a186b5900cec9/40b94/docs/tutorials/kubernetes-basics/public/images/module_01_cluster.svg" alt="img" style="zoom: 80%;" />

### 1-1. 대화형 튜토리얼

```
minikube version

minikube start

kubectl version
# client version : kubectl version
# server version : Kubernetes version installed on the master

kubectl cluster-info
# Kubernetes master 노드가 어디서 실행되고 있는지 addr과 port가 나온다.

kubectl get nodes
# 애플리케이션 호스팅에 쓰이고 있는 node를 보여준다.
# Name, Status, Roles, Age, Version 컬럼
```





## 2. kubectl을 사용해서 디플로이먼트 생성하기



* 쿠버네티스 디플로이먼트

  * 쿠버네티스 클러스터를 구동하고, 컨테이너화된 애플리케이션을 배포하기 위한 설정
  * 인스턴스 생성 및 업데이트에 대한 지시사항을 쿠버네티스 마스터에게 전달

* 쿠버네티스 디플로이먼트 컨트롤러

  * 생성된 인스턴스 모니터링
  * 자동 복구 매커니즘( self-healing )
    * 구동 중인 노드가 다운되거나 삭제되면, 컨트롤러가 인스턴스를 클러스터 내부의 다른 노드의 인스턴스로 교체
    * 머신의 장애나 정비에 대응할 수 있는 매커니즘

  ![img](https://d33wubrfki0l68.cloudfront.net/152c845f25df8e69dd24dd7b0836a289747e258a/4a1d2/docs/tutorials/kubernetes-basics/public/images/module_02_first_app.svg)

* **kubectl**

  * 쿠버네티스 CLI
  * 디플로이먼트를 생성 및 관리
  * 클러스터와 상호작용 하기 위해 쿠버네티스 API 사용

  

* 애플리케이션이 쿠버네티스에서 배포되려면?
  
  * 지원되는 컨테이너 형식 중 하나로 패키지 되어야 함



### 2-1. kubectl basics

```
(kubectl create deployment)
# 쿠버네티스 상에 디플로이먼트 생성하는 명령어

$ kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1
# deployment.apps/kubernetes-bootcamp created

# 명령어 실행 시, 애플리케이션이 구동될 수 있는 적합한 노드를 찾는다.
# 노드를 찾은 뒤에, 애플리케이션이 노드에서 실행될 수 있도록 스케줄링한다
# 필요할 때 새로운 노드에서 인스턴스를 다시 스케줄링할 수 있도록 클러스터를 구성.

$ kubectl get deployments
#	NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
	kubernetes-bootcamp   1/1     1            1           71s
	
# 현재 구동시킨 디플로이먼트를 확인하는 명령어

$ kubectl proxy
# Starting to serve on 127.0.0.1:8001

# Pod는 쿠버네티스의 독립된 프라이빗 네트워크 내에서 실행된다.
# kubectl 명령어는 프라이빗 네트워크로 통신을 포워딩할 수 있도록 해주는 프록시를 생성할 수 있다.
# 즉, 호스트는 kubectl 명령어를 통해 API endpoint를 이용하여 Pod와 통신할 수 있는 것이다.
# 생성된 프록시는 어떤 출력도 나타내지 않으며, Ctrl+C로 종료할 수 있다.
# 프록시를 통해 Pod와 직접적으로 통신할 수 있다.

예 :
$ curl http://localhost:8001/version
# {
  "major": "1",
  "minor": "17",
  "gitVersion": "v1.17.3",
  "gitCommit": "06ad960bfd03b39c8310aaf92d1e7c12ce618213",
  "gitTreeState": "clean",
  "buildDate": "2020-02-11T18:07:13Z",
  "goVersion": "go1.13.6",
  "compiler": "gc",
  "platform": "linux/amd64"
   }
   
# curl 명령어로 Pod의 version을 알아내는 예시

```





## 3. Pod와 Node



* 쿠버네티스 Pod란?
  * 하나 또는 그 이상의 애플리케이션 컨테이너( 도커 또는 rkt )들의 그룹을 나타내는 추상적인 개념
  * 컨테이너에 대한 자원을 공유
    * 공유 스토리지
      * 볼륨
    * 네트워킹
      * 클러스터 IP 주소
    * 컨테이너가 동작하는 방식에 대한 정보
      * 컨테이너 이미지 버전
      * 컨테이너 사용 포트
  * "로컬 호스트" 애플리케이션 모형 생성
    * 밀접하게 결합된 상이한 애플리케이션 컨테이너들을 수용할 수 있게 된다.
    * 예시 : 
      * Node.js 앱 + Node.js 웹서버에 필요한 DB 컨테이너
      * Pod 내 컨테이너는 IP 주소와 포트 정보를 공유하고, 동일 컨텍스트를 공유하여 동작, 스케줄링 된다.
  * 쿠버네티스 플랫폼 상에서의 최소 단위
  * Pod의 생성 단계
    1. 쿠버네티스에서 디플로이먼트 생성
    2. 컨테이너 내부에서 컨테이너와 함께 Pod 생성
    3. 각 Pod는 스케줄 되어진 노드에 묶임
    4. 소멸, 삭제 전까지 노드에 유지
       * 노드에 문제가 생기면, 클러스터 내의 다른 노드들을 대상으로 스케줄링 된다.

![img](https://d33wubrfki0l68.cloudfront.net/fe03f68d8ede9815184852ca2a4fd30325e5d15a/98064/docs/tutorials/kubernetes-basics/public/images/module_03_pods.svg)

* 쿠버네티스 Node란?

  * 쿠버네티스의 워커 머신
  * 클러스터에 따라 가상, 또는 물리 머신
  * 각 노드는 마스터에 의해 관리
  * 하나의 노드는 여러 개의 파드를 가질 수 있음
  * 마스터는 클러스터 내 노드를 통해서 파드에 대해 스케줄링을 자동 처리
  * 노드와 관련된 동작
    * kubelet
      * 쿠버네티스 마스터와 노드 간 통신을 책임지는 프로세스
      * 하나의 머신 상에서 동작하는 파드와 컨테이너 관리
    * 컨테이너 런타임( 도커, rkt )
      * 레지스트리에서 컨테이너 이미지를 pull 
      * 이미지를 풀고, 애플리케이션을 동작시키는 책임을 맡음

  <img src="https://d33wubrfki0l68.cloudfront.net/5cb72d407cbe2755e581b6de757e0d81760d5b86/a9df9/docs/tutorials/kubernetes-basics/public/images/module_03_nodes.svg" alt="img" style="zoom: 10%;" />

* 배포된 애플리케이션과 환경에 대한 정보를 얻기 위한 kubectl 기본 명령어

  ```
  $ kubectl get
  # 자원을 나열한다
  
  $ kubectl describe
  # 자원에 대해 상세한 정보를 보여준다.
  
  $ kubectl logs 
  # 파드 내 컨테이너의 로그들을 출력한다
  
  $ kubectl exec 
  # 파드 내 컨테이너에 대한 명령을 실행한다.
  ```

  

### 3-1. 명령어로 애플리케이션 조사하기



```
$ kubectl describe pods
# pod 내의 컨테이너나 사용중인 이미지에 대해 조사할 때 명령어 이용
# *********** 명령어 이용 시 알 수 있는 pod 정보 ***********
Name:         kubernetes-bootcamp-765bf4c7b4-trk9g
Namespace:    default
Priority:     0
Node:         minikube/172.17.0.33
Start Time:   Wed, 26 Aug 2020 02:05:39 +0000
Labels:       pod-template-hash=765bf4c7b4
              run=kubernetes-bootcamp
Annotations:  <none>
Status:       Running
IP:           172.18.0.3
IPs:
  IP:           172.18.0.3
Controlled By:  ReplicaSet/kubernetes-bootcamp-765bf4c7b4
Containers:
  kubernetes-bootcamp:
    Container ID:   docker://bd33ae91b41fac4b167c39dff949d787fc3a952b582bc7c60c866040f87e6440
    Image:          gcr.io/google-samples/kubernetes-bootcamp:v1
    Image ID:       docker-pullable://jocatalin/kubernetes-bootcamp@sha256:0d6b8ee63bb57c5f5b6156f446b3bc3b3c143d233037f3a2f00e279c8fcc64af
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Wed, 26 Aug 2020 02:05:42 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-xcv7p (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-xcv7p:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-xcv7p
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason            Age    From               Message
  ----     ------            ----   ----               -------
  Warning  FailedScheduling  3m19s  default-scheduler  0/1 nodes are available: 1 node(s) had taints that the pod didn't tolerate.
  Normal   Scheduled         3m13s  default-scheduler  Successfully assigned default/kubernetes-bootcamp-765bf4c7b4-trk9g to minikube
  Normal   Pulled            3m11s  kubelet, minikube  Container image "gcr.io/google-samples/kubernetes-bootcamp:v1" already present on machine
  Normal   Created           3m11s  kubelet, minikube  Created container kubernetes-bootcamp
  Normal   Started           3m10s  kubelet, minikube  Started container kubernetes-bootcamp
  
*****************************************************************************************
# 이외에도 node, pod, deployment에 대한 정보를 얻고자 할 때, describe 명령어를 이용할 수 있다.



$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range.items}}{{.metadata.name}}{{"\n"}}{{end}}')
$ echo Name of the Pod: $POD_NAME
# Name of the Pod:kubernetes-bootcamp-765bf4c7b4-trk9g

# Pod의 이름을 통해서도 프록시를 이용할 수 있다. 
# Pod의 이름을 알아내어 환경변수로 저장하도록 하는 명령어
# go의 text/template을 이용하여 이름 값을 모두 가져오도록 한다.

$ curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/
# Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-765bf4c7b4-trk9g | v=1

# 이후 curl 요청을 보내면 다음과 같은 결과를 얻을 수 있다.
# curl 요청을 보내는 url은 Pod의 API에 대한 경로이다.


$ kubectl logs $POD_NAME
# Kubernetes Bootcamp App Started At: 2020-08-26T02:05:42.114Z | Running On:  kubernetes-bootcamp-765bf4c7b4-trk9g

Running On: kubernetes-bootcamp-765bf4c7b4-trk9g | Total Requests: 1 | App Uptime: 808.112 seconds | Log Time: 2020-08-26T02:19:10.226Z


# 애플리케이션이 일반적으로 STDOUT로 전송하는 모든 것은 포드 내의 컨테이너에 대한 로그로 남는다.
# kubectl logs 명령어로 이러한 로그들을 확인할 수 있다.

$ kubectl exec $POD_NAME
# Pod의 이름을 파라미터 값으로 하여 해당 pod와 관련된 명령어를 실행할 수 있다.

$ kubectl exec $POD_NAME env
#   PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
	HOSTNAME=kubernetes-bootcamp-765bf4c7b4-lcr6h
	KUBERNETES_SERVICE_HOST=10.96.0.1
	KUBERNETES_SERVICE_PORT=443
	KUBERNETES_SERVICE_PORT_HTTPS=443
    KUBERNETES_PORT=tcp://10.96.0.1:443
    KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
    KUBERNETES_PORT_443_TCP_PROTO=tcp
    KUBERNETES_PORT_443_TCP_PORT=443
    KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
    NPM_CONFIG_LOGLEVEL=info
    NODE_VERSION=6.3.1
    HOME=/root
# 위의 명령어같은 경우는 해당 pod의 환경변수 설정값을 불러온다.

$ kubectl exec -ti $POD_NAME bash 
# root@kubernetes-bootcamp-765bf4c7b4-lcr6h:/#
# -it 옵션으로 터미널에서 입력 가능한 옵션을 주고, pod의 bash에 접속한다.
# exit 명령어로 터미널 종료 가능




```









## 4. 외부 트래픽에 애플리케이션 노출

> Pod에게는 생명주기가 있다.
>
> 워커 노드가 죽으면, 노드 상에 동작하는 Pod도 함께 종료된다.
>
> 이러한 Pod의 생명주기에 영향 받지 않고, 애플리케이션이 지속적인 동작을 할 수 있도록 돕는 **ReplicaSet**이 존재한다.
>
> * Replicaset
>   * 새로운 Pod들을 생성하여 클러스터를 미리 지정해둔 상태로 동적으로 되돌린다.
>
> 



* 쿠버네티스에서 **서비스**란?

  * 하나의 논리적인 Pod set을 정의	

    * 종속적인 Pod 결합성을 느슨하게 해줌

  * 외부 트래픽 노출, 로드 밸런싱, Pod들에 대한 서비스 디스커버리를 가능하게 해주는 추상적 개념

  * YAML, JSON을 이용하여 정의됨

  * 서비스가 대상으로 하는 Pod 셋은 보통 LabelSelector에 의해 결정

  * ServiceSpec

    * 각 파드들이 갖고 있는 고유의 IP들을 클러스터 외부로 노출
    * => 애플리케이션들에게 트래픽이 실릴 수 있도록 허용

    * `type`의 종류
      * *ClusterIP* (기본값) 
        * 클러스터 내에서 내부 IP 에 대해 서비스를 노출
        * 이 방식은 오직 클러스터 내에서만 서비스가 접근될 수 있도록 해준다.
      * *NodePort* 
        * NAT가 이용되는 클러스터 내에서 각각 선택된 노드들의 동일한 포트에 서비스를 노출
        * `<NodeIP>:<NodePort>`를 이용하여 클러스터 외부로부터 서비스가 접근할 수 있도록 해준다. 
        * ClusterIP의 상위 집합
      * *LoadBalancer* - (지원 가능한 경우) 
        * 기존 클라우드에서 외부용 로드밸런서를 생성하고 서비스에 고정된 공인 IP를 할당
        * NodePort의 상위 집합
      * *ExternalName* 
        * 이름으로 CNAME 레코드를 반환함으로써 임의의 이름(스펙에서 `externalName`으로 명시)을 이용하여 서비스를 노출
        * 프록시는 사용되지 않는다.
        * 이 방식은 `kube-dns` 버전 1.7 이상에서 지원 가능하다.

  * `selector`없이 생성된 서비스는 상응하는 엔드포인트 오브젝트를 생성하지 않는다.

    * => 하나의 서비스를 특정한 엔드포인트에 매핑 시킬 수 있도록 해줌
    *  `type`을 *ExternalName*으로 설정한 경우, selector을 생략하게 됨



* 서비스와 레이블

  

  ![img](https://d33wubrfki0l68.cloudfront.net/cc38b0f3c0fd94e66495e3a4198f2096cdecd3d5/ace10/docs/tutorials/kubernetes-basics/public/images/module_04_services.svg)

  * **서비스**

    * Pod Set에 걸쳐서 트래픽을 라우트

    * discovery와 routing은 서비스에 의해 처리된다.

    * 레이블과 셀렉터

      > 레이블이란?
      >
      > * 오브젝트들에 붙여진 key/value 쌍
      > * 오브젝트의 생성 시점, 또는 이후 시점에 붙여지며, 언제든지 수정 가능

      * 쿠버네티스의 객체들에 대해 논리 연산을 허용해주는 기본 그룹핑 단위

      * 이를 이용하여 쿠버네티스의 객체들과 Pod Set을 매치

      * 서비스 활용 예시

        * 개발, 테스트, 그리고 상용환경에 대한 객체들의 지정

        * 임베디드된 버전 태그들

        * 태그들을 이용하는 객체들에 대한 분류

          ![img](https://d33wubrfki0l68.cloudfront.net/b964c59cdc1979dd4e1904c25f43745564ef6bee/f3351/docs/tutorials/kubernetes-basics/public/images/module_04_labels.svg)



### 4-1. 앱 노출하기

```
$ kubectl get services
# NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
  kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   33s
# minikube 실행 시, default로 생성되는 서비스를 확인


$ kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
# service/kubernetes-bootcamp exposed
# kubernetes-bootcamp Pod를 8080포트의 NodePort 타입으로 트래픽 노출


# 다시 get services를 하면 방금 노출시킨 Pod의 서비스가 목록에 나타난다.
# NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
  kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP          2m37s
  kubernetes-bootcamp   NodePort    10.109.173.59   <none>        8080:30551/TCP   6s
  
  
$ kubectl describe services/kubernetes-bootcamp
#	Name:                     kubernetes-bootcamp
	Namespace:                default
    Labels:                   run=kubernetes-bootcamp
    Annotations:              <none>
    Selector:                 run=kubernetes-bootcamp
    Type:                     NodePort
    IP:                       10.109.173.59
    Port:                     <unset>  8080/TCP
    TargetPort:               8080/TCP
    NodePort:                 <unset>  30551/TCP
    Endpoints:                172.18.0.4:8080
    Session Affinity:         None
    External Traffic Policy:  Cluster
    Events:                   <none>
# 외부 트래픽이 잘 할당되었는지 확인하기 위한 명령어

$ export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
$ echo NODE_PORT=$NODE_PORT
# NODE_PORT=30551
# kubernetes-bootcamp가 할당받은 port값을 갖고 있는 NODE_PORT라는 환경변수를 새로이 생성

$ curl $(minikube ip):$NODE_PORT
# Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-765bf4c7b4-z8k9r | v=1
# 저장한 환경변수 값과 minikube의 ip 값을 이용하여 curl 명령어를 실행하면,
# kubernetes-bootcamp 노드가 외부에 트래픽 노출이 되었는지 확인할 수 있다.


# deployment를 실행하면, Pod의 label이 자동으로 생성된다.
$ kubectl describe deployment
#   Name:                   kubernetes-bootcamp
    Namespace:              default
    CreationTimestamp:      Wed, 26 Aug 2020 08:29:45 +0000
    Labels:                 run=kubernetes-bootcamp
    ...
    Pod Template:
      Labels:  run=kubernetes-bootcamp
    ...
# describe 명령어를 이용해서 label값을 확인할 수 있다.


$ kubectl get pods -l run=kubernetes-bootcamp
#   NAME                                   READY   STATUS    RESTARTS   AGE
    kubernetes-bootcamp-765bf4c7b4-z8k9r   1/1     Running   0          21m
# -l 옵션으로 Pod의 label 값을 함께 넘겨주면 해당 pod의 상태값을 받을 수 있다.

$ kubectl get services -l run=kubernetes-bootcamp
#   NAME                  TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    kubernetes-bootcamp   NodePort   10.109.173.59   <none>        8080:30551/TCP   22m
# service에 대해서도 똑같이 label값으로 해당 service를 검색할 수 있다.


$ kubectl label pod $POD_NAME app=v1
# pod/kubernetes-bootcamp-765bf4c7b4-z8k9r labeled
# $POD_NAME은 이전에 환경변수로 미리 저장해둔 kubernetes-bootcamp
# app=v1이라는 이름으로 kubernetes-bootcamp를 라벨링

# 라벨링 후, pod를 다시 describe 명령어로 살펴보면, 
$ kubectl describe pods $POD_NAME
#   Name:         kubernetes-bootcamp-765bf4c7b4-z8k9r
    Namespace:    default
    Priority:     0
    Node:         minikube/172.17.0.35
    Start Time:   Wed, 26 Aug 2020 08:29:50 +0000
    Labels:       app=v1
                  pod-template-hash=765bf4c7b4
                  run=kubernetes-bootcamp
                  ...
# 다음과 같이 라벨링 된 것을 확인할 수 있다.


$ kubectl delete service -l run=kubernetes-bootcamp
# service "kubernetes-bootcamp" deleted
# 서비스 삭제 시에도 라벨을 이용할 수 있다.
# 서비스 삭제 이후에는 클러스터 외부에서 접속할 수 없게된다.

# deployment는 여전히 실행 중이므로, exec 명령어를 이용하여 애플리케이션이 실행 중인것은 확인할 수 있다.
# 애플리케이션을 완전히 중지하고 싶다면, deployment 또한 삭제해야한다.
```







## 5. 앱 스케일링하기

* deployment의 replica 갯수를 변경하면 스케일링이 수행된다.
* <img src="https://d33wubrfki0l68.cloudfront.net/043eb67914e9474e30a303553d5a4c6c7301f378/0d8f6/docs/tutorials/kubernetes-basics/public/images/module_05_scaling1.svg" alt="img" style="zoom: 25%;" />

![img](https://d33wubrfki0l68.cloudfront.net/30f75140a581110443397192d70a4cdb37df7bfc/b5f56/docs/tutorials/kubernetes-basics/public/images/module_05_scaling2.svg)

* 스케일링을 통해 인스턴스를 복수로 구동하게 되면, 트래픽 분산이 필요해진다
  * => 로드 밸런서  서비스
* 다수의 애플리케이션 구동을 통해 다운타임 없이 롤링 업데이트가 가능해짐



### 5-1. 앱 스케일링하기

```
$ kubectl get deployments
# NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
  kubernetes-bootcamp   1/1     1            1           11m
# 활성화 되어있는 deployments의 목록

$ kubectl get rs
# 생성되어있는 ReplicaSet 목록 출력
# replicaset의 이름 형태 : [DEPLOYMENT-NAME]-[RANDOM-STRING]

$ kubectl scale deployments/kubernetes-bootcamp --replicas=4
# replicaset의 갯수를 scale 명령어로 조정, 
# 원하는 갯수 만큼 replica를 생성한다.


$ export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
# 샘플 노드의 포트 정보를 환경 변수로 저장

# 동일한 명령어를 이용해서 replicaset의 갯수를 줄일 수도 있다.

$ curl $(minikube ip):$NODE_PORT
# Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-765bf4c7b4-t5rfm | v=1
# Hello Kubernetes bootcamp! | Running on: kubernetes-bootcamp-765bf4c7b4-znkvf | v=1
# 로드밸런싱이 되고있음을 확인할 수 있다.
```





## 6. 앱 업데이트 수행하기

* 롤링 업데이트

  * Pod 인스턴스를 점진적으로 새로운 것으로 업데이트

  * deployment의 무중단 업데이트

  * 이전 버전으로의 롤백

  * 로드 밸런스

    * 여러 개의 인스턴스가 있을 때, 업데이트가 이루어지고 있는 Pod를 제외하고, 오직 사용 가능한 Pod에게만 트래픽을 로드 밸런싱 한다.

    1. ![img](https://d33wubrfki0l68.cloudfront.net/30f75140a581110443397192d70a4cdb37df7bfc/fa906/docs/tutorials/kubernetes-basics/public/images/module_06_rollingupdates1.svg)
    2. ![img](https://d33wubrfki0l68.cloudfront.net/678bcc3281bfcc588e87c73ffdc73c7a8380aca9/703a2/docs/tutorials/kubernetes-basics/public/images/module_06_rollingupdates2.svg)
    3. ![img](https://d33wubrfki0l68.cloudfront.net/9b57c000ea41aca21842da9e1d596cf22f1b9561/91786/docs/tutorials/kubernetes-basics/public/images/module_06_rollingupdates3.svg)
    4. ![img](https://d33wubrfki0l68.cloudfront.net/6d8bc1ebb4dc67051242bc828d3ae849dbeedb93/fbfa8/docs/tutorials/kubernetes-basics/public/images/module_06_rollingupdates4.svg)

### 6-1. 앱 업데이트 하기

```
$ kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2
# v1 테스트 이미지를 v2 테스트 이미지로 업데이트 하는 명령어

$ kubectl get pods
# 위의 명령어를 실행한 뒤 pods의 정보를 불러오면, 이전의 pods들이 모두 중지되고, 새롭게 생성된 pods들이 실행되고 있는 것을 확인할 수 있다.

$ kubectl rollout undo deployments/kubernetes-bootcamp
# rollout undo 명령어를 이용해서 Pod를 업데이트한 내역을 롤백 시킬 수 있다.
```







