# Docker 설치



* 이전에 redash 서비스를 설치하며, docker가 의존성 프로그램으로 함께 설치가 되었었다.

* redash 서비스를 삭제하였음에도 불구하고 어딘가 남아있는 docker 설정 파일로 인해 지속적으로 redash 관련 컨테이너가 띄워졌다.

* docker-compose.yml 을 이용한 테스트를 진행할 때도, default가 redash관련 docker-compose.yml 파일이라서 문제가 되었다.

* docker를 모두 삭제하고 다시 설치하기 위해 

  `sudo snap install docker` 명령어를 이용하였으나,

  docker에서 정식으로 서비스하는 것이 아닌, snap 스토어를 이용하여 우분투에서 서비스하는 것을 확인

  ![스크린샷, 2020-08-14 17-10-27](https://user-images.githubusercontent.com/58680504/90228423-2a342880-de51-11ea-82c0-63ac719c4f67.png)





* `docker version`으로 더 확인해본 결과, `runC` 모듈도 이용하고 있지 않던 docker임을 확인,
* 삭제 후 docker docs에서 안내하는 가이드를 따라 
  * https://docs.docker.com/engine/install/ubuntu/#install-docker-engine (설치 실패)
  * [https://zetawiki.com/wiki/%EC%9A%B0%EB%B6%84%ED%88%AC16_docker_%EC%84%A4%EC%B9%98](https://zetawiki.com/wiki/우분투16_docker_설치) (설치 실패) 
  * [https://blog.cosmosfarm.com/archives/248/%EC%9A%B0%EB%B6%84%ED%88%AC-18-04-%EB%8F%84%EC%BB%A4-docker-%EC%84%A4%EC%B9%98-%EB%B0%A9%EB%B2%95/](https://blog.cosmosfarm.com/archives/248/우분투-18-04-도커-docker-설치-방법/) (설치 실패)
* `sudo apt-get install docker-ce`
* 명령어를 실행하였지만, 

![스크린샷, 2020-08-14 17-13-25](https://user-images.githubusercontent.com/58680504/90228611-7e3f0d00-de51-11ea-9839-b35f90d77937.png)

* 다음과 같은 오류에 계속 부딪혔다.
* 찾아본 결과 docker 내에서 storage-driver를 제대로 인식하지 못하여서( /dev/null ) 발생하는 버그임을 확인했다.
  * 참고한 해결 방안
    * https://support.plesk.com/hc/en-us/articles/360012448554-Unable-to-start-Docker-service-Dependency-failed-for-Docker-Application-Container-Engine (해결 X)
    * https://forums.docker.com/t/failed-to-start-docker-application-container-engine/35594
    * https://stackoverflow.com/questions/49110092/failed-to-start-docker-application-container-engine
    * https://dev-note.tistory.com/6
    * https://github.com/docker/for-linux/issues/162
    * [https://www.it-swarm.dev/ko/linux/%EC%9A%B0%EB%B6%84%ED%88%AC-1604%EC%97%90%EC%84%9C-docker-%EC%84%9C%EB%B9%84%EC%8A%A4%EB%A5%BC-%EC%8B%9C%EC%9E%91%ED%95%A0-%EC%88%98-%EC%97%86%EC%8A%B5%EB%8B%88%EB%8B%A4/824426329/](https://www.it-swarm.dev/ko/linux/우분투-1604에서-docker-서비스를-시작할-수-없습니다/824426329/)
* 이전 버전이 문제가 될까 모두 삭제를 진행하였지만 같은 오류가 지속적으로 났다. 



* 버전을 지정하지 않고 설치 시, Docker 19.03.12 버전이 설치되는데, 
* 최신 버전이 문제가 될까 싶어서 모두 삭제 후 19.03. 9 버전을 기준으로 설치를 시도
* ![스크린샷, 2020-08-14 17-25-48](https://user-images.githubusercontent.com/58680504/90229661-3d47f800-de53-11ea-87e2-3464d66fc5fd.png)

* 이전에 다른 해결 방법을 따라서 만들어 놓았던 overlay.conf 파일이 문제를 일으키는 것으로 보여서, 
* `rm /etc/systemd/system/docker.service.d/overlay.conf`를 진행한뒤,
* 다시  `sudo apt-get install docker-ce="5:19.03.9~3-0~ubuntu-bionic" docker-ce-cli="5:19.03.9~3-0~ubuntu-bionic" containerd.io` 명령어를 진행
* ![스크린샷, 2020-08-14 17-27-58](https://user-images.githubusercontent.com/58680504/90229837-8ac46500-de53-11ea-99f4-7279e8ef2078.png)

* dpkg 오류가 발생하여서, 
* `sudo dpkg --configure -a`
* `sudo apt-get install -f`
* `sudo apt-get update`
* `sudo apt-get upgrade`
* 를 진행하였더니 무사히 설치되었다. 



