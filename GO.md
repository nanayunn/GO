## GO

* 2012년 3월 구글에서 발표된 프로그래밍 언어
* 로버트 그리즈머, 롭 파이크, 케네스 톰슨
* 2015년 Go lang 컴파일러 개발
* 단순함과 실용성을 지향



* 장점
  * 컴파일 언어지만 컴파일 속도가 매우 빠름
  * 간결한 코드 작성법
  * 풍부한 라이브러리
  * 인터프리터 언어의 즉흥성만큼 빠르진 않지만 컴파일 언어인 만큼 퍼포먼스가 확실



* GoRoutine ( 멀티 스레드 메커니즘 )
  * 비동기 메커니즘
  * 각각의 GoRoutine은 병렬로 동작하며 메시지 채널을 통해 값을 주고 받는다.
  * 동기화 방법 :
    * GoRoutine으로부터 반환값을 받는 코드를 메인 스레드에 추가



* Go는 바이너리 컴파일러

  * 다른 플랫폼에 배포시 환경변수 재설정 필요

  

***

###  GO lang 설치(Ubuntu 18.04)

* https://golang.org/doc/install



* `sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
  * `sudo` 명령어 필요했음
  * 해당 명령어 실행 시 `/usr/local/go` 에 tar 파일이 압축 해제, 자동 생성
* GO PATH 설정 
  * `sudo vi /etc/profile`
    * `sudo` 로 오픈 필요
  * `export PATH=$PATH:/usr/local/go/bin`
    * 맨 마지막 줄에 GO PATH를 추가로 입력
  * `source $HOME/.profile`
    * profile에 추가된 설정을 적용시키기 위한 command이지만 따로 적용되지 않음
    * 컴퓨터 다시시작 후 go build 가능



***

### GO VScode 환경설정

* VScode에서 go 플러그인 설치
* 추가적으로 설치하라고 알림창이 뜨면 모두 설치해준다.



* gofmt와 goimports 도구 연동

  * `ctrl + ,`

    * 설정 탭으로 이동 후 GO 메뉴에서 setting.json 편집

    * ```json
      {
          "go.buildOnSave": true,
          "go.lintOnSave": true,
          "go.vetOnSave": true,
          "go.buildTags": "",
          "go.buildFlags": [],
          "go.lintFlags": [],
          "go.vetFlags": [],
          "go.coverOnSave": false,
          "go.useCodeSnippetsOnFunctionSuggest": false,
          "go.formatOnSave": true,
          // goreturns 은 goimports(자동 임포트), gofmt(자동 포맷팅)를 사용하고 리턴코드도 자동으로 채워준다.
          "go.formatTool": "goreturns",
          "go.gocodeAutoBuild": false,
          // 맥,리눅스 기준
          //"go.goroot": "/usr/local/go",
          //"go.gopath": "/home/nam/projects/go"
          // 윈도우 기준
          "go.goroot": "c:\\go",
          "go.gopath": "d:\\projects\\Go"
      }
      ```

    * `go env` 명령어로 root와 path 확인 가능



***

### GO 문법



#### Package

* 모든 Go 프로그램은 패키지로 구성
* 프로그램은 main 패키지에서 시작
* import할 패키지 이름은 디렉토리 경로의 맨 마지막 이름을 사용하는 것이 원칙

```go
package main
```



#### Import

* 여러개의 `"package"` 를 소괄호로 감싸서 import를 표현
* 패키지를 `Import` 하면 패키지가 외부로 `export `한 것들(메서드나 변수, 상수 등)에 접근할 수 있다.

```go
import (
    "fmt"
    "math"
)
```

* import한 패키지는 단축어를 붙일 수 있다.

```go
import (
  "fmt"
  t "time"
)
```



#### Export

* exported name : 

  * 첫 문자는 대문자로.

    > 예를 들어 `Foo` 와 `FOO` 는 외부에서 참조할 수 있지만 
    >
    > `foo` 는 참조 할 수 없다.

  * 패키지를 사용하는 곳에서 접근할 수 있게 됨.



#### 함수

*  매개변수(인자)를 가진다.
* 매개변수의 타입은 변수명 *뒤에* 명시

```go
func add(x int, y int) int {
    return x + y
}
```

* 두 개 이상의 매개변수가 같은 타입(type)일 때, 같은 타입을 취하는 마지막 매개변수에만 타입을 명시하고 나머지는 생략할 수 있음

```go
func add(x, y int) int {
    return x + y
}
```

* 하나의 함수에서 **여러개**의 반환값을 가질 수 있음

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

```

* 함수의 이름과 선언된 매개변수는 함수 내의 변수로 쓰일 수 있음
* Named result :
  * return 값을 지정해주지 않으면 현재 구한 값들에 대해 반환

```go
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```



#### 변수

* `var` 을 붙여서 선언
* 타입은 변수의 이름 뒤에

```go
package main

import "fmt"

var x, y, z int
var c, python, java bool

func main() {
    fmt.Println(x, y, z, c, python, java)
}
```

* 변수는 선언과 동시에 초기화할 수 있다.
* 타입 생략 가능
  * 초기화 값에 따라 타입이 결정되도록 할 수 있다.

* 변수 선언 줄임문:
  * `:=` 
    * `var` 과 명시적인 타입(e.g. int, bool) 을 생략
    * 함수 안에서만 사용 가능

```go
package main

import "fmt"

func main() {
    var x, y, z int = 1, 2, 3
    c, python, java := true, false, "no!"

    fmt.Println(x, y, z, c, python, java)
}
```



#### 상수

*  `const` 키워드와 함께 변수처럼 선언
  * 문자(character)
  * 문자열(string)
  * 부울(boolean)
  * 숫자

```go
package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "안녕"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
```



#### 숫자형 상수 (Numeric Constants)

* *정밀한 값(values)* 을 표현
* 타입을 지정하지 않은 상수는 문맥(context)에 따라 타입을 가지게 됨

```go
package main

import "fmt"

const (
    Big   = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
```



#### For 반복문

* Go lang에서 유일한 반복문

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

* for문에서 선언과 ++ 부분을 생략할 수 있다.

```go
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
```

* 무한 루프

```go
func main() {
    for {
    }
}
```



#### If 반복문

* `()` 괄호는 안써도 `{}` 중괄호는 사용 필요

```go
package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}
```

* if 문 실행 전, 짧은 구문을 실행할 수 있다.

  ```go
  /* 짧은 실행문을 통해 선언된 변수는 if 안쪽 범위(scope) 에서 만 사용할 수 있다.*/
  func pow(x, n, lim float64) float64 {
      if v := math.Pow(x, n); v < lim {
          return v
      }
      return lim
  }
  ```

  

#### Switch 문

* `if ~else`와 같이 조건문으로 쓰인다.

```go
package main

import "fmt"

func main() {
	name := "H. J. Simp"

	// Add your switch statement below:
	switch name{
    case "Butch":
      fmt.Println("Head to Robbers Roost.")
    case "Bonnie":
      fmt.Println("Stay put in Joplin")
    default:
      fmt.Println("Just hide!")
  } 
}
```



#### 주석

* `//`
* `/* ~~~ */`



#### Print 문

* `Println`, `Print`, `Printf`

```go
/* Printf 예시 */
func main() {
  animal1 := "cat"
  animal2 := "dog"
  
  // Add your code below:
  fmt.Printf("Are you a %v or a %v person?", animal1, animal2)
}
```

* print문에서 쓸 수 있는 서식문자
  * `%v` : 
    * var 변수
  * `%T` : 
    * 변수의 Type
  * `%d`:
    * int 타입
  * `%f` :
    * float 타입
    * `%.2f` : 소수점 자릿수를 정할 수 있음

* `Sprintln`, `Sprintf`

  * ```go
    package main
    
    import "fmt"
    
    func main() {
      step1 := "Breathe in..."
      step2 := "Breathe out..."
      
      // Add your code below:
      meditation := fmt.Sprintln(step1, step2)
    
      fmt.Print(meditation)
    }
    ```

  * ```go
    package main
    
    import "fmt"
    
    func main() {
      template := "I wish I had a %v."
      pet := "puppy"
      
      var wish string
      
      // Add your code below:
      wish=fmt.Sprintf(template, pet)
      
      fmt.Println(wish)
    }
    ```



#### Scan

* 사용자에게 입력을 받는 함수

```go
package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")
  
  // Add your code below:
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.", food)
}
```

* `&` 연산자로 변수에 입력받은 값을 넣는다.

***

#### Go 실행

* 1. `go build XXX.go`
  2. `./XXX`
* 1. `go run main.go`



#### Go documentation 확인

* `go doc time.Now`

  * ```
    func Now() Time
        Now returns the current local time.
    ```

    - 이렇게 도움말이 나온다.



#### Go의 기본 자료형

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 다른 이름(alias)

rune // int32의 다른 이름(alias)
     // 유니코드 코드 포인트 값을 표현합니다. 

float32 float64

complex64 complex128
```



#### 구조체

* `struct` 는 필드(데이터)들의 조합입니다.

```go
/*type 선언으로 struct의 이름을 지정*/
type Vertex struct {
    X int
    Y int
}

/*구조체에 속한 필드(데이터)는 dot(.) 으로 접근*/
func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```



#### 구조체 Literal

* 필드와 값을 나열해서 구조체를 새로 할당하는 방법
* {Name: value}
* `&` 를 사용하면 구조체 리터럴에 대한 포인터를 생성

```go
type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 2}  // has type Vertex
    q = &Vertex{1, 2} // has type *Vertex
    r = Vertex{X: 1}  // Y:0 is implicit
    s = Vertex{}      // X:0 and Y:0
)
```



#### 포인터

* Go에는 포인터가 있지만 포인터 연산은 불가능
* 구조체 변수는 구조체 포인터를 이용해서 접근 ( 영향 미침 )

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    p := Vertex{1, 2}
    q := &p
    q.X = 1e9
    fmt.Println(p)
}
```



#### New 함수

* `new(T)` 는 모든 필드가 *0(zero value)* 이 할당된 `T` 타입의 포인터를 반환
* `var t *T = new(T)`  또는 `t := new(T)`

```go
type Vertex struct {
    X, Y int
}

func main() {
    v := new(Vertex) /* v는 초기화후 생성된 Vertex 구조체 포인터 */
    fmt.Println(v)
    v.X, v.Y = 11, 9
    fmt.Println(v)
}
```



#### Slice 함수

* 슬라이스는 배열의 값을 가리키며(point),  배열의 길이를 가지고 있다.

* `[]T` 는 타입 `T` 를 가지는 요소의 슬라이스(slice)

```go
func main() {
    p := []int{2, 3, 5, 7, 11, 13} /* 배열을 가리키는 p */
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",
            i, p[i])
    }
}
```

