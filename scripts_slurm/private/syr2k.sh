#!/bin/bash
cd samples
cd syr2k
nohup srun -J syr2k -w compasslab5 --output=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/syr2k.out --error=/home/sbin/dockerVol_mcm/mgvm/scripts_slurm/private/output/syr2k.err ./syr2k -timing -no-progress-bar -report-all -scheduling lasp -platform-type privatetlb -mem-allocator-type lasp -use-lasp-mem-alloc -sched-partition Xdiv -max-inst 3000000 -ni=16384 -nj=16384 &