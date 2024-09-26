#!/bin/bash
cd samples
cd gemm
nohup srun -J gemm -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gemm.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gemm.err ./gemm -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -num_i=16384 -num_j=16384 -num_k=16384 &