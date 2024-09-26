#!/bin/bash
cd samples
cd matrixtranspose
nohup srun -J matrixtranspose -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/matrixtranspose.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/matrixtranspose.err ./matrixtranspose -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 10000000 -width=24576 &