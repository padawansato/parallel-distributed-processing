#.shでないファイルを入ってから渡すときgnuplot>load 'fuga’とする．                                                                        
#$file="result/100-100-1000000.csv"
#ファイル名はplot"result/hoge/mergesort\_"が無い方がいい

#空行抜き
gnuplot << EOF
   set datafile separator ","
   set xlabel 'size'
   set xtics 0,200000,1000000
   set ylabel 'time'
   set style fill solid border lc rgb "black"
   plot 'result/mergesort_pointer_100-100-1000000.csv' using 1:2 with lines lw 2  title "mergesort_pointer.go",\
        'result/pmergesort_pointer_buildin_100-100-1000000.csv' using 1:2 with lines lw 2 title "pmergesort\_pointer\_buildin.go"
   set terminal svg
   set output 'result-c.svg'
   set terminal png
   set output "result-c.png"
   replot
   set output
EOF
