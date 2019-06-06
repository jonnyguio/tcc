#!/bin/bash
go build -o go_matrix || exit 1;
final_message="Ended execution. Following times:"
step_file_header="sample_number,cores,matrix_size,time"
date=`date '+%Y-%m-%d-%X'`
#for matrix_size in "1000";
for matrix_size in "100" "1000" "5000" "10000" "25000" "50000" "90000";
do
	#for cores in {1..2}
	for cores in "1" "2" "4";
	do
		rm execution-$cores-$matrix_size.csv
		echo "Step: number of cores is $cores size of matrix is ${matrix_size}x${matrix_size}"
		final_message="\n\nStep: number of cores is $cores size of matrix is ${matrix_size}x${matrix_size}"
		echo $step_file_header >> execution-$cores-$matrix_size.csv
		for step in {1..1000};
		#for step in {1..2};
		do
			if [ "$cores" -eq "1" ];
			then
				sample=`./go_matrix $matrix_size false`
			else
				sample=`GOMAXPROCS=$cores ./go_matrix $matrix_size true`
			fi
			step_file_content="$step,$cores,$matrix_size,$sample"
			echo $step_file_content >> execution-$cores-$matrix_size.csv
		done
		echo "Sending email"
		echo $final_message | mutt -s "[TCC][$date][Cores $cores][Matrix $matrix_size] Relatório execução multiplicação de matrizes go" -a execution-$cores-$matrix_size.csv -- joaoluisguio@gmail.com
	done
done
