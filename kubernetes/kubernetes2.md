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


$ export POD_NAME=$(kubectl get pods -o go-template --template '{{range.items}}{{.metadata.name}}{{"\n"}}{{end}}')
$ echo Name of the Pod: $POD_NAME
# Name of the Pod:

# Pod의 이름을 통해서도 프록시를 이용할 수 있다. 
# Pod의 이름을 알아내어 환경변수로 저장하도록 하는 명령어

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
  * 

