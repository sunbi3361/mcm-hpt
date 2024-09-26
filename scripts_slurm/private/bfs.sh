#!/bin/bash
cd samples
cd bfs
nohup srun -J bfs -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs.err ./bfs -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 300000000 -node 32768 -degree 32768 -depth 0 &