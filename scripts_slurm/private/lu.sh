#!/bin/bash
cd samples
cd lu
nohup srun -J lu -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/lu.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/lu.err ./lu -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -n=32768 -k=1 &