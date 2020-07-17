# Effective Go

### Go 코딩 스타일 가이드



### 1. 포맷팅(Formatting)

* `gofmt` (또는 `go fmt` : 패키지 레벨에서 실행되는 단위)

  * Go 프로그램을 읽은 뒤, 표준 스타일의 들여쓰기와 수직정렬, 유지 그리고 필요시 주석을 재포맷팅한 소스를 내놓는다.

  * formatting이 되어있지 않는 아래 코드를 작성한 이후에

    ```go
    /* formatting 안한 코드 */
    package main
    
    import "fmt"
    
    func main() {
    fmt.Println("Hello, playground")
    }
```
    
`gofmt`를 사용하면 format에 맞게 수정해줍니다.
    
* 실행 명령어 : $ `gofmt -w hello.go`
    
    ```go
    /* gofmt로 설정한 코드! */
    package main
    
    import "fmt"
    
    func main() {
    	fmt.Println("Hello, playground")
    }
    ```

* `-w` 옵션 : 정렬된 내용을 원본 소스 파일에 다시 저장



* 포맷팅에 대한 상세한 가이드

  * 들여쓰기

    > 들여쓰기를 위해 탭(tabs)을 사용.
    >
    >  `gofmt`도 기본값으로 탭을 사용.
    >
    > 꼭 써야하는 경우에만 스페이스(spaces)를 사용할 것.

  * 한 줄 길이

    > 한 줄 길이에 제한은 없다.
    >
    > 만약 라인 길이가 너무 길게 느껴진다면, 별도의 탭을 가지고 들여쓰기를하여 감싸기
    >
    > 긴 코드가 더 가독성이 좋다면 줄 길이를 짧게 하기 위해 줄을 끊지 말기

  * 괄호

    > Go는 C와 Java에 비해 적은 수의 괄호가 필요하다. 
    >
    > 제어 구조들(`if`, `for`, `switch`)의 문법엔 괄호가 없다. 
    >
    > 또한 연산자 우선순위 계층이 간단하며 명확하다. 
    >
    > 대신, 다른 언어와는 다르게 스페이스의 사용이 함축하는 바가 크다.

    ```go
    x<<8 + y<<16
    ```



### 2. 주석(Commentary)

* Go에서 제공하는 주석

  1.  C언어 스타일의 `/* */` 블럭주석

     > 주로 패키지(package)주석에 사용( 필수 ), 
     >
     > 많은 코드를 주석 처리하는 데에 사용

  2.  C++스타일의 `//` 한줄(line) 주석

     > 일반적으로 사용



* 주석 사용에 유의해야하는 이유?
  * `godoc`은 패키지의 내용에 대한 문서를 추출하도록 Go 소스 파일을 처리
  * 즉, 주석의 스타일과 유형은 godoc이 만들어내는 **문서의 질**을 결정



* **패키지 주석 **

  * 모든 패키지(package)는 패키지 구문 이전에 블럭주석형태의 패키지 주석이 있어야 한다.
  * 패키지 주석은 패키지를 소개해야하고, 전체 패키지에 관련된 정보를 제공해야한다. 
  * 패키지 주석은 godoc 문서의 처음에 나타나게 되니 이후의 자세한 사항도 작성해야 한다.

  ```go
  /* 패키지 주석의 예 */
  
  /*
  Package regexp implements a simple library for regular expressions.
  
  The syntax of the regular expressions accepted is:
  
      regexp:
          concatenation { '|' concatenation }
      concatenation:
          { closure }
      closure:
          term [ '*' | '+' | '?' ]
      term:
          '^'
          '$'
          '.'
          character
          '[' [ '^' ] character-ranges ']'
          '(' regexp ')'
  */
  package regexp
  ```



* 패키지 주석 작성 시 유의점

  * 생성된 출력이 고정폭 폰트로 주어지지 않을 수도 있으므로, 스페이스나 정렬등에 의존하지 말 것

  * 정확한 철자, 구두법, 문장구조를 사용하고 긴문장을 줄여야한다.

  * 주석의 시작은 `This function..` 보다는 설명하고자 하는 패키지나 함수의 이름으로 시작하는 것이 좋다.

    * grep문을 이용하는데 유용하기 때문!

    * ```go
      /* 이렇게 주석이 달린 패키지가 있을 때, */
      
      // Compile parses a regular expression and returns, if successful,
      // a Regexp that can be used to match against text.
      func Compile(str string) (*Regexp, error) {
      ```

    * `$ godoc regexp | grep parse`  명령어 실행 시,

    * ```
      $ godoc regexp | grep parse
          Compile parses a regular expression and returns, if successful, a Regexp
          parsed. It simplifies safe initialization of global variables holding
          cannot be parsed. It simplifies safe initialization of global variables
      ```

    * 원하는 주석 설명을 쉽게 찾을 수 있다.



### 3. 명칭

* 명칭 첫글자의 대소문자 여부에 따라 패키지 밖에서의 노출여부가 결정되기 때문에 명칭에 신경써야 한다.



#### 3-1. 패키지명

* 편의성 지향
  * 패키지명은 소문자, 한 단어로만 부여
  * 언더바(_)나 대소문자 혼용에 대한 필요가 없어야한다.
  * 오직 Import를 위한 이름!
    * util, common, misc, api, types, interfaces 같은 의미가 모호한 이름은 사용 X
    * `패키지 이름.xxx` 로 사용하기 때문에 `chubby.ChubbyFile` 보다는 `chubby.File` 방식으로 명명

* 패키지명은 소스 디렉토리 이름 기반으로.

  > `src/encoding/base64`의 `base64` 패키지를 import 한다면, 
  >
  > `import "base64"` 와 같은 방식으로 패키지 이름을 명명, 사용한다.

* 패키지에서 외부로 노출되는 타입들은 항상 패키지와 함께 불리고, 사용되기 때문에 다른 패키지와의 충돌은 걱정하지 않아도 된다.

* 다른 패키지에 있는 구조체를 복사할 때 주의

  * 복사본 구조체의 변수가 현재 패키지 내의 변수에 영향을 끼칠 수 있음

  

* **Import**

  ```go
  /* 패키지 그룹핑 방법 */
  
  package main
  
  import (
  	"fmt"
  	"hash/adler32"
  	"os"
  
  	"appengine/foo"
  	"appengine/user"
  
          "github.com/foo/bar"
  	"rsc.io/goversion/version"
  )
  ```

  * Import Dot

    ```go
    import (
    	"bar/testutil" // also imports "foo"
    	. "foo" 
    )
    ```

    * `import .` 을 하게 되면 import임에도 불구하고, 
    * 실행을 하려는 패키지가 `.`뒤의 패키지의 일부분인 것처럼 실행된다.
    * 주로 패키지 테스팅을 시도할 때 쓰인다.
    * 이외의 경우에는 가독성을 떨어뜨리므로 쓰기를 지양해야한다.

    

  * ##### goimports

    * import 경로를 추론하여 소스 상단에 추가를 해주는 도구

    * 아래 예제에서는 위와 달리 import “fmt”가 빠져있습니다.

      ```go
      package main
      
      func main() {
      	fmt.Println("Hello, playground")
      }
      ```

      * 그대로 실행하면 에러가 발생!

      ```go
      $ go run hello.go 
      # command-line-arguments
      ./hello.go:4: undefined: fmt in fmt.Println
      ```

      * 아래와 같이 goimports를 사용하면 fmt package가 자동 추가

      ```go
      $ goimports -w hello.go
      ```



#### 3-2. Getter & Setter

* Go는 getters와 setters를 자체적으로 제공 X

* 패키지밖으로 노출하기 위해 *대문자* 이름을 사용하는 것은 메서드로부터 필드를 식별할 수 있는 **훅(hook)**을 제공

  ```go
  /* Setter 예시 */
  
  owner := obj.Owner()
  if owner != user {
      obj.SetOwner(user)
  }
  ```

  

#### 3-3. 인터페이스명

* Method 이름을 기준으로, Method 이름에  `-er` 접미사를 붙이거나 에이전트 명사를 구성하는 유사한 변형에 의해 지정

  > 예 : Reader, Writer, Formatter, CloseNotifier..

* 사용될 "예정"인 인터페이스는 미리 정의하지 말것..!
  * 인터페이스를 미리 정의해두면 나중에 그 인터페이스의 사용여부 구별에 어려움을 겪을 수 있음



#### 3-4. Method 명

* Read, Write, Close, Flush, String 등등은 각각의 용법과 의미를 가지고 있음.

* 위와 같이 통용적인 용법과 의미를 뜻하는 Method의 이름은 같은 용법과 의미를 가질 Method에만 이름을 같게 쓰는 것이 좋다.

  > 예 : 문자 변환 메서드를 만든다면 `ToString`이 아닌 `String`을 사용.



#### 3-5. 대소문자 혼합

* Go에서의 Naming 규칙:

  *  여러 단어로된 이름을 명명할 때 언더바(_) 대신 대소문자 혼합을 사용

    > 예 : 
    >
    > `Mixed_Caps`( △ )
    >
    > `MixedCaps`나 `mixedCaps` ( ○ )

* 두문자어는 대문자로

  > 예 :  
  >
  > Url 보다는 URL, 
  >
  > ServeHttp 보다는 ServeHTTP 

#### 3-6. 변수 이름

* 변수 이름은 긴 것보다는 짧은 것을 선호
* `lineCount` 보다는 `c`, `sliceIndex` 보다는 `i `



#### 3-7. Receiver 이름

* 리시버 이름은 하나 혹은 두 문자로 축약
  * 메소드 인자이므로 서술성 X, 문서화 목적 X
* 모든 메소드에 걸쳐 사용될 수 있으니 짧을수록 좋다!



### 4. 세미콜론(Semicolons)

* Go의 정식문법은 구문을 종료하기 위하여 세미콜론을 사용. 
* 또는 for loop 구문에서 변수 초기화와 조건, 그리고 진행 변수를 구분할때에만 사용 
* *하지만* C언어와는 달리 세미콜론은 소스상에 나타나지 않는다..!
  * 구문분석기(lexer)
    *  간단한 규칙을 써서 스캔을 하는 과정에 자동으로 세미콜론을 삽입



* 세미콜론 삽입 규칙

  > 새로운 라인 앞의 마지막 토큰이 (int나 float64와 같은 단어를 포함한) 식별자이거나, 
  >
  > 숫자, 문자열과 같은 기본 리터럴, 혹은 다음의 토큰들중 하나 일 경우에, 
  >
  > 구문 분석기(lexer)는 토큰 다음에 세미콜론을 추가.

* 세미콜론 자동 삽입을 위해 지양해야하는 코드 작성법

  ```go
  /* 잘못된 코드 작성법 */
  
  if i < f()  // wrong! 중괄호 맨 끝이 아니라, f() 다음에 세미콜론이 삽입되게 된다.
  {           // wrong!
      g()
  }
  
  /* 옳은 코드 작성법 */
  
  if i < f() {
      g()
  }
  ```

  



### 5. 제어구조

* `do` 나 `while` 반복문은 없다.

*  `for`,  `if` ,`switch`가 존재

* 제어구조문에서는 초기화문을 이용할 수 있다.

  ```go
  if err := file.Chmod(0664); err != nil {
      log.Print(err)
      return err
  }
  ```

  > 초기화문 : `err := file.Chmod(0664);`
  >
  > if문 내의 지역변수 `err`

* 괄호는 필요하지 않다.
* body는 중괄호 사용 필수**( 의무 )**



* **재선언과 재할당**
* 변수의 단축선언, `v :=` 에서 변수 v 는 이미 선언되었더라도 다음의 경우 재선언이 가능하다.
  - 이 선언이 기존의 선언과 같은 스코프에 있어야 하고 (만약 v가 이미 외부 스코프에 선언되었다면, 이 선언은 새 변수를 만들것 이다. §),
  - 초기화 표현내에서 상응하는 값은 v에 할당할 수 있고,
  - 적어도 하나 이상의 새로운 변수가 선언문 안에 함께 있어야 한다.



*  for문 안에서 여러개의 변수를 사용하려면 *병렬 할당(parallel assignment)*을 사용

  ```go
  /* 병렬할당의 예시 */
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
      a[i], a[j] = a[j], a[i]
  }
  ```



* Go 언어에서 더 지향하는 조건문 : 
  * `if-else-if-else` 형태 보다는 **`switch`**
  * Go에서도 break 구문으로 switch를 일찍 종료할 수 있지만, 보편적이지 않다.



### 6. 메모리 할당

* Go에서 메모리를 할당하는 두 가지 기본 방식

  1. new
  2. make

  

* Go에서는 배열보다는 슬라이스 사용을 권장한다.



*  이차원의 slice를 메모리에 할당할 경우:

  1. 각 slice를 독립적으로 할당

  2. slice 하나를 할당하고 각 slice로 자른 다음 포인터를 주는 방식

     * **if** slice가 자라거나 줄어들 수 있다면, 

       * 독립적으로 할당해서 다음 줄을 덮어쓰는 일을 방지

     * **else**, 

       * 객체를 생성하기 위한 메모리 할당을 한번에

       

* Map을 사용할 때 Key로 정의할 수 있는 자료형 : 

  *  equality연산이 정의되어 있는 타입

  * ```
    integers, 
    floating point, 
    복소수(complex numbers), 
    strings, 포인터(pointers), 
    인터페이스(equality를 지원하는 동적 타입에 한해서), 
    structs,
    배열(arrays)
    ```

  * Slice는 map의 key로 사용될 수 없음.

    *  equality가 정의되어 있지 않기 때문 



#### 6-1. new

* 내장 함수
* 메모리를 초기화하지 않고, 단지 값을 제로화
* 제로값으로 할당된 타입 T를 가리키는 포인터를 반환



#### 6-2. 생성자와 합성 리터럴

* 제로값 이외에 생성자로 초기화하는 것

  ```go
  /* File 구조체를 초기화하는 것 */'
  func NewFile(fd int, name string) *File {
      if fd < 0 {
          return nil
      }
      f := File{fd, name, nil, 0}
      return &f /* C와는 달리, 로컬 변수의 주소를 반환해도 아무 문제가 없음 */
  } 
  ```

  * 변수에 연결된 저장공간은 함수가 반환해도 살아 남는다.
  * 따라서, 마지막 'return &f' 는 이렇게 변경할 수도 있음
  * => 'return &File{fd, name, nil, 0}'



#### 6-3. make

* slices, maps, 그리고 channels에만 사용
  * 각각의 데이터이 초기화되기 전까지 갖는 값 : nil(NULL)
*  (`*T`가 아닌) 타입 T의 (제로값이 아닌) 초기화된 값을 반환



```go
var p *[]int = new([]int)       // slice 구조체를 할당한다; *p == nil; 거의 유용하지 않다
var v  []int = make([]int, 100) // slice v는 이제 100개의 int를 갖는 배열을 참조한다

// 불필요하게 복잡한 경우:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Go 언어다운 경우:
v := make([]int, 100)
```



#### 6-4. 빈 슬라이스 선언

```go
/* 권장 선언 방법 */
 var t []string   //nil slice 값
 
 /* 권장하지 않는 방법 */
 t := []string{}  // non-nil에 len과 cap이 0인 slice
```

* JSON 오브젝트를 한정,
  * 권장 방법은 null, 권장하지 않는 방법은 []array 를 가지므로,
  * JSON에 대해서는 `t := []string{}` 선언 방식이 효율적





### 7. Receiver Type 결정

* value receiver와 pointer receiver의 사용에 대한 가이드

  |                                                              | 포인터 사용 여부 |
  | ------------------------------------------------------------ | ---------------- |
  | Receiver가 Map, func, channel일때                            | X                |
  | Receiver가 Slice, reslice나 realloc을 하지 않는 method       | X                |
  | Receiver가 작은 배열이나 구조체일 때                         | X                |
  | Method로 Receiver의 변경을 원할 때                           | O                |
  | Receiver가 구조체이고, sync.Mutex와 같은 동기화 필드를 가지고 있을 때 | O                |
  | Receiver를 사용하는 다른 곳에도 영향을 주고 싶을 때          | O                |
  | Receiver가 구조체, 배열, Slice이거나 포인터를 요소로 갖고 있을 때 | O                |
  | 무엇을 이용할지 불분명할 때                                  | O                |

  * 단순히 몇 바이트를 아끼기 위해서 포인터를 함수 인자로 넘기지 말것.
  * 함수 전반적으로 *x 와 같은 방식을 사용할 거면 포인터가 필요 X
  * 큰 구조체나 앞으로 커질 구초제에는 해당사항 X



### 8. Error 핸들링



#### 8-1. Don't Panic!

* 일반적인 에러를 다룰 때 panic을 사용 X
  * 대신  **built-in interface error** 를 활용



#### 8-2. Error

- error string은 소문자로 시작

- ```go
  /* 권장 에러 스트링 작성*/
  fmt.Errorf(“something bad”)
  
  /* 지양하는 에러 스트링 작성 */
  fmt.Errorf(“Something bad”)
  ```



* error 값을 받을 때 _ 로 받아 무시하지 말고 꼭 변수로 받아서 에러 처리를 확실하게 



#### 8-3. In-Band Errors

- Go는 여러개의 반환 값을 지원
  - client가 반환 값을 검사해서 오류 검사를 하도록 추가적인 값을 반환하도록 하는게 좋다.

```go
// Lookup returns the value for key or ok=false if there is no mapping for key.
func Lookup(key string) (value string, ok bool)value, ok := Lookup(key)
if !ok  {
    return fmt.Errorf("no value for %q", key)
}
return Parse(value)
```



#### 8-4. Indent Error Flow

- indentation을 최소화하면 가독성이 좋기 때문에 아래 보다는

```go
if err != nil {
	// error handling
} else {
	// normal code
}
```

- 아래와 같은 식으로 짜는게 좋다.

```go
if err != nil {
	// error handling
	return // or continue, etc.
}
// normal code
```



- 초기화 코드에서도 아래 보다는

```go
if x, err := f(); err != nil {
	// error handling
	return
} else {
	// use x
}
```

- 아래처럼 짜는게 좋습니다.

```go
x, err := f()
if err != nil {
	// error handling
	return
}
// use x
```



* Go에서는 else문을 최소화



#### 8-5. 테스트 케이스 Error

- 테스트가 실패한다면 그 이유를 알 수 있도록 충분한 정보가 있는 에러 메세지를 출력

```go
if got != tt.want {
	t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want) // or Fatalf, if test can't test anything more past this point
}
```





### 9. 동시성



#### 9-1. 고루틴 수명주기 관리

- 고루틴을 생성한 뒤 혹은 종료할 때 정리 작업은 필수!
- 동시성 코드는 고루틴 수명을 명확히 해서 단순하게 설계하는 것이 좋다. 
- 복잡한 고루틴은 언제 그리고 왜 고루틴이 종료하는지 문서화를 할 필요가 있다.