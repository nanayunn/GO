# Docker 도커

<img src="https://subicura.com/assets/article_images/2017-01-19-docker-guide-for-beginners-1/animals.jpg" alt="Tux(linux) - Moby Dock(docker) - Gopher (golang)" style="zoom: 50%;" />

[TOC]

## 1. Docker란?

* 2013년 3월 Solomon Hykes가 The future of Linux Containers라는 세션을 발표하며 나타남



* 컨테이너 기반의 오픈소스 가상화 플랫폼
  * 컨테이너란?
    
    * **격리**된 공간에서 프로세스가 동작하는 기술
    
    * 가상화 기술 중 하나!
    
      * 하이버 바이저를 추가로 로드할 필요가 없음
      * 호스트 시스템의 커널 내에서 직접 실행됨
    
      > 기존의 가상화 방식 : 
      >
      > * 전가상화 : 
      >   * **OS를 가상화**
      >   * KVM( Kernel-based Virtual Machine )
      >
      > * 반가상화 :
      >   * Xen
      >
      > => OS를 새로이 설치한다는 점에서 성능문제 발생
      >
      > <img src="https://subicura.com/assets/article_images/2017-01-19-docker-guide-for-beginners-1/vm-vs-docker.png" alt="가상머신과 도커" style="zoom: 33%;" />
    
    * 실행중인 컨테이너에 접속하여 명령어 입력 가능
    * `apt-get`이나 `yum`으로 패키지 설치 가능



### Docker의 발전

* LXC( Linux container )으로 시작
* 0.9 버전 + ( libcontainer 기술 )
* runC 기술에 합쳐짐



### Docker에서 이미지( Image )란?

* 컨테이너 실행에 필요한 파일과 설정값 등을 포함하고 있는 것

* 상태값을 가지지 않으며 변하지 않음( immutable )

  * 즉, 컨테이너를 만들기 위한 읽기 전용 템플릿
  * 변경값이나 상태값은 이미지를 실행한 컨테이너가 가지는 것 

  <img src="https://subicura.com/assets/article_images/2017-01-19-docker-guide-for-beginners-1/docker-image.png" alt="Docker image" style="zoom: 33%;" />



* 이미지 저장 방식

  * **레이어 저장 방식**

  * 도커의 이미지

    * 컨테이너를 실행하기 위한 모든 정보를 포함 ( 수백 메가의 용량 )
      * **레이어 ( Layer )**라는 개념 사용
    * 여러개의 읽기 전용 레이어로 구성,
      * 파일이 추가, 수정되면 새로운 레이어 생성
    * 컨테이너 생성 때도 레이어 방식 사용
      * 기존의 이미지 레이어 위에 read-write 레이어 추가
      * 컨테이너 실행 중 생성, 변경된 내용을 read-write 레이어에 추가

    <img src="https://subicura.com/assets/article_images/2017-01-19-docker-guide-for-beginners-1/image-layer.png" alt="Docker Layer" style="zoom: 33%;" />



* 이미지 관리 방식
  * **url 방식**으로 관리
  * 태그를 붙일 수 있음
  * 예 )
    * `docker.io/library/ubuntu:14.04`에서
    * `docker.io/library/`는 생략 가능



### Docker File

```dockerfile
# vertx/vertx3 debian version
FROM subicura/vertx3:3.3.1
MAINTAINER chungsub.kim@purpleworks.co.kr

ADD build/distributions/app-3.3.1.tar /
ADD config.template.json /app-3.3.1/bin/config.json
ADD docker/script/start.sh /usr/local/bin/
RUN ln -s /usr/local/bin/start.sh /start.sh

EXPOSE 8080
EXPOSE 7000

CMD ["start.sh"]
```

* 이미지를 만들기 위해 존재하는 파일
  * 이미지 생성 과정을 적기 위한 파일
* DSL( Domain-Specific-Language ) 언어를 이용하여 작성



### 이 외 Docker의 장점

* Docker hub
  * 공개 이미지를 무료로 관리
* Command & API
  * 명령어
    * 직관적, 사용하기 좋음
  * http 기반의 Rest API
    * 확장성이 좋음
* 발전속도가 빠른 오픈소스







## 2. Docker 엔진 및 아키텍쳐

* Docker 엔진과 아키텍쳐?

  * 클라이언트-서버 애플리케이션
  * 클라이언트-서버 아키텍쳐

  

  * `<Docker engine>`

  ![도커 엔진 구성 요소 흐름](https://docs.docker.com/engine/images/engine-components-flow.png)

  

  * `<Docker Architecture>`

  <img src="https://docs.docker.com/engine/images/architecture.svg" alt="도커 아키텍처 다이어그램" style="zoom:50%;" />

  

  * 구성 요소

    * docker daemon

      * 장기 실행 프로그램 유형의 서버
      * 이미지, 컨테이너, 네트워크 및 볼륨과 같은 Docker 객체를 만들고 관리

      

    * Rest API

      * 프로그램이 데몬과 통신하고, 수행할 작업을 지정해줄 때 사용하는 인터페이스를 지정해줌

      

    * Client

      * command Line( CLI )
      * 둘 이상의 daemon과 통신 가능

      

    * docker registry

      * docker의 이미지를 저장
      * Docker hub( default registry )
        * 공개 registry
      * 개인 registry도 실행 가능
      * 이미지 불러오기:
        * `docker pull`
      * 이미지로 컨테이너 생성:
        * `docker run`
      * 이미지가 있는 registry로 푸쉬
        * `docker push`

      

    * docker 객체

      *  이미지, 컨테이너, 네트워크, 볼륨, 플러그인 및 기타 객체를 생성 및 사용
      * 

  

  * 작동 방식
    * CLI로 docker daemon을 제어 및 상호작용
      * Docker Rest API를 사용하여 스크립팅
      * 직접적인 CLI 명령



* Docker로 작업하였을 때..
  1. 애플리케이션의 빠르고 일관된 제공
     * 개발자가 표준화된 환경( 컨테이너 )에서 작업할 수 있도록 함
       * 개발 수명주기를 간소화
     * CI / CD ( Continuous Integratoin / Continuous Delivery ) 워크 플로우에 적합
  2. 반응형 배포 및 확장 
     * 휴대성이 뛰어난 워크로드
       * 다양한 환경에서 실행 가능
       * 업무 요구에 따라 실시간으로 워크로드를 동적 관리, 확장, 분리 가능
  3. 동일한 하드웨어에서 더 많은 워크로드 실행
     * 다른 가상화 대비 더 많은 컴퓨팅 용량을 효율적으로 사용
     * 중소 규모 배포에 적합
       * 고밀도 환경과 적은 리소스









