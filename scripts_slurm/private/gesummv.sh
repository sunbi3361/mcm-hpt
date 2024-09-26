#!/bin/bash
cd samples
cd gesummv
nohup srun -J gesummv -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gesummv.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/gesummv.err ./gesummv -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -n=24576 &