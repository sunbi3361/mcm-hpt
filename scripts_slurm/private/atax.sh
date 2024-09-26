#!/bin/bash
cd samples
cd atax
nohup srun -J atax -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/atax.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/atax.err ./atax -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -x=32768 -y=32768 &