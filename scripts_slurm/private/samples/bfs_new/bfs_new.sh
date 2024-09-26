#!/bin/bash
# ./bfs_new -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 300000000 -load-graph Slashdot0902.txt -depth 0 &
nohup srun -J bfs_new -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs_new.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs_new.err ./bfs_new -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -node 268435456 -degree 524288 -depth 3 &
# nohup srun -J bfs_new -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs_new.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/bfs_new.err ./bfs_new -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 1000000000 -node 524288 -degree 8192 -depth 0 &