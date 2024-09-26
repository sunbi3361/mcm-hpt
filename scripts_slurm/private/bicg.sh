#!/bin/bash
cd samples
cd bicg
nohup srun -J bicg -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bicg.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bicg.err ./bicg -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -x=32768 -y=32768 &