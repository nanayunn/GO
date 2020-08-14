# Docker 심화

![docker](https://i.imgur.com/5DJjNOW.png)



[TOC]

## Docker 전체 구성도

![img](http://www.studytrails.com/wp-content/uploads/2018/12/Docker_Architecture_hierarchy.png)

## 1. Docker 작동 원리



### 1. Cgroup & namespace

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



### 2. Docker가 이용한 컨테이너 기술의 구현체

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





### 3. Docker 실행 시 내부에서 사용하는 Protocol과 Registry

> CLI에서 명령 시, `docker run`, `docker pull`, `docker create` 등의 명령어는 
>
> 어떠한 형태로 docker engine에 전달되고 수행되는가?



<img src="http://www.studytrails.com/wp-content/uploads/2018/12/docker_architecture.jpg" alt="Docker Architecture" style="zoom:50%;" />

* Docker 명령어 수행

  * => Docker CLI( Command Line Interface ) 를 이용하는 것

    * Docker CLI는 내부적으로 **unix domain socket**을 사용

      > unix domain socket이란?
      >
      > * UDP 통신 이용
      >
      >   * Unix domain에서의 UDP는 일종의 *PIPE*와 같은 형태
      >   * 동일한 시스템 영역에서의 데이터 교환
      >     * 즉, Internet socket의 UDP와 달리 패킷 유실이나 순서가 바뀌는 문제가 없다.
      >
      > * 메세지 전달 방식
      >
      >   * 패킷 전달 형식
      >   * 헤더 + 바디 형태의 데이터
      >   * "지정된 파일"을 통해서 이루어짐
      >     * 헤더영역 : 통신을 위한 파일 정보
      >
      > * 서버/클라이언트 구성 방식
      >
      >   * Domain 영역의 UDP는 경로 지정이 불가능
      >     * 기본적으로 가질 수 있는 정보다 "파일 경로"뿐이기 때문
      >   * => 다중 클라이언트를 대상으로 하는 서버의 경우 문제!!
      >   * 해결 방안 : 각 클라이언트마다 통신을 위한 각각의 소켓 지정 파일을 생성
      >     * TCP 소켓이 다중 연결을 받아들이기 위해 endpoint(:12) 소켓을 두는 것 처럼, 
      >     * 하나의 endpoint 파일을 두는 것..!
      >
      >   ![스크린샷, 2020-08-14 10-52-33](https://user-images.githubusercontent.com/58680504/90203681-53869180-de1c-11ea-9a9e-56179f842cdf.png)
      >
      >   

  * 또는 tcp socket을 이용하여 원격으로 명령 가능

  * 결론적으로, **Docker는 HTTP 프로토콜을 이용하여 json으로 통신**

<img src="http://blog.daocloud.io/wp-content/uploads/2014/12/003_docker_server.jpg" alt="img" style="zoom:33%;" />

* Docker daemon과 Registry 사이의 인터페이스도 **HTTP**로 통신
  * Docker register API를 이용
* Docker 이미지를 관리하는 Registry
  * Docker hub : public registry ( `docker pull` 수행 시 default registry )
  * Other public registry
  * Private registry





### 4. Binary와 daemon process



* Docker Binary

  



## 2. Docker 네트워크 구성









## 3. Docker 파일 시스템



### 1. Image Layer와 UFS

* Docker가 사용하는 파일 시스템

  * **UFS( Union File System )**

  * 여러 개의 File system을 하나로 결합하여 취급할 수 있도로 해주는 것

    > * Union Mount
    >
    >   * 읽기 전용 파일을 실행할 경우,
    >
    >     1. 해당 파일에 대해 쓰기가 가능한 임시 파일을 생성하여, 읽기 파일을 그대로 복제하여 실행
    >
    >     2. 수정된 내용에 대해 모두 사용하였다면, 쓰기 작업 진행
    >     3. 모두 완료되면 기존의 읽기 파일 대체

    * 프로그램 Binary : Docker Image
    * Process : Conatainer

  * Container 내부에서 바라보는 File System => 여러 개의 File System이 겹쳐져 있는 상태

    * Image를 이용하여 가동되는 Container은 여러 개의 Image가 겹쳐져 있는 상태
    * =>  특정 Image를 이용하여 container을 기동한다는 것은?
      * 차곡차곡 쌓은 Image layer을 하단에( Read-Only Image layer ) 깔아둔다.
      * 최종적으로 맨 위에 Container Layer을 쌓는다. ( Writable Layer )
      * 모든 Layer을 **결합**하여 사용자에게 하나의 **File System View( root filesystem )**으로 제공.

<img src="https://docs.docker.com/storage/storagedriver/images/sharing-layers.jpg" alt="img" style="zoom: 80%;" />

* Container에서 발생된 Diff는 Writable Layer에 쌓이고, Container가 제거되면 함께 제거됨

  * Diff를 유지하고자 한다면

    1. Container 상에서 변경을 수행한 후, image로 만들어 유지 

       ( 관리적인 측면 : Image의 변경은 없을 수록 좋음 )

    2. Volume을 구성하여 변경사항을 해당 Volume에 저장하도록 하여 유지

       ( 일반적으로 좋은 방법 )

    

  * Writable Layer가 아닌, Read-only Layer를 변경하고자 할 때는?

    * **Copy-on-Write** 기법 이용
    * image의 기존 내용을 변경할 경우, 
      * 해당 data를 Writable Layer로 copy한 후 이를 변경!
    * Copy-on-Write 기법의 단점
      1. 성능 하락 ( Performance Overhead )
         * data를 복제한 후 변경을 수행하므로..
      2. 용량 차지 ( disk space )
         * 하위 layer에 있는 원본 data + 상위 layer에 변경된 data 
         * 두 가지의 data를 모두 유지해야하므로 디스크 용량을 많이 차지하게 된다.



### 2. Storage Driver

> Docker에서 UFS를 구현하기 위해 어떠한 기술을 이용하는가?



* Docker Storage Driver

  * UFS를 구현

  * container 내의 파일 I/O 처리를 담당

    * 위의 Copy-on-write 기법을 수행하는 것도 이 드라이버

  * Pluggable한 구조

    * 원하는 Storage Driver를 선택할 수 있음

    * 각 OS 별 지원하는 Storage Driver

      <img src="http://cloudrain21.com/wordpress/wp-content/uploads/2019/09/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA-2019-09-04-%E1%84%8B%E1%85%A9%E1%84%92%E1%85%AE-10.28.11-1024x473.png" alt="img" style="zoom:80%;" />

  * Host filesystem을 Backing fileSystem으로 사용할 수 있음

    * 각 Storage Driver 별로 지원하는 Backing filesystem

      ![img](http://cloudrain21.com/wordpress/wp-content/uploads/2019/09/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA-2019-09-04-%E1%84%8B%E1%85%A9%E1%84%92%E1%85%AE-10.30.27-1024x392.png)





### 3. Container Size

> UFS와 Image Layer를 이용하여 실행되는 Container의 크기는 어떻게 결정되는가?

* Container Size 계산 시 고려해야할 것 
  1. virtual size : 
     * read-only size + writable size
  2. size : 
     * writable size

* 특정 image 를 이용하여 기동된 다수의 container 들이 사용하는 공간의 합

  * **(virtual size – size) + (size of all containers)**

    > 공통적으로 사용하는 read-only layer 크기 + 각 container 에서 변경한 사항들(writable layer)의 총합

  * container들이 사용하는 image와 변경사항의 저장 장소

    * 호스트 File System의 `/var/lib/docker` 디렉토리 내
      * Docker area / Backing Filesystem 이라고도 부른다.

















