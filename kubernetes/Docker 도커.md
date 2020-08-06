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
  * ![img](https://tech.osci.kr/assets/images/91690167/20.png)
  
* `docker build -t [이미지 이름]`

  * Dockerfile경로에서 해당 명령어를 입력하여 이미지 생성

* DSL( Domain-Specific-Language ) 언어를 이용하여 작성

  

* DSL 문법

  * FROM 
    * 기반이 되는 이미지 레이어
    * <이미지 이름>:<태그> 형식으로 작성 

    > ex) ubuntu:14.04

    

  * MAINTAINER 
    
  * 메인테이너 정보
  

  
  * RUN 
  * 도커이미지가 생성되기 전에 수행할 쉘 명령어
  
  
  
  * VOLUME 
    * 디렉터리의 내용을 컨테이너에 저장하지 않고 호스트에 저장하도록 설정
  * 데이터 볼륨을 호스트의 특정 디렉터리와 연결하려면 docker run 명령에서 -v 옵션을 사용
  
  > ex) -v /root/data:/data
  
  
  
  * CMD
    * 컨테이너가 시작되었을 때 실행할 실행 파일 또는 셸 스크립트
  * 해당 명령어는 DockerFile내 1회만 사용 가능
  
  
  
  * WORKDIR  
  
    * CMD에서 설정한 실행 파일이 실행될 디렉터리

  
  
  * EXPOSE 
    * 호스트와 연결할 포트 번호



* .dockerignore
  * Docker 이미지 생성 시 이미지안에 들어가지 않을 파일을 지정하는 파일





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

      * 이미지, 컨테이너, 네트워크, 볼륨, 플러그인 및 기타 객체를 생성 및 사용
      
        

  

  ![img](https://tech.osci.kr/assets/images/91690167/7.png)

  * 작동 방식

    * 코드개발
      
    * Dockerfile 생성

    * **Dockerfile Image** 생성

    * Container Orchestrator를 통한 배포

    * Container run

    * Container **image Push**

      

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





## 3. Docker 설치 및 기본 설정

### 1. Docker 설치하기

* 리눅스
* `curl -fsSL https://get.docker.com/ | sudo sh`



### 2. Docker 권한 설정

* `sudo` 권한 없이 Docker 명령어 이용하기

  * 다시 로그인 후 적용된다.

  * ```
    sudo usermod -aG docker $USER # 현재 접속중인 사용자에게 권한주기 
    sudo usermod -aG docker your-user # your-user 사용자에게 권한주기 
    ```

  * 권한 적용이 안되었다면 다음 명령어 실행

  * `sudo gpasswd -a $USER docker`

  * 또 `permission denied` 에러가 발생한다면

  * `sudo chown "$USER":"$USER" /home/"$USER"/.docker -R` 



### 3. Docker version 확인

* `docker version`  명령어 실행

* Docker Engine에서 앞서 본 Docker 엔진과 아키텍쳐 설명과 같이 Client & Server 버전이 함께 나옴.

  <img src="https://user-images.githubusercontent.com/58680504/89364254-11b76600-d70d-11ea-93aa-91a42a9a0c95.png" alt="스크린샷, 2020-08-05 10-54-25" style="zoom: 80%;" />



### 4. Hello-World Docker 테스트

* `docker run hello-world`

  ![스크린샷, 2020-08-05 11-10-30](https://user-images.githubusercontent.com/58680504/89364255-12e89300-d70d-11ea-89e0-118fb9540e97.png)

  * hello-world 이미지가 존재하지 않으므로 Docker hub에서 pull 해온다.

    * `tag` 옵션을 안주면 default 값은 `:latest`

  * 이미지 pulling이 모두 끝나면 다음과 같이 Hello가 뜬다.

    ![스크린샷, 2020-08-05 11-10-44](https://user-images.githubusercontent.com/58680504/89364256-13812980-d70d-11ea-8745-8550c0feefe7.png)

  * `docker ps -a`로 실행 후 종료된 `hello-world` 컨테이너를 확인 가능

    * `name`을 따로 지정해주지 않아서 랜덤으로 지어졌다..

    <img src="https://user-images.githubusercontent.com/58680504/89364258-1419c000-d70d-11ea-8f78-bfcc2ff6fc47.png" alt="스크린샷, 2020-08-05 11-11-07" style="zoom: 400%;" />

    

    

### 5. 프로젝트 이미지 빌드 테스트

* Node.js로 작성된 간단한 게시판 애플리케이션, `nod-bullentin-board`프로젝트를 다운로드 하여 실습

* 리눅스 버전

  * ```
    # 샘플 프로젝트 다운로드
    curl -LO https://github.com/dockersamples/node-bulletin-board/archive/master.zip
    # 압축 해제
    unzip master.zip
    # 해당 폴더에서 이미지를 빌드한 뒤, docker run을 실행할 것
    cd node-bulletin-board-master/bulletin-board-app
    ```

  * `docker build --tag bulletinboard:1.0 .` 명령어 이용

    * 이미지 빌드

    ![스크린샷, 2020-08-05 11-12-19](https://user-images.githubusercontent.com/58680504/89364265-15e38380-d70d-11ea-9b99-2566f011a136.png)

  * 이미지가 단계별로 빌드되는 모습을 볼 수 있다.

  * ㅇ

  * `docker run --publish 8000:8080 --detach --name bb bulletinboard:1.0`로 이미지 컨테이너로 실행

    * `--publish`
      * 호스트의 포트 8000에서 들어오는 트래픽을 컨테이너의 포트 8080으로 전달하도록 요청
      * 컨테이너에는 고유한 개인 포트 세트가 있으므로 네트워크의 포트에 도달하려면 이러한 방식으로 트래픽을 전달해야함 
        * 그렇지 않으면 방화벽 규칙으로 인해 모든 네트워크 트래픽이 기본 보안 상태로 컨테이너에 도달하지 못한다.
    * `--detach` 
      * 이 컨테이너를 백그라운드에서 실행하도록 요청
    * `--name`
      * 컨테이너를 참조할 수있는 이름을 지정
        * 예 :  `--name bb`

    <img src="https://user-images.githubusercontent.com/58680504/89364268-15e38380-d70d-11ea-8705-80b90ec90d53.png" alt="스크린샷, 2020-08-05 11-12-33" style="zoom:200%;" />

    * `docker ps`로 컨테이너 생성 및 실행 여부 확인

    ![스크린샷, 2020-08-05 11-12-55](https://user-images.githubusercontent.com/58680504/89364271-167c1a00-d70d-11ea-8e63-00e7a5b9efd2.png)

    

  * 이후 `localhost:8000`으로 접속하면, 컨테이너에 띄운 애플리케이션에 접속할 수 있다!

  ![스크린샷, 2020-08-05 12-39-56](https://user-images.githubusercontent.com/58680504/89369343-cb680400-d718-11ea-864d-8ba885ba20cc.png)

  

  



### 6. 배포한 컨테이너 중지 및 삭제

* 컨테이너를 중지하고 삭제
  * `docker stop <컨테이너 이름 혹은 id>`
  * `docker rm <컨테이너 이름 혹은 id>`
* 컨테이너를 중지하지 않고 바로 삭제
  * `docker rm --force <컨테이너 이름 혹은 id>`
    * `--force` 옵션은 실행중인 컨테이너를 중지하는 옵션





## 4. Docker Hub 설정 및 이용하기

#### 1. Docker hub 홈페이지에서 repository 생성

![스크린샷, 2020-08-06 14-51-07](https://user-images.githubusercontent.com/58680504/89495985-684b9f80-d7f4-11ea-9c4f-d79a16d5905d.png)



#### 2. CLI에서 `docker login`

![스크린샷, 2020-08-06 14-36-52](https://user-images.githubusercontent.com/58680504/89495975-641f8200-d7f4-11ea-849b-461634a6f107.png)



#### 3.  테스트를 위해 docker_test 디렉토리 생성, 이미지 빌드를 위해 Dockerfile 생성

![스크린샷, 2020-08-06 14-37-08](https://user-images.githubusercontent.com/58680504/89495977-6550af00-d7f4-11ea-9538-3b13180ec18e.png)



#### 4. Dockerfile 작성 내용

![스크린샷, 2020-08-06 14-37-19](https://user-images.githubusercontent.com/58680504/89495978-65e94580-d7f4-11ea-92cd-a22e496a2953.png)



#### 5.  작성된 Dockerfile로 이미지 빌드

![스크린샷, 2020-08-06 14-37-32](https://user-images.githubusercontent.com/58680504/89495979-65e94580-d7f4-11ea-8121-2e4773babc18.png)



#### 6.  빌드된 이미지 `run`

![스크린샷, 2020-08-06 14-37-43](https://user-images.githubusercontent.com/58680504/89495980-6681dc00-d7f4-11ea-863d-1f28e9a5df52.png)



#### 7. 이미지를 Docker hub에 push

![스크린샷, 2020-08-06 14-37-53](https://user-images.githubusercontent.com/58680504/89495981-671a7280-d7f4-11ea-94e2-31f2056a5a3a.png)



#### 8. Docker hub에 업로드된 모습

![스크린샷, 2020-08-06 14-42-54](https://user-images.githubusercontent.com/58680504/89495983-67b30900-d7f4-11ea-8191-896bc7dea14a.png)

* 따로 tag를 지정해주지 않았기 때문에, `latest` 버전으로 이미지가 push 되었다. 







## 5. Production 환경에서 Docker를 활용한 애플리케이션 실행



### 1.  오케스트레이션

* 오케스트레이터란( Orchestrator )?

  * 컨테이너화 된 응용프로그램을 관리, 확장 및 유지 관리하는 도구 

  > 예 : Kubernetes, Docker Swarm





