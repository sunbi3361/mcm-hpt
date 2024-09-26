#!/bin/bash
cd samples
cd fir
nohup srun -J fir -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fir.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/fir.err ./fir -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -length=536870912 &