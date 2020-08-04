* kubernetes 명령어
* kubectl



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





## 2. 