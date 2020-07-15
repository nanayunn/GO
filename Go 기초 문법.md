# Go 문법

## 1. Go 코드의 형태

```go
/* 패키지 선언(package declaration) */
package main

/* 다른 패키지 import 선언 */
import "fmt" 

/* 여러개의 패키지 선언 */
import (
    "fmt" 
    "strings" 
)

/* main 함수 선언 */
func main() {
    fmt.Println("Hello World\n")
}
```

- package
  - 모든 Go 프로그램은 패키지로 구성
  - 프로그램은 main 패키지에서 시작
  - import할 패키지 이름은 디렉토리 경로의 맨 마지막 이름을 사용하는 것이 원칙

- Import

  - 여러개의 "package" 를 소괄호로 감싸서 import를 표현

  - 패키지를 Import 하면 패키지가 외부로 export한 것들(메서드나 변수, 상수 등)에 접근할 수 있다.

    - import한 패키지는 단축어를 붙일 수 있다.

      ```go
      import (
        "fmt" 
        t "time" 
      )
      ```



## 2. 함수의 형태

- 매개변수의 타입은 변수명 뒤에 명시

  ```go
  func 함수의 이름(매개 변수) (반환형 : 선택사항) {
      본문(body)
  }
  /* => func main() {}  */
  
  /* 함수 예시 */
  func add(x, y int) int {
      return x + y
  }
  ```



## 3. 자료형

#### 3-1. 숫자

- 정수

- - int, uint(unsigned int)
    - uint8, uint16, uint32, uint64, int8, int16, int32, int64
    - 뒤에 붙은 숫자는 사용하는 비트 수를 나타냄
    - uint8 = byte
    - unit32 = rune

* 부동 소수점
  * float, complex

    - float32, float64
    - 부동 소수점 타입
    - complex64, complex 128
    - 복소수 타입

    

#### 3-2. 문자열

- 텍스트를 표현하는 데 사용되는 길이가 정해진 문자의 나열
- 각 문자는 개별 바이트로 구성 ( 한 문자마다 한 바이트 )



#### 3-3. boolean

- 참(true)과 거짓(false)를 나타내는 데 사용되는 특별한 1비트 정수 타입
  - && : and
  - || : or
  - ! : not



## 4. 변수

- 변수란?
  - 구체적인 타입과 이름을 가진 저장 공간
  - var 을 붙여서 선언
  - 타입은 변수의 이름 뒤에

- 변수 선언 줄임문:
  - :=
    - var 과 명시적인 타입(e.g. int, bool) 을 생략
    - 함수 안에서만 사용 가능

- 변수 선언 방법

  ```go
    var 변수명 타입 = 변수의 값
  
    /* 작성 예시 */
    var x string = "Hello World" 
  
    var x string
    x = "Hello World" 
  
    /* 축약형 */
    x := "Hello World" 
  
    var x = "Hello World" 
  ```

- 여러개의 변수 선언 방법

  ```go
    var (
        a = 5
        b = 10
        c = 15
    )
  
    /* enum식의 const 변수 선언 */
    const (
            Jan = 1
            Feb = 2
            Mar = 3
            Apr = 4
            May = 5
            Jun = 6
        )
  ```

- 변수의 형변환

  ```go
    /* int형 변수를 float형변환 */
    float64(len(x))
  ```

- 변수의 생략

  - `_`(언더스코어)

  - Go 컴파일러는 사용하지 않는 변수를 생성하지 못하게 한다.

    ```go
      var total float64 = 0
    
      /* 반복문에서 index가 사용되지 않으므로 _ 언더스코어로 표기 */
      for _, value := range x {
          total += value
      }
    ```



## 5. 상수

- const 키워드와 함께 변수처럼 선언
  - 문자(character)
  - 문자열(string)
  - 부울(boolean)
  - 숫자

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



### 5-1. 숫자형 상수 (Numeric Constants)

- 정밀한 값(values) 을 표현
- 타입을 지정하지 않은 상수는 문맥(context)에 따라 타입을 가지게 됨

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



## 6. 반복문과 조건문

#### 6-1. 반복문

- For

  - for문에서 선언과 ++ 부분을 생략할 수 있다.

    ```go
        /* For문 기본형 */
        func main() {
            for i := 1; i <= 10; i++ {
                fmt.Println(i)
            }
        }
    
        /* for문에서 추가적으로 문장을 실행할 수 있음 */
        for index, value := range x {
            total += value
        }
    ```

- 무한 루프

```go
func main() {
    for {
    }
}
```

#### 6-2. 조건문

- If

  - if ~ else if ~ else

  - '()'괄호는 안써도 '{}' 중괄호는 사용 필요

    ```go
        /* if else 기본형 */
        if i % 2 == 0 {
            // 짝수
        } else {
            // 홀수
        }
    ```

- switch

  - `if ~else`와 같이 조건문으로 쓰인다.

  - case의 코드 실행을 마치면 알아서 break

    ```go
        /* switch 기본형 */
        switch i {
        case 0: fmt.Println("영")
        case 1: fmt.Println("일")
        case 2: fmt.Println("이")
        case 3: fmt.Println("삼")
        case 4: fmt.Println("사")
        case 5: fmt.Println("오")
        default: fmt.Println("알 수 없는 숫자")
        }
    ```



## 7. 배열, 슬라이스, 맵

#### 7-1. 배열

- 길이가 고정된, 번호가 매겨진 단일 타입 원소의 나열

```go
     /* 배열의 기본형 */
     var 배열이름 [배열크기] 배열의 자료형

     var x [5]int

     /* 배열의 생성 및 값 할당 예시 */
     x := [5]float64{ 98, 93, 77, 82, 83 }
```



#### 7-2. 슬라이스

- 슬라이스란?
  - 길이가 변경될 수 있는, 배열의 일부
  - 슬라이스는 배열의 값을 가리키며(point), 배열의 길이를 가지고 있다.
  - `[]T` 는 타입 `T` 를 가지는 요소의 슬라이스(slice)

```go
   /* 슬라이스 생성 예시 */
   var 슬라이스 이름 [] 슬라이스 자료형

   var x []float64 /* 슬라이스는 길이 0인 상태로 생성 */

   /* 슬라이스 길이 설정 */
   make 함수를 이용

   x := make([]배열 자료형, 길이(총 길이에서 기저 배열이 점유하는 공간), 총 길이(선택 사항)) /* 세번째 매개변수로 용량(capacity)를 제한 */

   x := make([]float64, 5)

   /* 슬라이스 추가 예시 */
func main() {
    p := []int{2, 3, 5, 7, 11, 13} /* 배열을 가리키는 p */
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",
            i, p[i])
    }
}

/* 슬라이싱의 또다른 예 */
func main() {
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)
    fmt.Println("p[1:4] ==", p[1:4])

    // missing low index implies 0
    fmt.Println("p[:3] ==", p[:3])

    // missing high index implies len(s)
    fmt.Println("p[4:] ==", p[4:])

    q := p[4:]
    fmt.Println(q[:1]) /* 슬라이싱의 슬라이싱도 된다! */

}
```

- 슬라이스와 관련된 함수
  - append

```go
     /* append로 새로운 슬라이스에 이전 슬라이스를 추가 */
     func main() {
         slice1 := []int{1,2,3}
         slice2 := append(slice1, 4, 5)
         fmt.Println(slice1, slice2)
     }
     /* 결과 값 */
     slice2 = [1,2,3,4,5]
```

- copy

  ```go
       /* copy로 source slice에서 dest slice로 dest slice 길이만큼 복사 */
       func main() {
           slice1 := []int{1,2,3}
           slice2 := make([]int, 2)
           copy(slice2, slice1)   /* slice2가 길이가 2 이므로 2개의 요소만 복사 */
           fmt.Println(slice1, slice2)
       }
  ```



#### 7-3. 맵

- 순서가 없는 키-값(key-value) 쌍의 집합
- 반드시 사용하기 전에 `make` 를 명시 ( `new` 가 아니라는 것에 주의 )
  - `make` 를 수행하지 않은 `nil` 에는 값을 할당할 수 없다

```go
     /* 맵의 기본형 */
     var x map[string]int /* => x는 int에 대한 string의 맵이다*/

     /* 맵의 생성과 key, value 값 접근에 대한 예시 */
     var x map[string]int
     x["key"] = 10
     fmt.Println(x["key"]) /* 10 */
```

- 맵에 대한 추가 코드

  ```go
       package main
  
  import "fmt" 
  
  type Vertex struct {
      Lat, Long float64
  }
  
  /* 위에서 선언한 구조체에 대해 key값을 지정할 것을 선언 */
  var m map[string]Vertex
  
  func main() {
      m = make(map[string]Vertex) /* key에 대한 value 값을 가지도록 make */
      m["Bell Labs"] = Vertex{  /* string key 값에 구조체를 할당 */
          40.68433, -74.39967,
      }
      fmt.Println(m["Bell Labs"]) /* 할당된 구조체 값이 println */
  }
  
  /* map literal */
  var m2 = map[string]Vertex{
      "Bell Labs": Vertex{      /* key값을 반드시 지정해주어야 한다. */
          40.68433, -74.39967,
      },
      "Google": Vertex{         /* 타입명이 명시된 map literal의 경우, value 값의 타입명( 예 : Vertex)는 제거하여도 됨 */
          37.42202, -122.08408,
      },
  }
  ```

- 맵의 활용

  ```go
  m := make(map[string]int) 일때, 
  
  /* 맵 m 의 요소를 삽입하거나 수정하기: */
  m[key] = value
  
  /*요소 값 가져오기: */
  value = m[key]
  
  /*요소 지우기: */
  delete(m, key)
  
  /*키의 존재 여부 확인하기: */
  value, ok = m[key] /* key가 존재하면 true, 없으면 false 반환 */
  ```

- 맵을 이용한 wordcount 예시

  ```go
  package main
  
  import (
      "code.google.com/p/go-tour/wc" 
      "fmt" 
      "strings" 
  )
  
  func WordCount(s string) map[string]int {
  
      m := make(map[string]int) /* map을 이용하기 위해 make */
  
      f := strings.Fields(s)  /* strings.Fields함수 : 
      문자열을 입력받으면 '단어'를 추출하여 리스트를 반환*/
  
      for _, word:= range f{ /* f 안에 있는 word를 기준으로 for문 */
  
          m[word] += 1    /* key : 단어(word), value : 단어가 쓰인 횟수 */
          fmt.Println(word)
      }
  
      return m /* map 객체를 반환. */
  }
  
  func main() {
      wc.Test(WordCount)
  }
  ```



## 8. 함수

- 다중 함수 작성

```go
  func f(x int) {
      fmt.Println(x)
  }
  func main() {
      x := 5
      f(x)
  }
```

- 다중값 반환

- Named result :

  - return 값을 지정해주지 않으면 현재 구한 값들에 대해 반환

    ```go
      /* 결과값과 함께 오류 값을 반환하거나(x, err := f()) 
      성공을 나타내는 불린값을 반환(x, ok := f())하는 데 자주 사용 */
    
      func f() (int, int) {
          return 5, 6
      }
    
      func main() {
          x, y := f()
      }
    ```

- 가변 함수

```go
  /* '...' : 0개 이상의 매개변수 */
  func add(args ...int) int {
      total := 0
      for _, v := range args {
           total += v
  }
      return total
  }

  func main() {
      fmt.Println(add(1,2,3))
  }
```

- 클로저

```go
  /* 함수 안의 함수, 함수 내의 지역함수에서 return */
  func main() {
      add := func(x, y int) int {
          return x + y
      }                    /* add와 x + y가  클로저 형성 */
      fmt.Println(add(1,1))
  }
  
```

- 재귀

```go
  /* 팩토리얼 계산에 대한 함수 예시 */
  func factorial(x uint) uint {
      if x == 0 {
          return
      }
      return x * factorial(x-1)
  }
```

- 지연, 패닉, 복구

  - defer

    ```go
        package main
    
        import "fmt" 
    
        func first() {
            fmt.Println("1st")
        }
        func second() {
            fmt.Println("2nd")
        }
        func main() {
            defer second() /* first()보다 second()가 앞에 있지만, defer이 붙었기 때문에 지연 실행되어 더 늦게 실행된다. */
            first()
        }
    
        /* 주로 어떤 식으로든 자원을 해제해야할 때 자주 사용됨 */
        /* 파일 오픈 예시 */
        f, _ := os.Open(filename)
        defer f.Close()
    ```

- panic & recover
  - panic?
  - 프로그래밍 오류(예 : 범위를 벗어난 배열 인덱스 접근, 맵 초기화 안하고 이용..) 등 손쉽게 복구할 수 없는 예외적인 상황을 나타낸다

```go
    import "fmt" 

    func main() {
        defer func() { /* panic이 실행되면 defer로 지연되어 있던 recover 함수 실행  */
            str := recover()
            fmt.Println(str)
        }()
        panic("PANIC")
    }
```



## 9. 포인터

- 값이 저장된 메모리상의 위치를 가리킨다.

- Go에는 포인터가 있지만 포인터 연산은 불가능

  - ' * ' 연산자 :

    - 역참조 ( dereference )
    - 포인터가 가리키는 값을 불러온다.

  - ' & ' 연산자 :

    - 연산 대상에 대한 포인터 반환

    - 원본 변수의 값에 접근할 수 있다.

      ```go
        func one(xPtr *int) {
           *xPtr = 1
        }
        func main() {
            xPtr := new(int)
            one(xPtr)
            fmt.Println(*xPtr) // x는 1
        }
      ```

- new 함수

  - 인자로 타입을 하나 받아 해당 타입의 값에 맞는 충분한 메모리를 할당한 후 그것에 대한 포인터를 반환

  - ' & '와는 다르다.

  - `new(T)` 는 모든 필드가 **0(zero value)** 이 할당된 `T` 타입의 포인터를 반환

  - `var t *T = new(T)` 또는 `t := new(T)`

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



## 10. 구조체와 인터페이스

### 10-1. 구조체

- `struct` 는 필드(데이터)들의 조합

- 구조체의 필드는 보통 has-a 관계를 나타낸다

  ```go
  /* 구조체 기본형 */
  type 구조체 이름 struct{
      구조체 변수1, 
      구조체 변수2..
  }
  
  /* '원'에 대한 구조체 예시 */
  type Circle struct {
      x float64
      y float64
      r float64
  }
  
  /*구조체에 속한 필드(데이터)는 dot(.) 으로 접근*/
  func main() {
      v := Vertex{1, 2}
      v.X = 4
      fmt.Println(v.X)
  }
  ```

- 구조체 초기화

  ```go
  /* 구조체 변수 선언 */
  var c Circle
  
  /* 구조체 메모리 할당 */
  c := new(Circle)  /* => 각 필드는 0 값으로 설정된 후, 포인터가 반환(*Circle) 
                          즉, c는 Circle 구조체 객체에 대한 포인터*/
  
  /* 구조체 값 대입 예시 */
  /* 1 */
  c := Circle{x: 0, y: 0, r: 5}
  
  /* 2 */
  c := Circle{0, 0, 5}
  
  /* 3 : 필드값에 접근하여 대입 */
  c.x = 10
  c.y = 5
  ```

- 함수에서 구조체 값 변경하기

  ```go
  /* GO에서는 무조건 인자가 복사되므로 구조체 값을 함수에서 변경을 원할 시 포인터 필요 */
  
  func circleArea(c *Circle) float64 {
      return math.Pi * c.rc.r
  }
  
  /* 해당 함수에 인자값은 C에서 포인터 넘겨줄 때처럼 '&'로 넘긴다. */
  circleArea(&c)
  
  /* ================================================================== */
  /* '&'연산자 말고 수신자(receiver) 메소드 이용하기 */
  func (c *Circle) area() float64 {
      return math.Pi * c.r*c.r
  }
  
  /* 함수 이용법 */
  c.area()
  ```

- 구조체의 포함관계

  ```go
  /* 기본 구조체 예시 */
  type Person struct {
      Name string
  }
  func (p *Person) Talk() {
      fmt.Println("안녕, 내 이름은 ", p.Name, "야")
  }
  
  /* 새로운 구조체 */
  type Android struct {
      Person
      Model string
  }
  
  /* Android로 Person 구조체의 Talk()을 이용 가능 */
  a := new(Android)
  a.Talk()
  ```



### 10-2. 구조체 literal

- 필드와 값을 나열해서 구조체를 새로 할당하는 방법

- {Name: value}

- `&` 를 사용하면 구조체 리터럴에 대한 포인터를 생성

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



### 10-3. 메소드

- Go에는 클래스가 없는 대신 **메소드** 가 있다!

- 구조체에 메소드를 붙여 사용

  ```go
  package main
  
  import (
      "fmt" 
      "math" 
  )
  
  type Vertex struct {
      X, Y float64
  }
  
  /* method receiver : v *Vertex */
  func (v *Vertex) Abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
      v := &Vertex{3, 4}
      fmt.Println(v.Abs())
  }
  
  /* 구조체 외에도 type에 메소드를 붙일 수 있다. */
  type MyFloat float64
  
  func (f MyFloat) Abs() float64 {
  
      fmt.Println(f)
  
      if f < 0 {
          return float64(-f)
      }
      return float64(f)
  }
  
  /* 값에 직접 접근하여 변경이 필요한 경우가 아니라면, 포인터가 아닌 구조체로만 접근도 가능(값의 복사) */
  ```



### 10-4. 인터페이스

- 메소드의 집합

- struct가 interface형으로 구현될려면 interface에 정의된 메소드가 struct에 존재해야한다.

- 인터페이스의 구현?

  - 인터페이스를 정의한 패키지로 부터 구현 패키지를 분리(decouple)

  - 의존성이 없다

  - **덕타이핑** : 

    - 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
    
    ```go
    /* 인터페이스 기본형 */
    type 인터페이스명 interface {
        메서드1() 리턴값_자료형
        메서드2()                  // 리턴값이 없을 때
    }
    
    /* 인터페이스 예시 */
    type Shape interface {
        area() float64
    }
    ```

- 인터페이스 구현 예제

  ```go
  /* Rect 정의 */
  type Rect struct {
      width, height float64
  }
  
  /* Circle 정의 */
  type Circle struct {
      radius float64
  }
  
  /* Rect 타입에 대한 Shape 인터페이스 구현  */
  func (r Rect) area() float64 { return r.width * r.height }
  func (r Rect) perimeter() float64 {
       return 2 * (r.width + r.height)
  }
  
  /* Circle 타입에 대한 Shape 인터페이스 구현  */
  func (c Circle) area() float64 { 
      return math.Pi * c.radius * c.radius
  }
  func (c Circle) perimeter() float64 { 
      return 2 * math.Pi * c.radius
  }
  
  /* 인터페이스에 입력된 shape을 기준으로 for문을 통해 메소드 호출 */
  func showArea(shapes ...Shape) {
      for _, s := range shapes {
          a := s.area() 
          println(a)
      }
  }
  
  func main() {
      r := Rect{10., 20.}
      c := Circle{10}
  
      showArea(r, c)
  }
  ```



## 11. 동시성

- 동시성이란?
  - 하나 이상의 작업을 동시에 진행하는 것.
  - Go : 고루틴, 채널



### 11-1. Goroutine(고루틴)

- 다른 함수를 동시에 실행할 수 있는 함수

- 고루틴 생성 방법 :

  - `go 함수`

  - 프로그램 내 첫번째 고루틴 : main 함수

  - 고루틴이 실행되면, 즉시 다음 줄로 실행 흐름이 반환된다.

  - 함수 호출이 완료되기까지 기다리지 않는다.

  - 따라서 main 함수 안에서 고루틴을 생성하게 되면 두번째로 생성된 고루틴이 실행 및 종료가 되기 전에 main 함수가 종료될 수 있다.

    ```go
    /* 고루틴 main 안의 새로운 고루틴 */
    package main
    
    import "fmt" 
    
    func f(n int) {
        for i := 0; i < 10; i++ {
            fmt.Println(n, ":", i)
        }
    }
    
    func main() {
        go f(0)
        var input string
        fmt.Scanln(&input) /* Scanln이 없다면 go f(i)가 모두 실행되기 전에 main문이 종료될 것 */
    }
    ```

    ```go
    /* 고루틴 예시2 */
    
    package main
    
    import (
        "fmt" 
        "time" 
        "math/rand" 
    )
    
    func f(n int) {
        for i := 0; i < 10; i++ {
            fmt.Println(n, ":", i)
            amt := time.Duration(rand.Intn(250))
            time.Sleep(time.Millisecond * amt)
        }
    }
    
    func main() {
        for i := 0; i < 10; i++ {
            go f(i)
        }
        var input string
        fmt.Scanln(&input)
    }
    ```

- 위의 코드의 실행 결과

  ```go
  /* 고루틴 예시 2 실행 결과 */
  1 : 0
  0 : 0
  5 : 0
  8 : 0
  3 : 0
  4 : 0
  6 : 0
  7 : 0
  2 : 0
  9 : 0
  2 : 1
  9 : 1
  8 : 1
  9 : 2
  4 : 1
  3 : 1
  1 : 1
  5 : 1
  3 : 2
  0 : 1
  9 : 3
  6 : 1
  7 : 1
  8 : 2
  8 : 3
  2 : 2
  7 : 2
  0 : 2
  7 : 3
  8 : 4
  8 : 5
  1 : 2
  5 : 2
  4 : 2
  7 : 4
  8 : 6
  3 : 3
  4 : 3
  0 : 3
  6 : 2
  9 : 4
  6 : 3
  5 : 3
  2 : 3
  8 : 7
  1 : 3
  7 : 5
  2 : 4
  6 : 4
  4 : 4
  9 : 5
  2 : 5
  7 : 6
  3 : 4
  1 : 4
  1 : 5
  4 : 5
  0 : 4
  8 : 8
  5 : 4
  0 : 5
  6 : 5
  6 : 6
  5 : 5
  4 : 6
  7 : 7
  4 : 7
  9 : 6
  3 : 5
  6 : 7
  2 : 6
  1 : 6
  3 : 6
  3 : 7
  8 : 9
  4 : 8
  7 : 8
  2 : 7
  0 : 6
  5 : 6
  7 : 9
  4 : 9
  1 : 7
  9 : 7
  6 : 8
  3 : 8
  5 : 7
  1 : 8
  2 : 8
  5 : 8
  5 : 9
  3 : 9
  9 : 8
  0 : 7
  9 : 9
  2 : 9
  0 : 8
  1 : 9
  0 : 9
  6 : 9
  ```



### 11-2. Channel(채널)

- 두 고루틴이 서로 통신하고 실행흐름을 동기화하는 수단을 제공

- 데이터를 전달하는 통로 ( 변수의 역할X )

  - 채널 연산자 `<-` 를 이용

    - (데이터가 화살표 방향에 따라 흐릅니다.)

  - 타입이 존재하는 파이프

  - 맵이나 슬라이스처럼, 채널은 사용되기 전에 생성되어야 함

  - 송/수신은 상대편이 준비될 때까지 블록

    ```go
    /* 채널 생성의 기본형 */
    ch := make(chan int)
    
    /* 채널 이용의 간단한 예시 */
    ch <- v    // v 를 ch로 보냅니다.
    v := <-ch  // ch로부터 값을 받아서
               // v 로 넘깁니다.
    
    /* channel 객체 생성 : 버퍼의 크기가 0인 채널 변수 */
    var c chan string = make(chan string)
    
    /* ======================================================= */
    /* 3개의 go 루틴이 chan을 이용하여 ping, pong을 출력하는 예제 코드 */
    package main
    
    import (
        "fmt" 
        "time" 
    )
    
    func pinger(c chan string) {
        for i := 0; ; i++ {
            c <- "ping" /* channel 객체 c에 "ping"을 전달 */
        }
    }
    
    func ponger(c chan string) {
        for i := 0; ; i++ {
            c <- "pong" /* channel 객체 c에 "pong"을 전달 */
        }
    }
    
    func printer(c chan string) {
        for {
            msg := <-c  /* msg 객체에 c에 저장된 string을 전달 */
            fmt.Println(msg)
            time.Sleep(time.Second * 1) /* msg 출력 후 sleep */
        }
    }
    func main() {
        var c chan string = make(chan string)
    
        go pinger(c)
        go ponger(c)
        go printer(c)
    
        var input string
        fmt.Scanln(&input) /* main 고루틴이 끝나는 것을 방지하기 위해 Scanln() */
    }
    ```

- 채널의 방향성
  
  - 단방향, 양방향 채널

```go
  chan<- string // 송신 전용 채널
  <-chan string // 수신 전용 채널
```

- 채널 버퍼링

  - `make` 에 두번째 인자로 버퍼 용량을 넣음으로써 해당 용량만큼 버퍼링되는 채널을 생성

    ```go
    ch := make(chan int, 100)
    ```

  - 버퍼드 채널은 비동기 방식으로 동작

  - 버퍼링되는 채널로의 송신은 버퍼가 꽉 찰 때까지 블록

  - 지정된 버퍼를 넘어가는 데이터를 채널에 할당 시 나오는 오류 메세지

    ```go
    fatal error: all goroutines are asleep - deadlock!
    
    goroutine 1 [chan send]:
    main.main()
        /tmp/sandbox994386935/prog.go:9 +0x9b
    ```

- 채널 close

  - 데이터 송신측은 더이상 보낼 값이 없다는 것을 알리기 위해 채널을 `close`

  - 수신측은 'ok'로 채널 오픈 여부 확인

    ```go
        /* 값이 false가 나오면 닫힌 것 */
        v, ok := <-ch
    ```

  - 오로지 송신측만 채널을 닫을 수 있다. 이미 닫힌 채널에 데이터를 보내면 패닉이 일어난다.

  - 채널은 항상 닫을 필요는 없다. `range` 루프를 종료시켜야 할 때와 같은 때에만 닫는 것을 명시



### 11-3. Select

- Goroutines이 다수의 통신 동작으로부터 수행 준비를 기다릴 수 있게 함
- 다수의 채널이 동시에 준비되면 그 중 하나를 무작위로 선택

```go
/* 마치 switch문처럼.. */
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}
```

- `select` 의 `default` 케이스는 현재 수행 준비가 완료된 케이스가 없을 때 수행

------



## etc. 그 외



### 주석

- 아래의 기호로 주석을 사용할 수 있다.
- //
- /* ~~~ */



### Print 문

- `Println`, `Print`, `Printf`

```go
/* Printf 예시 */
func main() {
  animal1 := "cat" 
  animal2 := "dog" 

  // Add your code below:
  fmt.Printf("Are you a %v or a %v person?", animal1, animal2)
}
```

- print문에서 쓸 수 있는 서식문자
  - `%v` :
    - var 변수
  - `%T` :
    - 변수의 Type
  - `%d`:
    - int 타입
  - `%f` :
    - float 타입
    - `%.2f`: 소수점 자릿수를 정할 수 있음

- `Sprintln`, `Sprintf`

```go
    package main

    import "fmt" 

    func main() {
      step1 := "Breathe in..." 
      step2 := "Breathe out..." 

      // Add your code below:
      meditation := fmt.Sprintln(step1, step2)

      fmt.Print(meditation)
    }
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



### Scan

- 사용자에게 입력을 받는 함수

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

- & 연산자로 변수에 입력받은 값을 넣는다.



### ERROR문

* Go에서는 error와 관련하여 기본 interface와 메소드를 제공한다.

  ```go
  type error interface {
      Error() string
  }
  ```

* 에러 코드는 다음과 같이 코드에서 이용할 수 있다.

```go
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}

/* 실행 결과 : => open /etc/passwx: no such file or directory */
```

