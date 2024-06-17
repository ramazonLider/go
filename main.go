package main

import (
	// "fmt"
	// "time"
	// "sync"
	// "time"
	// "math"
	// "sort"
	// "unicode"
	// "unicode/utf16"
	// "unicode/utf8"
	// "strings"
	// "time"
	// "github.com/k0kubun/pp"
)

// type SCounter struct {
// 	mymutex sync.Mutex
// 	v map[string]int
// }

// func (s *SCounter) Increment(key string) {
// 	s.mymutex.Lock()
// 	s.v[key]++
// 	s.mymutex.Unlock()
// }

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64{
// 	return math.Pi * c.Radius * c.Radius
// }

// type Rectangle struct{
// 	width float64
// 	height float64
// }

// func (r Rectangle) Area() float64{
// 	return r.width * r.height
// }

// func Example() {
// 	// var s string = "Hello world test"
// 	// bs :=  []byte("Ramazon")
// 	// fmt.Printf("We can see it: %v\n", bs)
// 	// for i := range bs {
// 	// 	if bs[i] % 2 == 0 {
// 	// 		bs[i] = bs[i] + 1
// 	// 		continue
// 	// 	}
// 	// 	bs[i] = bs[i] - 1
// 	// }
// 	// fmt.Printf("This is decode: %s\n", bs)
// }

// func myStrings() {
// 	var s string
// 	s = "Hello world" // boshqa tillarda masalan rus tilida 2 barobar ko'p oladi
// 	fmt.Println(s)
// 	fmt.Printf("Matnning uzunligi %d\n", len(s))
// 	fmt.Printf("Kesib olingan qiymat %v\n", s[7:])

// 	s = s + " Hello world"
// 	fmt.Println(s)

// 	for _, b := range s{
// 		// fmt.Printf("%v ", b)
// 		fmt.Printf("%s ", string(b))
// 	}
// 	fmt.Println("")
// }

// func printPerson(name string, surname string, age int) {
// 	fmt.Printf("First name %s\n", name)
// 	fmt.Printf("Surnamename %s\n", surname)
// 	fmt.Printf("Age %d\n", age)
// } // pastda struck bilan yozilgan

// type Person struct {
// 	name int
// 	surname int
// 	age int
// }

// type Employee interface {
// 	PrintName(name string)
// 	PrintSalary(b int, t int) int
// }

// type Emp int

// func (e Emp) PrintName(name string) {
// 	fmt.Println("Employee id: ", e)
// 	fmt.Println("Employee name: ", name)
// }
// func (e Emp) PrintSalary(basic int, tax int) int {
// 	var salary = (basic * tax) / 100
// 	return basic - salary
// }

// func (p *Person) ScoreOverall() int {
// 	all := (p.name + p.surname) / 2
// 	return all;
// }

// func hello(a int)  {
// 	fmt.Println(a*a)
// }

// func Reverse(n int) int {
// 	sum := 0
// 	for  n != 0 {
// 		sum = sum * 10 + n % 10
// 		n = n / 10
// 	}
// 	return sum
// }

// func Numlength(n int) int{
// 	l := 0
// 	for n != 0{
// 		n = n / 10
// 		l++
// 	}
// 	return l
// }

// func myPrint(n ...interface{})  {
// 	for _, elem := range n{
// 		fmt.Println(elem)
// 	}
// }

// type Describer interface{
// 	Describer() string
// }

// type Person struct {
// 	Name string
// 	Age int
// }

// func (p Person) Describer() {
// 	fmt.Printf("Name %v, Age %v\n", p.Name, p.Age)
// }

// type Book struct {
// 	Title string
// 	Author string
// }

// func (b Book) Describer() {
// 	fmt.Printf("Title %s, Author %s\n", b.Title, b.Author)
// }

func main() {
	// hello(4)
	// fmt.Println(Numlength(1234567))
	// fmt.Println("Hello world!")
	// fmt.Println(time.Now())
	// pp.Println("Hello")
	// fmt.Println(Reverse(2345))

	// var a int16 = 89;
	// var b int16 = 9;

	// var c int16 = a + b;
	// fmt.Println(c);

	// var s string = "Hi";
	// g  := "Hi"; // qanday typedagi ma'lumot ekanligini o'zi bilib oladi
	// fmt.Println(s);
	// fmt.Println(g);

	// var f, j string // birdaniga 2 ta berish
	// f = "Salom"
	// j = "Ramazon"
	// fmt.Println(f, j)
	// var a,b = 1,4
	// sum := a + b
	// fmt.Println(sum)
	// var a = 5
	// a++
	// fmt.Println(a)
	// var name string
	// var age int
	// fmt.Print("Input your name:")
	// fmt.Scanf("%s", &name) // input
	// fmt.Println("Input your age:") // PrintF formatlab chop etish
	// fmt.Scanf("%d", &age) // input d bo'lgani sababi age int
	// fmt.Printf("Name: %s\n", name)
	// fmt.Printf("Age: %d\n", age)
	// var myVar bool = true
	// var a,b int = 10, 20
	// if a < b {
	// 	fmt.Println(a, "a kichik")
	// } else if a > b {
	// 	fmt.Println(b)
	// } else if a == b {
	// 	fmt.Println("a = b")
	// } else if a != b {
	// 	fmt.Println("a != b")
	// } else {
	// 	fmt.Println("Ular son emas")
	// }

	// var n int
	// fmt.Scan(&n)
	// if n <= 7 {
	// 	fmt.Println("7 dan kichkina")
	// } else if n <= 16 {
	// 	fmt.Println("16dan kichik")
	// } else {
	// 	fmt.Println("0")
	// }

	// var n int
	// fmt.Scan(&n)
	// if n % 2 == 0 {
	// 	fmt.Println("Juft")
	// } else {
	// 	fmt.Println("toq")
	// }

	// if n < 100 || n > 0 {
	// 	fmt.Println("Togri")
	// } else if n >= 75{
	// 	fmt.Println("75 dan katta")
	// } else {
	// 	fmt.Println("A")
	// }

	// if n > 0 && n <= 10 {
	// 	fmt.Println("Togri")
	// } else {
	// 	fmt.Println("notogri")
	// }

	// num := 5
	// if num == 1 {
	// 	fmt.Println(1)
	// } else if num == 2{
	// 	fmt.Println(2)
	// } else if num == 3{
	// 	fmt.Println(3)
	// } else if num == 4{
	// 	fmt.Println(4)
	// } else if num == 5{
	// 	fmt.Println(5)
	// } else if num == 6{
	// 	fmt.Println(6)
	// } else {
	// 	fmt.Println("Katta qiymat")
	// }
	// switch num {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// case 3:
	// 	fmt.Println(3)
	// case 4:
	// 	fmt.Println(4)
	// case 5:
	// 	fmt.Println(5)
	// case 6:
	// 	fmt.Println(6)
	// default:
	// 	fmt.Println("Katta qiymat")
	// }
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i * i)
	// }
	// for i := 0; i <= 10; i = i + 2 {
	// 	fmt.Println(i)
	// }
	// i := 0
	// for i <= 10{
	// 	fmt.Println(i)
	// 	i = i + 2
	// }

	// var i int
	// for {
	// 	fmt.Scan(&i)
	// 	if i == 0 {
	// 		break
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	if i % 2 == 0 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// var a [4]int = [4]int{1, 2, 3}
	// b := [3]int{1,2, 3}
	// a[3] = 9
	// fmt.Println(a[2])
	// fmt.Println(a[3])
	// // fmt.Println(b)

	// c := [...]int{7,5,6}
	// fmt.Println(c)
	
	// b := [3]int{1, 3} // yo'q bo'lsa o'rniga 0 yozib qo'yadi
	// fmt.Println(b)

	// if a == b {
	// 	fmt.Println("A va B arraylar bir xil")
	// } else {
	// 	fmt.Println("A va B bir xil emas")
	// }

	// if c == b {
	// 	fmt.Println("C va B arraylar bir xil")
	// } else {
	// 	fmt.Println("C va B bir xil emas")
	// }

	// for i := 0; i < len(a); i++ {
	// 	a[i] = a[i] * a[i]
	// 	fmt.Println(a[i])
	// }
	// fmt.Println(a)

	// for index, value := range a {
	// 	fmt.Printf("Index %d, value %d\n", index, value)
	// }

	// for index := range a {
	// 	fmt.Printf("Index %d\n", index)
	// }

	// for _, value := range a {
	// 	value = 10
	// 	fmt.Println(value)
	// }
	// fmt.Println(a)
	// var a []int
	// var b []int = []int{1,2,3}
	// c := []int{1,2,3}
	// d := []int{1:10}

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// var n int
	// fmt.Scan(&n)
	// muar := make([]int, n, n)
	// fmt.Println(muar)

	// users := [5]string{"Tony", "Bob", "Curls", "Jon", "Poal"}
	// users2 := users[:3]
	// users2[0] = "Bek"
	// // fmt.Println(users[0], users[1])
	// // fmt.Println(users[1:2])
	// // fmt.Println(users[:3])
	// // fmt.Println(users[3:])
	// fmt.Println(users2)

	// m := []int{1,2,3}
	// // m = append(m, 4,5,6,7)
	// fmt.Println(m)
	// fmt.Println(len(m))
	// fmt.Println(cap(m))

	// pointer := fmt.Sprintf("%p", m)
	// fmt.Println(pointer)

	// a := []int{1,2,3,4,5,6,7,8,9}
	// a = append(a[:5], a[6:]...)
	// n := copy(a, m)
	// fmt.Println(n)
	// fmt.Println(a)

	// a := []int{1,2,3}
	// b := make([]int, 3,3)
	// fmt.Println(b)
	// myPrint(213457899)
	// name := "Ramazon"
	// surname := "Ergashev"
	// age := 14

	// printPerson(name, surname, age)

	// p := Person{
	// 	name: "Ramazon",
	// 	surname: "Ergashev",
	// 	age: 14,
	// }
	// p := Person{
	// 	name: 88,
	// 	surname: 33,
	// 	age: 14,
	// }

	// fmt.Printf("First name %s\n", p.name)
	// fmt.Printf("Surnamename %s\n", p.surname)
	// fmt.Printf("Age %d\n", p.age)

	// sscore := p.ScoreOverall()
	// fmt.Println(sscore)
	// myStrings()

	// var s string  = "hello worldtest pro"
	// fmt.Println(strings.Contains(s, "es")) // contains es borligini tekshirayapti
	// fmt.Println(strings.Count(s, "s")) // count nechta s borligini tekshirayapti
	// fmt.Println(strings.HasPrefix(s, "s")) // boshida kelganini tekshiradi
	// fmt.Println(strings.Index(s, "s")) // s nechanchi indeksda ekanligini aytadi
	// fmt.Println(strings.Join([]string{"Tony", "Jekson"}, "s")) // 2ta arrayni birlashtiradi va oraga s harfini qo'shib qo'yadi
	// fmt.Println(strings.Repeat(s, 2)) // s ni 2 marta qaytaradi
	// fmt.Println(strings.Replace("Mysss", "sss", " ", -1)) // Mysss ni tekshirib sss ni olib o'rniga probel qo'yadi
	// fmt.Println(strings.Split("a-b-c-d", "-")) // a-b-c lardan - ni olib tashlaydi va array holida qaytaradi
	// fmt.Println(strings.ToLower("AWESDFGVH"))
	// fmt.Println(strings.ToUpper("dfxfgchcg"))
	// Example()
	// fmt.Println(unicode.IsDigit('1'))  // 1 raqam ekanligini tekshiradi
	// fmt.Println(unicode.IsLetter('1'))  // 1 string ekanligini tekshiradi
	// fmt.Println(unicode.IsLower('1'))  // 1 kichkina harf ekanligini tekshiradi
	// fmt.Println(unicode.IsUpper('1'))  // 1 katta harf ekanligini tekshiradi
	// fmt.Println(unicode.IsSpace('1'))  // 1 bo'sh joy ekanligini tekshiradi
	// fmt.Println(unicode.Is(unicode.Latin, 's')) // s lotin ekanligini teshirayapti
	// var uz string = "Hello world"
	// var ru string = "ㅁㄴㅇㄹㅎ ㅂㅈㄷㄱㅅ"
	// fmt.Println(utf8.RuneCountInString(uz), utf8.RuneCountInString(ru)) // lendan yaxshi variant

	// map[key]value{}
	// a := map[string]string{"one":"Hello", "two":"World", "three":"Simple"}
	// a := map[int]string{1:"Hello", 2:"World", 3:"Simple"}
	// fmt.Println(a)
	// fmt.Println(a["one"])
	// for key, value := range a {
	// 	fmt.Printf("%d is key for value %q\n", key, value)
	// }
	// a := map[int]string{1:"Hello", 2:"World", 3:"Simple"}
	// a[2] = "Dunyo"
	// // delete(map, key)
	// delete(a, 2)
	// keys := []int{}
	// values := []string{}
	// for key := range a {
	// 	keys = append(keys, key)
	// }
	// for _, value := range a {
	// 	values = append(values, value)
	// }
	// fmt.Println(keys)
	// fmt.Println(values)
	// sort.Strings(values)  // alifbe tartibida valuelarni joylashtirib beradi
	// fmt.Println(values)
	// sort.Ints(keys) // keylarni (sonlarni) ketma-ket joylashtirib beradi
	// fmt.Println(keys)

	// lang := map[string]string{
	// 	"Hello" : "Salom",
	// }
	// var input string
	// fmt.Scan(&input)
	// for key, value := range lang{
	// 	if key == input {
	// 		fmt.Println(value)
	// 		break
	// 	}
	// }
	// var el Employee
	// el = Emp(1)
	// el.PrintName("Tony Jackson")
	// fmt.Println("Employee salary: ", el.PrintSalary(200, 70))
	// circle := Circle(Radius: 5)
	// rectangle := Rectangle(
	// 	width: 10,
	// 	height: 7
	// )
	// fmt.Println(circle)
	// fmt.Println(rectangle)
	// printArea(circle)
	// printArea(rectangle)
	// person := Person(Name: "Alice", Age: 18) // fix it missing ',' in argument listsyntax
	// book := Book(Title: "Amir Temur", Author: "Temurbek")

	// person.Describer()
	// book.Descripber()
	// go myFunc()
	// time.Sleep(time.Second * 60) // myfunc ni ishlatib 60 sekund kutib turadi
	// go myPrint("Salom") // go kalit so'zini qo'ymasam salom so'zi cheksiz davom etadi
	// myPrint("Alik") // go myPrint("Alik") uchun pastga time.Sleep(time.Second*5)

	// var wg sync.WaitGroup
	// wg.Add(2) // shu wait groupga nechta goroutine ni kutishini berib o'tamiz
	// go myPrint("Salom", &wg)
	// go myPrint("Alik", &wg)

	// wg.Wait()


	// Anonim funksiya 
	// var wg sync.WaitGroup
	// wg.Add(2) // shu wait groupga nechta goroutine ni kutishini berib o'tamiz
	// go func ()  {
	// 	myPrint("Salom", &wg)
	// 	wg.Done()
	// }()
	// go func ()  {
	// 	myPrint("Alik", &wg)
	// 	wg.Done()
	// }()
	
	// // wg.Wait()
	// channel := make(chan int, 3) // Channel yasab olish
	// channel <- 1 // Channelga qiymat berish
	// channel <- 4 // Channelga qiymat berish
	// channel <- 5 // Channelga qiymat berish
	// // a := <- channel  // Channeldagi qiymatni oz'garuvchiga olish
	// fmt.Println(<-channel) // Channeldagi qiymatni chiqarib olish
	// fmt.Println(<-channel) // Channeldagi qiymatni chiqarib olish
	// fmt.Println(<-channel) // Channeldagi qiymatni chiqarib olish
	// fmt.Println(a) // Channeldagi qiymatni chiqarib olish

	// messageChannel := make(chan int)
	// go sendMessage(messageChannel)
	// for i := 1; i <= 5; i++ {
	// 	message := <-messageChannel
	// 	fmt.Println("Received: ", message)
	// }

	// close(messageChannel)

	// numbers := make(chan string, 3)
	// numbers <- "one"
	// numbers <- "two"
	// numbers <- "three"
	// close(numbers)
	// // numbers <- "four"
	// for v := range numbers {
	// 	fmt.Println(v)
	// }

	// c := SCounter{v : make(map[string]int)}
	// for i := 0; i < 100; i++ {
	// 	go c.Increment("key")
	// }
	// time.Sleep(time.Second * 3)
	// fmt.Println(c.v["key"])

	// channel1 := make(chan string)
	// channel2 := make(chan string)

	// go func ()  {
	// 	for {
	// 		channel1 <- "Fast"
	// 		time.Sleep(time.Nanosecond*1)
	// 	}
	// }()
	
	// go func ()  {
	// 	for {
	// 		channel2 <- "Slow"
	// 		time.Sleep(time.Second*2)
	// 	}
	// }()

	// for {
	// 	fmt.Println(<-channel1)
	// 	fmt.Println(<-channel2)
	// }
	// for {
	// 	select {
	// 		case msg1 := <- channel1:
	// 			fmt.Println(msg1)
	// 		case msg2 := <- channel2:
	// 			fmt.Println(msg2)
	// 		default:
	// 			fmt.Println("Default message")
	// 	}
	// }
}

// func sendMessage(ch chan int)  {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Sending ", i)
// 		ch <- i
// 		time.Sleep(time.Second)
// 	}

// 	close(ch)
// }

// func printArea(shape Shape)  {
// 	fmt.Printf("Area %f\n", shape.Area())
// }

// func myFunc()  {
// 	fmt.Println("Hello")
// }

// func myPrint(str string)  {
// 	for i := 1; true; i++ {
// 		fmt.Println(i, str)
// 		time.Sleep(1 * time.Second)
// 	} // cheksiz sikldagi for
// }

// func myPrint(str string, wg *sync.WaitGroup)  {
// 	for i := 1; i<=5; i++ {
// 		fmt.Println(i, str)
// 		time.Sleep(1 * time.Second)
// 	} // cheksiz sikldagi for
// 	wg.Done()
// }