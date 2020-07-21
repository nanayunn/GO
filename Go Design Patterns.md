# Go Design Patterns



## 1. 생성 패턴( Creational Patterns )

* 객체 생성에 관련된 패턴
* 객체의 생성과 조합을 캡슐화
* 특정 객체가 생성되거나 변경되어도 프로그램 구조에 영향을 크게 받지 않도록 유연성을 제공
  



### 1-1. 추상 팩토리 패턴 ( Abstract factory pattern )

> 관련된 객체들의 집단(군, family)를 생성하는 인터페이스를 제공하되, 
>
> 생성되는 객체의 구체적인 클래스를 알 필요가 없도록 한다.

* 공통의 테마를 가진 팩토리의 그룹을 캡슐화하는 방법을 제공하는 소프트웨어 디자인 패턴



#### 예시 코드

* `신발(shoe)`과 `반바지(short)`로 구성된 스포츠 키트( concrete products )를 사고싶을 때, 

* 이러한 스포츠 키트를 제공하는 `nike`와 `adidas`( abstract factory )가 있다고 예시를 든다.

* ```go
  nike와 adidas가 implement하는 인터페이스 : iSportsFactory
  
  구체적인 타입을 가지는 구조체 nikeShoe, adidasShoe가
  implement하는 인터페이스 : iShoe
  
  구체적인 타입을 가지는 구조체 nikeShort, adidasShort가
  implement하는 인터페이스 : iShort
  ```

* **iSportsFactory.go**

  ```go
  package main
  
  import "fmt"
  
  /* 통합 인터페이스 iSportsFactory 
  adidas와 nike는 인터페이스 내의 메소드를 공통적으로 갖고 있다.*/
  type iSportsFactory interface {
      makeShoe() iShoe
      makeShort() iShort
  }
  
  func getSportsFactory(brand string) (iSportsFactory, error) {
      if brand == "adidas" {
          return &adidas{}, nil
      }
      if brand == "nike" {
          return &nike{}, nil
      }
      return nil, fmt.Errorf("Wrong brand type passed")
  }
  ```

  **adidas.go**

  ```go
  package main
  
  /* adidas 구조체 */
  type adidas struct {
  }
  
  /* 함수를 실행하면 shoe 구조체를 변수로 가지는 adidasShoe 구조체에 값을 
  입력한 후, 포인터를 반환*/
  func (a *adidas) makeShoe() iShoe {
      return &adidasShoe{
          shoe: shoe{
              logo: "adidas",
              size: 14,
          },
      }
  }
  
  /* 함수를 실행하면 short 구조체를 변수로 가지는 adidasShort 구조체에 값을 
  입력한 후, 포인터를 반환*/
  func (a *adidas) makeShort() iShort {
      return &adidasShort{
          short: short{
              logo: "adidas",
              size: 14,
          },
      }
  }
  ```

  **nike.go**

  ```go
  package main
  
  /* nike 구조체 */
  type nike struct {
  }
  
  /* 함수를 실행하면 shoe 구조체를 변수로 가지는 nikeShoe 구조체에 값을 
  입력한 후, 포인터를 반환*/
  func (n *nike) makeShoe() iShoe {
      return &nikeShoe{
          shoe: shoe{
              logo: "nike",
              size: 14,
          },
      }
  }
  
  /* 함수를 실행하면 short 구조체를 변수로 가지는 nikeShort 구조체에 값을 
  입력한 후, 포인터를 반환*/
  func (n *nike) makeShort() iShort {
      return &nikeShort{
          short: short{
              logo: "nike",
              size: 14,
          },
      }
  }
  ```

  **iShoe.go**

  ```go
  package main
  
  /* getter와 setter 함수를 포함하는 인터페이스 iShoe */
  type iShoe interface {
      setLogo(logo string)
      setSize(size int)
      getLogo() string
      getSize() int
  }
  
  /* shoe 구조체, logo와 size 변수를 갖는다. */
  type shoe struct {
      logo string
      size int
  }
  
  func (s *shoe) setLogo(logo string) {
      s.logo = logo
  }
  
  func (s *shoe) getLogo() string {
      return s.logo
  }
  
  func (s *shoe) setSize(size int) {
      s.size = size
  }
  
  func (s *shoe) getSize() int {
      return s.size
  }
  ```

  **adidasShoe.go**

  ```go
  package main
  
  type adidasShoe struct {
  	shoe
  }
  ```

  **nikeShoe.go**

  ```go
  package main
  
  type nikeShoe struct {
      shoe
  }
  ```

  **iShort.go**

  ```go
  package main
  
  /* getter와 setter 함수를 포함하는 인터페이스 iShort */
  type iShort interface {
      setLogo(logo string)
      setSize(size int)
      getLogo() string
      getSize() int
  }
  
  /* short 구조체, logo와 size 변수를 갖는다. */
  type short struct {
      logo string
      size int
  }
  
  func (s *short) setLogo(logo string) {
      s.logo = logo
  }
  
  func (s *short) getLogo() string {
      return s.logo
  }
  
  func (s *short) setSize(size int) {
      s.size = size
  }
  
  func (s *short) getSize() int {
      return s.size
  }
  ```

  **adidasShort.go**

  ```go
  package main
  
  type adidasShort struct {
      short
  }
  ```

  **nikeShort.go**

  ```go
  package main
  
  type nikeShort struct {
      short
  }
  ```

  **main.go**

  ```go
  package main
  
  import "fmt"
  
  func main() {
      adidasFactory, _ := getSportsFactory("adidas") //반환값 : adias 구조체
      nikeFactory, _ := getSportsFactory("nike") // 반환값 : nike 구조체
      nikeShoe := nikeFactory.makeShoe()  // 반환값 : nikeShoe 구조체
      nikeShort := nikeFactory.makeShort() // 반환값 : nikeShort 구조체
      adidasShoe := adidasFactory.makeShoe() // 반환값 : adidasShoe 구조체
      adidasShort := adidasFactory.makeShort() // 반환값 : adidasShort 구조체
      printShoeDetails(nikeShoe)
      printShortDetails(nikeShort)
      printShoeDetails(adidasShoe)
      printShortDetails(adidasShort)
  }
  
  func printShoeDetails(s iShoe) {
      fmt.Printf("Logo: %s", s.getLogo())
      fmt.Println()
      fmt.Printf("Size: %d", s.getSize())
      fmt.Println()
  }
  
  func printShortDetails(s iShort) {
      fmt.Printf("Logo: %s", s.getLogo())
      fmt.Println()
      fmt.Printf("Size: %d", s.getSize())
      fmt.Println()
  }
  ```



### 1-2.  빌더 패턴( Builder Pattern )

> 복잡한 객체의 생성을 표현으로부터 분리시키는 것
>
> * => 같은 생성 과정(process),  다른 표현

* 객체의 생성의 단계(steps)들을 추상화해서, 이런 생성 단계의 구현이 달라지면 객체의 다른 표현이 나오도록 하는 것

  

* ![img](https://t1.daumcdn.net/cfile/tistory/171FBC434F9FF8A935)
  * **Builder** - 객체(제품)를 생성하는 추상 인터페이스
  * **Concrete Builder** - Builder의 구현 클래스. 다른 객체를 생성할 수 있도록 하는 구체적인 클래스. 객체를 만들기 위해 부분(부품)을 생성하고 조립한다.
  * **Director** - Director클래스는 객체 생성의 정확한 순서(sequenct)를 다루는 부분에 책임이 있다. 이 클래스는 ConcreteBuilder를 인자로 받아서 필요한 동작을 수행한다.
  * **Product** - Builder를 이용해서 Director가 만들어낸 최종 객체



#### 예시 코드



![img](https://i2.wp.com/golangbyexample.com/wp-content/uploads/2019/09/Builder-Design-Patter-1.jpg?w=640&ssl=1)

| Director           | director.go      |
| ------------------ | ---------------- |
| Builder Interface  | iBuilder.go      |
| Concrete Builder 1 | normalBuilder.go |
| Concrete Builder 2 | iglooBuilder.go  |
| Concrete Product   | house.go         |



**iBuilder.go**

```go
package main

/* house 구조체 변수에 대한 Setter, Getter 메소드의 집합 인터페이스 */
type iBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() house
}

/* string값 builderType에 따라 iBuilder 타입의 구조체 포인터를 반환하는 함수 */
func getBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalBuilder{}
    }
    if builderType == "igloo" {
        return &iglooBuilder{}
    }
    return nil
}
```

**normalBuilder.go**

```go
package main

/* 구조체 house와 같은 변수를 가지는 normalBuilder*/
type normalBuilder struct {
    windowType string
    doorType   string
    floor      int
}

/* 구조체 포인터 객체 생성자 */
func newNormalBuilder() *normalBuilder {
    return &normalBuilder{}
}
/* window 값 Setter */
func (b *normalBuilder) setWindowType() {
    b.windowType = "Wooden Window"
}
/* door 값 Setter */
func (b *normalBuilder) setDoorType() {
    b.doorType = "Wooden Door"
}
/* floor 값 Setter */
func (b *normalBuilder) setNumFloor() {
    b.floor = 2
}
/* 
매개변수 : normalBuilder 구조체 포인터
반환값 : house 구조체 타입의 객체*/
func (b *normalBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}
```

**iglooBuilder.go**

```go
package main

/* 구조체 house와 같은 변수를 가지는 iglooBuilder*/
type iglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

/* 구조체 포인터 객체 생성자 */
func newIglooBuilder() *iglooBuilder {
    return &iglooBuilder{}
}

/* window 값 Setter */
func (b *iglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

/* door 값 Setter */
func (b *iglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

/* floor 값 Setter */
func (b *iglooBuilder) setNumFloor() {
    b.floor = 1
}

/* floor 값 Setter */
func (b *iglooBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}
```

**house.go**

```go
package main

/* window, door, floor 변수를 가진 구조체 house */
type house struct {
    windowType string
    doorType   string
    floor      int
}
```

**director.go**

```go
package main

/* iBuilder 인터페이스를 implement하는 builder 객체가 변수인 director 구조체 */
type director struct {
    builder iBuilder
}

/* director 구조체 내의 변수값을
매개변수로 받는 iBuilder 인터페이스로 할당해주는 함수
반환값 : director 구조체 포인터*/
func newDirector(b iBuilder) *director {
    return &director{
        builder: b,
    }
}

/* 이미 생성되어있는 director 구조체 객체 내의 builder 변수값의 Setter*/  
func (d *director) setBuilder(b iBuilder) {
    d.builder = b
}

/* 구조체 내의 builder 변수 내의 메소드 실행하는 함수 */
func (d *director) buildHouse() house {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    normalBuilder := getBuilder("normal") //normalBuilder 구조체 포인터 반환
    iglooBuilder := getBuilder("igloo")  //iglooBuilder 구조체 포인터 반환

    director := newDirector(normalBuilder) //normalBuilder 타입의 builder 변수를 갖는 director 선언 및 정의
    normalHouse := director.buildHouse() //normalBuilder 구조체 내의 생성자 값으로 house 구조체 타입의 normalHouse 객체 선언 및 정의

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder) //director내의 builder 값 재정의
    iglooHouse := director.buildHouse() //iglooBuilder 구조체 내의 생성자 값으로 house 구조체 타입의 iglooHouse 객체 선언 및 정의

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
```



### 1-3. 팩토리 패턴( Factory pattern )

> 객체의 생성에 관련된 인터페이스를 정의하고, 
>
> 인터페이스를 구현하는 클래스에서 어떤 클래스를 인스턴스화(생성)할 것인지 정하게 두는 것이다. 
>
> 즉, 팩토리 메서드는 서브클래스에서 인스턴스화를 하도록 미루는 것

* 생성할 객체에 대한 정확한 클래스를 구분할 필요없이 객체를 생성
* 객체를 생성할 때 별도의 메서드를 정의
* 서브클래스에서 생성될 객체의 종류를 정할 수 있게 함



#### 예시코드



![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Factory-Design-Pattern-1.jpg?w=640&ssl=1)



| ProductFactory      | gunFactory |
| ------------------- | ---------- |
| iProduct            | iGun       |
| Product             | gun        |
| Concrete iProduct 1 | ak47g      |
| Concrete iProduct 1 | maverick   |
| Client              | main       |



```go
package main

import "fmt"

/* gun 구조체에 대한 Setter, Getter 메소드들의 집합 인터페이스 */
type iGun interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

/* gun 구조체 */
type gun struct {
    name  string
    power int
}

/* gun 구조체 메소드 Setter, Getter */
func (g *gun) setName(name string) {
    g.name = name
}

func (g *gun) getName() string {
    return g.name
}

func (g *gun) setPower(power int) {
    g.power = power
}

func (g *gun) getPower() int {
    return g.power
}

/* gun 구조체 embed하는 ak47 구조체 */
type ak47 struct {
    gun
}

/* ak47에 대한 구조체 포인터 생성자, iGun 타입의 gun 구조체의 변수값을 담아서 반환 */
func newAk47() iGun {
    return &ak47{
        gun: gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}

/* gun 구조체 embed하는 maverick 구조체 */
type maverick struct {
    gun
}

/* maverick에 대한 구조체 포인터 생성자, iGun 타입의 gun 구조체의 변수값을 담아서 반환 */
func newMaverick() iGun {
    return &maverick{
        gun: gun{
            name:  "Maverick gun",
            power: 5,
        },
    }
}

/* iGun 타입의 해당 concrete gun 타입의 객체 생성 및 반환 */
func getGun(gunType string) (iGun, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "maverick" {
        return newMaverick(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
    ak47, _ := getGun("ak47") //ak47 구조체 포인터 반환
    maverick, _ := getGun("maverick") //maverick 구조체 포인터 반환
    printDetails(ak47)
    printDetails(maverick)
}

func printDetails(g iGun) {
    fmt.Printf("Gun: %s", g.getName())
    fmt.Println()
    fmt.Printf("Power: %d", g.getPower())
    fmt.Println()
}
```



### 1-4. 오브젝트 풀 패턴( Object Pool Pattern )

> 객체를 매번 할당, 해제하지 않고 
>
> 고정 크기 풀에 들어 있는 객체를 **재사용**함으로써 메모리 사용 성능을 개선
>
> 주로 쓰이는 용도 : 
>
> - 객체를 빈번하게 생성/삭제할 때
> - 객체들의 크기가 비슷할 때
> - 객체를 힙에 생성하기가 느리거나, 메모리 단편화가 우려됨
> - 데이터베이스 연결, 네트워크 연결처럼 접근 비용이 비싸면서 재사용 가능한 자원을 객체가 캡슐화할 때



* 풀에 속한 객체는 현재 자신이 **'사용 중'인지 여부**를 알 수 있는 방법을 제공해야 한다.

* 초기화될 때 사용할 객체들을 미리 생성하고(보통 같은 종류의 객체를 연속된 배열에 넣는다), 이들 객체를 '사용 안 함' 상태로 초기화한다.

* 새로운 객체가 필요하면 풀에 요청한다. 

  * 풀은 사용 가능한 객체를 찾아 '사용 중'으로 초기화한 뒤 반환한다. 
  * 객체를 더 이상 사용하지 않는다면 '사용 안 함' 상태로 되돌린다.

  

- 객체 초기화 시 주의점은?
  -  객체를 풀 안에서 초기화한다면, 
    - 객체를 완전히 캡슐화 가능하다. 풀 클래스가 객체 초기화 때문에 복잡해진다.
  -  객체를 풀 바깥에서 초기화한다면, 
    - 풀의 인터페이스가 단순해진다. 외부 코드에서 객체 생성 실패 시의 처리를 할 수 있다.

#### 예시코드

**iPoolObject.go**

```go
package main

/* 두 개의 다른 pool 객체를 비교할 때 어떤 id 값을 가져오던지 이용할 수 있는 함수 */
type iPoolObject interface {
    getID() string 
}
```

**pool.go**

```go
package main

import (
    "fmt"
    "sync"
)

/* pool 구조체 */
type pool struct {
    idle   []iPoolObject
    active []iPoolObject
    capacity int
    mulock   *sync.Mutex
}

/* pool 구조체를 초기화하는 함수 */
func initPool(poolObjects []iPoolObject) (*pool, error) {
    if len(poolObjects) == 0 { //pool 객체가 비었을 때에 대한 예외 처리
        return nil, fmt.Errorf("Cannot craete a pool of 0 length")
    }
    active := make([]iPoolObject, 0) //iPollObject 배열을 생성
    pool := &pool{
        idle:     poolObjects, //pool을 만들 때 매개변수로 들어온 배열
        active:   active,
        capacity: len(poolObjects),
        mulock:   new(sync.Mutex),
    }
    return pool, nil // 생성된 pool 구조체 포인터를 반환
}

/* pool 객체를 꺼내는 함수 */
func (p *pool) loan() (iPoolObject, error) {
    p.mulock.Lock()
    defer p.mulock.Unlock() // lock을 걸어둔 뒤, 함수가 종료되면 unlock 하도록
    if len(p.idle) == 0 {
        return nil, fmt.Errorf("No pool object free. Please request after sometime")
    }
    obj := p.idle[0] //꺼낼 pool 객체 obj
    p.idle = p.idle[1:] //obj를 제외한 나머지를 pool에서 이용가능한 객체로 둔다.
    p.active = append(p.active, obj) //꺼낸 obj는 현재 이용중인 active 객체 리스트에 추가한다.
    fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
    return obj, nil //obj 객체 반환
}

/* pool 객체를 돌려받는 함수 */
func (p *pool) receive(target iPoolObject) error {
    p.mulock.Lock()
    defer p.mulock.Unlock() // lock을 걸어둔 뒤, 함수가 종료되면 unlock 하도록
    err := p.remove(target) //obj로 pool에서 나왔을 객체 target을 remove한다. 
    if err != nil {
        return err
    }
    p.idle = append(p.idle, target) // 다시 이용가능한 객체로 pool에 추가한다.
    fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
    return nil
}
/* pool 내에 target은 active 리스트 안에 존재하므로, active 리스트 안의 target이 remove되는 것. */
func (p *pool) remove(target iPoolObject) error {
    currentActiveLength := len(p.active) 
    for i, obj := range p.active { // 현재 이용중인 active 리스트 판단
        if obj.getID() == target.getID() { //없애려는 target과 active 리스트 내 obj의 ID값이 같다면,
            p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1] // 그 obj의 index값을 기준으로 순서를 swap 한다.(맨 마지막 값과 obj의 index 위치를 바꾼다.)
            p.active = p.active[:currentActiveLength-1] //obj가 swap된 마지막 index를 제외한 만큼 active 리스트 슬라이스로 갱신
            return nil // 갱신에 성공하면 nil(NULL)을 반환
        }
    }
    return fmt.Errorf("Targe pool object doesn't belong to the pool")
}
```

**connection.go**

```go
package main

/* pool의 객체의 id 값을 담을 구조체 */
type connection struct { 
    id string
}

/* 함수의 메소드 */
func (c *connection) getID() string {
    return c.id
}
```

**main.go**

```go
package main

import (
    "log"
    "strconv"
)

func main() {
    connections := make([]iPoolObject, 0) //iPoolObject 타입의 배열 정의
    for i := 0; i < 3; i++ {
        c := &connection{id: strconv.Itoa(i)} //id 값을 대입한다.(Itoa : 숫자를 문자열로 변환하는 함수 )
        connections = append(connections, c) //connections 배열에 c를 추가
    }
    pool, err := initPool(connections) //connetioins 배열로 pool 객체 생성
    if err != nil {
        log.Fatalf("Init Pool Error: %s", err)
    }
    conn1, err := pool.loan() //pool obj 객체 할당
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    conn2, err := pool.loan() //pool obj 객체 할당 2
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    pool.receive(conn1) //pool obj 객체 반환
    pool.receive(conn2) //pool obj 객체 반환
}
```





### 1-5. 프로토 타입 패턴( Prototype Pattern )

> 객체에 의해 생성될 객체의 타입이 결정되는 생성 디자인 패턴
>
> 이 패턴은 새로운 객체를 생성하기 위해 clone(복제)을 이용. 



* 클라이언트 어플리케이션에서 객체 생성자의 서브 클래싱을 피한다.
  * 반대로 추상 팩토리 패턴에서는 객체 생성자를 서브 클래싱해야만 한다.

* 주어진 애플리케이션에서 일반적인 방법(new 키워드를 이용한 방법)으로 객체를 생성할 때에 필요한 비용이 엄청난 경우에 이 비용을 없앨 수 있다.



![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/10/Prototype-Pattern.jpg?w=640&ssl=1)



#### 예시코드

| prototype interface  | inode.go  |
| -------------------- | --------- |
| Concrete Prototype 1 | file.go   |
| Concrete Prototype 2 | folder.go |
| client               | main.go   |



**inode.go**

```go
package main

/* inode 인터페이스 - print와 clone 메소드 포함 */
type inode interface {
    print(string)
    clone() inode
}
```



**file.go**

```go
package main

import "fmt"

/* file 구조체 */
type file struct {
    name string
}

/* file 구조체 포인터를 받으면 indent를 포함하여 이름을 출력하는 함수 */
func (f *file) print(indentation string) {
    fmt.Println(indentation + f.name + "_clone")
}

/* file 구조체에 f.name 값으로 초기화한 후 포인터를 반환하는 함수 */
func (f *file) clone() inode {
    return &file{name: f.name}
}
```



**folder.go**

```go
package main

import "fmt"

/* folder 구조체 */
type folder struct {
    childrens []inode // inode 배열 타입을 가지는 childrens 변수
    name      string // 폴더 이름
}

/* folder 구조체 포인터를 받으면 indent를 포함하여 이름을 출력하는 함수 */
func (f *folder) print(indentation string) {
    fmt.Println(indentation + f.name + "_clone)
    for _, i := range f.childrens {
        i.print(indentation + indentation)
    }
}

/* folder 구조체 포인터를 받으면 cloneFolder라는 새로운 구조체 생성, 
복사 대상 폴더의 이름 값으로 초기화한다.
for문으로 복사 대상 폴더 안의 file 값들을 복사, cloneFolder 안의 childrens 값이 될
tempChildren 배열에 append한다.
복사가 완료된 새로운 구조체 반환*/             
func (f *folder) clone() inode {
    cloneFolder := &folder{name: f.name}
    var tempChildrens []inode
    for _, i := range f.childrens {
        copy := i.clone()
        tempChildrens = append(tempChildrens, copy)
    }
    cloneFolder.childrens = tempChildrens
    return cloneFolder
}
```



**main.go**

```go
package main

import "fmt"

func main() {
    file1 := &file{name: "File1"} //이름값 초기화한 file1 구조체
    file2 := &file{name: "File2"} //이름값 초기화한 file2 구조체
    file3 := &file{name: "File3"} //이름값 초기화한 file3 구조체
    folder1 := &folder{
        childrens: []inode{file1}, //file1을 배열 원소로 갖는 folder1 구조체
        name:      "Folder1",
    }
    folder2 := &folder{ //folder1, file2, file3을 배열 원소로 갖는 folder2 구조체
        childrens: []inode{folder1, file2, file3},
        name:      "Folder2",
    }
    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.print("  ")
    cloneFolder := folder2.clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.print("  ")
}
```



**결과**

```
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
```



### 1-6. 싱글톤 패턴( Singleton Design Pattern)

> 가장 많이 쓰이는 디자인 패턴 중 하나
>
> 오직 하나의 구조체 인스턴스가 존재해야 할 때 사용
>
> 하나의 인스턴스 (single instance) = singleton object



* 싱글톤 패턴이 쓰이는 곳
  * DB instance
  * Logger instance 
* 구조체가 처음 초기화될 때 생성된다.
* 주로  getInstance()와 같은 메소드를 이용하여 생성되며, 한 번 생성된 이후로는 함수를 호출하더라도 이미 생성되어있는 동일한 객체가 반환된다.



* Go lang의 경우, 다수의 고루틴이 하나의 인스턴스에 접촉하려고 시도하더라도, 하나의 고루틴만 인스턴스에 접촉할 수 있도록 제한해야하는 특징이 있다.

##### 싱글톤 패턴 구현 기본 코드

```go
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creting Single Instance Now")
            singleInstance = &single{}
        } else {
            fmt.Println("Single Instance already created-1")
        }
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

1. 첫번째 singleInstance에 대한 nil 체크는 getInstance() 메소드가 호출될 때마다 lock을 걸어 낭비하지 않도록 하기 위함이다.

   체킹이 실패하면 singleInstance가 이미 만들어졌다는 뜻이다.

2. singleInstance는 lock이 걸린 뒤 생성되게 된다.

3. 생성 후 또다시 nil을 체크하는 이유는 오직 하나의 고루틴만이 singleInstance를 생성했는가를 체크하기 위한 것이다. 



## 2. 행동 디자인 패턴( Behavioral Design Pattern )



### 2-1. 책임 연쇄 패턴( Chain-responsibility pattern )

> 명령 객체와 일련의 처리 객체를 포함하는 디자인 패턴



* 각각의 처리 객체는 명령 객체를 처리할 수 있는 연산의 집합이고, 
* 체인 안의 처리 객체가 핸들할 수 없는 명령은 다음 처리 객체로 넘겨진다. 
* 이 작동방식은 새로운 처리 객체부터 체인의 끝까지 다시 반복된다.
* 병원에 가게 되었을 때, 예약->의사검진->약국->캐셔 의 형태로 방문하게 되는 일련의 과정을 예로 들 수 있다.



* 같은 요청을 처리할 수 있는 여러 개의 핸들러가 존재할 때 쓰인다.

* Client가 자신의 request가 어떤 핸들러에 의해 처리되는지 알 수 없게 하고싶을 때.

  * 이 경우, Client는 체인 처리 객체의 첫번째 객체만 알 수 있다.

  

![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Chain-of-Responsibility-Design-Pattern-1.jpg?w=640&ssl=1)



#### 예시코드

| handler            | department.go |
| ------------------ | ------------- |
| Concrete Handler 1 | account.go    |
| Concrete Handler 2 | doctor.go     |
| Concrete Handler 3 | medical.go    |
| Concrete Handler 4 | cashier.go    |
| Client             | main.go       |



**department.go**

```go
package main
/* 각 핸들러의 메소드를 모아놓은 인터페이스 객체 */
type department interface {
    execute(*patient)
    setNext(department)
}
```

**reception.go**

```go
package main

import "fmt"

/* department 타입의 next 변수를 갖는 reception 구조체 */
type reception struct {
    next department
}

/*  */
func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p) // 이미 Done이 체크되어 있다면 다음 단계의 execute()를 시도
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true // execute() 실행 시 입력받은 patient 구조체의 변수값 변경 
    r.next.execute(p) //미리 setNext로 연결해둔 next 변수의 execute() 실행
}

func (r *reception) setNext(next department) {
    r.next = next //reception 다음 처리단계의 핸들러와 연결
}
```

**doctor.go**

```go
package main

import "fmt"

type doctor struct {
    next department
}

func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *doctor) setNext(next department) {
    d.next = next
}
```

**medical.go**

```go
package main

import "fmt"

type medical struct {
    next department
}

func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *medical) setNext(next department) {
    m.next = next
}
```

**cashier.go**

```go
package main

import "fmt"

type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}

```

**patient.go**

```go
package main

/* 이름과 각각의 핸들러 통과여부를 변수로 갖는 patient 구조체 */
type patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}
```

**main.go**

```go
package main

func main() {
    cashier := &cashier{}
    //medical의 next : cashier
    medical := &medical{}
    medical.setNext(cashier)
    //doctor의 next : medical
    doctor := &doctor{}
    doctor.setNext(medical)
    //reception의 next : doctor
    reception := &reception{}
    reception.setNext(doctor)
    patient := &patient{name: "abc"}
    //patient 구조체 생성 및 초기화
    reception.execute(patient)
}
```





### 2-2. 커맨드 패턴( Command Design Pattern )

> 요청을 객체의 형태로 *캡슐화*하여 사용자가 보낸 요청을 나중에 이용할 수 있도록 
>
> *매서드 이름, 매개변수* 등 요청에 필요한 정보를 *저장 또는 로깅, 취소*할 수 있게 하는 패턴



* 용어 정리
  * 명령(command), 
  * 수신자(receiver), 
  * 발동자(invoker), 
  * 클라이언트(client)

* 커맨드 객체는 수신자 객체를 가지고 있다.
  * 수신자의 메서드를 호출하고, 이에 수신자는 자신에게 정의된 메서드를 수행한다.
*  커맨드 객체는 별도로 발동자 객체에 전달되어 명령을 발동하게 한다. 
  * 발동자 객체는 필요에 따라 명령 발동에 대한 기록을 남길 수 있다. 
  * 한 발동자 객체에 다수의 커맨드 객체가 전달될 수 있다.
*  클라이언트 객체는 발동자 객체와 하나 이상의 커맨드 객체를 보유한다. 
  * 클라이언트 객체는 어느 시점에서 어떤 명령을 수행할지를 결정한다. 
  * 명령을 수행하려면, 클라이언트 객체는 발동자 객체로 커맨드 객체를 전달한다.



* TV로 예를 들자면, 
  * **수신자**는 TV, 
  * **명령**은 '전원을 켜라'는 명령
  * **발동자**는 '전원을 켜라'는 명령을 가지는 리모컨의 전원 버튼, 혹은 TV의 전원버튼



![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Command-Design-Pattern.jpg?w=640&ssl=1)



#### 예시코드

| Invoker            | button.go     |
| ------------------ | ------------- |
| Command Interface  | command.go    |
| Concrete Command 1 | onCommand.go  |
| Concrete Command 2 | offCommand.go |
| Receiver Interface | device.go     |
| Concrete Receiver  | tv.go         |
| Client             | main.go       |



**button.go**

```go
package main

/* button 구조체 */
type button struct {
    command command
}

/* press() 메소드 실행시 command의 execute() 실행 */
func (b *button) press() {
    b.command.execute()
}
```

**command.go**

```go
package main

/* execute() 메소드를 갖는 command 인터페이스 */
type command interface {
    execute()
}
```

**onCommand.go**

```go
package main

/* 전원을 키는 명령을 갖는 구조체 */
type onCommand struct {
    device device
}

func (c *onCommand) execute() {
    c.device.on()
}
```

**offCommand.go**

```go
package main

type offCommand struct {
    device device
}

func (c *offCommand) execute() {
    c.device.off()
}
```

**device.go**

```go
package main

/* on, off 메소드의 집합 인터페이스 */
type device interface {
    on()
    off()
}
```

**tv.go**

```go
package main

import "fmt"

/* 켜져있는지 여부를 변수로 갖는 tv 구조체 */
type tv struct {
    isRunning bool
}

/* device 객체의 on, off에 따라 값의 변경 */
func (t *tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}
```

**main.go**

```go
package main

func main() {
    tv := &tv{}
    onCommand := &onCommand{
        device: tv,
    }
    offCommand := &offCommand{
        device: tv,
    }
    onButton := &button{
        command: onCommand,
    }
    onButton.press()
    offButton := &button{
        command: offCommand,
    }
    offButton.press()
}
```





### 2-3. 반복자 패턴( Iterator Design Pattern )

> 객체 지향 프로그래밍에서 반복자를 사용하여 컨테이너를 가로지르며 컨테이너의 요소들에 접근하는 디자인 패턴



* 기반이 되는 표현을 노출시키지 않고 연속적으로 객체 요소에 접근하는 방법을 제공하는 것



* Iterator Interface(반복자 인터페이스) :
  * hasNext(), getNext()와 같은 기본 동작을 제공
* Collection interface( 객체 모음 인터페이스 ):
  * iterator 타입을 반환하는 createIterator() 메소드 정의
  * 거쳐야하는 객체들의 모음을 나타냄
* Concrete Iterator, Concrete Collection:
  * 인터페이스들을 implement하는 보다 구체적인 변수들



![img](https://i0.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Iterator-Design-Pattern-1.jpg?w=640&ssl=1)



#### 예시코드

| Collection          | collection.go     |
| ------------------- | ----------------- |
| Concrete Collection | userCollection.go |
| Iterator            | mac.go            |
| Concrete Iterator 1 | userIterator.go   |
| Client              | main.go           |



**collection.go**

```go
package main

/* 반복자 생성하는 colletion 인터페이스 */
type collection interface {
    createIterator() iterator
}
```

**userCollection.go**

```go
package main

/* user 구조체 포인터 배열을 갖는 userCollection 구조체 */
type userCollection struct {
    users []*user
}

/* 배열을 iterator에 대입 */
func (u *userCollection) createIterator() iterator {
    return &userIterator{
        users: u.users,
    }
}
```

**iterator.go**

```go
package main

/* 다음 배열이 있는지 탐색한 후 그 배열을 가져오는 메소드 집합 인터페이스 */
type iterator interface {
    hasNext() bool
    getNext() *user
}
```

**userIterator.go**

```go
package main

/* userCollection을 iterate 하는 구조체 */
type userIterator struct {
    index int
    users []*user
}

/* index가 총 길이보다 작다면 true */
func (u *userIterator) hasNext() bool {
    if u.index < len(u.users) {
        return true
    }
    return false
}

/* 다음 배열 값을 가져오는 함수 */
func (u *userIterator) getNext() *user {
    if u.hasNext() {
        user := u.users[u.index]
        u.index++
        return user
    }
    return nil
}
```

**user.go**

```go
package main

/* user 구조체 */
type user struct {
    name string
    age  int
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    user1 := &user{ //user1 = user구조체 포인터
        name: "a",
        age:  30,
    }
    user2 := &user{ //user2 = user구조체 포인터
        name: "b",
        age:  20,
    }
    userCollection := &userCollection{
        users: []*user{user1, user2}, 
    }
    iterator := userCollection.createIterator()
    for iterator.hasNext() {
        user := iterator.getNext()
        fmt.Printf("User is %+v\n", user)
    }
}
```



### 2-4. 중재자 패턴( Mediator Pattern )

> 객체들의 집합이 어떻게 상호작용하는지를 함축해놓은 객체를 정의
>
> 이 패턴은 프로그램의 실행 행위를 변경할 수 있기 때문에 행위 패턴(behavioral design pattern)으로 간주된다.



* 중재자 패턴을 사용하면 객체 간 통신은 중재자 객체 안에 함축된다. 
* 객체들은 더 이상 다른 객체와 서로 직접 통신하지 않으며 대신 중재자를 통해 통신한다.
  *  객체 간 의존성을 줄이며, 최종적으로 결합도를 감소



* train, platform, stationManager을 가지고 예를 들 수 있다.
* stationManager은 플랫폼과 열차 사이의 중재자로서 역할한다.
  * 각 플랫폼마다 도착하고 출발하는 열차들에 대해 관리한다.
  * 대기 중인 열차 큐를 관리한다.



#### 예시코드

**train.go**

```go
package main

/* train과 관련된 메소드들의 집합 */
type train interface {
    requestArrival()
    departure()
    permitArrival()
}
```

**passengerTrain.go**

```go
package main

import "fmt"

type passengerTrain struct {
    mediator mediator
}
/* train 인터페이스 내의 메소드들에 따라 구조체 내의 mediator 구조체 변수 값이 바뀐다. */
func (g *passengerTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("PassengerTrain: Landing")
    } else {
        fmt.Println("PassengerTrain: Waiting")
    }
}

func (g *passengerTrain) departure() {
    fmt.Println("PassengerTrain: Leaving")
    g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
    fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}
```

**goodsTrain.go**

```go
package main

import "fmt"

type goodsTrain struct {
    mediator mediator
}

func (g *goodsTrain) requestArrival() {
    if g.mediator.canLand(g) {
        fmt.Println("GoodsTrain: Landing")
    } else {
        fmt.Println("GoodsTrain: Waiting")
    }
}

func (g *goodsTrain) departure() {
    g.mediator.notifyFree()
    fmt.Println("GoodsTrain: Leaving")
}

func (g *goodsTrain) permitArrival() {
    fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
```

**mediator.go**

```go
package main

/* mediator 인터페이스 */
type mediator interface {
    canLand(train) bool
    notifyFree()
}
```

**stationManager.go**

```go
package main

import "sync"

/* mediator 인터페이스를 implement */
type stationManager struct {
    isPlatformFree bool
    lock           *sync.Mutex
    trainQueue     []train
}

/* 생성자, 첫 생성시 초기화값 */
func newStationManger() *stationManager {
    return &stationManager{
        isPlatformFree: true,
        lock:           &sync.Mutex{},
    }
}

func (s *stationManager) canLand(t train) bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.isPlatformFree { //플랫폼이 비었다면 train이 Land할 수 있고, 각각 bool값 변경 및 반환
        s.isPlatformFree = false 
        return true
    }
    s.trainQueue = append(s.trainQueue, t) //비어있지 않다면 train 대기 큐에 추가
    return false
}

func (s *stationManager) notifyFree() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if !s.isPlatformFree { //플랫폼이 비어있지 않다면 비어있게 변경,
        s.isPlatformFree = true
    }
    if len(s.trainQueue) > 0 { 
        firstTrainInQueue := s.trainQueue[0] // 대기 큐의 첫번째 열차를 꺼내온뒤
        s.trainQueue = s.trainQueue[1:] //슬라이싱
        firstTrainInQueue.permitArrival() // 첫번째 열차가 진입 허가 출력문 출력
    }
}
```

**main.go**

```go
package main

func main() {
    stationManager := newStationManger()
    passengerTrain := &passengerTrain{
        mediator: stationManager,
    }
    goodsTrain := &goodsTrain{
        mediator: stationManager,
    }
    passengerTrain.requestArrival()
    goodsTrain.requestArrival()
    passengerTrain.departure()
}
```



### 2-5. 메멘토 패턴( Memento Design Pattern )

> 객체를 이전 상태로 되돌릴 수 있는 기능을 제공하는 소프트웨어 디자인 패턴
>
> (롤백을 통한 실행 취소)



* 구성요소
  * 오리지네이터(originator), 
    * 오리지네이터는 내부 상태를 보유하고 있는(메멘토로 저장되어 있는) 일부 객체
    * **savememento()**:
      * 내부 상태를 메멘토 객체에 저장하는 오리지네이터 메소드
    * **restorememento()**:
      * 입력값 : 메멘토 객체, 
      * 입력받은 메멘토 객체 내의 state로 오리지네이터 복원
  * 케어테이커(caretaker), 
    * 다수의 메멘토를 저장하는 객체
    * index 값이 주어지면 그에 맞는 메멘토를 반환한다.
  * 메멘토(memento)
    * 메멘토 객체 자신은 불투명 자료형
    * 오리지네이터의 상태를 저장하고 있다.

* 주의할 점
  * 객체의 내부 상태가 외부적으로 저장되어야 이후에 객체가 그 상태를 복구할 수 있다.
  * 객체의 캡슐화가 훼손되지 않아야 한다



#### 예시코드

**originator.go**

```go
package main

/* 상태값을 가지는 originator */
type originator struct {
    state string
}

func (e *originator) createMemento() *memento {
    return &memento{state: e.state}
}

func (e *originator) restorememento(m *memento) {
    e.state = m.getSavedState()
}

func (e *originator) setState(state string) {
    e.state = state
}

func (e *originator) getState() string {
    return e.state
```

**memento.go**

```go
package main

/* 상태 저장하는 memento 구조체 */
type memento struct {
    state string
}

func (m *memento) getSavedState() string {
    return m.state
}
```

**caretaker.go**

```go
package main

/* 메멘토 구조체 포인터 배열을 갖는 케어 테이커 */
type caretaker struct {
    mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
    c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMenento(index int) *memento {
    return c.mementoArray[index]
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    caretaker := &caretaker{ 
        mementoArray: make([]*memento, 0), //메멘토 객체 저장위해서 배열 생성
    }
    originator := &originator{
        state: "A", //state를 A로 초기화
    }
    fmt.Printf("Originator Current State: %s\n", originator.getState())
    caretaker.addMemento(originator.createMemento()) //A를 메멘토 객체로 저장 
    
    originator.setState("B") // 상태 변경
    fmt.Printf("Originator Current State: %s\n", originator.getState())
    
    caretaker.addMemento(originator.createMemento()) //B를 메멘토 객체로 저장
    originator.setState("C") //상태 재 변경
    
    fmt.Printf("Originator Current State: %s\n", originator.getState())
    caretaker.addMemento(originator.createMemento()) //C를 메멘토 객체로 저장
    
    originator.restorememento(caretaker.getMenento(1)) //1번 index 값으로 복원
    fmt.Printf("Restored to State: %s\n", originator.getState())
    
    originator.restorememento(caretaker.getMenento(0)) //0번 index 값으로 복원
    fmt.Printf("Restored to State: %s\n", originator.getState())
}
```

```
/* 실행 결과 */
originator Current State: A
originator Current State: B
originator Current State: C
Restored to State: B
Restored to State: A
```



### 2-6. 옵저버 패턴( Observer Design Pattern )

> 객체의 상태 변화를 관찰하는 관찰자들, 즉 옵저버들의 목록을 객체에 등록
>
> 상태 변화가 있을 때마다 메서드 등을 통해 객체가 직접 목록의 각 옵저버에게 통지하도록 하는 디자인 패턴



* 주로 분산 이벤트 핸들링 시스템을 구현하는 데 사용

* 옵저버 또는 리스너(listener)라 불리는 하나 이상의 객체를 관찰 대상이 되는 객체에 등록
* 각각의 옵저버들은 관찰 대상인 객체가 발생시키는 이벤트를 받아 처리한다.



* 구성요소
  * Subject : 
    * 상태 변화가 감지되면 이벤트를 발생시키는 인스턴스
  * Observer :
    * Subject를 주시하다가 이벤트 알림을 받는다.



![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Observer-Design-Pattern-1.jpg?w=640&ssl=1)





#### 예시코드

| Subject             | subject.go  |
| ------------------- | ----------- |
| Concrete Subject    | item.go     |
| observer            | observer.go |
| Concrete Observer 1 | customer.go |
| Client              | main.go     |

**subject.go**

```go
package main

type subject interface {
    register(Observer observer)
    deregister(Observer observer)
    notifyAll()
}
```

**item.go**

```go
package main

import "fmt"

type item struct {
    observerList []observer //옵저버 배열 
    name         string
    inStock      bool
}

func newItem(name string) *item {
    return &item{
        name: name,
    }
}

func (i *item) updateAvailability() {
    fmt.Printf("Item %s is now in stock\n", i.name)
    i.inStock = true
    i.notifyAll()
}

func (i *item) register(o observer) {
    i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
    i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) notifyAll() {
    for _, observer := range i.observerList {
        observer.update(i.name)
    }
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
    observerListLength := len(observerList)
    for i, observer := range observerList {
        if observerToRemove.getID() == observer.getID() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
            return observerList[:observerListLength-1]
        }
    }
    return observerList
}
```

**observer.go**

```go
package main

/* 옵저버 */
type observer interface {
    update(string)
    getID() string
}
```

**customer.go**

```go
package main

import "fmt"

/* customer 구조체 */
type customer struct {
    id string
}

func (c *customer) update(itemName string) {
    fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
    return c.id
}
```

**main.go**

```go
package main

func main() {
    shirtItem := newItem("Nike Shirt") // 제품이 새로 등록되면 item과 옵저버 배열 생성
    observerFirst := &customer{id: "abc@gmail.com"} // 옵저버 구조체 생성, 초기화
    observerSecond := &customer{id: "xyz@gmail.com"} //옵저버 구조체 생성, 초기화
    shirtItem.register(observerFirst) //옵저버 배열에 append
    shirtItem.register(observerSecond) //옵저버 배열에 append
    shirtItem.updateAvailability() //item에 대한 stock값 업데이트, 등록된 옵저버 들에게 stock값이 변경되었음을 알리는 출력문이 출력된다.
}
```



### 2-7. 상태 패턴( State Design Pattern )

> 객체지향 방식으로 (유한)상태 기계를 구현하는 행위 소프트웨어 디자인 패턴
>
> 패턴의 인터페이스에 정의된 메소드들의 호출을 통해 현재의 전략을 전환할 수 있는 전략 패턴으로 해석할 수 있다.



* 자판기를 예로 상태 패턴을 설명할 수 있다.
* object가 다양한 상태에 놓일 수 있을 때 이용된다.
  * 현재 요구되는 사항에 따라 상태를 변경할 필요가 있을 때
  * 예 :  자판기
    * 물건 있음, 없음, 물건이 요청됨, 돈이 입금됨...

![img](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/09/State-Design-Pattern-1.jpg?w=640&ssl=1)



#### 예시 코드

| Context          | vendingMachine.go     |
| ---------------- | --------------------- |
| State Interface  | state.go              |
| Concrete State 1 | noItemState.go        |
| Concrete State 2 | hasItemState.go       |
| Concrete State 3 | itemRequestedState.go |
| Concrete State 4 | hasMoneyState.go      |



**vendingMachine.go**

```go
package main

import "fmt"

/* 각각의 state 타입 변수를 갖고 있는 vendingMachine 구조체 */
type vendingMachine struct {
    hasItem       state
    itemRequested state
    hasMoney      state
    noItem        state

    currentState state

    itemCount int
    itemPrice int
}

/* 구조체 생성자, 입력받은 값으로 초기화 */
func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
    v := &vendingMachine{
        itemCount: itemCount,
        itemPrice: itemPrice,
    }
    hasItemState := &hasItemState{
        vendingMachine: v,
    }
    itemRequestedState := &itemRequestedState{
        vendingMachine: v,
    }
    hasMoneyState := &hasMoneyState{
        vendingMachine: v,
    }
    noItemState := &noItemState{
        vendingMachine: v,
    }

    v.setState(hasItemState)
    v.hasItem = hasItemState
    v.itemRequested = itemRequestedState
    v.hasMoney = hasMoneyState
    v.noItem = noItemState
    return v
}

func (v *vendingMachine) requestItem() error {
    return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
    return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
    return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
    return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
    v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
    fmt.Printf("Adding %d items\n", count)
    v.itemCount = v.itemCount + count
}
```

**state.go**

```go
package main

/* state 인터페이스 */
type state interface {
    addItem(int) error
    requestItem() error
    insertMoney(money int) error
    dispenseItem() error
}
```

**noItemState.go**

```go
package main

import "fmt"

/* vemdingMachine 구조체의 포인터를 변수로 갖는다. */
type noItemState struct {
    vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
    return fmt.Errorf("Item out of stock")
}

func (i *noItemState) addItem(count int) error {
    i.vendingMachine.incrementItemCount(count)
    i.vendingMachine.setState(i.vendingMachine.hasItem)
    return nil
}

func (i *noItemState) insertMoney(money int) error {
    return fmt.Errorf("Item out of stock")
}
func (i *noItemState) dispenseItem() error {
    return fmt.Errorf("Item out of stock")
}
```

**hasItemState.go**

```go
package main

import "fmt"

type hasItemState struct {
    vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
    if i.vendingMachine.itemCount == 0 {
        i.vendingMachine.setState(i.vendingMachine.noItem)
        return fmt.Errorf("No item present")
    }
    fmt.Printf("Item requestd\n")
    i.vendingMachine.setState(i.vendingMachine.itemRequested)
    return nil
}

func (i *hasItemState) addItem(count int) error {
    fmt.Printf("%d items added\n", count)
    i.vendingMachine.incrementItemCount(count)
    return nil
}

func (i *hasItemState) insertMoney(money int) error {
    return fmt.Errorf("Please select item first")
}
func (i *hasItemState) dispenseItem() error {
    return fmt.Errorf("Please select item first")
}
```

**itemRequestedState.go**

```go
package main

import "fmt"

type itemRequestedState struct {
    vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
    return fmt.Errorf("Item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
    return fmt.Errorf("Item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(money int) error {
    if money < i.vendingMachine.itemPrice {
        fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
    }
    fmt.Println("Money entered is ok")
    i.vendingMachine.setState(i.vendingMachine.hasMoney)
    return nil
}

func (i *itemRequestedState) dispenseItem() error {
    return fmt.Errorf("Please insert money first")
}
```

**hasMoneyState.go**

```go
package main

import "fmt"

type hasMoneyState struct {
    vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() error {
    return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) addItem(count int) error {
    return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
    return fmt.Errorf("Item out of stock")
}

func (i *hasMoneyState) dispenseItem() error {
    fmt.Println("Dispensing Item")
    i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
    if i.vendingMachine.itemCount == 0 {
        i.vendingMachine.setState(i.vendingMachine.noItem)
    } else {
        i.vendingMachine.setState(i.vendingMachine.hasItem)
    }
    return nil
}
```

**main.go**

```go
package main

import (
    "fmt"
    "log"
)

/* Go lang에는 예외 처리가 없으므로 매번 err 객체를 받아와서 err 처리를 해주어야 한다.
각각 함수의 반환값에 error가 포함되어 있는 이유는 이것 때문. 
각각의 구조체 함수중 nil값을 반환하는 함수를 제외하고는 모두 에러상황 */
func main() {
    vendingMachine := newVendingMachine(1, 10) //구조체 포인터 반환
    err := vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }
    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()
    err = vendingMachine.addItem(2)
    if err != nil {
        log.Fatalf(err.Error())
    }

    fmt.Println()

    err = vendingMachine.requestItem()
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.insertMoney(10)
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        log.Fatalf(err.Error())
    }
}
```





### 2-8. 전략 패턴( Strategy Design Pattern )

> - 특정한 계열의 알고리즘들을 정의하고
> - 각 알고리즘을 캡슐화하며
> - 이 알고리즘들을 해당 계열 안에서 상호 교체가 가능하게 만든다.



* 이와 같은 패턴은 *메모리 캐시*에서 찾아볼 수 있음.

* 캐시 메모리에서 aging 기법으로 오래된 정보를 삭제하는 알고리즘
  * **LRU** – Least Recently Used: 
    * 가장 오래전에 쓰였던 엔트리를 삭제
  * **FIFO** – First In First Out: 
    * 가장 먼저 만들어진 엔트리를 삭제
  * **LFU** – Least Frequently Used: 
    * 가장 적게 쓰인 엔트리를 삭제



* object의 수행 방법을 런타임중에  object에 대한 사항을 변경하지 않고 변경하고 싶을 때
* 런타임 수행 방법을 고를 때 피하고 싶은 많은 조건들이 있을 때
* 오직 실행되었을 때 몇가지 수행방법만이 다른 비슷한 알고리즘들이 많이 있을 때



![img](https://i0.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Strategy-Design-Pattern-1.jpg?w=640&ssl=1)





#### 예시 코드

| Context                    | cache.go        |
| -------------------------- | --------------- |
| Strategy                   | evictionAlgo.go |
| Concrete Strategy Object 1 | lfu.go          |
| Concrete Strategy Object 2 | lru.go          |
| Concrete Strategy Object 3 | fifo.go         |
| Client                     | main.go         |

**evictionAlgo.go**

```go
package main

type evictionAlgo interface {
    evict(c *cache)
}
```

**fifo.go**

```go
package main

import "fmt"

type fifo struct {
}

func (l *fifo) evict(c *cache) {
    fmt.Println("Evicting by fifo strtegy")
}
```

**lru.go**

```go
package main

import "fmt"

type lru struct {
}

func (l *lru) evict(c *cache) {
    fmt.Println("Evicting by lru strtegy")
}
```

**lfu.go**

```go
package main

import "fmt"

type lfu struct {
}

func (l *lfu) evict(c *cache) {
    fmt.Println("Evicting by lfu strtegy")
}
```

**cache.go**

```go
package main

type cache struct {
    storage      map[string]string //저장되어있는 정보(map, key와 value의 형태로 저장)
    evictionAlgo evictionAlgo //용량 확보를 위해 선택할 알고리즘
    capacity     int //캐시 구조체 용량
    maxCapacity  int //최대 용량
}

/* 캐시 구조체 생성자, 기본 초기화 정보 */
func initCache(e evictionAlgo) *cache {
    storage := make(map[string]string)
    return &cache{
        storage:      storage,
        evictionAlgo: e,
        capacity:     0,
        maxCapacity:  2,
    }
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
    c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
    }
    c.capacity++
    c.storage[key] = value
}

func (c *cache) get(key string) {
    delete(c.storage, key)
}

func (c *cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}
```

**main.go**

```go
package main

func main() {
    lfu := &lfu{}
    cache := initCache(lfu)
    cache.add("a", "1")
    cache.add("b", "2")
    cache.add("c", "3")
    lru := &lru{}
    cache.setEvictionAlgo(lru)
    cache.add("d", "4")
    fifo := &fifo{}
    cache.setEvictionAlgo(fifo)
    cache.add("e", "5")
}
```





### 2-9. 템플릿 메소드 패턴( Template Method Design Pattern )

> 동작 상의 알고리즘의 프로그램 뼈대를 정의하는 행위 디자인 패턴
>
> 알고리즘의 구조를 변경하지 않고 알고리즘의 특정 단계들을 다시 정의할 수 있게 해준다.



* OTP(One Time Password) 로 템플릿 메소드의 예를 들 수 있다.
  * OTP 타입은 SMS나 EMAIL이 될 수 있다.
  * 어떤 타입이 되든, 생성 방법은 두 타입이 모두 같다.
  * 생성 방법
    * n개의 랜덤 숫자를 생성한다.
    * 나중에 유효값 인증을 위해 이 숫자를 저장
    * 값을 준비
    * 알림을 보내고
    * 암호 등록



#### 예시코드

**otp.go**

```go
package main

/* otp가 생성되기 위해 필요한 메소드의 집합 인터페이스 */
type iOtp interface {
    genRandomOTP(int) string
    saveOTPCache(string)
    getMessage(string) string
    sendNotification(string) error
    publishMetric()
}

type otp struct {
    iOtp iOtp
}

func (o *otp) genAndSendOTP(otpLength int) error {
    otp := o.iOtp.genRandomOTP(otpLength)
    o.iOtp.saveOTPCache(otp)
    message := o.iOtp.getMessage(otp)
    err := o.iOtp.sendNotification(message)
    if err != nil {
        return err
    }
    o.iOtp.publishMetric()
    return nil
}
```

**sms.go**

```go
package main

import "fmt"

/* sms 형식의 otp */
type sms struct {
    otp
}

func (s *sms) genRandomOTP(len int) string {
    randomOTP := "1234" //랜덤값 지정
    fmt.Printf("SMS: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
    fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
    return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
    fmt.Printf("SMS: sending sms: %s\n", message)
    return nil
}

func (s *sms) publishMetric() {
    fmt.Printf("SMS: publishing metrics\n")
}
```

**email.go**

```go
package main

import "fmt"

type email struct {
    otp
}

func (s *email) genRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *email) saveOTPCache(otp string) {
    fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
    return "EMAIL OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
    fmt.Printf("EMAIL: sending email: %s\n", message)
    return nil
}

func (s *email) publishMetric() {
    fmt.Printf("EMAIL: publishing metrics\n")
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    smsOTP := &sms{}
    o := otp{
        iOtp: smsOTP,
    }
    o.genAndSendOTP(4) //4자리의 sms OTP 생성 및 발급
    fmt.Println("")
    emailOTP := &email{}
    o = otp{
        iOtp: emailOTP,
    }
    o.genAndSendOTP(4) //4자리의 email OTP 생성 및 발급
}
```





### 2-10. 방문자 패턴( Visitor Design Pattern )

> 알고리즘을 객체 구조에서 분리시키는 디자인 패턴
>
> 개방 폐쇄 원칙을 적용하는 방법 중 하나



* 구조체를 수정하지 않고도 실질적으로 새로운 동작을 기존의 객체 구조에 추가가 가능해짐.

* shape을 기준으로 예를 들 수 있다.
  * Square
  * Circle
  * Triangle
  * 세 가지 모두 하나의 '모양'



![img](https://i2.wp.com/golangbyexample.com/wp-content/uploads/2019/11/Visitor-Design-Pattern-1.jpg?w=640&ssl=1)



#### 예시코드

| element            | shape.go             |
| ------------------ | -------------------- |
| Concrete Element A | square.go            |
| Concrete Element B | circle.go            |
| Concrete Element C | rectangle.go         |
| Visitor            | visitor.go           |
| Concrete Visitor 1 | areaCalculator.go    |
| Concrete Visitor 2 | middleCoordinates.go |
| Client             | main.go              |

**shape.go**

```go
package main

/* shape 인터페이스 */
type shape interface {
    getType() string
    accept(visitor)
}
```

**square.go**

```go
package main

type square struct {
    side int
}

func (s *square) accept(v visitor) {
    v.visitForSquare(s)
}

func (s *square) getType() string {
    return "Square"
}
```

**circle.go**

```go
package main

type circle struct {
    radius int
}

func (c *circle) accept(v visitor) {
    v.visitForCircle(c)
}

func (c *circle) getType() string {
    return "Circle"
}
```

**rectangle.go**

```go
package main

type rectangle struct {
    l int
    b int
}

func (t *rectangle) accept(v visitor) {
    v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
    return "rectangle"
}
```

**visitor.go**

```go
package main

type visitor interface {
    visitForSquare(*square)
    visitForCircle(*circle)
    visitForrectangle(*rectangle)
}
```

**areaCalculator.go**

```go
package main

import (
    "fmt"
)

type areaCalculator struct {
    area int
}

func (a *areaCalculator) visitForSquare(s *square) {
    //Calculate area for square. After calculating the area assign in to the area instance variable
    fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
    //Calculate are for circle. After calculating the area assign in to the area instance variable
    fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForrectangle(s *rectangle) {
    //Calculate are for rectangle. After calculating the area assign in to the area instance variable
    fmt.Println("Calculating area for rectangle")
}
```

**middleCoordinates.go**

```go
package main

import "fmt"

type middleCoordinates struct {
    x int
    y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
    //Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
    //Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitForrectangle(t *rectangle) {
    //Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for rectangle")
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    square := &square{side: 2}  //각각의 구조체 생성 및 값 초기화
    circle := &circle{radius: 3}
    rectangle := &rectangle{l: 2, b: 3}
   
    areaCalculator := &areaCalculator{} //구조체 포인터 생성
    square.accept(areaCalculator) //accept하면 areaCalculator안의 메소드 실행
    circle.accept(areaCalculator)
    rectangle.accept(areaCalculator)
   
    fmt.Println()
    middleCoordinates := &middleCoordinates{} //구조체 포인터 생성
    square.accept(middleCoordinates) //accept하면 middleCoordinates안의 메소드 실행
    circle.accept(middleCoordinates)
    rectangle.accept(middleCoordinates)
}
```







## 3. 구조 설계 패턴( Structural Design Pattern )

> 

