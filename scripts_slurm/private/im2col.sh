#!/bin/bash
cd samples
cd im2col
nohup srun -J im2col -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/im2col.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/im2col.err ./im2col -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -N=16 -C=3 -H=1024 -W=1024 -kernel-height=3 -kernel-width=3 &