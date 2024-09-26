#!/bin/bash
cd samples
cd sssp
nohup srun -J sssp -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/sssp.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/sssp.err ./sssp -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 300000000 -numNodes=134217728 -numItems=268435456 &