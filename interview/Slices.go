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
*/
