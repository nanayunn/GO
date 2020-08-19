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
  * 간접적으로 커널 기술들을 컨트롤

* 이후 libcontainer 이용

  * docker 진영에서 커널의 가상화 기술을 다루기 위한 인터페이스

* Docker 1.11 버전 이후 : 

  * runC( libcontainer의 리팩토링 구현체 )
    * system 에서 container 관련된 기능들에 대해 docker 가 쉽게 사용할 수 있도록 해주는 가볍고 이식가능한 container runtime
  * 자체 구현체를 갖게 됨!

  

### 3. Docker Engine 

* Docker Engine의 구성
  * binary
  * daemon process

![스크린샷, 2020-08-18 09-37-11](https://user-images.githubusercontent.com/58680504/90457455-88555a00-e136-11ea-98b7-40833c9496f7.png)



#### 각 실행파일 별 세부 설명

* **docker**

  * CLI( Command Line Interface )를 수행하기 위한 binary

    * 사용자가 docker engine으로 다양한 명령을 요청할 때 사용하는 프로그램

  * CLI 작동 방식

    * `docker CLI` `=>` `REST API wrapper` `=>` `docker daemone`

  * docker client와 server( docker daemon ) 간의 통신 방식

    * Docker 명령어 수행

      * => Docker CLI( Command Line Interface ) 를 이용하는 것

        * Docker CLI는 내부적으로 **unix domain socket**을 사용

          * `/var/run/docker.sock` 파일을 사용하여 통신

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

      * 또는 fd, tcp socket을 이용하여 원격으로 명령 가능

        * fd 방식 : 
          * systemd base인 system에서만 사용 가능
          * Systemd socket activation 기능 이용
        * tcp 방식 :
          * remote에서 접속을 해야할 경우



* **docker-init**
  * 컨테이너 내에서 첫번째로 기동,
    * HOST의 init 프로세스처럼 동작하도록 만들어진 프로그램
  * default binary : 
    * /bin/docker-init
  * `docker run` 수행 시,
    *  `-init` 옵션이 주어지지 않을 경우 :
      * container 내에서 init process를 별도로 기동하지 않음
      * 즉, docker run 수행 시 넘겨준 command가 **1번 프로세스**가 된다.
    * `-init` 옵션이 주어질 경우:
      * *init 프로세스*를 container 구동 후 **1번 프로세스**로 기동
  * **init process 존재 여부의 중요성?**
    * child process를 받아주어 resource의 누수나 zombie process의 생성 등을 방지



* **docker-proxy**
  * `-p` 옵션을 통해 특정 port를 외부로 publish 할 때 기동
    * host와 연결된 container의 포트로 `forwarding`
  * tcp 통신 시, port:port 옵션 내의 port 변환 수행이 필요.
    * `ip:port` 형식의 테이블로 forwarding할 수 있는 정보를 관리
  * 컨테이너를 시작할 때, 외부에 노출하도록 설정된 port의 숫자만큼 생성된다.
  * **docker host에서 컨테이너로 패킷 전송은 사실, docker iptables rules에 의해 docker-proxy 없이도 포트 포워딩이 가능하다.**
    * docker daemon이 자동으로 제어
    * docker-proxy를 사용하는 경우?
      * docker host가 iptables 의 NAT를 사용하지 못하는 상황에 대한 처리
        * 정책상의 이유로 docker host의 iptables나 ip_forward를 이용하지 못하는 경우
        * docker host의 internal network (172.17.0.0/16) 에서는 현재의 NAT 설정을 통해서 노출된 port(여기서는 8080) 로 포워딩이 불가능
  * docker host의 iptable 확인 방법
    * ![스크린샷, 2020-08-19 13-29-33](https://user-images.githubusercontent.com/58680504/90592246-1b1df380-e220-11ea-9e24-e01b86b06a4e.png)
  
  
  
  
  
  * 작동 예시
    * ![인터넷의 Docker-가상 머신이없는 네트워킹-localhost.svg에 게시](https://info.itemis.com/hubfs/Blog/Docker/Networking-ohne-VM-publish-to-localhost.svg)
    * `browser -> (중간에 docker-proxy) -> docker0(172.17.0.1)  -> eth0(172.17.0.2) 8080` 



* **dockerd와 containerd**

  <img src="https://miro.medium.com/max/3156/1*uEtAhWOHVMFo5kRUnhKNBg.png" alt="img" style="zoom: 33%;" />

  * docker 서비스가 맨 처음 시작하면 가동되는 **daemon** process들

  * => 두 개를 합쳐서 `Docker engine`

    

  * *dockerd* : 

    * `dockerd` 는 컨테이너를 지속적으로 관리하는 데몬 프로세스

    * client 로부터 **REST API** 형식의 요청을 수신하여 처리

      * unix, tcp, fd의 세 가지 소켓 유형을 통해, 도커 API 요청을 수신

      

  * *containerd* : 

    * 독립 실행형 고수준(high-level) 컨테이너 런타임

      * 이미지를 `push` 하고 `pull` 하고, 스토리지를 관리하고, 네트워크 기능을 정의

    * container 의 lifecycle 을 관리

    * client 로부터의 container 관리 관련 요청은 dockerd 를 거쳐 **gRPC** 통신을 통해 containerd 로 전달

      > gRPC란?
      >
      > * g (oogle) Remote Procedure Call
      >
      >   * 원격 프로세스의 함수를 마치 로컬의 함수처럼 사용할 수 있는 프로토콜 또는 인터페이스
      >
      > * Protocol Buffers를 사용하여 서비스를 정의, 사용
      >
      >   * protocol buffers란?
      >
      >     * 직렬화 데이터 구조( Serialized Data Structure )
      >
      >       1. proto file로 데이터 타입을 정의
      >
      >          * 특정 언어에 종속성이 없는 형태
      >
      >          * ```protobuf
      >            syntax = "proto2";
      >            
      >            package tutorial;
      >            
      >            message Person {
      >              required string name = 1;
      >              required int32 id = 2;
      >              optional string email = 3;
      >            
      >              enum PhoneType {
      >                MOBILE = 0;
      >                HOME = 1;
      >                WORK = 2;
      >              }
      >            
      >              message PhoneNumber {
      >                required string number = 1;
      >                optional PhoneType type = 2 [default = HOME];
      >              }
      >            
      >              repeated PhoneNumber phones = 4;
      >            }
      >            
      >            message AddressBook {
      >              repeated Person people = 1;
      >            }
      >            ```
      >
      >       2. 프로그래밍 언어에서 사용하기 위해 해당 언어에 맞는 형태의 데이터 클래스로 생성
      >
      >          * protoc 컴파일러로 proto file을 컴파일
      >
      >          * 컴파일 결과 : `<파일이름_pb2.py>`
      >
      >          * ```python
      >            class Person(message.Message):
      >              __metaclass__ = reflection.GeneratedProtocolMessageType
      >            
      >              class PhoneNumber(message.Message):
      >                __metaclass__ = reflection.GeneratedProtocolMessageType
      >                DESCRIPTOR = _PERSON_PHONENUMBER
      >              DESCRIPTOR = _PERSON
      >            
      >            class AddressBook(message.Message):
      >              __metaclass__ = reflection.GeneratedProtocolMessageType
      >              DESCRIPTOR = _ADDRESSBOOK
      >            ```
      >
      >       3. 각 언어에 맞는 형태의 데이터 클래스 파일 생성
      >
      >          * 응용 예제 : 
      >
      >          * ```python
      >            import addressbook_pb2
      >            
      >            person = addressbook_pb2.Person()
      >            person.id = 1234
      >            person.name = "John Doe"
      >            person.email = "jdoe@example.com"
      >            phone = person.phones.add()
      >            phone.number = "555-4321"
      >            phone.type = addressbook_pb2.Person.HOME
      >            ```
      >
      >          * 다음과 같은 Error가 일어날 수 있다.
      >            1. AttributeError:
      >               * proto file에서 정의되어있지 않은 타입에 값을 입력할 때
      >            2. TypeError:
      >               * proto file에서 정의된 변수의 타입과 일치하지 않는 값을 입력할 때
      >
      > * HTTP 2.0를 기반으로 구성
      >
      > * 장점
      >
      >   * gRPC 채널
      >     * **HTTP Connection 재사용**
      >   * 멀티플렉싱
      >     * gRPC는 기본적으로 하나의 Connection에서 여러 요청을 보낼 수 있다.
      >     * 데이터에 대한 우선순위 설정 가능
      >     * HTTP 1.0 버전에서의 Head of Line Blocking( HoLB ) 문제 해결 가능
      >   * JSON 기반의 통신보다 더 가볍고 그만큼 통신 속도가 빠름
      >     * HTTP 2.0의 헤더 압축 기능 사용 => 전송 데이터 크기 축소
      >   * **internal** 통신이 빈번한 마이크로 서비스 구조에서 성능의 이점
      >     * 서버 <=> 클라이언트 간 양방향 데이터 전송 가능
      >
      >   ![img](https://github.com/alicek106/naver-blog-img-repo/blob/master/pictures/grpc-polyglot.png?raw=true)

    * OCI 구현체( runC )를 이용해 컨테이너를 관리해주는 daemon

      

  * docker.service를 start하면:

    * docker daemon이 자동으로 containerd까지 기동
      * service 관련 설정은 `/etc/systemd/system/multi-user.target.wants/docker.service` 파일에서 확인하여 기동
    * stop은 각각 해야한다.

  * containerd도 **별도**로 구동이 가능

    * dockerd가 containerd와 통신할 수 있도록 containerd의 **socket path**를 알려주어야 함.
    * 예:
      * `dockerd --containerd /run/containerd/containerd.sock`
    * docker.sock 파일( docker CLI <=> docker daemon )과 containerd.sock( docker daemon <=> containerd ) 파일
      * /var/run/( /run 디렉토리에 대한 심볼릭 링크 ) , 혹은 /run 디렉토리 하위에 존재



* **containerd-shim, runC**

  <img src="http://cloudrain21.com/wordpress/wp-content/uploads/2019/09/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA-2019-09-05-%E1%84%8B%E1%85%A9%E1%84%92%E1%85%AE-10.11.10-1024x605.png" alt="img" style="zoom:50%;" />

  * runC 

    * OCI 런타임 스펙을 구현하고 있는 저수준 컨테이너 런타임

      * 리눅스 커널 버전마다 달라지는 cgroup, namespace 기술 이용 방법에 대처하기 위한 매개체
      * 과거에는 이것에 대응하기 위한 방법으로 Libvirt, LXC( Linux Container ).. 과 같은 중간 매개체 이용
      * 오직 실행 중인 컨테이너 관리에만 그 범위를 집중

    * 컨테이너를 생성(`spawning`)과 실행(`running`) 할 수있는 CLI로 구현되어 있음

    * runtime 확인하는 명령어

      * `docker info | grep -i runtime`

        

  * containerd-shim

    * `runc`를 실행하고, 컨테이너 프로세스를 제어하는 경량 데몬
    * 컨테이너와 `containerd` 의 모든 통신은 `containerd-shim` 을 통해서 이루어진다.
    * 맡는 역할 : 
      * 컨테이너의 stdout 및 stderr의 스트림을 제공
        * 로그 파일로 저장 가능
        * `containerd` 가 재시작 중에도 문제가 발생하지 않는다.
      * 컨테이너 프로세스를 `shim` 프로세스가 관리
        * `runc` 는 컨테이너 프로세스를 실행(fork)한 다음, 포그라운드 프로세스를 종료하여, 컨테이너 프로세스를 의도적으로 데몬화
          * => 컨테이너 프로세스를 호스트의 `init` 프로세스가 담당하게 된다.
            * => 컨테이너의 관리가 어려워짐
            * `shim` 프로세스를 `subreaper`로 만들어서 관리





* `docker run`으로 컨테이너 기동을 요청하면 일어나는 과정

  1. dockerd 는 요청을 gRPC 를 통해 containerd 로 전달한다.
  2. containerd 는 exec 을 통해 containerd-shim 을 자식으로 생성한다.
  3. containerd-shim 은 runc 를 이용하여 container 를 생성한다.
     (runc 는 container 가 정상적으로 실행되면 exit 한다.)
  4. containerd-shim 은 그대로 살아있으며, 이는 container 내에서 실행되는 process 들의 부모가 된다.

  

![img](https://cameronlonsdale.com/assets/img/docker-work/architecture_2019.svg)







* **Docker는 HTTP 프로토콜을 이용하여 json으로 통신**

<img src="http://blog.daocloud.io/wp-content/uploads/2014/12/003_docker_server.jpg" alt="img" style="zoom:33%;" />

* Docker daemon과 Registry 사이의 인터페이스도 **HTTP**로 통신
  * Docker register API를 이용
* Docker 이미지를 관리하는 Registry
  * Docker hub : public registry ( `docker pull` 수행 시 default registry )
  * Other public registry
  * Private registry



### 4. Docker 아키텍쳐 그려보기

![스크린샷, 2020-08-18 17-45-44](https://user-images.githubusercontent.com/58680504/90491306-b8bee780-e17a-11ea-9c75-faf59648d6c8.png)









## 2. Docker Network Drivers

* 도커를 설치하면, 기본적으로 `docker0` 이라는 가상 브리지(bridge)가 생성
  
  * `docker0`?
  
    *  연결된 다른 네트워크 인터페이스 간에 패킷을 자동으로 전달하는 가상 이더넷 브리지
  
  * CIRD 표기법으로 172.17.0.1/16 의 주소 범위
  
    *  172.17.0.1 부터 172.17.255.254 까지의 아이피
  
    

### 컨테이너 게이트웨이

* 컨테이너 안에서 `route` 명령어를 실행
*  컨테이너 내부의 모든 패킷은 `default` 인 172.17.0.1로 전송됨



### 1. Bridge Mode Networking

> 리눅스에서 Bridge 인터페이스란?
>
> * => 하드웨어 브릿지를 소프트웨어로 구현한 것
>
> * MAC( Media Access Address )을 기반으로 네트워크를 구성
>   * MAC과 IP를 맵핑하는 ARP 테이블을 만들어서 패킷을 전송
> * 일종의 L2 스위치 역할
>   * L2 ( 데이터 링크 계층 )의 네트워크를 연결
>   * L2 레벨의 패킷 전송



* 동일한 Docker daemon host 위의 컨테이너에 적용하기 좋음
* 서로 연관없는 컨테이너들을 하나의 브릿지 네트워크로 묶어두는 것은 좋지 않음
  * 어떠한 영향을 끼치게 될지 모르기 때문에
  * => User-defined 브릿지 네트워크를 생성하여 별도의 브릿지 네트워크로 관리하는 것이 좋음



* 특징
  * 컨테이너 기본 네트워크 설정
    * `docker0`을 이용하여 내부 네트워크에 연결
  * 컨테이너를 분리된 네트워크 네임스페이스에 배치
  * 네트워크 주소 변환을 사용하여 여러 컨테이너 간에 호스트의 외부 IP 주소를 공유
    * 브릿지에 연결된 컨테이너들은 private ip가 할당
    * 컨테이너 => 인터넷 통신을 하기 위해서는?
      * 마스커레이드( Masquerade ) == **NAT**
        * 패킷 포워딩
        * 내부 사설 IP, port를 외부 통신용 IP와 port로 변환
  * 브릿지가 다른 컨테이너끼리는 통신이 불가능
    * host에 이에 대한 rule을 자동적으로 설정



* 장점
  * 동일한 호스트에서 여러 컨테이너를 실행할 때 네트워크 포트 충돌을 일으키지 않는다.
    * 즉, 동일한 포트를 사용하는 다수의 컨테이너를 하나의 호스트에서 실행 가능
  * 각 컨테이너는 호스트와 분리된 전용 네트워크 네임스페이스를 소유



* 단점 : 
  * NAT의 사용으로 인해 네트워크 처리량과 지연 시간에 영향
  * 호스트와 컨테이너 간의 네트워크 포트 매핑을 제어해야 함



![img](http://postfiles12.naver.net/MjAxOTAyMjFfMTcy/MDAxNTUwNjc5MDc0NDg1.oeK6S2DX-Ki0lUtsOPQ9no0HJStukSJVFZtD7MwY-Tcg.JTGmT4YKTzFtu6u3QS7JEfJS2TEIJNfFpU2kqVzgIvIg.PNG.jevida/022019_1611_DockerNetwo1.png?type=w580)

* 브릿지 네트워크가 설정되는 순서
  1. 컨테이너 생성
  2. 컨테이너를 위해서 페어 인터페이스 생성
     * 페어 인터페이스 ( pair interface )
       * eth0과 veth
  3. 내부 namespace에는 eth0, `docker0`브리지에는 veth가 바인딩된다.



* User-Bridge 네트워크 생성하기
  * `docker network create --driver bridge <브리지 이름>`
  * 예 : 
    * `docker network create --driver bridge mybridge`
  * 생성 브릿지는 --net 설정을 통해 사용 가능
    * `docker run -i -t --name mybridge_container --net mybrdige ubuntu:18.04`





### 2. Host Mode Networking



* 함께 쓰이지 않는 컨테이너을 대상으로 컨테이너와 호스트의 네트워크 격리를 없애기 위한 모드
* 17.06 버전 이상의 Docker와 Swarm에서 이용할 수 있다.
* 컨테이너가 관리해야하는 port가 많고, 네트워킹에서 NAT를 필요로 하지 않을 때 유용



* 특징
  * 컨테이너가 호스트의 네트워킹 네임스페이스를 공유
  * 외부 네트워크에 직접 노출
  * 호스트의 IP 주소와 호스트의 TCP 포트 공간을 사용
    * 컨테이너 내부에서 실행 중인 서비스를 노출



* 장점
  * 간단한 네트워크 모드
    * 이해하기 쉽고 사용하기 쉬움



* 단점

  * 호스트 네트워크를 그대로 사용

    * 동일한 네트워크 포트 사용 시 충돌

      * 동일한 포트를 사용하는 다수의 컨테이너를 하나의 호스트에서 실행할 경우, 포트 충돌로 인한 서비스 장애 발생

      

<img src="https://www.kangwoo.kr/wp-content/uploads/2020/08/docker-host-mode-networking.png" alt="img" style="zoom: 67%;" />

* 사용 명령어 예
  * `docker run -i -t --rm --net=host --name network_host ubuntu:18.04`



### 3. None Mode Networking



* 네트워크 연결을 하지 않을 때 사용 
  * 이 연결을 한 컨테이너는 외부와 단절
    * 예:
      * 컨테이너에는 lo 스위치만 나타남
  * 주로 사용자의 커스텀 네트워크 드라이버를 이용하기 위해 사용되는 모드



* 사용 명령어 예:
  * `docker run -i -t --rm --name network_none --net none ubuntu:18.04`







### 4. Container Mode Networking



* 특징
  * 다른 컨테이너의 네트워크 환경을 공유할 수 있음
  * 공유되는 속성: 
    * 내부 IP, 네트워크 인터페이스의 MAC 주소 등
  * 따라서, 해당 네트워크 설정을 이용하여 생성되는 컨테이너에게는 별도의 네트워크 자원이 할당되지 않는다.



* 사용 명령어 예:
  * `--net container:<다른 컨테이너 이름 또는 아이디>`



### 5. Overlay Mode Networking



* 서로 다른 호스트 위에서 생성된 컨테이너들 간의 통신이 필요할 때
* Swarm과 같은 서비스를 바탕으로 다수의 애플리케이션 구동이 필요할 때.
* 일종의 '터널링'이라고 불리며, single network pool을 형성하였다고 본다.



* 특징
  * 다수의 Docker daemon 호스트들 간의 분산 네트워크 구축
  * **암호화**가 가능할 때, 컨테이너들이 안전하게 통신할 수 있도록 해준다.
  * Swarm 서비스와 연관이 깊은 네트워크
    * ingress
      * Swarm 서비스에서 데이터 트래픽을 컨트롤해주는 overlay network
    * docker_gwbridge
      * Swarm 서비스 내에서 각기 다른 Docker daemon을 연결해주는 브릿지 네트워크



### 6. Macvlan Mode Networking



* VM 환경설정을 그대로 옮길 때, 
* Docker 호스트의 네트워크 스택을 통해 라우팅되지 않고, 물리 네트워크에 바로 연결시키고자 할 때
* 리눅스 SubInterface와 가상화에서 제공되는 Guest OS VLAN Taggin을 조합하여 사용하는 방식



* 특징
  * 컨테이너에게 MAC 주소를 할당
    * 네트워크 상에서 물리적인 장치로 인식되게끔 한다.
  * Docker daemon이 컨테이너들의 MAC 주소를 기반으로 라우팅을 함
  * macvlan 드라이버를 사용하기 위해 지정해주어야 하는 값:
    * 물리 인터페이스
    * 서브넷
    * macvlan의 gateway



* 주의사항

  * 적절하지 않은 다수의 유니크한 MAC 주소를 네트워크 내에 존재할 경우

    * VLAN spread, 혹은 IP 주소 소진( IP address exhaustion )으로 인해 네트워크에 영향을 줄 수 있다.

  * 네트워킹 장비가 "promicuous mode( 무차별 모드 )"를 핸들링 할 수 있어야 한다.

    > promicuous mode란?
    >
    > * 랜 카드가 항상 broadcasting을 하는 모드
    >
    > * 다른 호스트의 주소로 전송되는 이더넷 프레임을 폐기하지 않고 상위계층으로 전달하는 모드
    >   * 일반적인 NIC 디바이스 드라이버는 자신이 아닌 MAC 주소로 보내진 이더넷 프레임을 확인하지 않고 폐기
    > * MAC 주소의 형태에 상관없이 이더넷 프레임을 수신
    >   * CPU의 부하를 불러옴

    * 하나의 물리 인터페이스에 다수의 MAC 주소가 할당되기 때문

  * 브릿지와 오버레이 네트워크를 이용중이라면, 장기간으로 보았을 때, MACvlan 네트워킹이 더 효율적일 수 있다.



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








