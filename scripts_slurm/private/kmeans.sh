#!/bin/bash
cd samples
cd kmeans
nohup srun -J kmeans -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/kmeans.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/kmeans.err ./kmeans -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 100000000 -points=4194304 -features=128 -clusters=20 -max-iter=1 &