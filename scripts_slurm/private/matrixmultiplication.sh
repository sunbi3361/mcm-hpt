#!/bin/bash
cd samples
cd matrixmultiplication
nohup srun -J matrixmultiplication -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/matrixmultiplication.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/matrixmultiplication.err ./matrixmultiplication -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -x=32768 -y=16384 -z=16384 &