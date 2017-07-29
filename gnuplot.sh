#.shでないファイルを入ってから渡すときgnuplot>load 'fuga’とする．                                                                        
#$file="result/100-100-1000000.csv"

#空行抜き
gnuplot << EOF
   set datafile separator ","
   set xlabel 'size'
   set ylabel 'time'
   set style fill solid border lc rgb "black"
   plot 'result/mergesort\_100-100-1000000.csv' using 1:2 with lines lw 2  title "mergesort.go"
   plot 'result/mergesort\_pointer\_100-100-1000000.csv' using 1:2 with lines lw 2  title "mergesort.go"
   plot 'result/pmergesort\_pointer\_buildin\_100-100-1000000.csv' using 1:2 with lines lw 2 title "pmergesort_pointer_buildin.go"
   plot 'result/pmergesort\_pointer\_bubble\_100-100-1000000.csv' using 1:2 with lines lw 2 title "pmergesort_pointer_bubble.go"
   set output 'result.svg'
   replot
   set output
   set terminal png
   set output "result.png"
   replot
EOF
