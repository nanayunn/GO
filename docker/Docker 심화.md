# Docker 심화

![docker](https://i.imgur.com/5DJjNOW.png)

## 1. Docker 작동 원리

* 컨테이너

  * 애플리케이션이 동작하기 위해서 필요한 요소들을 패키지화, 격리하는 기술

    * 필요한 요소 :  실행 파일, 애플리케이션 엔진 등등..

    

  * 커널 기반의 기술 ( Cgroup, namespace )을 이용해서 프로세스를 완벽하게 격리, 분리된 환경 조성

  * **Cgroup이란?**

    * Control Group

    * 시스템의 자원을 제한, 격리할 수 있는 커널 기능

      * 자원 : 

        * 시스템의 CPU 시간, 시스템 메모리, 네트워크 대역폭

        * ![스크린샷, 2020-08-13 09-36-53](https://user-images.githubusercontent.com/58680504/90081604-b9a1e480-dd48-11ea-91a8-164c9459b9b0.png)

        * 각 서브 시스템에 대한 설명

          * | 서브시스템 |                            설 명                             |
            | :--------: | :----------------------------------------------------------: |
            |   blkio    |               Block Device 의 입출력 접근 제한               |
            |    cpu     |      CPU에 cgroup 작업 액세스를 제공하기 위해 스케줄러       |
            |  cpuacct   | cgroup의 작업에 사용된 CPU 자원에 대한 보고서를 자동으로 생성 |
            |   cpuset   | 개별 CPU (멀티코어 시스템에서) 및 메모리 노드를 cgroup의 작업에 할당 |
            |  devices   |  cgroup의 작업 단위로 장치에 대한 액세스를 허용하거나 거부   |
            |  freezer   |          cgroup의 작업을 일시 중지하거나 다시 시작           |
            |  net_cls   |  특정 cgroup에서 발생하는 패킷을 식별하기 위해 태그를 지정   |
            |  net_prio  |  cgroup의 작업에서 생성되는 네트워크 트래픽의 우선순위 지정  |
            |   memory   |     cgroup의 작업에서 사용되는 메모리에 대한 제한을 설정     |

    

  * **namespace란?**

    * 시스템 리소스를 프로세스의 전용 자원처럼 보이게끔 함
    * 다른 프로세스와 격리시키는 기능
    * namespace 종류
      * Mount namespacaes : 
        * 파일시스템의 Mount 를 분할하고 격리합니다.
      * PID namespacaes : 
        * 프로세스를 분할 관리합니다.
      * Network namespacaes : 
        * Network 관련된 정보를 분할 관리합니다.
      * IPC namespacaes : 
        * 프로세스간 통신을 격리합니다.
      * UTS namespacaes : 
        * 독립적인 hostname 할당합니다
      * USER namespacaes : 
        * 독립적인 UID를 할당합니다.



### Docker가 이용한 컨테이너 기술의 구현체

* LXC, LIbContainer, runC

  * cgroups, namespaces를 표준으로 정의해둔 OCI( Open Container Initative ) 스펙을 구현한 컨테이너 기술의 구현체

    

* ~ Docker 1.8

  * LXC 를 이용하여 구현

* 이후 libcontainer 이용

* Docker 1.11 버전 이후 : 

  * runC( libcontainer의 리팩토링 구현체 )
    * system 에서 container 관련된 기능들에 대해 docker 가 쉽게 사용할 수 있도록 해주는 가볍고 이식가능한 container runtime
  * 자체 구현체를 갖게 됨!

  <img src="https://miro.medium.com/max/3156/1*uEtAhWOHVMFo5kRUnhKNBg.png" alt="img" style="zoom: 33%;" />

* 기본적으로 구동되는 Docker daemon

  * dockerd

  * containerd

  *  => 두 개를 합쳐서 `Docker engine`

    

  * containerd : 

    * OCI 구현체( runC )를 이용해 컨테이너를 관리해주는 daemon

  * Docker engine :

    * 이미지, 네트워크, 디스크 등의 관리 역할 

  * containerd와 Docker engine은 완전히 분리되어있다!

    * Docker engine 버전 업그레이드를 할 때 Docker engine만 재시작하면 됨

* `docker run`으로 컨테이너 기동을 요청하면 일어나는 과정

  1. dockerd 는 요청을 gRPC 를 통해 containerd 로 전달한다.

  2. containerd 는 exec 을 통해 containerd-shim 을 자식으로 생성한다.

  3. containerd-shim 은 runc 를 이용하여 container 를 생성한다.
     (runc 는 container 가 정상적으로 실행되면 exit 한다.)

  4. containerd-shim 은 그대로 살아있으며, 이는 container 내에서 실행되는 process 들의 부모가 된다.

     

  * containerd-shim이 필요한 이유?

    1. daemonless container 를 위해서!	

       * container 하나 뜬다고 해서 뭔가 계속 수행되는 runtime daemon 이 존재할 필요가 없다.

         (runc 는 이미 exit 했다.)

       * containerd-shim 의 경우는 껍데기나 마찬가지

       * container 를 위한 STDIO 및 fd 를 계속 유지

    2. dockerd 와 containerd 가 둘다 죽게 되는 상황이 되면 pipe 의 한쪽이 닫혀서 container 까지 죽을 수 있는데 이를 막아준다.
       * => 즉, docker daemon 들의 장애 상황이 container 까지 전파되지 않도록 해준다.
       * 이는 container 에게 영향을 주지 않고 docker engine 의 upgrade 작업도 수행할 수 있게 해준다.

    3. container 의 exit status 등을 higher level tool 로 보고
       * containerd-shim -> runc -> container 인 상황에서 runc 는 exit
         * container 의 parent 는 containerd-shim
         * 따라서, containerd-shim 은 자식의 상태를 파악하고 이를 어딘가로 보고해줄 수 있는 구조가 된다.

![img](http://cloudrain21.com/wordpress/wp-content/uploads/2019/09/257FE535595AF79817.png)

![img](https://cameronlonsdale.com/assets/img/docker-work/architecture_2019.svg)





## 2. Docker 네트워크 구성









## 3. Docker 파일 시스템



* Docker가 사용하는 파일 시스템
  * **UFS( Union File System )**

