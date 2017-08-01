package main

import (
	"fmt"
//	"log"
	"sync"
	"time"
	"sort"
	"os"
	"strconv"
	"math/rand"
)

func merge(values *[]int, start int, stop int) {
	tmp_value := make([]int, (start+stop)/2+1)
	copy(tmp_value,(*values)[start:(start+stop)/2+1])
	k := start
	i := start
	j := (start+stop)/2+1
	for ; (start+stop)/2-i >= 0 && (stop-j) >= 0; {
		if tmp_value[0] < (*values)[j] {
			(*values)[k] = tmp_value[0]
			tmp_value = tmp_value[1:]
			i++
			k++
		} else {
			(*values)[k] = (*values)[j]
			j++
			k++
		}
	}
	// 片方のスライスから追加する要素がなくなったら残りは単純に連結できる (各スライスは既にソートされているため)
	if (start+stop)/2-i < 0 {
		for k<=stop {
			(*values)[k] = (*values)[j]
			k++
			j++
		}
	}else{
		for k<=stop {
			(*values)[k] = tmp_value[0]
			tmp_value = tmp_value[1:]
			k++
			i++
		}
	}
}

func merge_sort(values *[]int, start int, stop int, id int) {
	// ふたつのスライスをそれぞれ再帰的にソートする
	if id < 4 {
		var wg sync.WaitGroup
		if (start < stop) {
			wg.Add(2)
			merge_sort(values, start , (start+stop)/2, id+1)
			merge_sort(values, (start+stop)/2+1 , stop, id+1)

			merge(values, start , stop)
		}
	} else {
		sort.Ints((*values)[start:stop+1])
	}
	return
}

func Sort(values *[]int, n int) {
	merge_sort(values, 0, n, 1)
	return
}

func main() {
	// UNIX 時間をシードにして乱数生成器を用意する
	t := time.Now().Unix()
	s := rand.NewSource(t)
	r := rand.New(s)

	// ランダムな値の入った配列を作る
	N , _ := strconv.Atoi(os.Args[1])
	fmt.Println("N=",N)
	values := make([]int, N)
	for i := 0; i < N; i++ {
		values[i] = r.Intn(N)
	}

//	log.Print("Start")
	start :=time.Now();
	// ソートして結果を出力する
	Sort(&values, N-1)
//	log.Print("end")
	//fmt.Println(values)
	end :=time.Now();
	fmt.Printf("%f\n",(end.Sub(start)).Seconds())//秒
}
