package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"math/rand"
	//"sort"
	"os"
	"strconv"
)

func merge(values *[]int, start int, stop int) {
	tmp_value := make([]int, (start+stop)/2+1)	// コピーをとって作業する
	copy(tmp_value,(*values)[start:(start+stop)/2+1])
	k := start
	i := start
	j := (start+stop)/2+1

	// どちらかの要素がなくなるまでマージする
	for (start+stop)/2-i >= 0 && (stop-j) >= 0 {
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
	// 片方の要素がなくなったら，もう片方を詰めていく
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
	// id が（ 2^~ まで）マージソートする
	if id < 4 {
		// 分割・再帰
		var wg sync.WaitGroup
		if (start < stop) {
			wg.Add(2)
			go func(){
				merge_sort(values, start , (start+stop)/2,id+1)
				wg.Done()
			}()
			go func(){
				merge_sort(values, (start+stop)/2+1 , stop,id+1)
				wg.Done()
			}()
			wg.Wait()

			merge(values, start , stop)
		}
	} else {
		bubble_sort(values,start,stop+1)	// 遅いソート
	}
	return
}


func merge_Sort(values *[]int, n int) {
	merge_sort(values, 0, n,1)
	return
}

func main() {

	// ランダムな値生成
	t := time.Now().Unix()
	s := rand.NewSource(t)
	r := rand.New(s)

	N , _ := strconv.Atoi(os.Args[1])
	fmt.Println("N=",N)
	values := make([]int, N)
	for i := 0; i < N; i++ {
		values[i] = r.Intn(N)
	}
	//fmt.Println(values)	// 結果確認用

	// ソート
	log.Print("Start")
	merge_Sort(&values, N-1)
	log.Print("end")
	//fmt.Println(values)	// 結果確認用
}


func bubble_sort(values *[]int, start int, stop int) {
	for i:=start; i<stop; i++ {
		for j:=i+1; j<stop; j++{
			if (*values)[i] > (*values)[j] {
				t := (*values)[i]
				(*values)[i] = (*values)[j]
				(*values)[j] = t
			}
		}
	}
}
