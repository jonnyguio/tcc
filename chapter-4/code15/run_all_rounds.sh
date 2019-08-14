#!/bin/bash
clang++ -lpthread --std=c++11 --stdlib=libc++ -Ofast main.cpp -o c_matrix || exit 1;
final_message="Ended execution. Following times:"
step_file_header="sample_number,cores,matrix_size,time"
date=`date '+%Y-%m-%d-%X'`
#for matrix_size in "1000";
for matrix_size in "100" "1000" "2000" "3000" "4000" "5000";
do
	#for cores in {1..2}
	for cores in "2" "4";
	do
		rm execution-$cores-$matrix_size.csv
		echo "Step: number of cores is $cores size of matrix is ${matrix_size}x${matrix_size}"
		final_message="\n\nStep: number of cores is $cores size of matrix is ${matrix_size}x${matrix_size}"
		echo $step_file_header >> execution-$cores-$matrix_size.csv
		for step in {1..50};
		#for step in {1..2};
		do
			sample=`./c_matrix $matrix_size $cores`
			step_file_content="$step,$cores,$matrix_size,$sample"
			echo $step_file_content >> execution-$cores-$matrix_size.csv
		done
		echo "Sending email"
		echo $final_message | mutt -s "[tcc][$date][cores $cores][matrix $matrix_size] relatório execução multiplicação de matrizes c" -a execution-$cores-$matrix_size.csv -- joaoguio@protonmail.com
	done
done
