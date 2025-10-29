package main

import "fmt"

/*func main() {
	s1 := make([]int, 0, 3)  // len = 0 cap = 3 s1 = [0,0,0]
	s1 = append(s1, 1, 2, 3) // len = 3 cap = 3 s1 = [1,2,3]

	s2 := s1[:2]         // len = 2 cap = 3 base on s1 s2 = [1,2],3,  s1 = [1,2,3]
	s3 := append(s2, 99) // len = 3 cap = 3 base on s1
	// s3 = [1,2,99]
	//s2 = [1,2],99
	//s1 = [1,2,99]

	s2[0] = 77
	// s3 = [77,2,99]
	//s2 = [77,2],99
	//s1 = [77,2,99]
	s3[1] = 88
	// s3 = [77,88,99]
	//s2 = [77,88],99
	//s1 = [77,88,99]

	s1 = append(s1, 100) // len = 4 cap = 6 new array s1 = [77,88,99,100]0,0,
	// s3 = [77,88,99]
	//s2 = [77,88],99

	s3 = append(s3, 200) //len = 4 cap = 6 new array s3 = [77,88,99,200],0,0
	fmt.Println("E:", "s1:", s1, "s2:", s2, "s3:", s3)
	//s1 = [77,88,99,100]
	// s2 = [77,88]
	// s3 = [77,88,99,200]

}


func main() {
	s := make([]int, 0, 2)
	s = append(s, 1, 2)

	fmt.Printf("main: len=%d cap=%d ptr=%p\n", len(s), cap(s), &s[0])
	grow(s)
	fmt.Printf("after grow: len=%d cap=%d ptr=%p\n", len(s), cap(s), &s[0])
	growPtr(&s)
	fmt.Printf("after growPtr: len=%d cap=%d ptr=%p\n", len(s), cap(s), &s[0])

}
func grow(s []int) {
	fmt.Printf("func: before append len=%d cap=%d ptr=%p\n", len(s), cap(s), &s[0])
	s = append(s, 3) // возможно создаст новый backing array
	fmt.Printf("func: after  append len=%d cap=%d ptr=%p\n", len(s), cap(s), &s[0])
}

func growPtr(ps *[]int) {
	*ps = append(*ps, 3)
}


func main() {
	a := make([]int, 0, 4)    // len=0 cap=4 a=[_,_,_,_]
	a = append(a, 10, 20, 30) //len=3 cap=4 a=[10,20,30,_]
	fmt.Println("A:", a, "len:", len(a), "cap:", cap(a))

	b := a[:2]               // len=2 cap=4 b -> a backing array b=[10,20,_,_]
	c := make([]int, len(b)) //len=2 cap=2 new backing array c=[_,_]
	copy(c, b)               //len=2 cap=2 c=[10,20]
	fmt.Println("B:", "a:", a, "b:", b, "c:", c)
	fmt.Println("B:", b, "len:", len(b), "cap:", cap(b))

	b[0] = 100 // b->[100,20,_,_] a->[100,20,30,_]
	c[1] = 200 //c->[10,200]
	fmt.Println("C:", "a:", a, "b:", b, "c:", c)

	a = append(a, 40) //len=4 cap=4 a->[100,20,30,40]
	fmt.Println("D:", "a:", a, "b:", b, "c:", c)

	c = append(c, 300) // len=3 cap=4 new backing array for c->[10,200,300,_]
	fmt.Println("E:", "a:", a, "b:", b, "c:", c)
}

Концептуальные вопросы (уровень middle)
Почему copy(c, b) не создаёт общую память, хотя внутри всё копируется поэлементно?
Что произойдёт, если изменить строку a = make([]int, 0, 3)?
Если сделать b = a[:3] перед append(a, 40) — изменится ли момент создания нового массива?
Как сделать так, чтобы c разделял память с a, но имел свою длину и cap?

Почему выгодно иногда использовать copy() после append() (намёк: чтобы отрезать старый cap и не удерживать большой backing array)?
1) Потому что копируются только значения и еще потому что c и b могут отличаться по длине и берется минимальная длина оного из 2-х массивов 2) Потеряется связь между a и b, ёмкость а увеличится вдвое до 6 и старый массив скопируется в новый и сылка в дескрипторе слайса изменится и ссылка дескриптора b уже бедет ссылаться на массив которого уже возможно нет потому что GC удалил его из памяти 3) если до этого было условие a = make([]int, 0, 3) то да произойдут изменения которые я описал в п.2 4) невозможно 5) чтобы отрезать старый cap и не удерживать большой backing array


Задача №3 — “Трёхуровневое влияние и полный срез”
Эта задача тренирует:
понимание, когда append делит и разрывает память,
работу с full slice expression [:len:len] (для “отрезания cap”),
и то, как изменение одного среза отражается (или не отражается) в других.


func main() {
	s1 := make([]int, 0, 5)  //s1->[_][_][_][_][_]
	s1 = append(s1, 1, 2, 3) // s1->[1][2][3][_][_]

	s2 := s1[:2]   // len=2 cap=5, общая память c s1 s2->[1][2][_][_][_]
	s3 := s2[:1:1] // s3->[1] len=1 cap=1 — отрезаем cap (новый слайс, но тот же backing array, ограниченный)

	s3 = append(s3, 100) // должен создать новый массив, т.к. cap(s3)=1
	//s3->[1][100] new backing array len=2 cap=2

	s2 = append(s2, 200) // cap(s2)=5, хватает места — изменит s1
	//s2->[1][2][200][_][_]
	//s1->[1][2][200][_][_]

	s1 = append(s1, 300, 400, 500) //будет len=6 cap=10
	//создаст новый backing array и отвяжется от s2
	//s1->[1][2][200][300][400][500]
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
//s1->[1][2][200][300][400][500]
//s2->[1][2][200][_][_]
//s3->[1][100]
}
Вопросы от интервьюера
Почему s3 не изменил s1, хотя был срезан из него?
Что делает третий параметр в выражении [:1:1]?
Как можно было бы “отрезать” cap без полного среза, но с помощью copy()?
Что произойдёт, если после строки s2 := s1[:2] добавить s1 = append(s1, 999)?
Если удалить строку [:1:1], останется ли s3 независимым?


func main() {
	base := []int{1, 2, 3, 4} // len=4 cap=4 base ->[1 2 3 4]
	group := [][]int{
		base[:2],
		base[2:],
	}
	//group len=2 cap=2
	//len[0]=2 cap[0]=4 group[0]->[1 2]
	//len[1]=2 cap[1]=4 group[1]->[3 4]

	group[0][1] = 100
	//len[0]=2 cap[0]=4 group[0]->[1 100]
	//base ->[1 100 3 4]
	group[1][0] = 200
	//len[1]=2 cap[1]=4 group[1]->[200 4]
	//base ->[1 100 200 4]
	newBase := append(base, 999)
	//создаётся новый массив для newBase так как в емкости Base не хватает append незатрагивает base
	//newBase->[1 100 200 4 999] len=5 cap=8
	fmt.Println("Base:", base)
	fmt.Println("newBase:", newBase)

	group = append(group, newBase[:2]) //
	//group len=3 cap=4
	//len[0]=2 cap[0]=4 group[0]->[1 2]
	//len[1]=2 cap[1]=4 group[1]->[3 4]
	//len[2]=2 cap[1]=2 group[1]->[1 2] new backing array

	group[2][1] = 500
	//len[0]=2 cap[0]=4 group[0]->[1 2]
	//len[1]=2 cap[1]=4 group[1]->[3 4]
	//len[2]=2 cap[1]=2 group[1]->[1 500] new backing array
	//newBase->[1 500 200 4 999] len=5 cap=8

	fmt.Println("base:", base)       //base ->[1 100 200 4]
	fmt.Println("newBase:", newBase) //newBase->[1 500 200 4 999]
	fmt.Println("group:", group)
	//len[0]=2 cap[0]=4 group[0]->[1 2]
	//len[1]=2 cap[1]=4 group[1]->[3 4]
	//len[2]=2 cap[1]=2 group[1]->[1 500]
}



// 🧩 **Задача 1 — Разделённый append**

func main() {
	s := []int{1, 2, 3} //len=3 cap=3
	a := s[:2]          //a->[1 2] len=2 cap=3
	b := append(a, 10)  //b->[1 2 10]
	a[0] = 99
	//s->[99 2 3]
	//a->[99 2]
	//b->[99 2 10]

	fmt.Println("s:", s, "a:", a, "b:", b)
}


//🧠 Вопросы:

* Делят ли `a` и `b` общий массив?
* Изменится ли `s` после `a[0]=99`?



// 🧩 **Задача 2 — Потеря связи при append**

func main() {
	s := make([]int, 0, 2) //l=0 c=2 []
	s = append(s, 5)       //l=1 c=2 s>[5]
	t := append(s, 6)      //l=2 c=2 t->[5 6]
	s = append(s, 7)       // l=2 c=2 s->[5 7]
	//t->[5 7]
	fmt.Println("s:", s, "t:", t)
}

//🧠 Вопрос: после второго `append`, делят ли `s` и `t` общий backing array?


// 🧩 **Задача 3 — Срез из среза и изменение через общий массив**

func main() {
	s1 := []int{10, 20, 30, 40} //len=4 cap=4 s1->[10 20 30 40]
	s2 := s1[1:3]               //len=2 cap=4 s2->[20 30]
	s3 := s2[1:]                //len=1 cap=4 s3->[30]
	s3[0] = 999                 //len=1 cap=4 s3->[999]
	//s1->[10 20 999 40]
	//s2->[20 999]
	//s3->[999]
	fmt.Println("s1:", s1, "s2:", s2, "s3:", s3)
}


//🧠 Определи:

//* Что будет напечатано?
//* Почему изменение `s3` повлияло на `s1`?

//

// 🧩 **Задача 4 — Полный срез и обрезание cap**

func main() {
	data := []int{1, 2, 3, 4} // len=4 cap=4 data->[1 2 3 4]
	left := data[:2:2]        //len = 2 cap=2 left->[1 2]
	right := append(left, 10) //len=3 cap=4 right->[1 2 10]
	//data->[1 2 3 4]
	//left->[1 2]
	//right->[1 2 10]
	fmt.Println("data:", data, "left:", left, "right:", right)
}

//🧠 Почему `right` не изменил `data`?


// 🧩 **Задача 5 — Copy vs общий массив**

func main() {
	a := []int{1, 2, 3}      //len=3 cap=3 a->[1 2 3]
	b := make([]int, len(a)) // len=3 cap=3 b->[]
	copy(b, a)               //len=3 cap=3 b->[1 2 3]
	b[0] = 100
	//a->[1 2 3]
	//[100 2 3]
	fmt.Println("a:", a, "b:", b)
}

//🧠 Почему `a` не изменился? (Подумай, что делает `copy` с backing array.)


// 🧩 **Задача 6 — Удаление элемента и reslice**

func main() {
	s := []int{1, 2, 3, 4, 5} //len=5 cap=5 s->[1 2 3 4 5]
	i := 2
	s = append(s[:i], s[i+1:]...) //s->[1 2 4 5]
	fmt.Println("s:", s)
	fmt.Println("s cap:", cap(s))
	s = s[:cap(s)] //s->[1 2 4 5 5]
	fmt.Println("s:", s)
}

//🧠 Что будет, если после этого сделать `s = s[:cap(s)]`?


// 🧩 **Задача 7 — Вложенные слайсы (slice of slice)**

func main() {
	a := []int{1, 2, 3, 4} //len=4 cap=4 a=[1 100 3 4]
	group := [][]int{a[:2], a[2:]}
	//group ->{
	//[1 100] len2 cap4
	//[3 4] len2 cap4
	//}
	group[0][1] = 100
	fmt.Println("a:", a, "group:", group)
	//a=[1 100 3 4]
	////group ->{
	////[1 100] len2 cap4
	////[3 4] len2 cap4
	////}
}

//🧠 Почему изменение `group[0]` изменило `a`?


// 🧩 **Задача 8 — Разделённая память и append в под-слайсе**

func main() {
	a := []int{10, 20, 30, 40} // len4 cap4 a->[10 20 30 40]
	b := a[:2]                 //len2 cap4 b->[10 20]
	c := append(b, 99)         // c->[10 20 99]
	fmt.Println("a:", a, "b:", b, "c:", c)
	//a->[10 20 99 40]
	//b->[10 20]
	//c->[10 20 99]
}

//🧠 Где произойдёт изменение — в `a` или только в `c`?


// 🧩 **Задача 9 — Слияние слайсов**



func main() {
	a := []int{1, 2}     //a->[1 2]
	b := []int{3, 4}     //b->[3 4]
	c := append(a, b...) //c->[1 2 3 4]
	b[0] = 999           //[99 4]
	fmt.Println("a:", a, "b:", b, "c:", c)
	//a->[1 2]
	//b->[99 4]
	//c->[1 2 3 4]

}


//🧠 Почему `c` не изменился после `b[0]=999`?

*/

// 🧩 **Задача 10 — Полное переполнение и утрата связи**

func main() {
	s1 := make([]int, 0, 2) // l2 c2 s1->[]
	s1 = append(s1, 1, 2)   // s1->[1 2]
	s2 := s1                //s2->[1 2]
	s1 = append(s1, 3)      //s1->[1 2 3] new backing array
	s1[0] = 100             //s1->[100 2 3]
	fmt.Println("s1:", s1, "s2:", s2)
	//s1->[100 2 3]
	//s2->[1 2]
}

//🧠 Определи:

//* Момент, когда создаётся новый backing array.
//* Почему `s2` остался “старым”?

// 📘 Как лучше их проходить
