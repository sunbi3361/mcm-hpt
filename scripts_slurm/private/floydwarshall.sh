#!/bin/bash
cd samples
cd floydwarshall
nohup srun -J floydwarshall -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/floydwarshall.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/floydwarshall.err ./floydwarshall -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -node=16384 -iter=0 &