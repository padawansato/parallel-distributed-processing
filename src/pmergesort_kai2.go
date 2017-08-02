package main

import (
	//"fmt"
	"log"
	"sync"
	"time"
	"math/rand"
)

func merge(values *[]int, start int, stop int) {
	tmp_value := make([]int, (start+stop)/2+1)
	copy(tmp_value,(*values)[start:(start+stop)/2+1])
	k := start
	i := start
	j := (start+stop)/2+1
	//for (start+stop)/2-start > 0 && stop-((start+stop)/2+1) > 0 {
	for ; (start+stop)/2-i >= 0 && (stop-j) >= 0; {
		//for i=start; (start+stop)/2-i > 0 && (stop-j) > 0; {
		// ソート済みのふたつのスライスからより小さいものを選んで追加していく (これがソート処理)
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

func sort(values *[]int, start int, stop int) {
	// ふたつのスライスをそれぞれ再帰的にソートする
	var wg sync.WaitGroup
	if (start < stop) {
		wg.Add(2)
		go func(){
			sort(values, start , (start+stop)/2)
			wg.Done()
		}()
		go func(){
			sort(values, (start+stop)/2+1 , stop)
			wg.Done()
		}()
		wg.Wait()

		// ソート済みのふたつのスライスをひとつにマージする
		//(*values)[start:stop] = merge((*values)[start:stop/2], (*values)[stop/2+1:stop])
		//*values = append((*values)[:start], append(merge((*values)[start:stop/2], (*values)[stop/2+1:stop]), (*values)[stop:]...)...)

		merge(values, start , stop)
	}
	return
}

//func split(values []int) (left, right []int) {
// スライスを真ん中でふたつに分割する
//left = values[:len(values) / 2]
//right = values[len(values) / 2:]
//return
//}

func Sort(values *[]int, n int) {
	sort(values, 0, n)
	return
}

func main() {
	// UNIX 時間をシードにして乱数生成器を用意する
	t := time.Now().Unix()
	s := rand.NewSource(t)
	r := rand.New(s)

	// ランダムな値の入った配列を作る
	N := 100000
	values := make([]int, N)
	for i := 0; i < N; i++ {
		values[i] = r.Intn(N)
	}

	//values := []int{6, 1, 4, 7, 0 ,9, 5, 2, 2, 5}
	log.Print("Start")

	// ソートして結果を出力する
	//sortedValues := Sort(&values, N)
	Sort(&values, N-1)
	//Sort(values)
	log.Print("end")
	//fmt.Println(values)
}
